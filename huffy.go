package huffy

import (
	"math/rand"
	"path/filepath"
	"testing"
	"time"
)

type Generator func(int) interface{}

type Unit func(*testing.T, interface{})

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

type TestCase struct {
	ID   int64       `json:"id"`
	Data interface{} `json:"data"`
}
