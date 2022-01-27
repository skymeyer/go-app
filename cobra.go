package app

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewVersionCmd creates a new cobra.Command printing app version information.
func NewVersionCmd() *cobra.Command {
	var short bool
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print version information",
		Run: func(cmd *cobra.Command, args []string) {
			var out = Info()
			if short {
				out = Version
			}
			fmt.Fprintf(cmd.OutOrStdout(), "%s\n", out)
		},
	}

	cmd.Flags().BoolVar(&short, "short", short, "If true, print just the version number.")

	return cmd
}
