package commons

import (
	"flag"
	"fmt"
	"os"
)

var (
	InputFile, OutputFile string
	HelloName, FromEmail  string
	MaxWorkers            int
)

func init() {
	help := flag.Bool("h", false, "show usage")
	flag.StringVar(&InputFile, "i", "addresses.txt", "input file name")
	flag.StringVar(&OutputFile, "o", "mailcheck.csv", "output CSV file name")
	flag.StringVar(&HelloName, "n", "no-reply.net", "name for SMTP HELO command")
	flag.StringVar(&FromEmail, "f", "mailcheck@no-reply.net", "email address for SMTP MAIL FROM command")
	flag.IntVar(&MaxWorkers, "j", 10, "maximum number of workers for concurrent processing")
	flag.Parse()

	if *help {
		fmt.Printf("mailcheck version %s\n", Version)
		flag.Usage()
		os.Exit(0)
	}
}
