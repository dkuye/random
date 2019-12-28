package random

import (
	"math/rand"
	"time"
)

// String struct
type String struct {
	Lower   bool
	Upper   bool
	Number  bool
	Special bool
}

/*
* Gen
* generate random string base on character set
*/
func (r String) Gen(length int) string {

	lower := "abcdefghijklmnopqrstuvwxyz"
	upper := "ABCDEFGHIJKLMNPQRSTUVWXYZ"
	number := "0123456789"
	special := "~!@#$%^&*()-_=+[]{}?/,.<>;:" //"@#$%&?_-+="
	charset := ""

	if r.Lower {
		charset += lower
	} else if r.Upper {
		charset += upper
	} else if r.Number {
		charset += number
	} else if r.Special {
		charset += special
	} else {
		charset += lower
	}

	return randomProcessor(length, charset)
}

func randomProcessor(length int, charset string) string {
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
