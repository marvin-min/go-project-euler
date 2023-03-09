package main

import "fmt"

func Question2() {
	max := 4_000_000
	first := 0
	second := 1
	sum := 0
	for second < max {
		first, second = second, second+first
		fmt.Printf("%v ", second)
		if second%2 == 0 {
			sum += second
		}
	}
	fmt.Println("total is:", sum)
}
