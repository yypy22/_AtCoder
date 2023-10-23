package main

import (
	"fmt"
)

func main() {
	var A, B int
	fmt.Scanf("%d %d", &A, &B)
	ans := 0
	outlet := 1

	for outlet < B {
		outlet--
		outlet += A
		ans++
	}

	fmt.Println(ans)
}
