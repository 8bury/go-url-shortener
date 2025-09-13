package base62

import (
	"strings"
)

const base62Digits = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func ConvertFromInt(n int) string {
	encoded := ""
	for n > 0 {
		remainder := n % 62
		encoded = string(base62Digits[remainder]) + encoded
		n = n / 62
	}
	return encoded
}

func ConvertToInt(s string) int {
	decoded := 0
	for _, char := range s {
		decoded = decoded*62 + strings.Index(base62Digits, string(char))
	}
	return decoded
}
