package main

import "fmt"

func minRemove( s string ) string {
	sArr := []int{}
	res := ""
	for i:=0;i<len(s) ;i++  {
		if string(s[i])==")"&& len(sArr)==0 {
			continue
		}else if string(s[i])==")" {
			res += string(s[i])
			sArr = sArr[:(len(sArr)-1)]
			continue
		}else if string(s[i])=="(" {
			sArr = append(sArr, i)
		}
		res += string(s[i])
	}
	if len(sArr)>0 {
		for i:=len(sArr)-1;i>=0;i--{
			if sArr[i] == len(res) {
				res = res[:(len(res)-1)]
			}
			res = res[:sArr[i]]+res[sArr[i]+1:]
		}
	}
	return res
}

func main() {
	fmt.Println(minRemove("(m(i(h)o)yo("))
}