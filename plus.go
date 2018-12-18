package main

import "fmt"

var (
	sum     = 0
	oddsum  = 0
	fracsum = 0.0
)

func main() {
	var (
		num     = 100
		numodd  = 99
		numfrac = 99
	)
	plus(num)
	fmt.Printf("从1加到%d的和是", num)
	fmt.Println(sum)

	plusodd(numodd)
	fmt.Printf("从1加到%d所有奇数的和是", numodd)
	fmt.Println(oddsum)

	plusfrac(numfrac)
	fmt.Printf("从1加到1/%d的和是", numfrac)
	fmt.Printf("%.4f", fracsum)
}

func plus(n int) {
	for n > 0 {
		sum += n
		n--
	}
}

func plusodd(n int) {
	for n > 0 {
		oddsum += n
		n = n - 2
	}
}

func plusfrac(n int) {
	for n > 0 {
		fracsum += 1.0 / float64(n)
		n = n - 2
	}
}
