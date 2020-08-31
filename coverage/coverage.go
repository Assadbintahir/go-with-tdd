package coverage

import (
	"fmt"
	"os"
	"testing"
)

/*
If any other work is required in TestMain, then `m.Run` & `os.Exit` should be part of
TestMain and only logic for enforcing the coverage residue here.
*/
func EnforceCoverage(m *testing.M, threshold float64, packageName string) {
	rc := m.Run()

	// rc 0 means we've passed,
	// and CoverMode will be non empty if run with -cover
	if rc == 0 && testing.CoverMode() != "" {
		c := testing.Coverage()
		if c < threshold {
			fmt.Printf("Coverage for package %v failed at %v \n", packageName, c)
			rc = -1
		}
		fmt.Printf("Coverage passed for package %v \n", packageName)
	}
	os.Exit(rc)
}
