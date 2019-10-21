package main

import "fmt"

func main() {
	s := "abcd"

	wordDict := []string{"a", "abc", "b", "cd"}
	b := wordBreaks(s, wordDict)
	//c := maximalRectangle(a)
	fmt.Println(b)
}

func wordBreak(s string, wordDict []string) bool {
	slen := len(s)
	res := make([][]bool, slen)
	for i := range res {
		res[i] = make([]bool, slen)
	}
	for _, k := range wordDict {
		klen := len(k)
		if klen > slen {
			continue
		}
		for j := 0; j < slen-klen+1; j++ {
			if s[j:j+klen] == k {
				res[j+klen-1][j] = true
			}
		}
		ok := checkOk(res, slen-1)
		if ok {
			return true
		}
	}

	return false
}
func checkOk(res [][]bool, i int) bool {
	for j, k := range res[i] {
		if k == true {
			if j == 0 {
				return true
			}
			ok := checkOk(res, j-1)
			if ok {
				return true
			}

		}
	}
	return false
}
func wordBreaks(s string, wordDict []string) bool {
	slen := len(s)
	res := []int{0}
	for {
		temp := []int{}
		hasWord := false
		hasAlphabet := false
		for _, key := range res {

			for _, dict := range wordDict {
				dlen := len(dict)
				end := key + dlen
				if end > slen {
					continue
				}
				hasAlphabet = hasAlphabet || checkAlphabet(s[key:end], dict)
				if s[key:end] == dict {
					if end == slen {
						return true
					}
					hasWord = true
					temp = append(temp, key+dlen)
				}
			}

		}
		if !hasAlphabet {
			return false
		}
		if !hasWord {
			return false
		}
		res = temp
	}
}
func checkAlphabet(s string, dict string) bool {
	temp := make(map[rune]bool)
	for _, k := range s {
		if _, ok := temp[k]; !ok {
			temp[k] = true
		}
	}
	for _, d := range dict {
		if _, ok := temp[d]; !ok {
			return false
		}
	}
	return true
}

func wordBreakDP(s string, wordDict []string) bool {
	slen := len(s)
	dp := make([]bool, slen+1)
	dp[0] = true
	wd := sliceToMap(wordDict)
	for i := 1; i <= slen; i++ {
		for j := 0; j < i; j++ {
			_, ok := wd[s[j:i]]
			if dp[j] && ok {
				dp[i] = true
			}
		}
	}
	return dp[slen]
}

func sliceToMap(w []string) map[string]bool {
	m := make(map[string]bool)
	for _, i := range w {
		m[i] = true
	}
	return m
}
