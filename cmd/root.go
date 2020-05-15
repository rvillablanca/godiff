package cmd

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/rvillablanca/godiff/internal/diff"
	"github.com/spf13/cobra"
)

const (
	version = "1.1.1"
)

var rootCmd = &cobra.Command{
	Use:           "godiff",
	Short:         "godiff <old-dir> <new-dir> <dest-dir>",
	SilenceErrors: true,
	SilenceUsage:  true,
	Version:       version,
	RunE:          run,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func run(_ *cobra.Command, args []string) error {
	if len(args) != 3 {
		return errors.New("invalid number of arguments")
	}
	return diff.Patch(args[0], args[1], args[2])
}
