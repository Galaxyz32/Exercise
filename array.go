package main

import "fmt"

func main() {
	var arr [10]int
	arr[1] = 1
	arr[2] = 2
	for i := 3; i < len(arr); i++ {
		arr[i] = arr[i-1] * arr[i-2]
	}

	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d	", arr[i])

	}

}
