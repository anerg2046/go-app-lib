package str

import (
	"bytes"
	"crypto/rand"
	"math/big"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

func Random(len int) string {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func VerifyMobileFormat(mobileNum string) bool {
	regular := `^1[3|4|5|6|7|8|9]\d{9}$`

	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

// 处理非可见字符
func Trim(content string) (str string) {
	defer func() {
		if err := recover(); err != nil {
			str = ""
		}
	}()
	newByte := make([]byte, 0)
	if len(content) > 0 {
		for _, b := range []byte(content) {
			if b < 32 { //小于32的字符都可以以这样的方式处理，本次只处理0
				newByte = append(newByte, 32)
			} else {
				newByte = append(newByte, b)
			}
		}
	}
	str = strings.TrimSpace(string(newByte))
	return
}

// 替换括号为中文括号
func CnBrackets(content string) (rsp string) {
	rsp = Trim(content)
	rsp = strings.ReplaceAll(rsp, "(", "（")
	rsp = strings.ReplaceAll(rsp, ")", "）")
	return rsp
}

// 替换括号为英文括号
func EnBrackets(content string) (rsp string) {
	rsp = Trim(content)
	rsp = strings.ReplaceAll(rsp, "（", "(")
	rsp = strings.ReplaceAll(rsp, "）", ")")
	return rsp
}

// 隐藏字符串中间的内容
func HiddenMiddle(content string, bothLen int) string {
	length := utf8.RuneCountInString(content)
	if length <= bothLen*2 {
		return content
	}
	contentRune := []rune(content)
	start := contentRune[0:bothLen]
	end := contentRune[length-bothLen:]
	stars := strings.Repeat("*", length-bothLen*2)
	return string(start) + stars + string(end)
}

// 定义一个映射表，存储英文符号和对应的中文符号
// var punctuationMap = map[rune]rune{
// 	'(':  '（',
// 	')':  '）',
// 	'[':  '【',
// 	']':  '】',
// 	'{':  '｛',
// 	'}':  '｝',
// 	'<':  '《',
// 	'>':  '》',
// 	'"':  '“',
// 	'\'': '‘',
// 	'?':  '？',
// 	'!':  '！',
// 	';':  '；',
// 	':':  '：',
// }

// 文本清理
//
//	如果英文符号前后都是中文，则替换为中文符号
//	并移除前后和中间的多余空字符串
func ClearChineseText(text string, punctuationMap map[rune]rune) string {
	var space = regexp.MustCompile(`\s+`) // 定义一个正则表达式，匹配连续的空格

	text = strings.TrimSpace(text)           // 移除前后的空字符串
	text = space.ReplaceAllString(text, " ") // 替换中间的多余空字符串为一个空格
	text = strings.ReplaceAll(text, "(", "（")
	text = strings.ReplaceAll(text, ")", "）")

	return text

	// var builder strings.Builder // 使用 strings.Builder 来提高字符串拼接的效率
	// var content []rune
	// for _, t := range text {
	// 	content = append(content, t)
	// }
	// count := len(content)
	// for idx, r := range content {
	// 	if unicode.IsPunct(r) { // 判断是否是标点符号
	// 		var (
	// 			prev rune
	// 			next rune
	// 		)
	// 		if idx > 1 && idx < count-1 {
	// 			prev = content[idx-1]
	// 			next = content[idx+1]
	// 		}
	// 		if unicode.Is(unicode.Han, prev) || unicode.Is(unicode.Han, next) { // 如果符号前或者后是中文，就替换
	// 			if c, ok := punctuationMap[r]; ok { // 如果映射表里有对应的中文符号，就替换
	// 				builder.WriteRune(c)
	// 			} else { // 否则保留原来的符号
	// 				builder.WriteRune(r)
	// 			}
	// 		} else { // 否则保留原来的符号
	// 			builder.WriteRune(r)
	// 		}
	// 	} else { // 不是标点符号就不变
	// 		builder.WriteRune(r)
	// 	}
	// }
	// return builder.String()

}

// 移除字符串中不可见字符
func RemoveNonPrintable(s string) string {
	var result []rune
	for _, r := range s {
		if unicode.IsPrint(r) {
			result = append(result, r)
		}
	}
	return strings.TrimSpace(string(result))
}

// 清理人名
func ClearPersonName(s string) string {
	// 清理主数据持有人名称，去掉身份证号之类的东西
	reg, _ := regexp.Compile(`[（|\d]\d+[\d|X|）]$`)
	if reg.MatchString(s) {
		s = reg.ReplaceAllString(s, "")
	}

	// 清理连续星号结尾的
	reg, _ = regexp.Compile(`\*+$`)
	if reg.MatchString(s) {
		s = reg.ReplaceAllString(s, "")
	}

	// 清理中文括号中间都是星号的
	reg, _ = regexp.Compile(`（\*+）$`)
	if reg.MatchString(s) {
		s = reg.ReplaceAllString(s, "")
	}

	// 所有人名中间的特殊点替换为英文点
	reg, _ = regexp.Compile(`[•|·|・|▪|●]`)
	if reg.MatchString(s) {
		s = reg.ReplaceAllString(s, ".")
	}

	return strings.TrimSpace(s)
}

// 是否包含中文
func HasHan(str string) bool {
	for _, s := range str {
		if unicode.Is(unicode.Han, s) {
			return true
		}
	}
	return false
}
