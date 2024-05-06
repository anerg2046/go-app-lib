package str

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
)

func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func HmacSha1(key, message string) string {
	// 将密钥转换为字节数组
	keyBytes := []byte(key)
	// 将消息转换为字节数组
	messageBytes := []byte(message)

	// 创建一个新的 HMAC 实例，使用 SHA1 散列算法和密钥
	h := hmac.New(sha1.New, keyBytes)
	// 写入消息到 HMAC 实例中
	h.Write(messageBytes)
	// 计算 HMAC 值
	hash := h.Sum(nil)

	// 返回base64
	return base64.StdEncoding.EncodeToString(hash)
}

func Sha1(message string) string {
	// 将消息转换为字节数组
	messageBytes := []byte(message)

	// 使用 SHA-1 哈希算法计算消息的哈希值
	hash := sha1.Sum(messageBytes)

	// 返回十六进制编码的哈希值
	return hex.EncodeToString(hash[:])
}

func Base64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}
