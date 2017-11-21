package main

func matrixReshape(nums [][]int, r int, c int) [][]int {
	rawRow := len(nums)
	rawColomn := len(nums[0])
	if r*c != rawRow*rawColomn {
		return nums
	}
	posX := 0
	posY := 0
	res := make([][]int, r)
	for i := 0; i < r; i++ {
		res[i] = make([]int, c)
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			res[i][j] = nums[posX][posY]
			posY++
			if posY == rawColomn {
				posX++
				posY = 0
			}
		}
	}
	return res
}
