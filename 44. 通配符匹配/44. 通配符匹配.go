package main

import "fmt"

func main() {
	a := "aa"
	p := "a"
	b := isMatch(a, p)
	fmt.Println(b)
}

func isMatch(s string, p string) bool {
	slen := len(s)
	plen := len(p)
	res := make([][]bool, slen)
	for i := range res {
		res[i] = make([]bool, plen)
	}
	for j:=0;j<plen;j++{
		for k:=0;k<slen;k++{
			if p[k]==s[j] || p[k]=='?'{
				if j==0&&k==0{
					res[0][0]=true
				}else{
					res[j][k]=res[j-1][k-1]
				}
				
			}
		}
	}
	//fmt.Println(res)
	return false
}
