package util

import (
	"bytes"
	"log"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	userName := "UserName"
	pass := Encode(userName, "Password")
	log.Println("pass:[" + pass + "]")
	require.Equal(t, "Password", Decode(userName, pass))
	// require.NotEqual(t, "Password", Decode(userName, pass))

	password := RandomString(6)

	hashedPassword1, err := HashPassword(userName, password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword1)

	err = CheckPassword(password, hashedPassword1)
	require.NoError(t, err)

	wrongPassword := RandomString(6)
	err = CheckPassword(wrongPassword, hashedPassword1)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hashedPassword2, err := HashPassword(userName, password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword2)
	require.NotEqual(t, hashedPassword1, hashedPassword2)
}

func TestString2SqlList(t *testing.T) {
	str := String2SqlList([]string{"one", "two"})
	log.Println(str)
	require.True(t, false)
}

func TestDecryptFile(t *testing.T) {
	// Decrypt the encrypted config file
	key := []byte("ruEanqdOKvgzoN9n") // Replace with your secret key
	plaintext, err := DecryptFile("temp_config.enc", key)
	require.NoError(t, err, "Error decrypting config")

	// if err != nil {
	// 	fmt.Println("Error decrypting config:", err)
	// 	return
	// }

	// Initialize Viper with the decrypted config data
	viper.SetConfigType("yaml") // Adjust to your config format
	err = viper.ReadConfig(bytes.NewBuffer(plaintext))
	require.NoError(t, err, "Error reading config")

	// if err != nil {
	// 	fmt.Println("Error reading config:", err)
	// 	return
	// }

	// Access configuration values
	value := viper.GetString("key")
	require.Equal(t, "expectedValue", value, "Config value mismatch")
	// require.True(t, false)
}

func TestEncryptFile(t *testing.T) {
	// Create a temporary file for testing
	tempFilename := "temp_config.enc"

	// // Cleanup the temporary file after the test
	// defer func() {
	// 	if err := os.Remove(tempFilename); err != nil {
	// 		t.Errorf("Error cleaning up temporary file: %v", err)
	// 	}
	// }()

	// Replace with your actual secret key and plaintext configuration data
	key := []byte("ruEanqdOKvgzoN9n")
	plaintext := []byte("DB_DRIVER=postgres\nDB_HOST=localhost\nDB_PORT=5432")

	err := EncryptFile(tempFilename, key, plaintext)
	if err != nil {
		t.Errorf("Error encrypting config: %v", err)
	}
}
