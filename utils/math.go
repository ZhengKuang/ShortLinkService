package utils

import (
	"strconv"
	"strings"
)

var (
	Tokens string
	Length int = 62
)

func init() {
	for i := 0; i <= 9; i++ {
		Tokens += strconv.Itoa(i)
	}

	for i := 0; i < 26; i++ {
		Tokens += string(byte('a') + byte(i))
	}

	for i := 0; i < 26; i++ {
		Tokens += string(byte('A') + byte(i))
	}
}

func IdToString(id int) string {
	var res string
	for id > 0 {
		d := id % Length
		res = string(Tokens[d]) + res
		id /= Length
	}
	return res
}

func StringToId(str string) int {
	var res = 0
	for _, s := range str {
		value := strings.Index(Tokens, string(s))
		res = res*Length + value

	}
	return res

}
