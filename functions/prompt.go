package functions

import (
	"errors"
	"log"

	"github.com/manifoldco/promptui"
)

type PromptContent struct {
	ErrorMsg string
	Label string
}

func PromptGetInput(pc PromptContent) string {
	prompt := promptui.Prompt{
		Label: pc.Label,
		Templates: &promptui.PromptTemplates{
			Prompt: "{{ . }}",
			Valid: "{{ . | green }}",
			Invalid: "{{ . | red }}",
		},
		Validate: func(input string) error {
			if len(input) <= 0 {
				return errors.New(pc.ErrorMsg)
			}
			return nil
		},
	}

	result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	log.Printf("Input: %s\n", result)
	return result
}

func PromptGetSelect(pc PromptContent, items []string) string {
	var (
		index = -1
		result string
		err error
	)

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label: pc.Label,
			Items: items,
		}

		index, result, err = prompt.Run()
		if index == -1 {
			items = append(items, result)
		}
	}

	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	log.Printf("Input: %s\n", result)
	return result
}
