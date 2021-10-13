package main

import "fmt"

func initArray() []rune {
	// s 正确解 "a b c dd e "
	s := "a  b  cc d  eee    "
	return []rune(s)
}

func main() {
	runeArray := initArray()
	checkRune := ' '
	len := len(runeArray)
	for i := 0; i < len; i++ {
		//检测到空格
		if checkRune == runeArray[i] {
			n := i + 1
			for ; n < len; n++ {
				//第二个是空格
				if runeArray[n] == checkRune {
					continue
				}
				//不相等 n=i+1
				if n == i+1 {
					break
				}
				//不相等 并且大于2
				runeArray = append(runeArray[:i+1], runeArray[n:]...)
				len = len - (n - i - 1)
				break
			}
			if n == len {
				len = i
			}
		}
	}
	fmt.Println(string(runeArray[:len]))
}
