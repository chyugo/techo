package utils

import (
	"fmt"
	"regexp"
	"strings"
)

// 纯数字
func IsMatchNumber(s string) bool {
	fmt.Println("s", s)
	reg := regexp.MustCompile(`^[0-9]*$`)
	return reg.MatchString(s)
}

// 含有非法字符
func IsHaveIllegalCharacters(sList []string) bool {
	have := false
	for i := 0; i < len(sList); i++ {
		var temp string
		fmt.Println("sList[i]", sList[i])
		temp = strings.ToLower(sList[i])
		fmt.Println("sList[i]", sList[i])
		reg := regexp.MustCompile(`\b(and|exec|insert|select|drop|grant|alter|delete|update|count|chr|mid|master|truncate|char|declare|or)\b|(\*|;|\+|'|%)`)
		if reg.MatchString(temp) {
			have = true
		}
	}
	return have

}
