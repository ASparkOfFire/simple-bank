package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCEFGHIJKLMNOPQRSTUVWXYZ"

var rng *rand.Rand

func init() {
	rng = rand.New(
		rand.NewSource(
			time.Now().UnixNano()),
	)
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(length int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < length; i++ {
		c := alphabet[rng.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomOwner Generates a random owner
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generates a random currency code
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "GBP", "JPY", "INR"}
	n := len(currencies)
	return currencies[rng.Intn(n)]
}
