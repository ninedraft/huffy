package huffy

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func (tester *Tester) runGeneratedTests(test *testing.T) {
	var caseFile, errOpen = os.OpenFile(tester.CaseFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if errOpen != nil {
		test.Fatalf("[huffy.Tester.runGeneratedTests] unable to open case file %q: %v", tester.CaseFile, errOpen)
	}
	defer caseFile.Close()
	var encoder = json.NewEncoder(caseFile)
	for i := range r(tester.N) {
		var name, tcase = tester.Generator(i)
		if tcase == nil {
			break
		}
		if name == "" {
			name = fmt.Sprintf("%d", i)
		}
		var ok = test.Run(name, func(test *testing.T) {
			tester.Unit(test, tcase)
		})
		if !ok {
			if err := encoder.Encode(TestCase{
				ID:   tester.Rnd.Int63(),
				Name: name,
				Data: tcase,
			}); err != nil {
				test.Fatalf("[huffy.Tester.runGeneratedTests] unable to encode test case %v", err)
			}
		}
	}
}
