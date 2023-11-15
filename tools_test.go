package toolkit

import (
	"strconv"
	"testing"
)

func TestTools_RandomString(t *testing.T) {
	var testTools Tools
	s := testTools.RandomString(10)
	if len(s) != 10 {
		t.Error("wrong length random string")
	}
}

func TestTools_RandomNumber(t *testing.T) {
	var testTools Tools
	n := testTools.RandomNumber(10)
	toString := strconv.Itoa(n)

	if len(toString) != 10 {
		t.Error("wrong length random number")
	}
}
