package csv

import (
	"encoding/csv"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/markkuit/mailcheck/internal/commons"
	"github.com/markkuit/mailcheck/internal/verifier"
)

func header() []string {
	var header []string

	t := verifier.CheckResult{}
	e := reflect.ValueOf(&t).Elem()
	for i := 0; i < e.NumField(); i++ {
		header = append(header, strings.ToLower(e.Type().Field(i).Name))
	}

	return header
}

func ExportResults(c <-chan verifier.CheckResult) {
	f, err := os.Create(commons.OutputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	w.Write(header())

	for r := range c {
		w.Write(r.StringSlice())
	}
	w.Flush()
}
