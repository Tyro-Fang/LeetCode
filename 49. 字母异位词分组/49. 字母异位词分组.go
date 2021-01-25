package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	wordDict := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	b := groupAnagrams(wordDict)
	fmt.Println(b)
}
func groupAnagrams(strs []string) [][]string {
	dict := make(map[string][]string)
	for _, v := range strs {
		//fmt.Println(SortString(v))
		i, ok := dict[SortString(v)]
		if ok {
			dict[SortString(v)] = append(i, v)
		} else {
			dict[SortString(v)] = []string{v}
		}
	}
	dlen := len(dict)
	res := make([][]string, dlen)
	i := 0
	for _, v := range dict {
		res[i] = v
		i++
	}
	return res
}
func SortString(s string) string {
	t := strings.Split(s, "")
	sort.Strings(t)
	return strings.Join(t, "")
}
