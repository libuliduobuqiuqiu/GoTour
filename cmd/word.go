package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"strings"
	"unicode"
)

const (
	ModeUpper = iota + 1
	ModeLower
	ModeUnderscoreToUpperCameClass
	ModeUnderscoreToLowerCameClass
	ModeCamelCaseToUnderscore
)

var desc = strings.Join([]string{
	"该子命令支持各种单词格式转换，模式如下：",
	"1：全部转大写",
	"2：全部转小写",
	"3：下划线转大写驼峰",
	"4：下划线转小写驼峰",
	"5：驼峰转下划线",
}, "\n")

var mode int8
var str string

var worCmd = &cobra.Command{
	Use:   "word",
	Short: "单词转换器",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = ToUpper(str)
		case ModeLower:
			content = ToLower(str)
		case ModeUnderscoreToUpperCameClass:
			content = UnderscoreToUpperCameClass(str)
		case ModeUnderscoreToLowerCameClass:
			content = UnderscoreToLowerCameClass(str)
		case ModeCamelCaseToUnderscore:
			content = CameClassToUnderscore(str)
		default:
			log.Fatalf("暂不支持该转换模式，请执行help word查看帮助文档")
		}
		log.Printf("输出结果： %s\n", content)
	},
}

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

func UnderscoreToUpperCameClass(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	s = strings.Replace(s, " ", "", -1)
	return s
}

func UnderscoreToLowerCameClass(s string) string {
	s = UnderscoreToUpperCameClass(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

func CameClassToUnderscore(s string) string {
	var output []rune
	for i, r := range s {

		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}

		if unicode.IsUpper(r) {
			output = append(output, '_')
		}

		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}

func init() {
	worCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	worCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换模式")
}
