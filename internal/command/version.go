package command

import "github.com/spf13/cobra"

// NewVersionCmd initializes version (sub) command
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
