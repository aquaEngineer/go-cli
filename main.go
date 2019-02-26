package main

import (
	    "os"
		"fmt"
		"github.com/urfave/cli"
		finder "github.com/b4b4r07/go-finder"
		"github.com/b4b4r07/go-finder/source"
		)

func main() {
	app := cli.NewApp()
	app.Name = "sampleApp"
	app.Usage = "This app echo input arguments"
	app.Version = "0.0.1"
	app.Action = func (context *cli.Context) error {
		if context.Bool("show") {
			fzf, err := finder.New("fzf")
			if err != nil {
				panic(err)
			}

			// If needed, install fzf to the path
			fzf.Install("/usr/local/bin")
			fmt.Printf("fzf obeject:   %#v\n", fzf)

			// Read files list within dir as data source of fzf
			fzf.Read(source.Dir(".", true))

			items, err := fzf.Run()
			if err != nil {
				panic(err)
			}
			fmt.Printf("selected items:%#v\n", items)
		}
		return nil
	}
	app.Flags = []cli.Flag{
		cli.BoolFlag {
				Name: "show, s",
				Usage: "show",
			},
	}
	app.Run(os.Args)
}
