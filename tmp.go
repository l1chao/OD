package main

// 1.1题目背景
// 消消乐，是一款休闲游戏，相同图案相邻成一条直线即可消除。本题中的字符消消乐，与此类似，作了简化。
// 1.2题目描述
// 给定一个字符串，从左往右扫描，如存在两个或两个以上的相同字符靠在一起，则消除这些字符。对每次消除后剩下的字符串，继续应用上述规则，直到不能再消除为止。

// 输入说明：任意字符串。
// 返回说明：消除后的字符串。

// aabbacaccd
// aaa ca d
// cad

// caaabbcd
func Ans1(s string) string {
	stack := []byte{}
	var k byte
	for p := 0; p < len(s); p++ {
		if len(stack) == 0 {
			stack = append(stack, s[p])
		} else if s[p] == stack[len(stack)-1] {
			if k != 0 {
				continue
			}
			stack = stack[:len(stack)-1]
			k = s[p]
		} else if s[p] != stack[len(stack)-1] {
			if k != 0 {
				k = 0
				stack = stack[:len(stack)-1]
				stack = append(stack, s[p])
			} else {
				stack = append(stack, s[p])
			}
		}
	}
	if k != 1 {
		stack = stack[:len(stack)-3]
	}
	return string(stack)
}

// a a a b b c d

// func main() {
// 	fmt.Println(Ans1("caaabbcd")) // d
// }
