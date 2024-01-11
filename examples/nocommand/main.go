package main

import (
	"github.com/seastart/cli-go"
)

func main() {
	// "" means the main app is one command
	app := cli.NewCliWholeApp("app desc", func(cmd *cli.Command, remaincmds []string) (err error) {
		start, err := cmd.OptVal("start")
		if err != nil {
			return
		}
		cmd.App().Infof("into default command start=%v\n", start)
		cmd.App().Warningf("into default command start=%v\n", start)
		cmd.App().Successf("into default command start=%v\n", start)
		cmd.App().Errorf("into default command start=%v\n", start)
		cmd.App().Exitf(0, "into default command start=%v\n", start)
		return
	}, &cli.Option{
		Name:  "start",
		Dft:   0,
		Usage: "begin no",
	})
	app.Run()
	// ./main
	// ./main -start=2
}
