package prompter

import (
	"regexp"
	"strings"
)

// Prompt simple prompting
func Prompt(message, defaultAnswer string) string {
	return (&Prompter{
		Message: message,
		Default: defaultAnswer,
	}).Prompt()
}

// YN y/n choice
func YN(message string, defaultToYes bool) bool {
	defaultChoice := "n"
	if defaultToYes {
		defaultChoice = "y"
	}
	input := (&Prompter{
		Message:    message,
		Choices:    []string{"y", "n"},
		IgnoreCase: true,
		Default:    defaultChoice,
	}).Prompt()

	return strings.ToLower(input) == "y"
}

// YesNo yes/no choice
func YesNo(message string, defaultToYes bool) bool {
	defaultChoice := "no"
	if defaultToYes {
		defaultChoice = "yes"
	}
	input := (&Prompter{
		Message:    message,
		Choices:    []string{"yes", "no"},
		IgnoreCase: true,
		Default:    defaultChoice,
	}).Prompt()

	return strings.ToLower(input) == "yes"
}

// Password asks password
func Password(message string) string {
	return (&Prompter{
		Message: message,
		NoEcho:  true,
	}).Prompt()
}

// Choose make a choice
func Choose(message string, choices []string, defaultChoice string) string {
	return (&Prompter{
		Message: message,
		Choices: choices,
		Default: defaultChoice,
	}).Prompt()
}

// Regexp checks the answer by regexp
func Regexp(message string, reg *regexp.Regexp, defaultAnswer string) string {
	return (&Prompter{
		Message: message,
		Regexp:  reg,
		Default: defaultAnswer,
	}).Prompt()
}
