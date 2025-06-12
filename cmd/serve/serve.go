package serve

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-tc-api/internal"
	"os"
)

type Serve struct{}

func (s Serve) Init(root *cobra.Command) {
	root.AddCommand(s.cmd())
}

func (s Serve) cmd() *cobra.Command {
	serve := &cobra.Command{
		Use:   "serve",
		Short: "Start Go TC API server",
		Args:  cobra.ExactArgs(0),
		Run:   s.run,
	}
	Config{}.Init(serve)
	return serve
}

func (Serve) run(cmd *cobra.Command, _ []string) {
	configFile, _ := cmd.Flags().GetString("config")
	if err := internal.GoTCAPIServer(configFile); err != nil {
		fmt.Printf("Server start failed: %v\n", err)
		os.Exit(1)
	}
}
