package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"

	// "crypto/rand"
	// "encoding/base64"
	// "io"
	"errors"
	"os"
)

func Encode(key string, text string) string {
	var nStart int32
	var nIncr int32
	var nNum int32
	// var i int32
	var Str1 string = ""
	var Str2 string = ""

	key = Replicate(key, len(text))

	if len(key) > 0 {
		nIncr = DigitCheckCode(ToInt32(key[0:1])+int32(len(key))+int32(len(text))) + 1
	} else {
		nIncr = 0
	}

	nNum = 32 + nIncr
	nStart = nNum
	for i := 32; i <= 255; i++ {
		Str1 = Str1 + string(nNum)
		nNum = nNum + nIncr
		if nNum > 255 {
			nStart = nStart - 1
			nNum = nStart
		}
	}

	for i := 0; i < len(text); i++ {
		// log.Printf("encode Ascii(text[i:i+1]): %v", Ascii(text[i:i+1]))
		nNum = Ascii(text[i:i+1]) - Ascii(key[i:i+1]) + 1 - int32(i)*nIncr
		for ok := true; ok; ok = nNum < 32 {
			nNum = nNum + 255 - 31
		}
		// log.Printf("3. for %v", nNum)
		Str2 = Str2 + string(byte(nNum))
		// log.Printf("Encode Str2: %v", Str2)
	}
	return Str2
}

func Decode(key string, text string) string {
	var nStart int32
	var nIncr int32
	var nNum int32
	var Str1 string
	var Str2 string
	Str1 = ""
	Str2 = ""
	key = Replicate(key, len(text))
	textLen := int32(len([]rune(text)))

	if len(key) > 0 {
		nIncr = DigitCheckCode(ToInt32(key[0:1])+int32(len(key))+textLen) + 1
	} else {
		nIncr = 0
	}

	nNum = 32 + nIncr
	nStart = nNum
	for i := 32; i <= 255; i++ {
		Str1 = Str1 + string(byte(nNum))
		nNum = nNum + nIncr
		if nNum > 255 {
			nStart = nStart - 1
			nNum = nStart
		}
	}

	// log.Printf("DeCode text [%v]", text)
	// log.Printf("DeCode textLen: %v", textLen)
	for i := 0; i <= int(textLen)-1; i++ {
		// log.Printf("DeCode int32([]rune(text)[i]) : %v", int32([]rune(text)[i]))
		nNum = int32([]rune(text)[i]) + Ascii(key[i:i+1]) - 1 + int32(i)*nIncr
		for ok := true; ok; ok = nNum > 255 {
			nNum = nNum - 255 + 31
		}
		// log.Printf("3. for %v", nNum)
		Str2 = Str2 + string(byte(nNum))
		// log.Printf("DeCode Str2: %v", Str2)
	}
	return Str2
}

// DecryptFile decrypts an encrypted file and returns the plaintext content.
func DecryptFile(filename string, key []byte) ([]byte, error) {
	cipherText, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(cipherText) < aes.BlockSize {
		return nil, errors.New("cipherText too short")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	plaintext := make([]byte, len(cipherText))
	stream.XORKeyStream(plaintext, cipherText)

	return plaintext, nil
}

func EncryptFile(filename string, key []byte, plaintext []byte) error {
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	cipherText := make([]byte, aes.BlockSize+len(plaintext))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plaintext)

	return os.WriteFile(filename, cipherText, 0644)
}

// DecryptData decrypts data using AES encryption and the provided key.
func DecryptData(encryptedData []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(encryptedData) < aes.BlockSize {
		return nil, errors.New("cipherText too short")
	}

	iv := encryptedData[:aes.BlockSize]
	cipherText := encryptedData[aes.BlockSize:]

	if len(cipherText)%aes.BlockSize != 0 {
		return nil, errors.New("cipherText is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(cipherText))
	mode.CryptBlocks(plaintext, cipherText)

	// Remove padding (if any)
	plaintext, err = unPad(plaintext)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func pad(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

func unPad(data []byte) ([]byte, error) {
	padding := int(data[len(data)-1])
	if padding > len(data) || padding == 0 {
		return nil, errors.New("invalid padding")
	}
	return data[:len(data)-padding], nil
}

// EncryptData encrypts data using AES encryption and the provided key.
func EncryptData(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Generate a random initialization vector (IV)
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	// Pad the plaintext to a multiple of the block size
	data = pad(data, aes.BlockSize)

	cipherText := make([]byte, aes.BlockSize+len(data))
	copy(cipherText[:aes.BlockSize], iv)

	// Encrypt the data using AES in CBC mode
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[aes.BlockSize:], data)

	return cipherText, nil
}
