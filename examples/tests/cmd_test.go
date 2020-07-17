package tests

import (
	"testing"

	"github.com/raralabs/hike/tests"
	"github.com/stretchr/testify/assert"
)

func TestCommands(t *testing.T) {
	csvT := tests.NewCsvTest()

	tsts := []struct {
		expFile string
		actFile string
		cmd     string
	}{
		{"output1.csv", "output1.tmp.csv", `file(input1.csv) | avg(age) as avg_age by last_name`},
	}

	for _, tt := range tsts {
		if errStr, ok := csvT.Test(tt.expFile, tt.actFile, tt.cmd); !ok {
			assert.Failf(t, "%s", errStr)
		}
	}
}
