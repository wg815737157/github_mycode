package main

import "fmt"

func filterRepeatElement() {
	inputArray := []int{1, 2, 3, 4, 5, 4, 4, 5, 5, 5, 5}
	end := 0
	for i := 1; i < len(inputArray); i++ {
		if inputArray[i-1] != inputArray[i] {
			inputArray[end] = inputArray[i-1]
			end++
		}
	}
	end++
	inputArray[end] = inputArray[len(inputArray)-1]
	fmt.Println(inputArray[:end])
}

func main() {
	filterRepeatElement()

}
