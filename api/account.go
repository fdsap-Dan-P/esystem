package api

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"time"

	db "simplebank/db/datastore/account"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/shopspring/decimal"
)

type createAccountRequest struct {
	CustomerId       int64           `json:"customerId"`
	Acc              string          `json:"acc"`
	AlternateAcc     sql.NullString  `json:"alternateAcc"`
	AccountName      string          `json:"accountName"`
	Balance          decimal.Decimal `json:"balance"`
	NonCurrent       decimal.Decimal `json:"nonCurrent"`
	ContractDate     sql.NullTime    `json:"contractDate"`
	Credit           decimal.Decimal `json:"credit"`
	Debit            decimal.Decimal `json:"debit"`
	Isbudget         sql.NullBool    `json:"isbudget"`
	LastActivityDate sql.NullTime    `json:"lastActivityDate"`
	OpenDate         time.Time       `json:"openDate"`
	PassbookLine     int16           `json:"passbookLine"`
	PendingTrnAmt    decimal.Decimal `json:"pendingTrnAmt"`
	Principal        decimal.Decimal `json:"principal"`
	ClassId          int64           `json:"classId"`
	TypeId           int64           `json:"TypeId"`
	BudgetAccountId  sql.NullInt64   `json:"budgetAccountId"`
	Currency         string          `json:"currency" binding:"required,currency"`
	OfficeId         int64           `json:"officeId"`
	ReferredbyId     sql.NullInt64   `json:"referredbyId"`
	StatusId         int64           `json:"statusId"`
	Remarks          sql.NullString  `json:"remarks"`
	OtherInfo        sql.NullString  `json:"otherInfo"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	log.Println("API:createAccount..1", ctx)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Println("API:createAccount..1A", req.Currency)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	log.Println("API:createAccount..1B", req.Currency)

	// authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.AccountRequest{
		CustomerId:       req.CustomerId,
		Acc:              req.Acc,
		AlternateAcc:     req.AlternateAcc,
		AccountName:      req.AccountName,
		Balance:          req.Balance,
		NonCurrent:       req.NonCurrent,
		ContractDate:     req.ContractDate,
		Credit:           req.Credit,
		Debit:            req.Debit,
		Isbudget:         req.Isbudget,
		LastActivityDate: req.LastActivityDate,
		OpenDate:         req.OpenDate,
		PassbookLine:     req.PassbookLine,
		PendingTrnAmt:    req.PendingTrnAmt,
		Principal:        req.Principal,
		ClassId:          req.ClassId,
		AccountTypeId:    req.TypeId,
		BudgetAccountId:  req.BudgetAccountId,
		Currency:         req.Currency,
		OfficeId:         req.OfficeId,
		ReferredbyId:     req.ReferredbyId,
		StatusCode:       req.StatusId,
		Remarks:          req.Remarks,
		OtherInfo:        req.OtherInfo,
	}

	log.Println("API:createAccount..2", req)

	// bal := util.RandomMoney()
	// Acc, _ := uuid.NewRandom()
	// arg := db.CreateAccountParams{
	// 	// LoginName:    user.LoginName,
	// 	Balance:       bal,
	// 	Currency:      req.Currency,
	// 	OtherInfo:     []byte{},
	// 	CustomerID:    1,
	// 	TypeID: 1,
	// 	ClassID:       1,
	// 	OfficeID:      1,
	// 	StatusID:      1,
	// 	Acc:           Acc.String()[1:15],
	// 	AlternateAcc:  sql.NullString(sql.NullString{String: Acc.String()[1:15], Valid: true}),
	// }

	account, err := server.store.createAccount(ctx, arg)
	log.Println("API:createAccount..3A", account)
	if err != nil {
		log.Println("API:createAccount..3", err)
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	log.Println("API:createAccount..4", http.StatusOK, account)

	ctx.JSON(http.StatusOK, account)
}

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getAccount(ctx *gin.Context) {
	var req getAccountRequest
	log.Println("API:getAccount..1")
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	log.Println("API:getAccount..2")
	account, err := server.store.GetAccount(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	log.Println("API:getAccount..3")

	// authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if account.CustomerId != account.CustomerId { // authPayload.Username {
		err := errors.New("account doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type listAccountRequest struct {
	PageID   int32 `form:"pageId" binding:"required,min=1"`
	PageSize int32 `form:"pageSize" binding:"required,min=5,max=10"`
}

func (server *Server) listAccounts(ctx *gin.Context) {
	var req listAccountRequest
	log.Println("API:listAccounts..1", ctx)
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	log.Println("API:listAccounts..2", req)

	// authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.ListAccountParams{
		CustomerId: 1, //authPayload.Username,
		Limit:      req.PageSize,
		Offset:     (req.PageID - 1) * req.PageSize,
	}

	accounts, err := server.store.ListAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}
