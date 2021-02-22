package command

import "github.com/spf13/cobra"

func NewVersionCmd(version string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "display version",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println(version)
		},
	}
	return cmd
}
