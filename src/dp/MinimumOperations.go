package dp

import "fmt"

//Link
//https://www.geeksforgeeks.org/minimum-cost-to-reach-a-point-n-from-0-with-two-different-operations-allowed/
// to reach  from o to num operation allowed double a number or add + 1 in number
// Solution go from num  to zero

// Main5 ...
func Main5() {

	var t int
	fmt.Scan(&t)

	for t > 0 {
		var num int
		fmt.Scan(&num)
		fmt.Println(findCost(num))
		t--
	}
}

func findCost(num int) int {

	count := 0
	for num > 0 {
		if num%2 == 0 {
			num = num / 2
		} else {
			num = num - 1
		}

		count++

	}

	return count
}
