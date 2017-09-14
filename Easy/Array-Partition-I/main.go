// 位运算

package main

import "fmt"
import "sort"

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	res := 0
	for i := 0; i < len(nums); i += 2 {
		res += nums[i]
	}
	return res
}

func main() {
	fmt.Print(arrayPairSum([]int{1, 2, 3, 4}))
}
