package actions

import (
	"github.com/alipourhabibi/work-progress/repository"
	"github.com/urfave/cli/v2"
)

func DrawJob(c *cli.Context) error {
	name := c.StringSlice("n")
	time := c.StringSlice("t")

	if len(name) == 0 {
		name = []string{""}
	}
	if len(time) == 0 {
		time = []string{""}
	}

	job := repository.NewJob(name[0], "", 0, time[0])
	job.Draw()

	return nil
}
