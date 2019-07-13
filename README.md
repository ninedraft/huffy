# huffy

Huffy is a small library for unit testing in symbiosis with test data generators. It allows you to remember the test arguments that caused the fail test. At the next test run, these arguments will be used first.

## Example

```go
// calculate.go

package countchars

// Div must return result of division, rounded to biggest value
func Div(x, y int) int {
	return x / y
}

```

```go
// calculate_test.go
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
```

```bash
go test -timeout 30s github.com/ninedraft/huffy/examples/calculate -run ^(TestDiv)$ -race
```

```
--- FAIL: TestDiv (0.00s)
    huffy/examples/calculate/calculate_test.go:34: 44/25: expected 2, got 1
FAIL
FAIL	github.com/ninedraft/huffy/examples/calculate	0.041s
Error: Tests failed.
```