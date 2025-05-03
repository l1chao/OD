package main

func maxVowels(s string, k int) int {
	cnt := 0
	maxNum := 0
	for i, v := range s {
		// 增
		if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' {
			cnt++
		}
		if i < k-1 {
			continue
		}

		// 检
		maxNum = max(maxNum, cnt)

		// 删
		out := s[i-k+1]
		if out == 'a' || out == 'e' || out == 'i' || out == 'o' || out == 'u' {
			cnt++
		}
	}
	return maxNum
}

func maxVowels1(s string, k int) int {
	maxN := 0
	stat := 0
	j := 0
	for i, v := range s {
		if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' {
			stat++
		}
		if i < k-1 {
			continue
		}
		maxN = max(maxN, stat)

		if s[j] == 'a' || s[j] == 'e' || s[j] == 'i' || s[j] == 'o' || s[j] == 'u' {
			stat--
			j++
		}
	}
	return maxN
}
