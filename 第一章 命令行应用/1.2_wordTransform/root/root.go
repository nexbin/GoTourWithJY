package root

import (
	"github.com/spf13/cobra"
	"wordTransform/cmd"
)

var rootCmd = cobra.Command{}

func init() {
	rootCmd.AddCommand(&cmd.WordCmd)
}

func Execute() {
	rootCmd.Execute()
}
