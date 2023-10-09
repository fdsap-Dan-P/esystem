package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	dsUsr "simplebank/db/datastore/user"
	mockdb "simplebank/db/mock"
	model "simplebank/model"
	"simplebank/util"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

type eqCreateUserParamsMatcher struct {
	arg      dsUsr.UserRequest
	password string
}

func (e eqCreateUserParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(dsUsr.UserRequest)
	if !ok {
		return false
	}

	err := util.CheckPassword(e.password, []byte(arg.Password))
	if err != nil {
		return false
	}

	e.arg.Password = arg.Password
	return reflect.DeepEqual(e.arg, arg)
}

func (e eqCreateUserParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.arg, e.password)
}

func EqCreateUserParams(arg dsUsr.UserRequest, password string) gomock.Matcher {
	return eqCreateUserParamsMatcher{arg, password}
}

func TestCreateUserAPI(t *testing.T) {

	userInfo, password := randomUser(t)
	user := info2User(userInfo)
	password = "secret"
	user.UserPassword = password

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"uuid":         user.Uuid,
				"loginName":    user.LoginName,
				"password":     password,
				"iiid":         user.Iiid,
				"accessRoleId": user.AccessRoleId,
				"statusId":     user.StatusId,
				"otherInfo":    user.OtherInfo,
				// FullName:       util.RandomOwner(),
				// Email:          util.RandomEmail(),
				// DateGiven:
				// DateExpired:
				// DateLocked:
				// UserPassword:
				// Attempt:
				// Isloggedin:

				// "full_name": user.FullName,
				// "email":     user.Email,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := dsUsr.UserRequest{
					Uuid:         user.Uuid,
					LoginName:    user.LoginName,
					Iiid:         user.Iiid,
					AccessRoleId: user.AccessRoleId,
					StatusCode:   user.StatusId,
					OtherInfo:    user.OtherInfo,
				}

				log.Println("store.EXPECT().")

				store.EXPECT().
					CreateUser(gomock.Any(), EqCreateUserParams(arg, password)).
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				log.Println("Rhick:EqCreateUserParams..1", http.StatusOK, recorder.Code)
				fmt.Printf("Get by user 1:%+v\n", recorder.Body)
				fmt.Printf("Get by user 2: LoginName-->%+v\n", user.LoginName)
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchUser(t, recorder.Body, user)
			},
		},
		{
			name: "InternalError",
			body: gin.H{
				"loginName": user.LoginName,
				"password":  password,
				// "full_name": user.FullName,
				// "email":     user.Email,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(model.User{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		/*
			 			{
							name: "DuplicateLoginName",
							body: gin.H{
								"loginName": user.LoginName,
								"password":  password,
								// "full_name": user.FullName,
								// "email":     user.Email,
							},
							buildStubs: func(store *mockdb.MockStore) {
								store.EXPECT().
									CreateUser(gomock.Any(), gomock.Any()).
									Times(1).
									Return(model.User{}, &pq.Error{Code: "23505"})
							},
							checkResponse: func(recorder *httptest.ResponseRecorder) {
								require.Equal(t, http.StatusForbidden, recorder.Code)
							},
						},
						{
							name: "InvalidLoginName",
							body: gin.H{
								"loginName": "invalid-user#1",
								"password":  password,
								// "full_name": user.FullName,
								// "email":     user.Email,
							},
							buildStubs: func(store *mockdb.MockStore) {
								store.EXPECT().
									CreateUser(gomock.Any(), gomock.Any()).
									Times(0)
							},
							checkResponse: func(recorder *httptest.ResponseRecorder) {
								require.Equal(t, http.StatusBadRequest, recorder.Code)
							},
						},
						// {
						// 	name: "InvalidEmail",
						// 	body: gin.H{
						// 		"loginName":    user.LoginName,
						// 		"password": password,
						// 		// "full_name": user.FullName,
						// 		"email": "invalid-email",
						// 	},
						// 	buildStubs: func(store *mockdb.MockStoreUser) {
						// 		store.EXPECT().
						// 			CreateUser(gomock.Any(), gomock.Any()).
						// 			Times(0)
						// 	},
						// 	checkResponse: func(recorder *httptest.ResponseRecorder) {
						// 		require.Equal(t, http.StatusBadRequest, recorder.Code)
						// 	},
						// },
						{
							name: "TooShortPassword",
							body: gin.H{
								"loginName": user.LoginName,
								"password":  "123",
								// "full_name": user.FullName,
								// "email":     user.Email,
							},
							buildStubs: func(store *mockdb.MockStore) {
								store.EXPECT().
									CreateUser(gomock.Any(), gomock.Any()).
									Times(0)
							},
							checkResponse: func(recorder *httptest.ResponseRecorder) {
								require.Equal(t, http.StatusBadRequest, recorder.Code)
							},
						},*/
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/users"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
		// break
	}
}

func TestLoginUserAPI(t *testing.T) {
	userInfo, password := randomUser(t)
	user := info2User(userInfo)

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"loginName": user.LoginName,
				"password":  password,
			},
			buildStubs: func(store *mockdb.MockStore) {
				log.Println("TTT", user.LoginName)
				store.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(user.LoginName)).
					Times(1).
					Return(userInfo, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		/*
			{
				name: "UserNotFound",
				body: gin.H{
					"loginName": "NotFound",
					"password":  password,
				},
				buildStubs: func(store *mockdb.MockStore) {
					store.EXPECT().
						GetUser(gomock.Any(), gomock.Any()).
						Times(1).
						Return(model.User{}, sql.ErrNoRows)
				},
				checkResponse: func(recorder *httptest.ResponseRecorder) {
					require.Equal(t, http.StatusNotFound, recorder.Code)
				},
			},
			{
				name: "IncorrectPassword",
				body: gin.H{
					"loginName": user.LoginName,
					"password":  "incorrect",
				},
				buildStubs: func(store *mockdb.MockStore) {
					store.EXPECT().
						GetUser(gomock.Any(), gomock.Eq(user.LoginName)).
						Times(1).
						Return(user, nil)
				},
				checkResponse: func(recorder *httptest.ResponseRecorder) {
					require.Equal(t, http.StatusUnauthorized, recorder.Code)
				},
			},
			{
				name: "InternalError",
				body: gin.H{
					"loginName": user.LoginName,
					"password":  password,
				},
				buildStubs: func(store *mockdb.MockStore) {
					store.EXPECT().
						GetUser(gomock.Any(), gomock.Any()).
						Times(1).
						Return(model.User{}, sql.ErrConnDone)
				},
				checkResponse: func(recorder *httptest.ResponseRecorder) {
					require.Equal(t, http.StatusInternalServerError, recorder.Code)
				},
			},
			{
				name: "InvalidLoginName",
				body: gin.H{
					"loginName": "invalid-user#1",
					"password":  password,
					// "full_name": user.FullName,
					// "email":     user.Email,
				},
				buildStubs: func(store *mockdb.MockStore) {
					store.EXPECT().
						GetUser(gomock.Any(), gomock.Any()).
						Times(0)
				},
				checkResponse: func(recorder *httptest.ResponseRecorder) {
					require.Equal(t, http.StatusBadRequest, recorder.Code)
				},
			},
		*/
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/users/login"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func randomUser(t *testing.T) (user dsUsr.UserInfo, password string) {
	password = util.RandomString(6)
	userPassword, err := util.HashPassword(password)
	require.NoError(t, err)

	user = dsUsr.UserInfo{
		// Uuid:         uuid.MustParse("26f922de-da14-4f89-8337-355b2d216aa3"),
		LoginName:      util.RandomOwner(),
		HashedPassword: userPassword,
		// FullName:     util.RandomOwner(),
		// Email:        util.RandomEmail(),
		AccessRoleId: 1,
		StatusCode:   2467,
		// DateGiven:
		// DateExpired:
		// DateLocked:
		// UserPassword:
		// Attempt:
		// Isloggedin:
		// OtherInfo: []byte{},
	}
	return
}

func info2User(usr dsUsr.UserInfo) model.User {
	return model.User{
		Id:                usr.Id,
		Uuid:              usr.Uuid,
		Iiid:              usr.Iiid,
		LoginName:         usr.LoginName,
		DisplayName:       usr.DisplayName,
		AccessRoleId:      usr.AccessRoleId,
		StatusId:          usr.StatusId,
		DateGiven:         usr.DateGiven,
		DateExpired:       usr.DateExpired,
		DateLocked:        usr.DateLocked,
		PasswordChangedAt: usr.PasswordChangedAt,
		UserPassword:      usr.UserPassword,
		Attempt:           usr.Attempt,
		Isloggedin:        usr.Isloggedin,
		OtherInfo:         usr.OtherInfo,
		Thumbnail:         usr.Thumbnail,
	}
}
func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, user model.User) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotUser model.User
	err = json.Unmarshal(data, &gotUser)

	require.NoError(t, err)
	fmt.Printf("Get by user A:%+v\n", user)
	fmt.Printf("Get by user B:%+v\n", gotUser)
	require.Equal(t, user.LoginName, gotUser.LoginName)
	// require.Equal(t, user.FullName, gotUser.FullName)
	// require.Equal(t, user.Email, gotUser.Email)
	require.Empty(t, gotUser.UserPassword)
}
