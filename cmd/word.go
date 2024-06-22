package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"strings"
	"tour/internal/word"
)

const (
	ModeUpper                      = iota + 1 // 转换为大写
	ModeLower                                 // 转换为小写
	ModeUnderscoreToUpperCamelCase            // 下划线转大驼峰
	ModeUnderscoreToLowerCamelCase            // 下划线转小驼峰
	ModeCamelCaseToUnderscore                 // 驼峰转下划线
)

var (
	mode int8
	str  string
	desc = strings.Join([]string{
		"该子命令支持各种单词格式转换， 模式如下:",
		"1: 全部转为大写",
		"2: 全部转为小写",
		"3: 下划线转大驼峰",
		"4: 下划线转小驼峰",
		"5: 驼峰转下划线",
	}, "\n")
)

var wordCmd = &cobra.Command{
	Use:   "word",   // 子命令标识
	Short: "单词格式转换", // 简短说明
	Long:  desc,     // 完整说明
	Run: func(cmd *cobra.Command, args []string) {
		var content string = str
		switch mode {
		case ModeUpper:
			content = word.ToUpper(content)
		case ModeLower:
			content = word.ToLower(content)
		case ModeUnderscoreToUpperCamelCase:
			content = word.UnderscoreToUpperCameCase(content)
		case ModeUnderscoreToLowerCamelCase:
			content = word.UnderscoreToLowerCameCase(content)
		case ModeCamelCaseToUnderscore:
			content = word.CamelCaseToUnderscore(content)
		default:
			log.Fatalf("暂不支持该转换模式， 请执行 help word 查看帮助文档")
		}
		log.Printf("输出结果: %s", content)
	},
}

func init() {
	// VarP, 第一个为绑定变量， 第二个为完整命令参数， 第三个为短标识， 第四个是默认值， 第五个为使用说明
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换的模式")
}

/*
go run . word -s=fwf -m=1
go run . word -s=FwF -m=2
go run . word -s=fff -m=3
go run . word -s=fw_f -m=4
go run . word -s=Fwf -m=5
*/
