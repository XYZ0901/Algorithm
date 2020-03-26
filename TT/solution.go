package main

import "fmt"

// 找出第一个重复字符
func getFirstRepeatChar(strSrc string) byte {
	// write code here
	m := map[byte]int{}
	var b byte
	for i := 0; i < len(strSrc); i++ {
		if _, ok := m[strSrc[i]]; ok {
			b = strSrc[i]
			break
		}
		m[strSrc[i]] += 1
	}
	return b
}

// 删除在strPat中出现过的字符
func stringFilter( strSrc string ,  strPat string ) string {
	// write code here
	str := strSrc
	n := len(strSrc)
	for i:=0;i<n;i++{
		for j:=0;j<len(strPat);j++{
			if str[i] == strPat[j] {
				str = str[:i] + str[i+1:]
				i--
				n--
				break
			}
		}
	}
	fmt.Println(len(str))
	return str
}

func main() {
	fmt.Println(stringFilter("welcome to amazing seasunwwaecdf", ""))
}
