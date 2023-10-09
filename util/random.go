package util

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomNullInt(min, max int64) sql.NullInt64 {
	return SetNullInt64(RandomInt(min, max))
}

func RandomInt32(min, max int32) int32 {
	return min + rand.Int31n(max-min+1)
}

func RandomInt16(min, max int16) int16 {
	return min + int16(rand.Int31n(int32(max)))
}

func RandomNullInt32(min, max int32) sql.NullInt32 {
	return SetNullInt32(RandomInt32(min, max))
}

func RandomFloat64() float64 {
	return rand.Float64()
}

func RandomNullFloat64(min, max float64) sql.NullFloat64 {
	return SetNullFloat64(RandomFloat64())
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomNullString generates a random string of length n
func RandomNullString(n int) sql.NullString {
	return SetNullString(RandomString(n))
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random amount of money
func RandomMoney() decimal.Decimal {
	return decimal.NewFromInt(RandomInt(0, 1000))
}

// RandomCurrency generates a random currency code
func RandomCurrency() string {
	currencies := []string{USD, EUR, CAD}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

// RandomProduct generates a random Product code
func RandomProduct() string {
	prod := []string{"Savings", "Loan", "Inventory", "Donation", "Remittance", "Rental"}
	n := len(prod)
	return prod[rand.Intn(n)]
}

// RandomAccountType generates a random AccountType code
func RandomAccountType() string {
	prod := []string{
		"Sikap 1",
		"Sikap 2",
		"Sikap 3",
		"Sikap 4",
		"Sipag Term Loan",
		"Sipag Flex"}
	n := len(prod)
	return prod[rand.Intn(n)]
}

// RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}

func RandomDate() time.Time {
	return DateValue(fmt.Sprintf("%v-04-14", RandomInt(1930, 2021)))
}

func SetDate(s string) time.Time {
	d, _ := time.Parse("2006-01-02", s)
	return d
}

func SetNullDate(s string) sql.NullTime {
	return sql.NullTime{Time: SetDate(s), Valid: true}
}

func RandomNullDate() sql.NullTime {
	return SetNullTime(RandomDate())
}

func NewUUID() string {
	u, _ := uuid.NewUUID()
	return u.String()
}

func UUID() uuid.UUID {
	u, _ := uuid.NewUUID()
	return u
}

func SetUUID(s string) uuid.UUID {
	u, _ := uuid.Parse(s)
	return u
}
