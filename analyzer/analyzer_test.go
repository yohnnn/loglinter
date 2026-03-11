package analyzer

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestLoglinter(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, New(), "logtest")
}
