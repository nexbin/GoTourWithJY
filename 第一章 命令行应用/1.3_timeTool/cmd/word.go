package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"timeTool/internal/word"
)

/*
	添加word子命令
*/

const (
	ModeUpper                      = iota + 1 // 全部转大写
	ModeLower                                 // 全部转小写
	ModeUnderscoreToUpperCamelCase            // 下划线转大写驼峰
	ModeUnderscoreToLowerCamelCase            // 下线线转小写驼峰
	ModeCamelCaseToUnderscore                 // 驼峰转下划线
)

var (
	str  string
	mode int8

	WordCmd = cobra.Command{
		Use:   `w`,
		Short: `word transform`,
		Long:  `print 'hello, $yourName'`,
		Run: func(cmd *cobra.Command, args []string) {
			var content string
			switch mode {
			case ModeUpper:
				content = word.ToUpper(str)
			case ModeLower:
				content = word.ToLower(str)
			case ModeUnderscoreToUpperCamelCase:
				content = word.UnderscoreToUpperCamelCase(str)
			case ModeUnderscoreToLowerCamelCase:
				content = word.UnderscoreToLowerCamelCase(str)
			case ModeCamelCaseToUnderscore:
				content = word.CamelCaseToUnderscore(str)
			default:
				log.Fatalf("暂不支持该转换模式，请执行 help word 查看帮助文档")
			}
			log.Printf("输出结果: %s", content)
		},
	}
)

func init() {
	WordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	WordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入转换模式")
}
