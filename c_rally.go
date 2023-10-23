package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Scan(&N)
	x := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&x[i])
	}

	ans := math.MaxInt

	var a, b int
	a, b = findMinMax(x)

	for i := a; i <= b; i++ {
		var l []int
		for _, j := range x {
			l = append(l, (j-i)*(j-i))
		}
		ans = intMin(ans, sum(l))
	}

	fmt.Println(ans)
}

func findMinMax(arr []int) (int, int) {
	min := arr[0]
	max := arr[0]

	for _, val := range arr {
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}

	return min, max
}

func sum(arr []int) int {
	total := 0
	for _, val := range arr {
		total += val
	}
	return total
}

func intMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
