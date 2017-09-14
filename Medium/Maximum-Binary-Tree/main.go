// 树的遍历

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findLargestIndex(nums []int, lower int, higher int) int {
	largest := -1
	index := 0
	for i := lower; i <= higher; i++ {
		if nums[i] > largest {
			largest = nums[i]
			index = i
		}
	}
	return index
}

func helper(nums []int, lower int, higher int) *TreeNode {
	if lower > higher {
		return nil
	}
	largestIndex := findLargestIndex(nums, lower, higher)
	tree := &TreeNode{
		Val:   nums[largestIndex],
		Left:  helper(nums, lower, largestIndex-1),
		Right: helper(nums, largestIndex+1, higher),
	}
	return tree
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return helper(nums, 0, len(nums)-1)
}
