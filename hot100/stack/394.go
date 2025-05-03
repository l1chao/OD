package main

import (
	"strconv"
	"strings"
)

// 也就是说，嵌套括号问题，都是考虑是否内层嵌套可以被“消除”（如例子"3[a2[c]]"里面的2[c]，c左右的括号是最内层的，而且可以被处理后，再用同样的处理方法处理外层字符串）

// 对于具体代码，要点如下：
// 1. 遍历字符串
// 2. 分类讨论，遇到哪个入栈，遇到哪个后出栈
// 3. （可能）将出栈后的内容逆序输出。
func decodeString(s string) string {
	ptr := 0
	stack := []string{}
	for ptr < len(s) {
		if s[ptr] >= '0' && s[ptr] <= '9' {
			digits := getDigits(s, &ptr)
			stack = append(stack, digits)
		} else if s[ptr] >= 'a' && s[ptr] <= 'z' {
			letters := getLetters(s, &ptr)
			stack = append(stack, letters)
		} else if s[ptr] == '[' {
			stack = append(stack, string(s[ptr]))
			ptr++
		} else { // cur == ']'的情况
			ptr++
			i := len(stack) - 1
			for ; stack[i] != "["; i-- {
			}
			letters := strings.Builder{}
			for j := i + 1; j < len(stack); j++ {
				letters.Write([]byte(stack[j]))
			}
			stack = stack[:i]

			reaptTime, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-1]

			stack = append(stack, strings.Repeat(letters.String(), reaptTime))
		}
	}
	ans := strings.Builder{}
	for _, v := range stack {
		ans.Write([]byte(v))
	}
	return ans.String()
}

func getDigits(s string, ptr *int) string { // 如果有一位是数字，那么获取所有的数字
	numsBuilder := strings.Builder{}
	for ; s[*ptr] >= '0' && s[*ptr] <= '9'; (*ptr)++ {
		numsBuilder.WriteByte(s[*ptr])
	}
	return numsBuilder.String()
}

func getLetters(s string, ptr *int) string {
	letterBuilder := strings.Builder{}
	for ; *ptr <= len(s)-1 && s[*ptr] >= 'a' && s[*ptr] <= 'z'; *ptr++ {
		letterBuilder.WriteByte(s[*ptr])
	}
	return letterBuilder.String()
}

// 有哪些循环方式？
func temp() {
	s := []int{}

	// c循环
	i := 0
	for ; i < len(s); i++ {

	}
	for j := 0; j < len(s); j++ {
	}

	// 单循环
	for i < len(s) {

	}

	// range 循环
	for i, v := range s {

	}

	// range的次数循环（仅用于控制循环次数）
	for i := range 10 {

	}

}

func temp1(s string, ptr *int) string {
	// 有指针肯定是用c循环
	//从ptr开始获取连续的数字并返回
	nums := strings.Builder{}
	for ; s[*ptr] >= '0' && s[*ptr] <= '9'; *ptr++ {
		nums.WriteByte(s[*ptr])
	}
	return nums.String()
}

func temp2(s string, ptr *int) string {
	// 一直收集直到[
	ans := strings.Builder{}
	for ; s[*ptr] != '['; *ptr++ {

	}
}
