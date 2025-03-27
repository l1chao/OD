package main

import (
	"fmt"
)

// 找到二元组个数

func Ans6() {
	var len1, len2 int
	var list1, list2 []int

	_, err := fmt.Scanf("%d\n", &len1)
	if err != nil {
		return
	}
	list1 = make([]int, 0, len1)
	for range len1 {
		var value int
		_, err := fmt.Scanf("%d", &value)
		if err != nil {
			return
		}
		list1 = append(list1, value)
	}
	fmt.Scanln()

	_, err = fmt.Scanf("%d\n", &len2)
	if err != nil {
		return
	}

	list2 = make([]int, 0, len2)
	for range len2 {
		var value int
		_, err := fmt.Scanf("%d", &value)
		if err != nil {
			return
		}
		list2 = append(list2, value)
	}
	fmt.Scanln()

	map1 := make(map[int]int)
	map2 := make(map[int]int)

	for _, v := range list1 {
		map1[v]++
	}
	for _, v := range list2 {
		map2[v]++
	}

	total := 0
	for k, v := range map1 {
		if value, ok := map2[k]; ok {
			total += v * value
		}
	}
	fmt.Println(total)
}
