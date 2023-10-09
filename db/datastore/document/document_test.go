package db

import (
	"context"
	"database/sql"
	"log"
	"path"

	"fmt"
	"testing"

	"encoding/json"
	common "simplebank/db/common"

	"simplebank/model"
	"simplebank/util"
	"simplebank/util/images"

	"github.com/stretchr/testify/require"
)

func TestDocument(t *testing.T) {

	// Test Data
	d1 := RandomDocument()
	d1.Uuid = util.SetUUID("6f22923e-6d83-4e60-b3c3-e42db606da3c")
	d1.Code = d1.Uuid.String()

	fname1 := "me.jpg"
	log.Printf("homePath: %v", homePath)
	img := images.NewImageDocument(path.Join(sorsImgPath, fname1), homePath, d1.Uuid, targetPath)
	log.Printf("Image: %v", img)
	errImg := img.CreateThumbnail()
	require.NoError(t, errImg)

	d1.FilePath = img.ImageData().NewFile.FullPath()
	d1.Thumbnail = util.ImageToBytes(img.Thumbnail())

	fmt.Printf("FilePath %v", d1.FilePath)

	d2 := RandomDocument()
	d2.Uuid = util.SetUUID("b5a9de53-1719-457a-92a8-fdda2db8886c")
	d2.FilePath = img.ImageData().NewFile.FullPath()
	d2.Thumbnail = util.ImageToBytes(img.Thumbnail())
	d2.Code = d2.Uuid.String()

	log.Printf("Struct2Json: %v", util.Struct2Json(d1))
	// Test Create
	CreatedD1 := createTestDocument(t, d1)
	CreatedD2 := createTestDocument(t, d2)

	// Get Data
	getData1, err1 := testQueriesDocument.GetDocument(context.Background(), CreatedD1.Id)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.ServerCode, getData1.ServerCode)
	require.Equal(t, d1.FilePath, getData1.FilePath)
	require.Equal(t, d1.DocDate.Time.Format("2006-01-02"), getData1.DocDate.Time.Format("2006-01-02"))
	require.Equal(t, d1.Thumbnail, getData1.Thumbnail)
	require.Equal(t, d1.DoctypeId, getData1.DoctypeId)
	require.Equal(t, d1.Description, getData1.Description)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesDocument.GetDocument(context.Background(), CreatedD2.Id)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.ServerCode, getData2.ServerCode)
	require.Equal(t, d2.FilePath, getData2.FilePath)
	require.Equal(t, d2.DocDate.Time.Format("2006-01-02"), getData2.DocDate.Time.Format("2006-01-02"))
	require.Equal(t, d2.Thumbnail, getData2.Thumbnail)
	require.Equal(t, d2.DoctypeId, getData2.DoctypeId)
	require.Equal(t, d2.Description, getData2.Description)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesDocument.GetDocumentbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Id, getData.Id)
	fmt.Printf("Get by UUId%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Id = getData2.Id
	updateD2.Description.String = updateD2.Description.String + "Edited"

	getData, err = testQueriesDocument.GetDocumentbyUuid(context.Background(), updateD2.Uuid)
	// log.Println(updateD2)
	updatedD1 := updateTestDocument(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.ServerCode, getData.ServerCode)
	require.Equal(t, updateD2.FilePath, getData.FilePath)
	require.Equal(t, updateD2.DocDate.Time.Format("2006-01-02"), getData.DocDate.Time.Format("2006-01-02"))
	require.Equal(t, updateD2.Thumbnail, getData.Thumbnail)
	require.Equal(t, updateD2.DoctypeId, getData.DoctypeId)
	require.Equal(t, updateD2.Description, getData.Description)
	require.JSONEq(t, updateD2.OtherInfo.String, getData.OtherInfo.String)

	testListDocument(t, ListDocumentParams{
		DocTypeId: updatedD1.DoctypeId,
		Limit:     5,
		Offset:    0,
	})

	// Delete Data
	// testDeleteDocument(t, getData1.Id)
	// testDeleteDocument(t, getData2.Id)
}

func testListDocument(t *testing.T, arg ListDocumentParams) {

	document, err := testQueriesDocument.ListDocument(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", document)
	require.NotEmpty(t, document)

}

func RandomDocument() DocumentRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	docType, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "DocumentType", 0, "Profile Picture")
	log.Println("docType")
	log.Println(docType)
	arg := DocumentRequest{
		ServerCode:  "Local-Image",
		DocDate:     util.SetNullTime(util.RandomDate()),
		Thumbnail:   []byte{},
		DoctypeId:   docType.Id,
		Description: util.SetNullString(util.RandomString(10)),

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestDocument(
	t *testing.T,
	d1 DocumentRequest) model.Document {

	getData1, err := testQueriesDocument.CreateDocument(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	getData, err := testQueriesDocument.GetDocumentbyUuid(context.Background(), getData1.Uuid)
	require.Equal(t, d1.ServerCode, getData.ServerCode)
	require.Equal(t, d1.FilePath, getData.FilePath)
	require.Equal(t, d1.DocDate.Time.Format("2006-01-02"), getData.DocDate.Time.Format("2006-01-02"))
	require.Equal(t, d1.Thumbnail, getData.Thumbnail)
	require.Equal(t, d1.DoctypeId, getData.DoctypeId)
	require.Equal(t, d1.Description, getData.Description)
	require.JSONEq(t, d1.OtherInfo.String, getData.OtherInfo.String)

	return getData1
}

func updateTestDocument(
	t *testing.T,
	d1 DocumentRequest) model.Document {

	getData1, err := testQueriesDocument.UpdateDocument(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	getData, err := testQueriesDocument.GetDocumentbyUuid(context.Background(), getData1.Uuid)
	require.Equal(t, d1.ServerCode, getData.ServerCode)
	require.Equal(t, d1.FilePath, getData.FilePath)
	require.Equal(t, d1.DocDate.Time.Format("2006-01-02"), getData.DocDate.Time.Format("2006-01-02"))
	require.Equal(t, d1.Thumbnail, getData.Thumbnail)
	require.Equal(t, d1.DoctypeId, getData.DoctypeId)
	require.Equal(t, d1.Description, getData.Description)
	require.JSONEq(t, d1.OtherInfo.String, getData.OtherInfo.String)

	return getData1
}

func testDeleteDocument(t *testing.T, id int64) {
	err := testQueriesDocument.DeleteDocument(context.Background(), id)
	require.NoError(t, err)

	ref1, err := testQueriesDocument.GetDocument(context.Background(), id)
	require.Error(t, err)
	// require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}

func TestCreateDocumentbyURL(t *testing.T) {

	d1 := RandomDocument()
	d1.Uuid = util.SetUUID("b5a9de53-1719-457a-92a8-fdda2db8886c")
	d1.Code = d1.Uuid.String()
	d1.FilePath = "app/images"

	getData1, err := testQueriesDocument.CreateDocumentImageFromURL(
		context.Background(), d1,
		"https://assets.suysing.com/upload/products/GLA98Y.png")
	require.NoError(t, err)
	require.NotEmpty(t, getData1)
}
