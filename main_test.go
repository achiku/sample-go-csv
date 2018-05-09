package main

import (
	"io"
	"strings"
	"testing"

	csv "github.com/JensRantil/go-csv"
)

var data = `1|data 1|value 1
2|data 2 \| this is value too|value2
3|data 3|value3`

var dataOK = `1|data 1|value 1
2|"data 2 \| this is value too"|value2
3|data 3|value3`

func TestParseBackslashedEscapeCSV(t *testing.T) {
	for _, source := range []string{data, dataOK} {
		d := strings.NewReader(source)
		dialect := csv.Dialect{
			Delimiter: '|',
			// EscapeChar:  '\\',
			DoubleQuote: csv.NoDoubleQuote,
		}
		r := csv.NewDialectReader(d, dialect)

		for {
			rec, err := r.Read()
			if err != nil {
				if err == io.EOF {
					break
				}
				t.Fatal(err)
			}
			t.Logf("%#+v", rec)
		}
	}
}
