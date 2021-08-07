package utils

import (
	"strings"
)

// VerifyPrefix  验证前缀
func VerifyPrefix(str string, verifyPrefix string) bool {
	return strings.HasPrefix(strings.TrimSpace(str), verifyPrefix)
}
func DefaultSkip(conf string) bool {
	if conf != "default" {
		return false
	}
	return true
}
