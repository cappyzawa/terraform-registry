package command

import (
	"io"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Opt has options for command
type Opt struct {
	Name    string
	Args    []string
	Version string
	Out     io.Writer
	Err     io.Writer
}

func NewBaseCmd(opt *Opt) *cobra.Command {
	cmd := &cobra.Command{
		Use:           opt.Name,
		Short:         "building terraform-registry in local",
		SilenceErrors: true,
		SilenceUsage:  true,
	}
	cmd.AddCommand(
		NewVersionCmd(opt.Version),
		NewServerCmd(),
	)
	cobra.OnInitialize(initConfig)
	viper.BindPFlags(cmd.PersistentFlags())
	cmd.SetOut(opt.Out)
	cmd.SetErr(opt.Err)
	return cmd
}

func initConfig() {
	viper.AutomaticEnv()
}

func Execute(opt *Opt) int {
	cmd := NewBaseCmd(opt)
	if err := cmd.Execute(); err != nil {
		cmd.PrintErr(err)
		return 1
	}
	return 0
}
