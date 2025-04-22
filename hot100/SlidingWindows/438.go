package main

// 输入: s = "cbaebabacd", p = "abc"
// 输出: [0,6]

func findAnagrams(s string, p string) []int {
	cnt := [26]int{}

	cntp := [26]int{}
	for _, v := range p {
		cntp[v-'a']++
	}

	ans := []int{}
	for right, v := range s {
		cnt[v-'a']++
		if right < len(p)-1 {
			continue
		}

		if cnt == cntp {
			ans = append(ans, right-len(p)+1)
		}

		cnt[s[right-len(p)+1]-'a']--
	}
	return ans
}
