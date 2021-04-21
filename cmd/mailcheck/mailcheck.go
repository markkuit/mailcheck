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
	commons.NewProgressBar(-1, "Reading input file")
	lines, err := util.ScanFile(commons.InputFile)
	if err != nil {
		log.Fatal(err)
	}

	pool := pond.New(commons.MaxWorkers, len(lines))
	c := make(chan verifier.CheckResult, len(lines))

	commons.NewProgressBar(len(lines), "Processing emails")
	for _, v := range lines {
		str := v // pesky races
		pool.Submit(func() {
			if err := verifier.Check(str, c); err != nil {
				log.Fatal(err)
			}
		})
	}
	pool.StopAndWait()
	close(c)

	commons.NewProgressBar(len(lines), "Exporting results")
	if err := csv.ExportResults(c); err != nil {
		log.Fatal(err)
	}
}
