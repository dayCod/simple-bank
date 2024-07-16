package util

import (
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

func init() {
	rand.NewSource(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// Generates a random owner
func RandomOwner() string {
	return gofakeit.Name()
}

// Generates a random money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// Generates a random currency
func RandomCurrency() string {
	currencies := []string{"USD", "EUR"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
