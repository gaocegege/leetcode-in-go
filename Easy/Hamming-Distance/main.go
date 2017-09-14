// 位运算

package main

import "fmt"

func hammingDistance(x int, y int) int {
	ir := x ^ y
	count := 0
	for ir != 0 {
		if ir%2 == 1 {
			count++
		}
		ir = ir >> 1
	}
	return count
}

func main() {
	fmt.Print(hammingDistance(1, 4))
}
