package dp

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//https://www.codechef.com/problems/ALTARAY

//Main1 ...
func Main1() {

	var t int
	//fmt.Scan(&t)

	reader := bufio.NewReader(os.Stdin)
	read, _ := reader.ReadString('\n')

	t, _ = strconv.Atoi(strings.TrimSpace(read))

	for i := 0; i < t; i++ {

		var n int

		read, _ := reader.ReadString('\n')
		n, _ = strconv.Atoi(strings.TrimSpace(read))

		ans := make([]int, n)
		read, _ = reader.ReadString('\n')
		str := strings.Split(strings.TrimSpace(read), " ")

		var p, c string

		for j := n - 1; j >= 0; j-- {

			a, _ := strconv.Atoi(str[j])
			if a < 0 {
				c = "N"
			} else {
				c = "P"
			}

			if p == c || j == n-1 {
				ans[j] = 1
			} else {
				ans[j] = ans[j+1] + 1
			}

			p = c
		}

		for k := range ans {
			fmt.Print(ans[k])
			if k != (len(ans) - 1) {
				fmt.Print(" ")
			}
		}

		if i != (t - 1) {
			fmt.Println("")
		}

	}

}
