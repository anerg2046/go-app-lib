package str

import (
	"bytes"
	"crypto/rand"
	"math/big"
	"regexp"
	"strings"
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
