package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// `func init`is a special function that will be called automatically when the package is first used

func init() {
	// `rand.Seed` called with `time.Now` ensures that, every time the code is called, the generated values will be different
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	// `rand.Int63n` will return an random integer between `0` and `max-min`
	// by adding `min` a random integer between `min` and `max` is generated
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length `size`
func RandomString(size int) string {
	var stringBuilder strings.Builder
	alphabetLength := len(alphabet)

	for index := 0; index < size; index++ {
		char := alphabet[rand.Intn(alphabetLength)]
		stringBuilder.WriteByte(char)
	}

	return stringBuilder.String()
}

// RandomOwner generates a random owner name with `6` characters
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random amount of money between `0` and `10000`
func RandomMoney() int64 {
	return RandomInt(0, 10000)
}

// RandomCurrency generates a random currency code from
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD", "BRL"}
	currenciesLength := len(currencies)

	return currencies[rand.Intn(currenciesLength)]
}