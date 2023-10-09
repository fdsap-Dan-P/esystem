package file

import (
	"log"
	"simplebank/util"
	"testing"

	"github.com/purnaresa/bulwark/encryption"
	"github.com/stretchr/testify/require"
)

func TestFile(t *testing.T) {

	config, err := util.LoadConfig("../..")

	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	encryptionClient := encryption.NewClient()
	secret := encryptionClient.GenerateRandomString(32)
	err = Encyrpt(
		config.HomeFolder+"static/uploads/images/me.jpg",
		config.HomeFolder+"static/uploads/images/me_Encrypted.jpg",
		secret)

	require.NoError(t, err)
	status, er := Exists(config.HomeFolder + "static/uploads/images/me_Encrypted.jpg")
	require.NoError(t, er)
	require.True(t, status)

	err = Decrypt(
		config.HomeFolder+"static/uploads/images/me_Encrypted.jpg",
		config.HomeFolder+"static/uploads/images/me_Decrypted.jpg",
		secret)
	require.NoError(t, err)
	status, er = Exists(config.HomeFolder + "static/uploads/images/me_Decrypted.jpg")
	require.NoError(t, er)
	require.True(t, status)
}
