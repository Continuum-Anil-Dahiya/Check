package sort

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

//Methods19 calls from main2/ch.go
func Methods19() {

	reader := bufio.NewReader(os.Stdin)
	read, _ := reader.ReadString('\n')

	str := strings.Split(strings.TrimSpace(read), " ")

	arr := make([]int, len(str))
	res := make([]int, len(str))
	max := math.MinInt64
	for v, k := range str {
		arr[v], _ = strconv.Atoi(k)

		if arr[v] > max {
			max = arr[v]
		}

	}

	arr1 := make([]int, max+1)
	for v := range str {
		arr1[arr[v]] = arr1[arr[v]] + 1
	}

	for i := 1; i < len(arr1); i++ {
		arr1[i] = arr1[i] + arr1[i-1]
	}

	for v := range str {
		arr1[arr[v]] = arr1[arr[v]] - 1
		res[arr1[arr[v]]] = arr[v]
	}

	fmt.Println(res)
}
