package helpers

// Thanks to this blog post for the quick solution: https://www.calhoun.io/creating-random-strings-in-go/

import (
	"math/rand"
	"time"
)

const (
	charset string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length  int    = 8
)

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

// Generate returns a string of 'length' characters, based on random
func Generate() string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
