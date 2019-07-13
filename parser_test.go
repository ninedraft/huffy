package huffy

import (
	"bytes"
	"math/rand"
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

func parse(str string) []string {
	var prevUppercase = false
	var tokens []string
	var acc = &bytes.Buffer{}
	for _, r := range str {
		var isUpper = unicode.IsUpper(r)
		var emptyAcc = acc.Len() == 0
		var uppecaseChain = isUpper && prevUppercase
		var startOfToken = isUpper != prevUppercase && !prevUppercase
		if (startOfToken || uppecaseChain) && !emptyAcc {
			tokens = append(tokens, acc.String())
			acc.Reset()
		}
		acc.WriteRune(r)
		prevUppercase = isUpper
	}
	if acc.Len() != 0 {
		tokens = append(tokens, acc.String())
	}
	return tokens
}

func TestParser(test *testing.T) {
	type TCase struct {
		Input    string
		Expected []string
	}
	var wordParts = strings.Fields(`Decode reads the next json encoded value from its input and stores it in the value pointed to by v`)
	(&Tester{
		Generator: func(rnd *rand.Rand, i int) interface{} {
			var n = len(wordParts)
			var tokens []string
			for range r(i%5 + 1) {
				var word = wordParts[rnd.Intn(n)]
				tokens = append(tokens, strings.Title(word))
			}
			return TCase{
				Input:    strings.Join(tokens, ""),
				Expected: tokens,
			}
		},
		Unit: func(test *testing.T, data interface{}) {
			var tcase = data.(TCase)
			assert.Equal(test, tcase.Expected, parse(tcase.Input))
		},
	}).R(test)
}
