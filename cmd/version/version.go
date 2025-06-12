package version

import (
	"fmt"
	"github.com/spf13/cobra"
)

type Version struct{}

const version = "0.0.1"

func (v Version) Init(root *cobra.Command) {
	root.AddCommand(v.cmd())
}

func (v Version) cmd() *cobra.Command {
	version := &cobra.Command{
		Use:   "version",
		Short: "Version information",
		Args:  cobra.ExactArgs(0),
		Run:   v.run,
	}
	return version
}

func (Version) run(_ *cobra.Command, _ []string) {
	fmt.Printf("%s\n", version)
}
