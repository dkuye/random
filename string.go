package random

import (
	"math/rand"
	"time"
)

type String struct {
	Lower                   bool
	LowerUpper              bool
	LowerUpperNumber        bool
	LowerUpperSpecial       bool
	LowerUpperNumberSpecial bool
	LowerNumber             bool
	LowerNumberSpecial      bool
	LowerSpecial            bool
	Upper                   bool
	UpperNumber             bool
	UpperNumberSpecial      bool
	UpperSpecial            bool
	Number                  bool
	NumberSpecial           bool
	Special                 bool
}

func (r String) Gen(length int) string {

	lower := "abcdefghijklmnopqrstuvwxyz"
	upper := "ABCDEFGHIJKLMNPQRSTUVWXYZ"
	number := "0123456789"
	special := "~!@#$%^&*()-_=+[]{}?/,.<>;:" //"@#$%&?_-+="
	characters := ""

	if r.Lower {
		return randomProcessor(length, lower)

	} else if r.LowerUpper {
		characters = lower + upper
		return randomProcessor(length, characters)

	} else if r.LowerUpperNumber {
		characters = lower + upper+ number
		return randomProcessor(length, characters)

	} else if r.LowerUpperSpecial {
		characters = lower + upper + special
		return randomProcessor(length, characters)

	} else if r.LowerUpperNumberSpecial {
		characters = lower + upper + number + special
		return randomProcessor(length, characters)

	} else if r.LowerNumber {
		characters = lower + number
		return randomProcessor(length, characters)

	} else if r.LowerNumberSpecial {
		characters = lower + number + special
		return randomProcessor(length, characters)

	} else if r.LowerSpecial {
		characters = lower + special
		return randomProcessor(length, characters)

	} else if r.Upper {
		return randomProcessor(length, upper)

	} else if r.UpperNumber {
		characters = upper + number
		return randomProcessor(length, characters)

	} else if r.UpperNumberSpecial {
		characters = upper + number + special
		return randomProcessor(length, characters)

	} else if r.UpperSpecial {
		characters = upper + special
		return randomProcessor(length, characters)

	} else if r.Number {
		return randomProcessor(length, number)

	} else if r.NumberSpecial {
		characters = number + special
		return randomProcessor(length, characters)

	} else if r.Special {
		return randomProcessor(length, special)

	}

	return ""
}

func randomProcessor(length int, charset string) string {
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
