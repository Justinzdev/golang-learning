package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	for idx, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", idx, value)
	}
}
