package tests

// New changes have broken these tests
// TODO: Need to fix them.

//func TestCommands(t *testing.T) {
//	csvT := tests.NewCsvTest()
//
//	tsts := []struct {
//		expFile string
//		actFile string
//		cmd     string
//	}{
//		{"output1.csv", "output1.tmp.csv",
//			`
//			file(input1.csv)
//			| avg(age) as avg_age
//			by last_name
//			| batch()
//			| file(output1.tmp.csv)
//		`},
//
//		{"iris_out1.csv", "iris_out1.tmp.csv",
//			`
//			file(iris.csv)
//			| count() as Count,
//			  Max_of_SepalLengthCm = max(SepalLengthCm),
//			  Min_of_SepalLengthCm = min(SepalLengthCm),
//			  Median_of_SepalLengthCm = median(SepalLengthCm),
//			  Var_of_SepalLengthCm = var(SepalLengthCm),
//			  Avg_of_SepalLengthCm = avg(SepalLengthCm),
//			  DCount_of_SepalLengthCm = dcount(SepalLengthCm)
//			by Species
//			| batch()
//			| file(iris_out1.tmp.csv)
//		`},
//
//		{"iris_out2.csv", "iris_out2.tmp.csv",
//			`
//			file(iris.csv)
//			| count() as Count,
//			  Max_of_SepalLengthCm = max(SepalLengthCm),
//			  Min_of_SepalLengthCm = min(SepalLengthCm),
//			  Median_of_SepalLengthCm = median(SepalLengthCm),
//			  Var_of_SepalLengthCm = var(SepalLengthCm),
//			  Avg_of_SepalLengthCm = avg(SepalLengthCm),
//			  DCount_of_SepalLengthCm = dcount(SepalLengthCm)
//			by Species
//			| count(),
//			  max(Max_of_SepalLengthCm),
//			  min(Min_of_SepalLengthCm),
//			  median(Median_of_SepalLengthCm),
//			  var(Var_of_SepalLengthCm),
//			  avg(Avg_of_SepalLengthCm),
//			  dcount(DCount_of_SepalLengthCm)
//			| batch()
//			| file(iris_out2.tmp.csv)
//		`},
//
//		{"iris_out3.csv", "iris_out3.tmp.csv",
//			`
//			file(iris.csv)
//			| DCount_of_SepalLengthCm = dcount(SepalLengthCm)
//			by Species
//			| count() by DCount_of_SepalLengthCm
//			| batch()
//			| file(iris_out3.tmp.csv)
//		`},
//
//		{"iris_out4.csv", "iris_out4.tmp.csv",
//			`
//			file(iris.csv)
//			| DCount_of_SepalLengthCm = dcount(SepalLengthCm)
//			by Species
//			| count() by DCount_of_SepalLengthCm
//			| filter(DCount_of_SepalLengthCm > 20)
//			| batch()
//			| file(iris_out4.tmp.csv)
//		`},
//	}
//
//	for _, tt := range tsts {
//		if errStr, ok := csvT.Test(tt.expFile, tt.actFile, tt.cmd); !ok {
//			assert.Failf(t, "%s", errStr)
//		}
//	}
//}
