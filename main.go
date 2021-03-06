package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"

	"github.com/feliixx/mgodatagen/datagen"
)

const version = "0.7.3"

func main() {
	var options datagen.Options
	p := flags.NewParser(&options, flags.Default&^flags.HelpFlag)
	_, err := p.Parse()
	if err != nil {
		fmt.Println("try mgodatagen --help for more informations")
		os.Exit(1)
	}
	if options.Help {
		fmt.Printf("mgodatagen version %s", version)
		p.WriteHelp(os.Stdout)
		os.Exit(0)
	}

	if options.Version {
		fmt.Printf("mgodatagen version %s", version)
		os.Exit(0)
	}

	err = datagen.Generate(&options, os.Stdout)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
