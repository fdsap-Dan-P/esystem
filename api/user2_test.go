package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	dsUsr "simplebank/db/datastore/user"
	mockdb "simplebank/db/mock"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestUse(t *testing.T) {

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
				// "uuid":         user.Uuid,
				"loginName":    user.LoginName,
				"password":     password,
				"iiid":         user.Iiid,
				"accessRoleID": user.AccessRoleId,
				"statusID":     user.StatusId,
				"otherInfo":    user.OtherInfo,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := dsUsr.UserRequest{
					// Uuid:         user.Uuid,
					LoginName:    user.LoginName,
					Iiid:         user.Iiid,
					AccessRoleId: user.AccessRoleId,
					StatusId:     user.StatusId,
					OtherInfo:    user.OtherInfo,
				}

				log.Println("store.EXPECT().")
				fmt.Printf("Get by user 1:%+v\n", arg)
				store.EXPECT().
					CreateUser(gomock.Any(), EqCreateUserParams(arg, password)).
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				// log.Println("Rhick:EqCreateUserParams..1", http.StatusOK, recorder.Code)
				// fmt.Printf("Get by user 1:%+v\n", recorder.Body)
				// fmt.Printf("Get by user 2: LoginName-->%+v\n", user.LoginName)
				require.Equal(t, http.StatusOK, recorder.Code)
				//requireBodyMatchUser(t, recorder.Body, user)
			},
		},
	}

	tc := testCases[0]

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
