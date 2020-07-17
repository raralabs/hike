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
		{"iris_out1.csv", "iris_out1.tmp.csv", `
		file(iris.csv)
		| count() as Count,
		 max(SepalLengthCm) as Max_of_SepalLengthCm,
		 min(SepalLengthCm) as Min_of_SepalLengthCm,
		 median(SepalLengthCm) as Median_of_SepalLengthCm,
		 var(SepalLengthCm) as Var_of_SepalLengthCm,
		 avg(SepalLengthCm) as Avg_of_SepalLengthCm,
		 dcount(SepalLengthCm) as DCount_of_SepalLengthCm
		by Species
		`},
		{"iris_out2.csv", "iris_out2.tmp.csv", `
		file(iris.csv)
		| count() as Count,
		 max(SepalLengthCm) as Max_of_SepalLengthCm,
		 min(SepalLengthCm) as Min_of_SepalLengthCm,
		 median(SepalLengthCm) as Median_of_SepalLengthCm,
		 var(SepalLengthCm) as Var_of_SepalLengthCm,
		 avg(SepalLengthCm) as Avg_of_SepalLengthCm,
		 dcount(SepalLengthCm) as DCount_of_SepalLengthCm
		by Species
		| count(),
		 max(Max_of_SepalLengthCm),
		 min(Min_of_SepalLengthCm),
		 median(Median_of_SepalLengthCm),
		 var(Var_of_SepalLengthCm),
		 avg(Avg_of_SepalLengthCm),
		 dcount(DCount_of_SepalLengthCm)
		`},
		{"iris_out3.csv", "iris_out3.tmp.csv", `
		file(iris.csv)
		| dcount(SepalLengthCm) as DCount_of_SepalLengthCm
		by Species
		| count() by DCount_of_SepalLengthCm
		| batch()
		| filter(Count > 0)
		`},
	}

	for _, tt := range tsts {
		if errStr, ok := csvT.Test(tt.expFile, tt.actFile, tt.cmd); !ok {
			assert.Failf(t, "%s", errStr)
		}
	}
}
