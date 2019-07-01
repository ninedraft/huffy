package huffy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit(test *testing.T) {
	type testCase struct {
		unit     interface{}
		args     []interface{}
		panic    bool
		expected bool
	}
	var tcases = map[string]testCase{
		"no args no panic": {
			unit:     func() bool { return false },
			expected: false,
		},
		"no args panic": {
			unit:  func() bool { panic("baccano!") },
			panic: true,
		},
		"invalid return types: int": {
			unit:  func() int { return 0 },
			panic: true,
		},
		"invalid return types: empty return": {
			unit:  func() {},
			panic: true,
		},
		"sum ints: true": {
			unit: func(sum, a, b, c int) bool {
				return sum == a+b+c
			},
			args:     []interface{}{42, 32, 8, 2},
			expected: true,
		},
		"sum ints: false": {
			unit: func(sum, a, b, c int) bool {
				return sum == a+b+c+1
			},
			args:     []interface{}{42, 32, 8, 2},
			expected: false,
		},
		"unit is not function": {
			unit:  666,
			panic: true,
		},
		"unit is nil": {
			panic: true,
		},
	}

	for name, tcase := range tcases {
		tcase := tcase
		test.Run(name, func(test *testing.T) {
			var tfunc = func() {
				var testUnit = unit(tcase.unit)
				assert.Equalf(test, tcase.expected, testUnit(tcase.args...), "expected unit to return %v on %v", tcase.expected, tcase.args)
			}
			if tcase.panic {
				assert.Panics(test, tfunc)
			} else {
				assert.NotPanics(test, tfunc)
			}
		})
	}
}
