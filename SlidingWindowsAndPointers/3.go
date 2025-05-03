package main

// 不定长滑动窗口，快指针是循环变量，一直变化，慢指针看情况变化。

// 输入: s = "abcabcbb"
// 输出: 3

func lengthOfLongestSubstring(s string) int {
	cnt := [128]int{} // 128是ascii码的总数，这里的string只限于ascii码。
	left := 0         // 慢指针，因为不是每次循环变量都要改变，所以要写在外面
	maxNum := 0
	for right, c := range s {
		cnt[c]++ // 快指针每一次都要变化
		for cnt[c] > 1 {
			cnt[s[left]]--
			left++ // 慢指针看情况才变化
		}
		maxNum = max(maxNum, right-left+1) // 内部的快指针的结果通过maxNum保存到了循环的外部。
	}
	return maxNum
}

// 输入: s = "abcabcbb"
// 输出: 3
// 快慢针+hashmap记录过程，过程中时时用maxNum保存最大值。
// 快针按照循环变量的方式移动，慢针根据条件移动。
func lengthOfLongestSubstring1(s string) int {
	left := 0
	cnt := [128]int{}
	maxNum := 0

	for rigtht, v := range s {
		cnt[v]++
		for cnt[v] > 1 {
			cnt[s[left]]--
			left++
		}
		maxNum = max(maxNum, rigtht-left+1)
	}

	return left
}

// 1. 左右针夹map 2.快针是循环变量，慢针根据条件移动 3.右进一个就比较一次最大值

// 输入: s = "abcabcbb"
// 输出: 3
func lengthOfLongestSubstring_(s string) int {
	left := 0
	cnt := [128]int{}
	maxNum := 0
	for right, v := range s {
		cnt[v]++
		for cnt[v] > 1 {
			cnt[s[left]]--
			left++
		}
		maxNum = max(maxNum, right-left+1)
	}
	return maxNum
}
