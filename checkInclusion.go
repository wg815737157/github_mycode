package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	count1 := [26]int{}
	count2 := [26]int{}
	for _, v := range s1 {
		count1[v-'a']++
	}
	for i := 0; i < len(s1); i++ {
		count2[s2[i]-'a']++
	}
	if count1 == count2 {

		return true
	}
	for i := len(s1); i < len(s2); i++ {
		count2[s2[i-len(s1)]-'a']--
		count2[s2[i]-'a']++
		if count1 == count2 {
			return true
		}
	}
	return false
}
func main() {
	s1 := "aa"
	s2 := "ab"
	fmt.Println(checkInclusion(s1, s2))
}
