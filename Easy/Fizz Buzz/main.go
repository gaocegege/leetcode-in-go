package main

import "strconv"

func fizzBuzz(n int) []string {
	fizzFlag := 1
	BuzzFlag := 1
	flag := false
	results := make([]string, n)
	for i := 1; i <= n; i++ {
		res := ""
		if fizzFlag == 3 {
			res += "Fizz"
			fizzFlag = 0
			flag = true
		}
		if BuzzFlag == 5 {
			res += "Buzz"
			BuzzFlag = 0
			flag = true
		}
		fizzFlag++
		BuzzFlag++
		if flag == true {
			results[i-1] = res
			flag = false
		} else {
			results[i-1] = strconv.Itoa(i)
		}
	}
	return results
}
