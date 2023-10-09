package file

import (
	"log"

	"github.com/purnaresa/bulwark/encryption"
	"github.com/purnaresa/bulwark/utils"
)

func Encyrpt(sorsFilePath string, targetPath string, key string) error {
	// step 1
	image := utils.ReadFile(sorsFilePath)
	// step 2
	encryptionClient := encryption.NewClient()
	// secret := encryptionClient.GenerateRandomString(32)
	// step 3
	cipherImage := encryptionClient.EncryptAES(image, []byte(key))
	// step 4
	err := utils.WriteFile(cipherImage, targetPath)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
}

func Decrypt(encryptedFilePath string, targetPath string, key string) error {
	// decryption start
	// 1
	encryptionClient := encryption.NewClient()
	encryptedImage := utils.ReadFile(encryptedFilePath)

	// 3
	plainImage := encryptionClient.DecryptAES(encryptedImage, []byte(key))
	err := utils.WriteFile(plainImage, targetPath)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
	// decryption end
}
