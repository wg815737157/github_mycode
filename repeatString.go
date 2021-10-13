package main

import (
	"fmt"
	"strconv"
)

func isInt(a rune) bool {
	if 49 <= a && a <= 57 {
		return true
	}
	return false
}
func isChar(a rune) bool {
	if a == '(' || a == ')' {
		return false
	}
	return !isInt(a)
}

func repeatString(s string) {

	runeS := []rune(s)

	stack1 := &Stack{}
	stack2 := &Stack{}

	res := []rune{}

	for i := 0; i < len(runeS); {
		fmt.Println(i)
		//字符
		if isChar(runeS[i]) {
			//fmt.Println(string(runeS[i]))
			res = append(res, runeS[i])
			i++
			continue
		}
		// 数值
		if isInt(runeS[i]) {
			n := i
			for ; isInt(runeS[n]) && n < len(runeS); n++ {
			}
			v, err := strconv.Atoi(string(runeS[i:n]))
			if err != nil {
				fmt.Println(err)
				return
			}
			stack1.Push(v)
			i = n
			continue
		}

		// 左括号
		if runeS[i] == '(' {
			//fmt.Println(string(runeS[i]))
			n := i + 1
			for ; isChar(runeS[n]) && n < len(runeS); n++ {
			}
			stack2.Push(runeS[i+1 : n])
			i = n
			continue
		}
		// 右括号
		if runeS[i] == ')' {
			//fmt.Println(string(runeS[i]))
			v := stack1.Pop().(int)
			str := stack2.Pop().([]rune)

			tmp := []rune{}
			for j := 0; j < v; j++ {
				tmp = append(tmp, str...)
			}
			fmt.Println(string(tmp))

			if stack2.Len() > 0 {
				lowStr := stack2.Pop().([]rune)
				lowStr = append(lowStr, tmp...)
				stack2.Push(lowStr)
				i++
				continue
			}

			//最后一个')'
			res = append(res, tmp...)
			i++
		}
	}
	fmt.Println(string(res))
}

func main() {

	str := "a2(b2(cb))"

	repeatString(str)

}
