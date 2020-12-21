package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
)

const (
	uvnan    = 0x7FF8000000000001
	uvinf    = 0x7FF0000000000000
	uvneginf = 0xFFF0000000000000
	mask     = 0x7FF
	shift    = 64 - 11 - 1
	bias     = 1023
)

func main() {

	//a := math.NaN()
	mapA := make(map[string]interface{})
	mapA["a"] = math.Inf(0)
	mapA["b"] = "b"
	jsonStr, err := json.Marshal(mapA)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(jsonStr)

}
