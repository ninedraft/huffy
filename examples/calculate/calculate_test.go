package countchars

import (
	"math/rand"
	"testing"

	"github.com/ninedraft/huffy"
)

func TestDiv(test *testing.T) {
	type TestCase struct {
		X, Y     int
		Expected int
	}

	huffy.Tester{
		Generator: func(rnd *rand.Rand, id int) interface{} {
			var x = rnd.Intn(100) + 2
			var y = rnd.Intn(x-1) + 1
			var expected = x / y
			if x%y != 0 {
				expected++
			}
			return TestCase{
				X:        x,
				Y:        y,
				Expected: expected,
			}
		},
		Unit: func(test *testing.T, v interface{}) {
			var tc = v.(TestCase)
			var got = Div(tc.X, tc.Y)
			if tc.Expected != got {
				test.Fatalf("%d/%d: expected %d, got %d", tc.X, tc.Y, tc.Expected, got)
			}
		},
	}.R(test)
}
