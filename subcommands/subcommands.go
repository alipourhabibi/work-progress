package subcommands

import (
	"os"

	"github.com/alipourhabibi/work-progress/actions"
	"github.com/alipourhabibi/work-progress/flags"
	"github.com/urfave/cli/v2"
)

type jobApp struct {
	app *cli.App
}

func NewJobApp(app *cli.App) *jobApp {
	return &jobApp{app: app}
}

func (w *jobApp) InitSubCommands() {
	job := jobSubCommand()
	chart := chartCommands()
	w.app.Commands = append(w.app.Commands, job...)
	w.app.Commands = append(w.app.Commands, chart...)
}

func jobSubCommand() []*cli.Command {

	jobCommands := []*cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add jobs to the jobs list",
			Action:  actions.AddJob,
			Flags:   flags.JobAddFlags(),
		},
		{
			Name:    "modify",
			Aliases: []string{"m"},
			Usage:   "modify a job based on given name",
			Action:  actions.ModifyJob,
			Flags:   flags.JobModifyFlags(),
		},
		{
			Name:    "delete",
			Aliases: []string{"d"},
			Usage:   "delete a job based on given name",
			Action:  actions.DeleteJob,
			Flags:   flags.JobDeleteFlags(),
		},
		{
			Name:    "get",
			Aliases: []string{"g"},
			Usage:   "get a job based on given name",
			Action:  actions.GetJob,
			Flags:   flags.JobGetFlags(),
		},
	}

	job := []*cli.Command{
		{
			Name:        "jobs",
			Usage:       "Do CRUD on the jobs list",
			Subcommands: jobCommands,
		},
	}

	return job
}

func chartCommands() []*cli.Command {

	chart := []*cli.Command{
		{
			Name:    "draw",
			Aliases: []string{"d"},
			Usage:   "draw charts",
			Action:  actions.DrawJob,
			Flags:   flags.JobDraw(),
		},
	}
	return chart
}

func (w *jobApp) Run() {
	w.app.Run(os.Args)
}
