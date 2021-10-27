package root

import (
	"github.com/spf13/cobra"
	"timeTool/cmd"
)

var rootCmd = cobra.Command{}

func init() {
	// 单词处理
	rootCmd.AddCommand(&cmd.WordCmd)
	// 时间处理
	rootCmd.AddCommand(&cmd.NowTimeCmd)
	rootCmd.AddCommand(&cmd.TimerCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
