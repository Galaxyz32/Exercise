package main

import "fmt"

var max int
var pre int

func main() {
	max = 60
	pre = 0
	f(max)
}
func f(num int) {
	i := 0
	fi := 1
	fmt.Printf("不大于%d的fib数列项有 \n", max)
	for ; fi < num; i++ {
		pre = fi
		fmt.Printf("第%d个数是%d \n", i-1, fi)
		fi, _ = fib(i)
	}
	_, k := fib(i - 2)

	fmt.Printf("其中最大的是第%d项，%d", k, pre)
}

func fib(i int) (int, int) {
	if i <= 1 {
		return 1, i
	} else {
		x, _ := fib(i - 1)
		y, _ := fib(i - 2)
		return x + y, i
	}
}
