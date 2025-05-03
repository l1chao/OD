package main

import (
	"strconv"
	"strings"
)

func decodeString1(s string) string {
	stk := []string{}
	ptr := 0
	for ptr < len(s) {
		cur := s[ptr]
		if cur >= '0' && cur <= '9' {
			digits := getDigits(s, &ptr)
			stk = append(stk, digits)
		} else if (cur >= 'a' && cur <= 'z' || cur >= 'A' && cur <= 'Z') || cur == '[' {
			stk = append(stk, string(cur))
			ptr++
		} else { // cur == ']'的情况
			ptr++ // 跳过']'
			sub := []string{}
			for stk[len(stk)-1] != "[" {
				sub = append(sub, stk[len(stk)-1])
				stk = stk[:len(stk)-1]
			}
			for i := 0; i < len(sub)/2; i++ { // 将字符串数组逆序
				sub[i], sub[len(sub)-i-1] = sub[len(sub)-i-1], sub[i]
			}
			stk = stk[:len(stk)-1] // pop '['
			repTime, _ := strconv.Atoi(stk[len(stk)-1])
			stk = stk[:len(stk)-1] // pop 数字
			t := strings.Repeat(getString(sub), repTime)
			stk = append(stk, t)
		}
	}
	return getString(stk)
}

func getDigits_(s string, ptr *int) string {
	ret := ""
	for ; s[*ptr] >= '0' && s[*ptr] <= '9'; *ptr++ {
		ret += string(s[*ptr])
	}
	return ret
}

func getString(v []string) string { // 正序遍历stack
	ret := ""
	for _, s := range v {
		ret += s
	}
	return ret
}
