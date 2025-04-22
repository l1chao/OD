package main

// 输入: nums = [0,1,0,3,12]
// 输出: [1,3,12,0,0]
func moveZeroes(nums []int) {
	k := 0 // 0的个数
	for i, v := range nums {
		if v == 0 {
			k++
		} else {
			nums[i-k] = nums[i]
		}
	}
	for i := range k {
		nums[len(nums)-1-i] = 0
	}
}

// 双指针
func moveZeroes1(nums []int) {
	j := 0 //慢指针：1.在外面 2.是关键 3.下一个可插入的位置
	for i := range nums {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}
	for ; j < len(nums); j++ {
		nums[j] = 0
	}
}

func moveZeroes1_(nums []int) {
	j := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}
	for ; j < len(nums); j++ {
		nums[j] = 0
	}
}
