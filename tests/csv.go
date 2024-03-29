package tests

import (
	"encoding/csv"
	"fmt"
	"github.com/raralabs/hike/parser/single"
	"io"
	"log"
	"os"
	"reflect"

	"github.com/raralabs/hike/cli"
	"github.com/raralabs/hike/parser"
)

type CsvTest struct {
	prsr parser.IParser
}

func NewCsvTest() *CsvTest {

	prsr := single.NewParser()
	// Initialize the parser
	prsr.Init()

	return &CsvTest{
		prsr: prsr,
	}
}

func (ct *CsvTest) run(cmd string) {

	cmds := []string{cmd + ";", "end"}
	i := 0

	cli.SingleCommandMode(func(prompt string) string {
		defer func() {
			i++
		}()
		return cmds[i]
	}, func(cmd string) {}, ct.prsr)
}

func (ct *CsvTest) Test(expFile, actFile, cmd string) (string, bool) {

	exp, err := os.Open(expFile)
	if err != nil {
		log.Panic(err)
	}
	defer exp.Close()

	ct.run(cmd)

	act, err := os.Open(actFile)
	if err != nil {
		log.Panic(err)
	}
	defer act.Close()

	actual := csv.NewReader(act)
	expected := csv.NewReader(exp)

	actual.TrimLeadingSpace = true
	expected.TrimLeadingSpace = true

	i := 0
	for {
		expRec, err := expected.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Panic(err)
		}

		actRec, err := actual.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Panic(err)
		}
		i++

		if !reflect.DeepEqual(actRec, expRec) {
			str := fmt.Sprintf("Want: %v \nGot: %v \nLine: %d", expRec, actRec, i)
			return str, false
		}
	}

	return "", true
}
