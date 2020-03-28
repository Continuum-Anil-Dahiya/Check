package array

import "fmt"

//https://practice.geeksforgeeks.org/problems/inversion-of-array/0/
//Inversion Count for an array indicates – how far (or close) the array is from being sorted
//Formally, two elements a[i] and a[j] form an inversion if a[i] > a[j] and i < j.

//InversionOfArray ...
func InversionOfArray() {
	var t int
	fmt.Scan(&t)

	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	//inversionOfArrayfirstApproach(arr, n)
	count := 0
	inversionOfArray(arr, 0, n-1, &count)
	fmt.Println(count)

}

//O(n square) Approach
func inversionOfArrayfirstApproach(arr []int, high int) {

	count := 0
	for i := 0; i < high-1; i++ {
		for j := i + 1; j < high; j++ {
			if arr[i] > arr[j] {
				count++
			}
		}

	}

	fmt.Println(count)

}

//O(n square) Approach
func inversionOfArray(arr []int, low int, high int, count *int) {

	if low < high {
		mid := (high + low) / 2
		inversionOfArray(arr, low, mid, count)
		inversionOfArray(arr, mid+1, high, count)
		inversionOfArrayUtil(arr, low, mid, high, count)
	}

}

func inversionOfArrayUtil(arr []int, low int, mid int, high int, count *int) {

	left := make([]int, mid-low+1)
	right := make([]int, high-mid)

	for i, j := low, 0; i <= mid; i, j = i+1, j+1 {
		left[j] = arr[i]
	}

	for i, j := mid+1, 0; i <= high; i, j = i+1, j+1 {
		right[j] = arr[i]
	}

	i := 0
	j := 0
	k := low
	for i < len(left) && j < len(right) {

		if left[i] <= right[j] {
			arr[k] = left[i]
			k++
			i++
		} else {
			arr[k] = right[j]
			k++
			j++

			*count = *count + mid - i + 1

		}
	}

	for i < len(left) {
		arr[k] = left[i]
		k++
		i++
	}

	for j < len(right) {
		arr[k] = right[j]
		k++
		j++
	}
}
