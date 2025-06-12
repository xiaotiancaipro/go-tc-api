package cmd

import (
	"github.com/spf13/cobra"
	"go-tc-api/cmd/serve"
	"go-tc-api/cmd/version"
)

type Root struct{}

func (r Root) Init() *cobra.Command {
	return r.cmd()
}

func (r Root) cmd() *cobra.Command {
	root := &cobra.Command{
		Use:   "go-tc-api",
		Short: "Go TC API Server runner",
		Run:   r.run,
	}
	version.Version{}.Init(root)
	serve.Serve{}.Init(root)
	return root
}

func (Root) run(cmd *cobra.Command, _ []string) {
	cmd.Help()
}
