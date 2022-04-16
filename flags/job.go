package flags

import (
	"github.com/urfave/cli/v2"
)

func JobAddFlags() []cli.Flag {
	addFlags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:     "name",
			Aliases:  []string{"n"},
			Usage:    "name of job",
			Required: true,
		},
		&cli.StringSliceFlag{
			Name:     "description",
			Aliases:  []string{"d"},
			Usage:    "description of the job",
			Required: false,
		},
	}

	return addFlags
}

func JobModifyFlags() []cli.Flag {
	modifyFlags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:     "name",
			Aliases:  []string{"n"},
			Usage:    "name of job you wish to modify",
			Required: true,
		},
		&cli.StringSliceFlag{
			Name:     "new-name",
			Aliases:  []string{"nn"},
			Usage:    "new name of job",
			Required: false,
		},
		&cli.StringSliceFlag{
			Name:     "new-description",
			Aliases:  []string{"nd"},
			Usage:    "new description",
			Required: false,
		},
	}

	return modifyFlags
}

func JobDeleteFlags() []cli.Flag {

	deleteFlags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:     "name",
			Aliases:  []string{"n"},
			Usage:    "name of job",
			Required: true,
		},
	}
	return deleteFlags
}

func JobGetFlags() []cli.Flag {
	getFlags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:     "name",
			Aliases:  []string{"n"},
			Usage:    "name of job",
			Required: true,
		},
	}
	return getFlags
}