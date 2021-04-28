package command

import (
	"os"

	"github.com/and-rad/ocis-books/pkg/version"
	"github.com/micro/cli/v2"
)

func Execute() error {
	app := &cli.App{
		Name:     "books",
		Version:  version.String,
		Usage:    "Books, an oCIS extension for managing your e-book library",
		Compiled: version.Compiled(),
		Authors: []*cli.Author{
			{
				Name:  "Andreas Raddau",
				Email: "and.rad@posteo.de",
			},
		},
	}

	return app.Run(os.Args)
}
