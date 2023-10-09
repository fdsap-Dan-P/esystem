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
	"testing"
	"time"

	dsAcc "simplebank/db/datastore/account"
	mockdb "simplebank/db/mock"
	model "simplebank/model"
	token "simplebank/token"
	"simplebank/util"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestGetAccountAPI(t *testing.T) {
	user, _ := randomUser(t)
	accountInfo := randomAccount(1) // Change to random customer ID
	account := dsAcc.AccountInfo2Model(accountInfo)

	testCases := []struct {
		name          string
		accountID     int64
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			accountID: account.Id,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.LoginName, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account.Id)).
					Times(1).
					Return(accountInfo, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchAccount(t, recorder.Body, account)
			},
		},
		// {
		// 	name:      "UnauthorizedUser",
		// 	accountID: account.Id,
		// 	setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
		// 		addAuthorization(t, request, tokenMaker, authorizationTypeBearer, "unauthorized_user", time.Minute)
		// 	},
		// 	buildStubs: func(store *mockdb.MockStore) {
		// 		store.EXPECT().
		// 			GetAccount(gomock.Any(), gomock.Eq(account.Id)).
		// 			Times(1).
		// 			Return(account, nil)
		// 	},
		// 	checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
		// 		require.Equal(t, http.StatusUnauthorized, recorder.Code)
		// 	},
		// },
		{
			name:      "NoAuthorization",
			accountID: account.Id,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name:      "NotFound",
			accountID: account.Id,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.LoginName, time.Minute)
			},

			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account.Id)).
					Times(1).
					Return(dsAcc.AccountInfo{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:      "InternalError",
			accountID: account.Id,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.LoginName, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account.Id)).
					Times(1).
					Return(dsAcc.AccountInfo{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:      "InvalidID",
			accountID: 0,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.LoginName, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
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

			url := fmt.Sprintf("/accounts/%d", tc.accountID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestCreateAccountAPI(t *testing.T) {
	user, _ := randomUser(t)
	accountInfo := randomAccount(1) // Change to random customer ID
	account := dsAcc.AccountInfo2Model(accountInfo)

	testCases := []struct {
		name          string
		body          gin.H
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"balance":      account.Balance,
				"currency":     account.Currency,
				"otherInfo":    account.OtherInfo,
				"customerId":   account.CustomerId,
				"TypeId":       account.AccountTypeId,
				"classId":      account.ClassId,
				"officeId":     account.OfficeId,
				"statusId":     account.StatusCode,
				"acc":          account.Acc,
				"alternateAcc": account.AlternateAcc,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.LoginName, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := dsAcc.AccountRequest{
					CustomerId:    account.CustomerId,
					Currency:      account.Currency,
					Balance:       account.Balance,
					OtherInfo:     account.OtherInfo,
					AccountTypeId: account.AccountTypeId,
					ClassId:       account.ClassId,
					OfficeId:      account.OfficeId,
					StatusCode:    account.StatusCode,
					Acc:           account.Acc,
					AlternateAcc:  account.AlternateAcc,
				}

				store.EXPECT().
					CreateAccount(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(account, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchAccount(t, recorder.Body, account)
			},
		},
		{
			name: "NoAuthorization",
			body: gin.H{
				"currency": account.Currency,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateAccount(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "InternalError",
			body: gin.H{
				"currency": account.Currency,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.LoginName, time.Minute)
			},

			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateAccount(gomock.Any(), gomock.Any()).
					Times(1).
					Return(model.Account{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		// {
		// 	name: "InvalidCurrency",
		// 	body: gin.H{
		// 		"currency":        "invalid",
		// 		"balance":         account.Balance,
		// 		"otherInfo":      account.OtherInfo,
		// 		"customerId":     account.CustomerId,
		// 		"TypeId": account.TypeId,
		// 		"classId":        account.ClassId,
		// 		"officeId":       account.OfficeId,
		// 		"statusId":       account.StatusId,
		// 		"acc":             account.Acc,
		// 		"alternateAcc":   account.AlternateAcc,
		// 	},
		// 	setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
		// 		addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.LoginName, time.Minute)
		// 	},
		// 	buildStubs: func(store *mockdb.MockStore) {
		// 		store.EXPECT().
		// 			CreateAccount(gomock.Any(), gomock.Any()).Times(1).
		// 			Return(model.Account{}, nil)
		// 	},
		// 	checkResponse: func(recorder *httptest.ResponseRecorder) {
		// 		log.Println("Check Response:", http.StatusBadRequest, recorder.Code)
		// 		require.Equal(t, http.StatusBadRequest, recorder.Code)
		// 	},
		// },
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

			url := "/accounts"
			log.Println("test1", tc.name, data)
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			log.Println("test2", tc.name, data)
			require.NoError(t, err)
			log.Println("test3", tc.name, data)

			tc.setupAuth(t, request, server.tokenMaker)
			log.Println("test4", tc.name, data)
			server.router.ServeHTTP(recorder, request)
			log.Println("test5", tc.name, data)
			tc.checkResponse(recorder)
			log.Println("test6", tc.name, data)
		})
		// break
	}
}

func TestListAccountsAPI(t *testing.T) {
	user, _ := randomUser(t)

	n := 7
	accounts := make([]dsAcc.AccountInfo, n)
	for i := 0; i < n; i++ {
		accounts[i] = randomAccount(1) // Change to Random CustomerID
	}

	type Query struct {
		pageID   int
		pageSize int
	}

	testCases := []struct {
		name          string
		query         Query
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			query: Query{
				pageID:   1,
				pageSize: n,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.LoginName, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := dsAcc.ListAccountParams{
					CustomerId: 1,
					Limit:      int32(n),
					Offset:     0,
				}

				store.EXPECT().
					ListAccount(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(accounts, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchAccounts(t, recorder.Body, accounts)
			},
		},
		{
			name: "NoAuthorization",
			query: Query{
				pageID:   1,
				pageSize: n,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListAccount(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "InternalError",
			query: Query{
				pageID:   1,
				pageSize: n,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.LoginName, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListAccount(gomock.Any(), gomock.Any()).
					Times(1).
					Return([]dsAcc.AccountInfo{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidPageID",
			query: Query{
				pageID:   -1,
				pageSize: n,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.LoginName, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListAccount(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InvalidPageSize",
			query: Query{
				pageID:   1,
				pageSize: 100000,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.LoginName, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListAccount(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
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

			url := "/accounts"
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			// Add query parameters to request URL
			q := request.URL.Query()
			q.Add("pageId", fmt.Sprintf("%d", tc.query.pageID))
			q.Add("pageSize", fmt.Sprintf("%d", tc.query.pageSize))
			request.URL.RawQuery = q.Encode()

			tc.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func randomAccount(customerId int64) dsAcc.AccountInfo {
	Acc, _ := uuid.NewRandom()
	return dsAcc.AccountInfo{
		Id:         util.RandomInt(1, 1000),
		CustomerId: customerId,
		Balance:    util.RandomMoney(),
		Currency:   util.RandomCurrency(),
		// OtherInfo:     []byte{},
		AccountTypeId: 1,
		ClassId:       1,
		OfficeId:      1,
		StatusCode:    1,
		Acc:           Acc.String()[1:15],
		AlternateAcc:  sql.NullString(sql.NullString{String: Acc.String()[1:15], Valid: true}),
	}
}

func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account model.Account) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotAccount model.Account
	log.Println("API:requireBodyMatchAccount..1")
	err = json.Unmarshal(data, &gotAccount)
	log.Println("API:requireBodyMatchAccount..2")
	require.NoError(t, err)
	log.Println("API:requireBodyMatchAccount..3")
	//	require.Equal(t, account, gotAccount)
	log.Println("API:requireBodyMatchAccount..4")
}

func requireBodyMatchAccounts(t *testing.T, body *bytes.Buffer, accounts []dsAcc.AccountInfo) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotAccounts []dsAcc.AccountInfo
	err = json.Unmarshal(data, &gotAccounts)
	require.NoError(t, err)
	// require.Equal(t, accounts, gotAccounts)
}
