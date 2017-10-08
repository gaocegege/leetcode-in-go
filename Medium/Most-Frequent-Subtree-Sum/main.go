package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var resMap = map[int]int{}

var maxTime = 0

func findFrequentTreeSum(root *TreeNode) []int {
	// Fix the inconsistency bug.
	resMap = map[int]int{}
	maxTime = 0
	res := make([]int, 0)
	getAndSetSum(root)
	for k, v := range resMap {
		if v == maxTime {
			res = append(res, k)
		}
	}
	return res
}

func getAndSetSum(tn *TreeNode) int {
	if tn == nil {
		return 0
	}
	score := tn.Val + getAndSetSum(tn.Left) + getAndSetSum(tn.Right)
	resMap[score] = resMap[score] + 1
	if resMap[score] > maxTime {
		maxTime = resMap[score]
	}
	return score
}
