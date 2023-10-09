package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"log"
)

func main() {

	value1 := "202204250222" //time.Now().Format("20060102150405")  //transactionTimestamp
	value2 := "101"          //branchid
	value3 := "647089"       //corporateid
	secret_key := "bc101831-2464-4d94-961b-d9b07480c21c"

	log.Println("step 1: " + value1 + " | " + value2 + " | " + value3)

	value1_hash := GetMD5Hash(value1)

	log.Println("step 2: " + value1_hash)

	value2_reverse := ReverseString(value2)

	log.Println("step 3: " + value2_reverse)

	hashValue := AppendValue(value1_hash, value2_reverse)

	log.Println("step 4: " + hashValue)

	hashValue1 := AppendValue(value1_hash, value2_reverse)

	log.Println("step 5: hashValue: " + hashValue + "| hashValue1: " + hashValue1)

	hashValue = AppendValue(hashValue, hashValue1)

	log.Println("step 6: hashValue: " + hashValue)

	h1 := AppendValue(hashValue, secret_key)

	log.Println("step 7: hashValue: " + h1)

	hashValue2 := GetMD5Hash(h1)

	log.Println("step 8: hashValue2: " + hashValue2)

	hashValue = AppendValue(h1, hashValue2)

	log.Println("step 9: hashValue: " + hashValue)

	value3_reverse := ReverseString(value3)

	log.Println("step 10: value3_reverse: " + value3_reverse)

	hashvalue3 := EncryptString2(value3_reverse)

	log.Println("step 11: value3_hash: " + hashvalue3)

	hashValue = AppendValue(hashValue, hashvalue3)

	log.Println("step 12: hashvalue: " + hashValue)

	transactionKey := EncryptString1(hashValue)

	log.Println("step 13: transactionkey: " + transactionKey)

}

func AppendValue(value1 string, value2 string) string {
	return value1 + "" + value2
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func ReverseString(str string) string {
	byte_str := []rune(str)
	for i, j := 0, len(byte_str)-1; i < j; i, j = i+1, j-1 {
		byte_str[i], byte_str[j] = byte_str[j], byte_str[i]
	}
	return string(byte_str)
}

// func EncryptString1(message string, secret string) string {
// 	key := []byte(secret)
// 	h := hmac.New(sha256.New, key)
// 	h.Write([]byte(message))
// 	return base64.StdEncoding.EncodeToString(h.Sum(nil))
// }

func EncryptString1(str string) string {
	hash := sha256.Sum256([]byte(str))
	return hex.EncodeToString(hash[:])
}

func EncryptString2(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	sha1_hash := hex.EncodeToString(h.Sum(nil))

	return sha1_hash
}
