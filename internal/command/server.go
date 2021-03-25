package command

import (
	"io"
	"log"
	"os"

	"github.com/cappyzawa/terraform-registry/internal/server"
	"github.com/spf13/cobra"
)

var (
	port       string
	configFile string
	pidFile    string
	logFile    string
)

func NewServerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "run terraform registry locally",
		RunE: func(cmd *cobra.Command, args []string) error {
			overWriteByEnv()

			var file *os.File
			if logFile != "" {
				f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
				if err != nil {
					return err
				}
				file = f
			}
			multiLogfile := io.MultiWriter(os.Stderr, file)
			logger := log.New(multiLogfile, "", 0)
			logger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
			so := &server.Opt{
				Port:       port,
				ConfigPATH: configFile,
				PIDPATH:    pidFile,
				Logger:     logger,
			}
			return server.Run(so)
		},
	}

	cmd.PersistentFlags().StringVarP(&port, "port", "p", "8080", "running port")
	cmd.PersistentFlags().StringVarP(&configFile, "config", "c", "config.yaml", "specify config file for registry")
	cmd.PersistentFlags().StringVarP(&pidFile, "pid-file", "", "", "write PID of registy to specified file")
	cmd.PersistentFlags().StringVarP(&logFile, "log-file", "", "", "the file to which the log is output")
	return cmd
}

func overWriteByEnv() {
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	if os.Getenv("CONFIG_FILE") != "" {
		configFile = os.Getenv("CONFIG_FILE")
	}
	if os.Getenv("PID_FILE") != "" {
		pidFile = os.Getenv("PID_FILE")
	}
	if os.Getenv("LOG_FILE") != "" {
		logFile = os.Getenv("LOG_FILE")
	}
}
