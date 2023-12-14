package cli

import "github.com/janeczku/go-spinner"

type Spinner struct {
	spinner *spinner.Spinner
}

func StartSpinner(title string) Spinner {
	s := spinner.NewSpinner(title)
	s.Start()
	return Spinner{spinner: s}
}

func (s *Spinner) Start() {
	s.spinner.Start()
}

func (s *Spinner) Stop() {
	s.spinner.Stop()
}
