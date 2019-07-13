package huffy

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"
)

func (tester *Tester) runGeneratedTests(test *testing.T) {
	var testdata, errOpen = os.OpenFile(tester.TestFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.ModePerm)
	if errOpen != nil {
		test.Fatalf("unable to run generated tests: unable to open testdata file: %v", errOpen)
	}
	defer func() {
		if err := testdata.Close(); err != nil {
			test.Fatalf("unable to run generated tests: unable to close testdata file: %v", err)
		}
	}()
	var encoder = json.NewEncoder(testdata)
	for i := range r(tester.N) {
		var arg = tester.Generator(i)
		var id = tester.Rnd.Int63()
		if !test.Run(fmt.Sprintf("%d", id), tester.unitTest(arg)) {
			if err := encoder.Encode(TestCase{ID: id, Data: arg}); err != nil {
				log.Fatalf("unable to run generated tests: unable to write testdata to file: %v", err)
			}
		}
	}
}
