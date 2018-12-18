package main

//冒泡排序从大到小
import "fmt"

func main() {
	var num = [10]int{1, 5, 7, 6, 8, 2, 3, 1, 9, 0}
	fmt.Print("原数组")
	fmt.Println(num)

	for i := 1; i < len(num); i++ {
		for j := 0; j < len(num)-i; j++ {
			if num[j] < num[j+1] {
				num[j], num[j+1] = num[j+1], num[j]
			}
		}
	}
	fmt.Print("排序后")
	fmt.Println(num)

}
