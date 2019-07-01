package huffy

import (
	"math/rand"
	"path/filepath"
	"testing"
)

const SEED = 42

type Generator func(int) (string, interface{})

type Unit func(*testing.T, interface{})

type Tester struct {
	N         int
	CaseFile  string
	Unit      Unit
	Generator Generator
	Rnd       *rand.Rand
}

func (tester *Tester) R(test *testing.T) {
	if tester.N <= 0 {
		tester.N = 400
	}
	if tester.CaseFile == "" {
		tester.CaseFile = filepath.Join("testdata", test.Name()+".json")
	}
	if tester.Rnd == nil {
		tester.Rnd = rand.New(rand.NewSource(SEED))
	}
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
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

func r(n int) []struct{} {
	return make([]struct{}, n)
}
