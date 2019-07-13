package huffy

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

func (tester *Tester) runMemorizedTests(test *testing.T) {
	var testdata, errOpen = os.Open(tester.TestFile)
	switch {
	case os.IsNotExist(errOpen):
		return
	case errOpen != nil:
		test.Fatalf("unable to run memorized tests: %v", errOpen)
	}
	defer testdata.Close()

	var decoder = json.NewDecoder(testdata)
	var newArg, elem = unitTestArgFactory(tester.Generator(tester.Rnd, 0))

	for decoder.More() {
		var tc = TestCase{
			Data: newArg(),
		}
		if err := decoder.Decode(&tc); err != nil {
			test.Fatalf("unable to run memorized tests: %v", errOpen)
		}
		tester.Unit(test, elem(tc.Data))
	}
}

type newArgFactory func() interface{}

type elem func(interface{}) interface{}

func unitTestArgFactory(paragon interface{}) (newArgFactory, elem) {
	var tt = reflect.TypeOf(paragon)
	var elem = func(v interface{}) interface{} {
		return reflect.ValueOf(v).Elem().Interface()
	}
	if tt.Kind() == reflect.Ptr {
		tt = tt.Elem()
		elem = func(v interface{}) interface{} { return v }
	}
	return func() interface{} {
		return reflect.New(tt).Interface()
	}, elem
}
