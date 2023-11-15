package toolkit

import (
	"crypto/rand"
	mrand "math/rand"
)

const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_+"

// Tools is the type used to instantiate this module. Any variable of this type will have access
// to all the methods with the receiver *Tools
type Tools struct{}

// RandomString returns a string of random characters of length n, using randomStringSource
// as the source for the string
func (t *Tools) RandomString(n int) string {
	s, r := make([]rune, n), []rune(randomStringSource)
	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x, y := p.Uint64(), uint64(len(r))
		s[i] = r[x%y]
	}

	return string(s)
}

func (t *Tools) RandomNumber(n int) int {
	min := 1
	for i := 0; i < n-1; i++ {
		min *= 10
	}
	max := min*10 - 1

	return mrand.Intn(max-min+1) + min
}
