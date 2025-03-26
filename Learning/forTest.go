package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortS(ss []string) []string {
	sort.Slice(ss, func(i, j int) bool {
		return ss[i][0] < ss[j][0]
	})
	return ss
}

func main() {
	n := 0
	fmt.Scanln(&n)

	slices := []string{}
	var s string
	_, err := fmt.Scanln(&s)
	if err != nil {
		return
	}
	list := strings.Split(s, " ")
	slices = append(slices, list...)

	slices = sortS(slices)
	fmt.Println(strings.Join(slices, " "))
}
