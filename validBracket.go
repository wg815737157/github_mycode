package main

import "fmt"

func validBracket(a string) bool {
	runeA := []rune(a)
	s := &Stack{}
	fmt.Println(len(runeA))
	for i := 0; i < len(runeA); i++ {
		if s.Len() == 0 {
			s.Push(runeA[i])
			continue
		}
		if s.Len() == 1 {
			if s.Peak().(rune) == runeA[i] {
				s.Pop()
				continue
			}
			fmt.Println(i)
			fmt.Println(runeA[i])

			fmt.Println(s.Peak().(rune))
			return false
		}
	}
	if s.Len() > 0 {
		return false
	}
	return true
}

func main() {
	a := "[]"
	fmt.Println(validBracket(a))
}
