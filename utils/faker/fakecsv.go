package faker

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"time"
)

func FakeCsvData(w io.Writer, m map[string][]interface{}, numRows int) {

	rand.Seed(time.Now().Unix())
	writer := csv.NewWriter(w)

	fields := make([]string, len(m))
	i := 0
	for k := range m {
		fields[i] = k
		i++
	}

	writer.Write(fields)

	for i := 0; i < numRows; i++ {
		record := make([]string, len(fields))
		for j := 0; j < len(fields); j++ {
			choices := m[fields[j]]
			if choices == nil {
				log.Panic("No choices to create fake data")
			}
			choice := choices[rand.Intn(len(choices))]
			record[j] = fmt.Sprintf("%v", choice)
		}
		writer.Write(record)
	}

	writer.Flush()
}
