package huffy

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

func (tester *Tester) runMemorizedTests(test *testing.T) {
	var caseFile, errOpen = os.Open(tester.CaseFile)
	if errOpen != nil && !os.IsNotExist(errOpen) {
		test.Fatalf("[huffy.Tester.runMemorizedTests] unable to open case file %q: %v", tester.CaseFile, errOpen)
	}
	defer caseFile.Close()
	var decoder = json.NewDecoder(caseFile)
	var _, testCaseExample = tester.Generator(0)
	var newTestCase = testCaseFactory(testCaseExample)
	for decoder.More() {
		var testCase = newTestCase()
		if err := decoder.Decode(&testCase); err != nil {
			test.Fatalf("[huffy.Tester.runMemorizedTests] unable to decode test case: %v", err)
		}
		var data = reflect.ValueOf(testCase.Data).Elem().Interface()
		test.Run(testCase.Name, tester.unitTest(data))
	}
}

func testCaseFactory(example interface{}) func() TestCase {
	var tt = reflect.TypeOf(example)
	return func() TestCase {
		return TestCase{
			Data: reflect.New(tt).Interface(),
		}
	}
}
