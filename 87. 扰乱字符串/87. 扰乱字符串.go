package main

import "fmt"

var m map[string]bool

func main() {
	s1 := "abc"
	s2 := "bca"

	b := isScrambleWithDP(s1, s2)
	fmt.Println(b)
}

func isScramble(s1 string, s2 string) bool {
	m = make(map[string]bool)
	return isScrambleWithMap(s1, s2)
}
func isScrambleWithMap(s1 string, s2 string) bool {
	k := s1 + "-" + s2
	i, ok := m[k]
	if ok {
		if !i {
			return false
		}
	}
	olen := len(s1)
	tlen := len(s2)
	if olen != tlen {
		key := s1 + "-" + s2
		m[key] = false
		return false
	}
	if s1 == s2 {
		key := s1 + "-" + s2
		m[key] = true
		return true
	}
	if !alphabetCount(s1, s2) {
		key := s1 + "-" + s2
		m[key] = false
		return false
	}
	for i := 1; i < olen; i++ {
		if isScramble(s1[0:i], s2[0:i]) && isScramble(s1[i:], s2[i:]) {
			key := s1 + "-" + s2
			m[key] = true
			return true
		}
		if isScramble(s1[olen-i:], s2[0:i]) && isScramble(s1[0:olen-i], s2[i:]) {
			key := s1 + "-" + s2
			m[key] = true
			return true
		}
	}
	key := s1 + "-" + s2
	m[key] = false
	return false
}

func isScrambleWithDP(s1 string, s2 string) bool {
	olen := len(s1)
	tlen := len(s2)
	if olen != tlen {
		return false
	}
	if s1 == s2 {
		return true
	}
	if !alphabetCount(s1, s2) {
		return false
	}
	res := make([][][]bool, olen+1)
	for i := range res {
		res[i] = make([][]bool, olen)
		for j := range res[i] {
			res[i][j] = make([]bool, olen)
		}
	}

	for le := 1; le <= olen; le++ {
		for i := 0; i+le <= olen; i++ {
			for j := 0; j+le <= olen; j++ {
				if le == 1 {
					res[le][i][j] = s1[i] == s2[j]
				} else {
					for q := 1; q < le; q++ {
						res[le][i][j] = (res[q][i][j] && res[le-q][i+q][j+q]) || (res[q][i][j+le-q] && res[le-q][i+q][j])
						if res[le][i][j] {
							break
						}
					}

				}

			}
		}
	}
	return res[olen][0][0]
}

func alphabetCount(s1 string, s2 string) bool {
	m := make(map[string]int)
	for _, v := range s1 {
		value := string(v)
		i, ok := m[value]
		if ok {
			m[value] = i + 1
		} else {
			m[value] = 1
		}
	}
	for _, v := range s2 {
		value := string(v)
		i, ok := m[value]
		if !ok {
			return false
		}
		if i > 0 {
			m[value] = i - 1
		} else {
			return false
		}
	}
	return true
}
