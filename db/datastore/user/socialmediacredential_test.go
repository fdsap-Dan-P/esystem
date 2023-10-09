package db

import (
	"context"
	"database/sql"

	"fmt"
	"testing"

	"encoding/json"
	common "simplebank/db/common"
	"simplebank/model"
	"simplebank/util"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestSocialMediaCredential(t *testing.T) {

	// Test Data
	d1 := randomSocialMediaCredential()
	usr, _ := testQueriesUser.GetUserbyName(context.Background(), "erick1421@gmail.com")
	d1.UserId = usr.Id

	d2 := randomSocialMediaCredential()
	usr, _ = testQueriesUser.GetUserbyName(context.Background(), "olive.mercado0609@gmail.com")
	d2.UserId = usr.Id

	fmt.Printf("Get by UUId%+v\n", usr)

	// Test Create
	CreatedD1 := createTestSocialMediaCredential(t, d1)
	CreatedD2 := createTestSocialMediaCredential(t, d2)

	// Get Data
	getData1, err1 := testQueriesUser.GetSocialMediaCredential(context.Background(), CreatedD1.Uuid)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.ProviderKey, getData1.ProviderKey)
	require.Equal(t, d1.ProviderType, getData1.ProviderType)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesUser.GetSocialMediaCredential(context.Background(), CreatedD2.Uuid)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.UserId, getData2.UserId)
	require.Equal(t, d2.ProviderKey, getData2.ProviderKey)
	require.Equal(t, d2.ProviderType, getData2.ProviderType)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesUser.GetSocialMediaCredentialbyUuId(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUId%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestSocialMediaCredential(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.UserId, updatedD1.UserId)
	require.Equal(t, updateD2.ProviderKey, updatedD1.ProviderKey)
	require.Equal(t, updateD2.ProviderType, updatedD1.ProviderType)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListSocialMediaCredential(t, ListSocialMediaCredentialParams{
		UserId: updatedD1.UserId,
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteSocialMediaCredential(t, CreatedD1.Uuid)
	testDeleteSocialMediaCredential(t, CreatedD2.Uuid)
}

func testListSocialMediaCredential(t *testing.T, arg ListSocialMediaCredentialParams) {

	socialMediaCredential, err := testQueriesUser.ListSocialMediaCredential(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", socialMediaCredential)
	require.NotEmpty(t, socialMediaCredential)

}

func randomSocialMediaCredential() SocialMediaCredentialRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	arg := SocialMediaCredentialRequest{
		// UserId:       util.RandomInt(1, 100),
		ProviderKey:  util.RandomString(10),
		ProviderType: model.SocialProviderTypeFacebook,

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestSocialMediaCredential(
	t *testing.T,
	d1 SocialMediaCredentialRequest) model.SocialMediaCredential {

	fmt.Printf("d1.UserId: %+v\n", d1.UserId)

	getData1, err := testQueriesUser.CreateSocialMediaCredential(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.ProviderKey, getData1.ProviderKey)
	require.Equal(t, d1.ProviderType, getData1.ProviderType)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestSocialMediaCredential(
	t *testing.T,
	d1 SocialMediaCredentialRequest) model.SocialMediaCredential {

	getData1, err := testQueriesUser.UpdateSocialMediaCredential(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.ProviderKey, getData1.ProviderKey)
	require.Equal(t, d1.ProviderType, getData1.ProviderType)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteSocialMediaCredential(t *testing.T, uuid uuid.UUID) {
	err := testQueriesUser.DeleteSocialMediaCredential(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesUser.GetSocialMediaCredential(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
