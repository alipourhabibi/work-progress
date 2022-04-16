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
	amount := c.StringSlice("a")
	amountFloat, err := strconv.ParseFloat(amount[0], 32)
	if err != nil {
		fmt.Printf("[ERROR] Invalid amont\n")
		return err
	}
	if len(description) <= 0 {
		description = []string{""}
	}
	job := repository.NewJob(name[0], description[0], amountFloat, "")
	job.Add()
	return nil
}

func ModifyJob(c *cli.Context) error {
	name := c.StringSlice("n")
	description := c.StringSlice("d")
	amount := c.StringSlice("a")
	newName := c.StringSlice("nn")
	var amountFloat float64
	var err error
	if len(amount) != 0 {
		amountFloat, err = strconv.ParseFloat(amount[0], 32)
		if err != nil {
			fmt.Printf("[ERROR] Invalid amont\n")
			return err
		}
	} else {
		amountFloat = -1
	}
	if len(description) == 0 {
		description = []string{""}
	}
	if len(newName) == 0 {
		newName = []string{""}
	}
	job := repository.NewJob(name[0], description[0], amountFloat, "")
	job.Modify(newName[0])
	return nil
}

func DeleteJob(c *cli.Context) error {
	name := c.StringSlice("n")
	time := c.StringSlice("t")

	job := repository.NewJob(name[0], "", 0, time[0])
	job.Delete()

	return nil
}

func GetJob(c *cli.Context) error {
	name := c.StringSlice("n")
	time := c.StringSlice("t")

	if len(name) == 0 {
		name = []string{""}
	}
	if len(time) == 0 {
		time = []string{""}
	}

	job := repository.NewJob(name[0], "", 0, time[0])
	job.Get()

	return nil
}
