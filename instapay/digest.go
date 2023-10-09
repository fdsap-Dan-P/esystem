package instapay

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func Digest(xmlMessage []byte) string {

	hasher := sha256.New()

	hasher.Write(xmlMessage)
	digest := hasher.Sum(nil)

	fmt.Printf("SHA256: %v", string(digest))
	bytes, _ := hex.DecodeString(hex.EncodeToString(digest))

	digestValue := base64.StdEncoding.EncodeToString(bytes)

	return digestValue
}
