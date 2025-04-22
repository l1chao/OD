package main

// 输入：[1,8,6,2,5,4,8,3,7]
// 输出：49
// 1.两端点指针往中间收拢。 2.收拢的时候，长端点收拢没有用，短端点才能够让结果更好。
func maxArea(height []int) int {
	i, j := 0, len(height)-1
	ans := 0
	for i < j {
		ans = max(ans, min(height[i], height[j])*(j-i))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return ans
}
