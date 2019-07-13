package huffy

import (
	"math/rand"
	"path/filepath"
	"testing"
	"time"
)

// Generator produces testdata for unit test
type Generator func(int) interface{}

// Unit represents unit test. Must not call t.Parallel()!
type Unit func(*testing.T, interface{})

// Tester tries to to find generated test cases, which cause failed tests and write correpsonding test data to Tester.TestFile
type Tester struct {
	N         int
	TestFile  string
	Unit      Unit
	Generator Generator
	Rnd       *rand.Rand
}

func (tester *Tester) init(test *testing.T) {
	if tester.N <= 0 {
		tester.N = 400
	}
	if tester.TestFile == "" {
		tester.TestFile = filepath.Join("testdata", test.Name()+".json")
	}
	if tester.Rnd == nil {
		var seed = time.Now().UnixNano()
		var source = rand.NewSource(seed)
		tester.Rnd = rand.New(source)
	}
}

// R runs memorized tests, when runs at most N tests with generated test args.
// If tester.TestFile is "", then uses default value "testdata/TESTNAME.json"
// If Tester.N <= 0, then uses default value 400.
// If Tester.Rnd is nil, then uses random generator initialized with time.Now().UnixNano().
// If any test fails, then Tester.R writes corresponding arg to tester.TestFile
func (tester Tester) R(test *testing.T) {
	tester.init(test)
	tester.runMemorizedTests(test)
	tester.runGeneratedTests(test)
}

func (tester *Tester) unitTest(data interface{}) func(*testing.T) {
	return func(test *testing.T) {
		tester.Unit(test, data)
	}
}

// TestCase is a storable data object, which plays role of container for generated test cases.
// Tester encoded
type TestCase struct {
	ID   int64       `json:"id"`
	Data interface{} `json:"data"`
}
