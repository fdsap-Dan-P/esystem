package util

import (
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword returns the bcrypt hash of the password
func HashPassword(userName string, password string) ([]byte, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to hash password: %w", err)
	}
	return hashedPassword, nil
}

// CheckPassword checks if the provided password is correct or not
func CheckPassword(password string, hashedPassword []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
}

func DigitCheckCode(nNum int32) int32 {
	var s string
	n := nNum
	s = "000"
	var j int32 = 0
	for {
		if len(strings.Trim(s, " ")) == 1 {
			break
		}
		s = strings.Trim(string(n), " ")
		n = 0
		for i := 0; i < len(s)-1; i++ {
			n = n + ToInt32(s[i:1])
		}
		j++
		if j > 10 {
			break
		}
	}
	return n
}

func Replicate(str string, nLen int) string {
	for {
		if len(str) < nLen && len(str) != 0 {
			str = str + str
		} else {
			break
		}
	}
	return str[0:nLen]
}

func Ascii(s string) int32 {
	return int32([]rune(s[0:1])[0])
}
