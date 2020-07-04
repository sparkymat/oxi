package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kingpin"
)

var (
	app = kingpin.New("oxi", "oxi is a commandline HTTP client").Version("0.0.1")
	url = app.Arg("url", "URL to fetch").String()
)

func main() {
	if _, err := app.Parse(os.Args[1:]); err != nil {
		panic(err)
	}

	if url != nil && *url != "" {
		fmt.Printf("GET %s", *url)
	} else {
		fmt.Println("Entering interactive mode")
	}
}
