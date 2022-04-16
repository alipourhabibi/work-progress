package main

import (
	"errors"
	"os"

	"github.com/alipourhabibi/work-progress/files"
	"github.com/alipourhabibi/work-progress/subcommands"
	"github.com/urfave/cli/v2"
)

func main() {
	files.Init()
	if _, err := os.Stat(files.WORK); errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(files.PARRENT, 0777)
		if err != nil {
			panic(err)
		}
		o, err := os.Create(files.WORK)
		if err != nil {
			panic(err)
		}
		defer o.Close()
	}

	app := cli.NewApp()
	jobApp := subcommands.NewJobApp(app)
	jobApp.InitSubCommands()
	jobApp.Run()
}
