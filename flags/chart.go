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
		&cli.StringSliceFlag{
			Name:     "chart",
			Aliases:  []string{"c"},
			Usage:    "give the type of chart",
			Required: true,
		},
		&cli.StringSliceFlag{
			Name:     "port",
			Aliases:  []string{"p"},
			Usage:    "port which chart will be shown",
			Required: false,
		},
	}
	return drawFlags
}
