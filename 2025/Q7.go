package main

import (
	"fmt"
	"strings"
)

// 总结：对于od里面的题目，很多都是描述看起来要求莫名其妙的，但是代码实际上很友善，有一种“命题者先写出来代码，再丑陋的用自然语言倒着拼出一个题目描述来”，
// 所以：
// 1.解题的时候，要沉稳慢慢读，理清楚题目里面说了什么
// 2.坚信解决的代码很简单。
// 当然，所有基本的知识还是大概要过一遍。
func Ans7() {
	n := 0
	_, err := fmt.Scanln(&n)
	if err != nil {
		return
	}
	operations := make([]string, n)

	result := 0
	is_ordered := true
	queue_size := 0

	for _, v := range operations {
		if strings.HasPrefix(v, "head") {
			if queue_size != 0 {
				is_ordered = false
			}
			queue_size++
		} else if strings.HasPrefix(v, "tail") {
			queue_size++
		} else { // remove
			if queue_size == 0 {
				continue
			}
			if !is_ordered {
				is_ordered = true
				result++
			}
			queue_size--
		}
	}
}
