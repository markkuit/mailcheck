package commons

import "github.com/schollz/progressbar/v3"

var progressBar *progressbar.ProgressBar

func NewProgressBar(max int, description string) {
	if !Quiet {
		progressBar = progressbar.Default(int64(max), description)
	} else {
		progressBar = progressbar.DefaultSilent(int64(max), description)
	}
}

func IncrementProgressBar() error {
	if err := progressBar.Add(1); err != nil {
		return err
	}
	return nil
}

func FinishProgressBar() error {
	if err := progressBar.Finish(); err != nil {
		return err
	}
	return nil
}
