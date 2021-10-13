package main

import (
	"fmt"
	"math"
)

func revertInt(chu int) (ret int) {
	for chu != 0 {
		yu := chu % 10
		chu = chu / 10
		ret = ret*10 + yu
		if ret < (-1<<31) || ret > (1<<31-1) {
			return 0
		}
	}
	return
}

func main() {
	a := -120
	fmt.Println(revertInt(a))
}

func reverse(x int) (rev int) {
	for x != 0 {
		if rev < math.MinInt32/10 || rev > math.MaxInt32/10 {
			return 0
		}
		digit := x % 10
		x /= 10
		rev = rev*10 + digit
	}
	return
}
