package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"simplebank/util"

	dsUsr "simplebank/db/datastore/user"
	"simplebank/model"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

// type createUserRequest struct {
// 	LoginName    string `json:"loginName" binding:"required,alphanum"`
// 	UserPassword string `json:"password" binding:"required,min=6"`

// 	Iiid         int64 `json:"iiid"`
// 	AccessRoleID int64 `json:"accessRoleId"`
// 	StatusID     int64 `json:"statusId"`
// 	// OtherInfo    []byte `json:"otherInfo"`
// }

// type userResponse struct {
// 	ID                int64          `json:"id"`
// 	Uuid              uuid.UUID      `json:"uuid"`
// 	Iiid              int64          `json:"iiid"`
// 	TitleID           int64          `json:"titleID"`
// 	TitleUuid         uuid.UUID      `json:"titleUuid"`
// 	Title             string         `json:"title"`
// 	LastName          string         `json:"lastName"`
// 	FirstName         sql.NullString `json:"firstName"`
// 	MiddleName        sql.NullString `json:"middleName"`
// 	MotherMaidenName  sql.NullString `json:"motherMaidenName"`
// 	Birthday          sql.NullTime   `json:"birthday"`
// 	Sex               sql.NullBool   `json:"sex"`
// 	GenderID          int64          `json:"genderID"`
// 	GenderUuid        uuid.UUID      `json:"genderUuid"`
// 	Gender            string         `json:"gender"`
// 	CivilStatusID     int64          `json:"civilStatusID"`
// 	CivilStatusUuid   uuid.UUID      `json:"civilStatusUuid"`
// 	CivilStatus       string         `json:"civilStatus"`
// 	AlternateID       sql.NullString `json:"alternateID"`
// 	Phone             sql.NullString `json:"phone"`
// 	Email             sql.NullString `json:"email"`
// 	IdentityMapID     sql.NullInt64  `json:"identityMapID"`
// 	SimpleName        sql.NullString `json:"simpleName"`
// 	LoginName         string         `json:"loginName"`
// 	StatusID          int64          `json:"statusID"`
// 	StatusUuid        uuid.UUID      `json:"statusUuid"`
// 	StatusDesc        string         `json:"statusDesc"`
// 	AccessRoleID      int64          `json:"accessRoleID"`
// 	AccessName        sql.NullString `json:"accessName"`
// 	DateGiven         sql.NullTime   `json:"dateGiven"`
// 	DateExpired       sql.NullTime   `json:"dateExpired"`
// 	DateLocked        sql.NullTime   `json:"dateLocked"`
// 	PasswordChangedAt sql.NullTime   `json:"passwordChangedAt"`
// 	Attempt           int16          `json:"attempt"`
// 	Isloggedin        sql.NullBool   `json:"isloggedin"`
// 	ModCtr            int64          `json:"modCtr"`
// 	OtherInfo         []byte         `json:"otherInfo"`
// 	Created           sql.NullTime   `json:"created"`
// 	Updated           sql.NullTime   `json:"updated"`
// }

func newUserResponse(user model.User) dsUsr.UserInfo {
	return dsUsr.UserInfo{
		LoginName:         user.LoginName,
		PasswordChangedAt: sql.NullTime{Time: user.PasswordChangedAt.Time, Valid: true},
		// FullName:          user.FullName,
		// Email:             user.Email,
		// CreatedAt:         user.CreatedAt,
	}
}

func newUserRowResponse(user dsUsr.UserInfo) dsUsr.UserInfo {
	return dsUsr.UserInfo{
		LoginName:         user.LoginName,
		PasswordChangedAt: sql.NullTime{Time: user.PasswordChangedAt.Time, Valid: true},
		// FullName:          user.FullName,
		// Email:             user.Email,
		// CreatedAt:         user.CreatedAt,
	}
}

func (server *Server) createUser(ctx *gin.Context) {
	var req dsUsr.UserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	userPassword, err := util.HashPassword(req.UserPassword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// log.Println("API: createUser.B")
	arg := dsUsr.UserRequest{
		LoginName:    req.LoginName,
		UserPassword: userPassword,
		Iiid:         req.Iiid,
		AccessRoleId: req.AccessRoleId,
		StatusId:     req.StatusId,
		// OtherInfo:    req.OtherInfo,
	}
	// log.Println("API: createUser.1A", dsUsr.UserRequest(arg).UserPassword, dsUsr.UserRequest(arg))
	fmt.Printf("Get by createUser A: %+v\n", arg)
	user, err := server.store.CreateUser(ctx, arg)
	fmt.Printf("Get by createUser B: %+v\n", user)
	// log.Println("API: createUser.1B", user.UserPassword)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	// log.Println("API: createUser.2")
	// u, err := server.store.GetUser(ctx, user.LoginName)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		ctx.JSON(http.StatusNotFound, errorResponse(err))
	// 		return
	// 	}
	// 	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	// 	return
	// }

	// log.Println("API: createUser.3")
	// fmt.Printf("Get by user%+v\n", user)
	rsp := newUserResponse(user)
	ctx.JSON(http.StatusOK, rsp)
}

type loginUserRequest struct {
	LoginName    string `json:"loginName" binding:"required,alphanum"`
	UserPassword string `json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	AccessToken string         `json:"accessToken"`
	User        dsUsr.UserInfo `json:"user"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	log.Println("loginName", req.LoginName)
	user, err := server.store.GetUserbyName(ctx, req.LoginName)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = util.CheckPassword(req.UserPassword, user.UserPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(
		user.LoginName,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := loginUserResponse{
		AccessToken: accessToken,
		User:        newUserRowResponse(user),
	}
	log.Println("API:GetUser 6..", req.LoginName)
	ctx.JSON(http.StatusOK, rsp)
}
