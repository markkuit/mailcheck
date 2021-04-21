package util

import (
	"bufio"
	"os"

	"github.com/markkuit/mailcheck/internal/commons"
)

func ScanFile(path string) ([]string, error) {
	var lines []string

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())

		if err := commons.IncrementProgressBar(); err != nil {
			return []string{}, err
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if err := commons.FinishProgressBar(); err != nil {
		return []string{}, err
	}
	return lines, nil
}
