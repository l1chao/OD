package main

// 输入：nums = [7,4,3,9,1,8,5,2,6], k = 3

// 输出：[-1,-1,-1,5,4,4,-1,-1,-1]

func getAverages(nums []int, k int) []int {
	cnt := 0
	length := 2*k + 1
	tmpArr := make([]int, 0)

	for i, v := range nums {
		cnt += v
		if i < length-1 {
			continue
		}
		tmpArr = append(tmpArr, cnt/length)
		cnt -= nums[i-length+1]
	}
	ansArr := make([]int, len(nums))
	for i := range len(ansArr) {
		if i < k || i+k > len(nums)-1 {
			ansArr[i] = -1
		} else {
			ansArr[i] = tmpArr[i-k]
		}
	}
	return ansArr
}
