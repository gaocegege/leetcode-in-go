package main

func numberOfArithmeticSlices(A []int) int {
	globalCounter := 0
	res := 0

	for i := 0; i <= len(A)-3; i++ {
		counter := 0
		if globalCounter > 0 {
			globalCounter--
			res += globalCounter
			continue
		}
		if A[i+1]-A[i] == A[i+2]-A[i+1] {
			counter++
			interval := A[i+1] - A[i]
			for j := i + 3; j < len(A); j++ {
				if A[j]-A[j-1] == interval {
					counter++
				} else {
					break
				}
			}
			globalCounter = counter
			res += counter
		} else {
			continue
		}
	}
	return res
}

// Simple Implementation
func numberOfArithmeticSlicesSimple(A []int) int {
	counter := 0

	for i := 0; i <= len(A)-3; i++ {
		if A[i+1]-A[i] == A[i+2]-A[i+1] {
			counter++
			interval := A[i+1] - A[i]
			for j := i + 3; j < len(A); j++ {
				if A[j]-A[j-1] == interval {
					counter++
				} else {
					break
				}
			}
		} else {
			continue
		}
	}
	return counter
}
