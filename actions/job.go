package actions

import (
	"fmt"
	"strconv"

	"github.com/alipourhabibi/work-progress/repository"
	"github.com/urfave/cli/v2"
)

func AddJob(c *cli.Context) error {
	name := c.StringSlice("n")
	description := c.StringSlice("d")
	amount := c.StringSlice("t")
	amountFloat, err := strconv.ParseFloat(amount[0], 32)
	if err != nil {
		fmt.Printf("[ERROR] Invalid time\n")
		return err
	}
	if len(description) <= 0 {
		description = []string{""}
	}
	job := repository.NewJob(name[0], description[0], amountFloat)
	job.Add()
	return nil
}
