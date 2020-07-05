package ui

import (
	"strings"

	"github.com/manifoldco/promptui"
)

//Option is the interface which an option passed to Select must adhere to
type Option interface {
	Display() string
}

type StringOption string

func (s StringOption) Display() string {
	return string(s)
}

// API defines the interface for the choice service
type API interface {
	GetUserChoice(question string, choices []Option) int
}

// New returns an instance of the choice service
func New() API {
	return &Service{}
}

// Service is a concrete type for the choice service
type Service struct {
}

// GetUserChoice prompts the user with a list of choices, and return the selected index (or -1 if an error occurs)
func (s *Service) GetUserChoice(question string, choices []Option) int {
	prompt := promptui.Select{
		Label: question,
		Items: choices,
		Searcher: func(query string, index int) bool {
			label := choices[index].Display()
			return strings.Contains(strings.ToLower(label), strings.ToLower(query))
		},
		Templates: &promptui.SelectTemplates{
			Label:    "{{.Display}}",
			Active:   "=> {{.Display}}",
			Inactive: "   {{.Display}}",
		},
	}

	index, _, err := prompt.Run()

	if err != nil || index == -1 {
		return -1
	}

	return index
}
