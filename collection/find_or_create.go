package collection

import (
	"errors"
	"fmt"

	"github.com/mgutz/ansi"
	"github.com/sparkymat/oxi/ui"
)

func FindOrCreate() (*Type, error) {
	options := []ui.Option{}

	bold := ansi.ColorFunc("white+b")
	greenBold := ansi.ColorFunc("green+b")

	options = append(options, ui.StringOption(fmt.Sprintf("[ %s ]", greenBold("CREATE"))))

	list := []Type{
		{Name: "V1"},
	}

	for index, item := range list {
		item := item
		label := item.Name

		options = append(options, ui.StringOption(fmt.Sprintf("%d. %s", index+1, bold(label))))
	}

	chooser := ui.New()
	chooser.GetUserChoice("Select collection", options)

	return nil, errors.New("x")
}
