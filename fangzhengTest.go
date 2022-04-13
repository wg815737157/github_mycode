package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	//fmt.Println(3 << 1)
	fmt.Println(myAtoi("  +  413"))
}
func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := dfs(root.Left)
	rightHeight := dfs(root.Right)
	if leftHeight == -1 || rightHeight == -1 || abs(leftHeight-rightHeight) > 1 {
		return -1
	}
	return max(leftHeight, rightHeight) + 1
}
func isBalanced(root *TreeNode) bool {
	if dfs(root) == -1 {
		return false
	}
	return true
}
func myAtoi(s string) int {
	max := math.MaxInt32
	min := math.MinInt32
	isdigital := false
	res := 0
	fmt.Println(unsafe.Sizeof(res))
	sign := ' '
	for i := 0; i < len(s); i++ {
		if isdigital == false {
			if s[i] == ' ' {
				if sign == ' ' {
					continue
				}
				return 0
			}

			if '0' <= s[i] && s[i] <= '9' {
				res = int(s[i] - '0')
				isdigital = true
				continue
			}

			if s[i] == '+' || s[i] == '-' {
				if sign == ' ' {
					sign = int32(s[i])
					continue
				}
				return 0
			}
			return 0
		}

		if '0' <= s[i] && s[i] <= '9' {
			units := int(s[i] - '0')
			if sign == '+' || sign == ' ' {
				tmp := 10 * res
				if max-tmp > units {
					res = 10*res + units
					continue
				}
				return max
			}
			if sign == '-' {
				tmp := -res * 10
				if tmp-min > units {
					res = 10*res + units
					continue
				}
				return min
			}
			continue
		}
		break
	}
	if sign == '-' {
		res = -res
	}
	return res
}
