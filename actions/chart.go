package actions

import (
	"fmt"
	"strconv"

	"github.com/alipourhabibi/work-progress/repository"
	"github.com/urfave/cli/v2"
)

func DrawJob(c *cli.Context) error {
	name := c.StringSlice("n")
	time := c.StringSlice("t")
	chart := c.StringSlice("c")
	port := c.StringSlice("p")

	if len(name) == 0 {
		name = []string{""}
	}
	var intPort int
	var err error
	if len(port) == 0 {
		intPort = 8080
	} else {
		intPort, err = strconv.Atoi(port[0])
		if err != nil {
			fmt.Printf("[ERROR] Invalid port given\n")
			return err
		}
	}
	if len(time) == 0 {
		time = []string{""}
	}

	job := repository.NewJob(name[0], "", 0, time[0])
	job.Draw(chart[0], intPort)

	return nil
}
