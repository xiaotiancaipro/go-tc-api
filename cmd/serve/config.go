package serve

import "github.com/spf13/cobra"

type Config struct{}

func (c Config) Init(serve *cobra.Command) {
	serve.Flags().StringP("config", "c", "./config.toml", "Configuration file path")
	serve.AddCommand(c.cmd())
}

func (c Config) cmd() *cobra.Command {
	config := &cobra.Command{
		Use:   "config",
		Short: "Configuration file path",
	}
	return config
}
