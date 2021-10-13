package main

import "fmt"

func StrContainAlpha(str []rune, b rune) bool {
	for i := 0; i < len(str); i++ {
		if str[i] == b {
			return true
		}
	}
	return false
}

func lengthOfLongestSubstring(str string) int {
	strArray := []rune(str)
	if len(strArray) == 0 {
		return 0
	}
	i, j := 0, 1
	maxStr := 1
	for ; j < len(strArray); j++ {
		//不包含
		if !StrContainAlpha(strArray[i:j], strArray[j]) {
			maxStr = twoMaxNum(maxStr, j-i+1)
			continue
		}
		//包含
		for i = i + 1; i < j; i++ {
			if !StrContainAlpha(strArray[i:j], strArray[j]) {
				break
			}
		}
	}
	return maxStr
}

func charInString(c byte, sBytes []byte) bool {
	n := len(sBytes)
	for i := 0; i < n; i++ {
		if c == sBytes[i] {
			return true
		}
	}
	return false
}

func lengthOfLongestSubstring2(s string) int {
	sBytes := []byte(s)
	n := len(sBytes)
	if n == 0 {
		return 0
	}
	maxLen := 1
	maxBytes := sBytes[:1]
	for i := 1; i < n; i++ {
		if !charInString(sBytes[i], maxBytes) {
			maxBytes = append(maxBytes, sBytes[i])
			if maxLen < len(maxBytes) {
				maxLen = len(maxBytes)
			}
		} else {
			for j := 1; ; j++ {
				if !charInString(sBytes[i], maxBytes[j:]) {
					maxBytes = maxBytes[j:]
					maxBytes = append(maxBytes, sBytes[i])
					break
				}
			}
		}
	}
	return maxLen
}

func lengthOfLongestSubstring3(s string) int {
	sBytes := []byte(s)
	n := len(sBytes)
	if n == 0 {
		return 0
	}
	maxLen := 1
	characterKey := make(map[byte]bool)
	for i, j := 0, 0; j < n; {
		for ; j < n && characterKey[sBytes[j]] == false; {
			characterKey[sBytes[j]] = true
			if j-i+1 > maxLen {
				maxLen = j - i + 1
			}
			j++
			continue
		}
		delete(characterKey, sBytes[i])
		i++
	}
	return maxLen
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func lengthOfLongestSubstring4(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := 0, 0
	for i := 0; i < n; i++ {
		for rk < n && m[s[rk]] == 0 {
			// 不断地移动右指针
			m[s[rk]]++
			// 第 i 到 rk 个字符是一个极长的无重复字符子串
			ans = max(ans, rk-i+1)
			rk++
		}
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
	}
	return ans
}

func main() {
	str := "pwwkew"
	res := lengthOfLongestSubstring4(str)
	fmt.Println(res)
}
