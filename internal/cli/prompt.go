package cli

import (
	"errors"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

func PromptGetInput(label, errorMsg string) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(errorMsg)
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }}: ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     label,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return result
}
