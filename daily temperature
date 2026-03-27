package main

import (
	"fmt"
)

func main() {
	var arr []int = []int{30, 38, 30, 36, 35, 40, 28}
	var out []int = make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		var j int
		for j = i + 1; j < len(arr); j++ {
			if arr[i] < arr[j] {
				break
			}
		}
		if j != len(arr) {
			out[i] = j - i
		}
	}
	fmt.Println(out)
}
