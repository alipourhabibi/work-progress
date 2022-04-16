package actions

import (
	"github.com/alipourhabibi/work-progress/repository"
	"github.com/urfave/cli/v2"
)

func AddJob(c *cli.Context) error {
	name := c.StringSlice("n")
	description := c.StringSlice("d")
	if len(description) <= 0 {
		description = []string{""}
	}
	job := repository.NewJob(name[0], description[0])
	job.Add()
	return nil
}
