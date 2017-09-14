package main

func containsNearbyDuplicate(nums []int, k int) bool {
	if nums == nil {
		return false
	}
	hashMap := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		if i > k {
			delete(hashMap, nums[i-k-1])
		}
		if hashMap[nums[i]] {
			return true
		}
		hashMap[nums[i]] = true
	}

	return false
}
