package utils

import (
	"github.com/manifoldco/promptui"
	"strings"
)

func GetPromptString(label string, mask bool) string {
	var prompt promptui.Prompt
	if mask {
		prompt = promptui.Prompt{
			Label: label,
			Mask: '*',
		}
	}else {
		prompt = promptui.Prompt{
			Label: label,
		}
	}

	result, _ := prompt.Run()
	return result
}

func PromptSearchStrings(names []string, label string) string {
	searcher := func(input string, index int) bool {
		projectName := strings.Replace(strings.ToLower(names[index]), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(projectName, input)
	}

	prompt := promptui.Select{
		Label:    label,
		Items:    names,
		Searcher: searcher,
	}

	_, result, _ := prompt.Run()
	return result
}