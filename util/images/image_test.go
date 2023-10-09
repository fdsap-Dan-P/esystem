package images

import (
	"log"
	"path/filepath"
	"simplebank/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDocument(t *testing.T) {

	TargetPath = "app/images"
	// // jpg
	// uuid := uuid.New()

	img := NewImageDocument(filepath.Join(Config.HomeFolder,
		"static/uploads/images/me.jpg"), Config.HomeFolder,
		util.SetUUID("e4613170-7a4c-45ba-8e5e-956d93c905d2"), TargetPath)
	log.Printf("%+v", img)

	err = img.CreateThumbnail()
	require.NoError(t, err)

	log.Printf("%+v", img)

	// gif
	img = NewImageDocument(filepath.Join(Config.HomeFolder,
		"static/uploads/images/iPAD.gif"), Config.HomeFolder,
		util.SetUUID("bcf149f5-adc3-4cd4-960b-37bd422ec978"), TargetPath)
	log.Printf("%+v", img)

	err = img.CreateThumbnail()
	require.NoError(t, err)

	log.Printf("%+v", img)

	// png
	img = NewImageDocument(filepath.Join(Config.HomeFolder,
		"static/uploads/images/backend-master.png"),
		Config.HomeFolder, util.SetUUID("a2099800-48ca-45a9-8269-c1f433f08480"), TargetPath)
	log.Printf("%+v", img)

}

func TestDocumenfromtURL(t *testing.T) {
	// url
	TargetPath = "app/images"

	img := NewImageDocumentFromURL("https://assets.suysing.com/upload/products/GLA01A.jpg",
		Config.HomeFolder, util.SetUUID("f9d2b21e-08fb-43a9-bd7a-dfa87c00776a"), TargetPath)

	log.Printf("%+v --> Config.HomeFolder: %v", img, Config.HomeFolder)

	err = img.CreateThumbnail()
	require.NoError(t, err)

	log.Printf("%+v", img)
	// require.Equal(t, d1.StorageId, getData1.StorageId)

}
