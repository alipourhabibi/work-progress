package flags

import "github.com/urfave/cli/v2"

func JobDraw() []cli.Flag {
	drawFlags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:     "name",
			Aliases:  []string{"n"},
			Usage:    "name of job",
			Required: false,
		},
		&cli.StringSliceFlag{
			Name:     "time",
			Aliases:  []string{"t"},
			Usage:    "time of job based on regex",
			Required: false,
		},
	}
	return drawFlags
}
