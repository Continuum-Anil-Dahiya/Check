package dp

import "fmt"

//Main4 ...
func Main4() {
	var t int
	fmt.Println("Enter the number of testcases:")
	fmt.Scan(&t)

	for t > 0 {
		var lengthArray int
		fmt.Scan(&lengthArray)
		arr := make([]int, lengthArray)

		for i := 0; i < lengthArray; i++ {
			fmt.Scan(&arr[i])
		}

		liss(arr)
		t--
	}

}

func liss(arr []int) {

	l := make([]int, len(arr))
	sub := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		l[i] = 1
		temp := 1
		for j := 0; j < i; j++ {
			if arr[j] <= arr[i] && (l[j]+1) >= temp {
				temp = l[j] + 1
				sub[i] = j
			}
		}
		l[i] = temp
	}

	fmt.Println("sub")
	fmt.Println(sub)
	fmt.Println("l")
	fmt.Println(l)

}
