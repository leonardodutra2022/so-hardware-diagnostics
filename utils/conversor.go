package utils

import (
	"strconv"
	"strings"
)

func StrToInt(str string) int64 {
	value, _ := strconv.ParseInt(strings.TrimSpace(str), 10, 64)
	return value
}

func IntToStr(inteiro int) string {
	return strconv.Itoa(inteiro)
}
