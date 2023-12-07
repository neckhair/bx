package cli

import "github.com/fatih/color"

func PrintError(msg string) {
	color.Red(msg)
}

func Bold(msg string) {
	c := color.New(color.Bold)
	c.Println(msg)
}
