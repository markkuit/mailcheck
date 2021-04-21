package main

import (
	"log"

	"github.com/alitto/pond"
	"github.com/markkuit/mailcheck/internal/commons"
	"github.com/markkuit/mailcheck/internal/csv"
	"github.com/markkuit/mailcheck/internal/util"
	"github.com/markkuit/mailcheck/internal/verifier"
)

func main() {
	lines, err := util.ScanFile(commons.InputFile)
	if err != nil {
		log.Fatal(err)
	}

	pool := pond.New(commons.MaxWorkers, len(lines))
	c := make(chan verifier.CheckResult, len(lines))

	for _, v := range lines {
		str := v // pesky races
		pool.Submit(func() {
			verifier.Check(str, c)
		})
	}
	pool.StopAndWait()
	close(c)

	csv.ExportResults(c)
}
