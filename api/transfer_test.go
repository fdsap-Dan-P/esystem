package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	ds "simplebank/db/datastore"
	mockdb "simplebank/db/mock"
	model "simplebank/model"

	"simplebank/token"
	"simplebank/util"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestTransferAPI(t *testing.T) {
	amount := decimal.NewFromInt(10)

	user1, _ := randomUser(t)
	// user2, _ := randomUser(t)
	user3, _ := randomUser(t)

	account1 := randomAccount(1) // TODO:Change to random CustomerID
	account2 := randomAccount(1) // TODO:Change to random CustomerID
	account3 := randomAccount(1) // TODO:Change to random CustomerID

	account1.Currency = util.USD
	account2.Currency = util.USD
	account3.Currency = util.EUR

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
				"fromAccountId": account1.Id,
				"toAccountId":   account2.Id,
				"amount":        amount,
				"currency":      util.USD,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user1.LoginName, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account1.Id)).Times(1).Return(account1, nil)
				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account2.Id)).Times(1).Return(account2, nil)

				arg := ds.TransferTxParams{
					FromAccountId: account1.Id,
					ToAccountId:   account2.Id,
					Amount:        amount,
				}
				store.EXPECT().TransferTx(gomock.Any(), gomock.Eq(arg)).Times(1)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		// {
		// 	name: "UnauthorizedUser",
		// 	body: gin.H{
		// 		"fromAccountId": account1.Id,
		// 		"toAccountId":   account2.Id,
		// 		"amount":          amount,
		// 		"currency":        util.USD,
		// 	},
		// 	setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
		// 		addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user2.LoginName, time.Minute)
		// 	},
		// 	buildStubs: func(store *mockdb.MockStore) {
		// 		store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account1.Id)).Times(1).Return(account1, nil)
		// 		store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account2.Id)).Times(0)
		// 		store.EXPECT().TransferTx(gomock.Any(), gomock.Any()).Times(0)
		// 	},
		// 	checkResponse: func(recorder *httptest.ResponseRecorder) {
		// 		require.Equal(t, http.StatusUnauthorized, recorder.Code)
		// 	},
		// },
		{
			name: "NoAuthorization",
			body: gin.H{
				"fromAccountId": account1.Id,
				"toAccountId":   account2.Id,
				"amount":        amount,
				"currency":      util.USD,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetAccount(gomock.Any(), gomock.Any()).Times(0)
				store.EXPECT().TransferTx(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "FromAccountNotFound",
			body: gin.H{
				"fromAccountId": account1.Id,
				"toAccountId":   account2.Id,
				"amount":        amount,
				"currency":      util.USD,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user1.LoginName, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account1.Id)).Times(1).Return(model.Account{}, sql.ErrNoRows)
				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account2.Id)).Times(0)
				store.EXPECT().TransferTx(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "ToAccountNotFound",
			body: gin.H{
				"fromAccountId": account1.Id,
				"toAccountId":   account2.Id,
				"amount":        amount,
				"currency":      util.USD,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user1.LoginName, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account1.Id)).Times(1).Return(account1, nil)
				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account2.Id)).Times(1).Return(model.Account{}, sql.ErrNoRows)
				store.EXPECT().TransferTx(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "FromAccountCurrencyMismatch",
			body: gin.H{
				"fromAccountId": account3.Id,
				"toAccountId":   account2.Id,
				"amount":        amount,
				"currency":      util.USD,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user3.LoginName, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account3.Id)).Times(1).Return(account3, nil)
				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account2.Id)).Times(0)
				store.EXPECT().TransferTx(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "ToAccountCurrencyMismatch",
			body: gin.H{
				"fromAccountId": account1.Id,
				"toAccountId":   account3.Id,
				"amount":        amount,
				"currency":      util.USD,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user1.LoginName, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account1.Id)).Times(1).Return(account1, nil)
				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account3.Id)).Times(1).Return(account3, nil)
				store.EXPECT().TransferTx(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InvalidCurrency",
			body: gin.H{
				"fromAccountId": account1.Id,
				"toAccountId":   account2.Id,
				"amount":        amount,
				"currency":      "XYZ",
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user1.LoginName, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetAccount(gomock.Any(), gomock.Any()).Times(0)
				store.EXPECT().TransferTx(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		// {
		// 	name: "NegativeAmount",
		// 	body: gin.H{
		// 		"fromAccountId": account1.Id,
		// 		"toAccountId":   account2.Id,
		// 		"amount":          amount,
		// 		"currency":        util.USD,
		// 	},
		// 	setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
		// 		addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user1.LoginName, time.Minute)
		// 	},
		// 	buildStubs: func(store *mockdb.MockStore) {
		// 		log.Println("API:transder-1--: ")
		// 		store.EXPECT().GetAccount(gomock.Any(), gomock.Any()).Times(0)
		// 		log.Println("API:transder-2--: ")
		// 		store.EXPECT().TransferTx(gomock.Any(), gomock.Any()).Times(0)
		// 		log.Println("API:transder-3--: ")
		// 	},
		// 	checkResponse: func(recorder *httptest.ResponseRecorder) {
		// 		log.Println("API:checkResponse", http.StatusBadRequest, recorder.Code)
		// 		require.Equal(t, http.StatusBadRequest, recorder.Code)
		// 	},
		// },
		{
			name: "GetAccountError",
			body: gin.H{
				"fromAccountId": account1.Id,
				"toAccountId":   account2.Id,
				"amount":        amount,
				"currency":      util.USD,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user1.LoginName, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetAccount(gomock.Any(), gomock.Any()).Times(1).Return(model.Account{}, sql.ErrConnDone)
				store.EXPECT().TransferTx(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "TransferTxError",
			body: gin.H{
				"fromAccountId": account1.Id,
				"toAccountId":   account2.Id,
				"amount":        amount,
				"currency":      util.USD,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user1.LoginName, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account1.Id)).Times(1).Return(account1, nil)
				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account2.Id)).Times(1).Return(account2, nil)
				store.EXPECT().TransferTx(gomock.Any(), gomock.Any()).Times(1).Return(ds.TransferTxResult{}, sql.ErrTxDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		log.Println("API:testCases-1: ", tc.name)
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			log.Println("API:testCases-2: ", tc.name)
			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			log.Println("API:testCases-3: ", tc.name)
			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			// Marshal body data to JSON
			log.Println("API:testCases-4: ", tc.name)
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/transfers"
			log.Println("API:testCases-5: ", tc.name)
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			log.Println("API:testCases-6: ", tc.name)
			tc.setupAuth(t, request, server.tokenMaker)
			log.Println("API:testCases-7: ", tc.name, recorder, request)
			server.router.ServeHTTP(recorder, request)
			log.Println("API:testCases-8: ", tc.name)
			tc.checkResponse(recorder)
			log.Println("API:testCases-9: ", tc.name)
		})
	}
}
