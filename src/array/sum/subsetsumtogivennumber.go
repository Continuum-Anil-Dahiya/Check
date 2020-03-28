package sum

import "fmt"

//https://practice.geeksforgeeks.org/problems/subarray-with-given-sum/0

//Check ...
func Check() {
	var t int
	fmt.Scan(&t)

	for t > 0 {

		var n, sum int
		fmt.Scan(&n)
		fmt.Scan(&sum)

		arr := make([]int, n)

		for i := 0; i < n; i++ {
			fmt.Scan(&arr[i])
		}

		findSubArray(arr, n, sum)

		t--
	}
}

func findSubArray(arr []int, n int, sum int) {

	map1 := make(map[int]int)
	currSum := 0

	for i := 0; i < n; i++ {
		currSum = currSum + arr[i]

		if currSum == sum {
			fmt.Println("StartingIndex")
			fmt.Println("0")
			fmt.Println("LastIndex")
			fmt.Println(i)
		}

		left := currSum - sum
		index, ok := map1[left]
		if ok {
			fmt.Println("StartingIndex")
			fmt.Println(index + 1)
			fmt.Println("LastIndex")
			fmt.Println(i)
		}

		map1[currSum] = i
	}
}
