package first

import (
	"fmt"
	"os"
	"strconv"
)

func sumDivisible(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err == nil {
		fmt.Println(sumDivisible(int(n)))
	}
}
