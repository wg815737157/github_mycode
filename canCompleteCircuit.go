package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	for i := 0; i < n; {
		count := 0
		sumGas, sumCost := 0, 0
		for ; count < n; count++ {
			tmp := (i + count) % n
			sumGas += gas[tmp]
			sumCost += cost[tmp]
			if sumGas < sumCost {
				break
			}
		}
		if count == n {
			return i
		}
		i += count + 1
	}
	return -1
}
func main() {
	fmt.Println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))

}
