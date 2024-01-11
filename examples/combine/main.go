package main

import (
	"fmt"

	"github.com/seastart/cli-go"
)

func main() {

	app := cli.NewCliApp("app desc", &cli.Option{
		Name: "env",
		Dft:  "dev",
		Desc: "envirioment",
	})
	// list
	app.AddCommand("list", "list command", func(cmd *cli.Command, remaincmds []string) (err error) {
		env := cmd.App().OptVal("env")
		page := cmd.OptVal("page")
		fmt.Printf("app env=%v\n", env)
		fmt.Printf("into list command page=%v\n", page)
		return
	}, &cli.Option{
		Name: "page",
		Dft:  1,
		Desc: "page no",
	})
	app.Run()
	// ./main
	// ./main -env=prod
	// ./main -env=prod list -page=3
}
