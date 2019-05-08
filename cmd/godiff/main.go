package main

import (
	"fmt"
	"os"

	"github.com/rvillablanca/godiff/internal/pkg/diff"

	"github.com/pkg/errors"

	"github.com/urfave/cli"
)

const (
	version = "1.0"
)

func main() {

	app := cli.NewApp()
	app.Name = "godiff"
	app.Usage = "it allows you to make a patch by comparing two directories"
	app.UsageText = "godiff <old-dir> <new-dir> <dest-dir>"
	app.Author = "Rodrigo Villablanca"
	app.Description = "Make your patch better"
	app.Version = version
	app.Action = func(c *cli.Context) error {

		if len(c.Args()) != 3 {
			return errors.New("invalid number of arguments")
		}

		oldDir := c.Args().Get(0)
		newDir := c.Args().Get(1)
		destDir := c.Args().Get(2)

		return diff.Patch(oldDir, newDir, destDir)
	}

	err := app.Run(os.Args)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
	}
}
