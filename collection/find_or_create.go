package collection

import (
	"fmt"

	"github.com/mgutz/ansi"
	"github.com/sparkymat/oxi/ui"
)

func FindOrCreate() (*Type, error) {
	options := []ui.Option{}

	bold := ansi.ColorFunc("white+b")
	greenBold := ansi.ColorFunc("green+b")

	options = append(options, ui.StringOption(fmt.Sprintf("[ %s ]", greenBold("CREATE"))))

	var list []Type
	var err error
	if list, err = FindAll(); err != nil {
		return nil, err
	}

	for index, item := range list {
		item := item
		label := item.Name

		options = append(options, ui.StringOption(fmt.Sprintf("%d. %s", index+1, bold(label))))
	}

	chooser := ui.New()
	choice := chooser.GetUserChoice("Select collection", options)
	if choice == 0 {
		return &Type{Name: "foo"}, nil
	}

	return &list[choice-1], nil
}
