package main

import "fmt"

func main() {
	a := "aa"
	p := "a*"
	b := isMatchFront(a, p)
	fmt.Println(b)
}

func isMatch(s string, p string) bool {
	slen := len(s)
	plen := len(p)
	res := make([][]bool, slen+1)
	for i := range res {
		res[i] = make([]bool, plen+1)
	}
	res[slen][plen] = true
	for i := plen - 1; i >= 0; i-- {
		if p[i] == '*' && res[slen][i+1] {
			res[slen][i] = true
		} else {
			res[slen][i] = false
		}
	}
	for i := 0; i < slen; i++ {
		res[i][plen] = false
	}
	for i := slen - 1; i >= 0; i-- {
		for j := plen - 1; j >= 0; j-- {
			if s[i] == p[j] || p[j] == '?' {
				res[i][j] = res[i+1][j+1]
				continue
			}
			if p[j] == '*' {
				res[i][j] = res[i][j+1] || res[i+1][j]
				continue
			}
			res[i][j] = false
		}
	}
	//fmt.Println(res)
	return res[0][0]
}

func isMatchFront(s string, p string) bool {
	slen := len(s)
	plen := len(p)
	res := make([][]bool, slen+1)
	for i := range res {
		res[i] = make([]bool, plen+1)
	}
	res[0][0] = true
	for i := 1; i < slen+1; i++ {
		res[i][0] = false
	}
	for j := 1; j < plen+1; j++ {
		if p[j-1] == '*' && res[0][j-1] {
			res[0][j] = true
		} else {
			res[0][j] = false
		}
	}
	for i := 1; i < slen+1; i++ {
		for j := 1; j < plen+1; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '?' {
				res[i][j] = res[i-1][j-1]
			} else if p[j-1] == '*' {
				res[i][j] = res[i-1][j] || res[i][j-1]
			} else {
				res[i][j] = false
			}
		}
	}
	return res[slen][plen]
}
