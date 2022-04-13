package main

import "fmt"

func myCalculate(s string) int {
	pre := '+'
	digitS := []int{}
	n := len(s)
	digit := 0
	for index, c := range s {
		if index != n-1 {
			if c == ' ' {
				continue
			}
			if '0' <= c && c <= '9' {
				digit = digit*10 + int(c-'0')
				continue
			}
			switch pre {
			case '+':
				digitS = append(digitS, digit)
			case '-':
				digitS = append(digitS, -digit)
			case '*':
				digitS[len(digitS)-1] = digitS[len(digitS)-1] * digit
			default:
				digitS[len(digitS)-1] = digitS[len(digitS)-1] / digit
			}
			digit = 0
			pre = c
		} else {
			if '0' <= c && c <= '9' {
				digit = digit*10 + int(c-'0')
			}
			switch pre {
			case '+':
				digitS = append(digitS, digit)
			case '-':
				digitS = append(digitS, -digit)
			case '*':
				digitS[len(digitS)-1] = digitS[len(digitS)-1] * digit
			default:
				digitS[len(digitS)-1] = digitS[len(digitS)-1] / digit
			}
		}
	}
	res := 0
	for _, v := range digitS {
		res = res + v
	}
	return res
}

func calculate(s string) int {
	preSign := '+'
	stack := []int{}
	num := 0
	for i, ch := range s {
		condition := ch <= '9' && ch >= '0'
		if condition {
			num = num*10 + int(ch-48)
		}
		if !condition && ch != ' ' || i == len(s)-1 {
			switch preSign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] = stack[len(stack)-1] * num
			case '/':
				stack[len(stack)-1] = stack[len(stack)-1] / num
			}
			num = 0
			if i != len(s)-1 {
				preSign = ch
			}
		}
	}
	res := 0
	for _, v := range stack {
		res += v
	}
	return res
}

func calculate2(s string) int {
	opt := []int{1}
	sign := 1
	n := len(s)
	num := 0
	res := 0
	for i := 0; i < n; {
		switch s[i] {
		case '+':
			sign = opt[len(opt)-1]
			i++
		case '-':
			sign = -opt[len(opt)-1]
			i++
		case '(':
			opt = append(opt, sign)
			i++
		case ')':
			opt = opt[:len(opt)-1]
			i++
		case ' ':
			i++
		default:
			for ; i < n && '0' <= s[i] && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			res += sign * num
			num = 0
		}
	}
	return res

}
func main() {
	s := "2-1+2"
	fmt.Println(calculate2(s))
}
