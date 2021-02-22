package command

import (
	"github.com/cappyzawa/terraform-registry/internal/server"
	"github.com/spf13/cobra"
)

var (
	port       string
	configFile string
	pidFile    string
)

func NewServerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "run terraform registry locally",
		RunE: func(cmd *cobra.Command, args []string) error {
			so := &server.Opt{
				Port:       port,
				ConfigPATH: configFile,
				PIDPATH:    pidFile,
			}
			return server.Run(so)
		},
	}

	cmd.PersistentFlags().StringVarP(&port, "port", "p", "8080", "running port")
	cmd.PersistentFlags().StringVarP(&configFile, "config", "c", "config.yaml", "specify config file for registry")
	cmd.PersistentFlags().StringVarP(&pidFile, "pid-file", "", "", "write PID of registy to specified file")
	return cmd
}