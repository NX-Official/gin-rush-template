package tools

import (
	"math/rand"
	"strings"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// RandString 生成一个随机字符串
func RandString(length int) string {
	var sb strings.Builder
	sb.Grow(length)

	// 随机数种子
	for i := 0; i < length; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}

	return sb.String()
}
