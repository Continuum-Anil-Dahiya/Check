package array

import "fmt"

//Link

//Method1
//Using Min heap for largest nth element
//USing Max Heap for smallest nth element
//https://www.youtube.com/watch?v=hGK_5n81drs complexity O(nlogk)

//Method2
//QuickSort in O(n) in average time complexity

//CheckQuick ...
func CheckQuick() {

	var t int
	fmt.Scan(&t)
	for t > 0 {
		var n, k, element int
		fmt.Scan(&n)
		fmt.Scan(&k)
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&arr[i])
		}

		quicksort(arr, 0, n-1, n-k, &element)
		fmt.Println(element)

		t--
	}

}

func quicksort(arr []int, low int, high int, nthlargestIndex int, element *int) {

	if low < high {

		index := partionLogic(arr, low, high)

		if index == nthlargestIndex {
			*element = arr[index]
			return
		} else if index > nthlargestIndex {
			quicksort(arr, low, index-1, nthlargestIndex, element)
		} else {
			quicksort(arr, index+1, high, nthlargestIndex, element)
		}

	}

}

func partionLogic(arr []int, low int, high int) int {

	pivoit := arr[high]
	partionIndex := low

	for i := low; i < high; i++ {

		if pivoit >= arr[i] {
			arr[i], arr[partionIndex] = swap(arr[i], arr[partionIndex])
			partionIndex++
		}
	}

	arr[partionIndex], arr[high] = swap(arr[partionIndex], arr[high])

	return partionIndex
}

func swap(a int, b int) (int, int) {
	return b, a
}
