package app

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewVersionCmd creates a new cobra.Command printing app version information.
func NewVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print version information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(cmd.OutOrStdout(), "%s\n", Info())
		},
	}
}
