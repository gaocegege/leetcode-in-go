package main

import "fmt"

func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0, 100)
	result := make([]int, k)

	calculate(&res, result, 0, 1, k, n)
	return res
}

func sum(result []int) int {
	sum := 0
	for _, v := range result {
		sum += v
	}
	return sum
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func calculate(res *[][]int, result []int, index, start, k, n int) {
	if index == k {
		if sum(result) == n {
			buf := make([]int, k)
			copy(buf, result)
			*res = append(*res, buf)
			return
		}
		return
	}
	for i := start; i < 10; i++ {
		fmt.Println(result)
		result[index] = i
		calculate(res, result, index+1, i+1, k, n)
	}
}

func main() {
	fmt.Println(combinationSum3(4, 24))
}
