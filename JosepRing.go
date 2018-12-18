package main

import "fmt"

type node struct {
	index int
	next  *node
	exist bool
}

var head, headline *node
var num, key int

func main() {

	var ptr *node
	println("请输入节点数，出列序号")
	fmt.Scanln(&num)
	fmt.Scanln(&key)

	ptr = new(node)
	head = ptr
	headline = head
	for i := 1; i < num; i++ {
		ptr.index = i
		ptr.exist = true
		ptr.next = new(node)
		ptr = ptr.next
	}
	ptr.index = num
	ptr.exist = true
	ptr.next = head

	printAllNode()

	for num != 0 {
		josepRing(head)
		fmt.Scanln()
	}
}

func printAllNode() {
	fmt.Print("当前全体结点")

	ptr := headline
	for i := 0; i < num; i++ {
		fmt.Print(ptr.index)
		fmt.Print("	")
		ptr = ptr.next
	}

	fmt.Println("———END———")
}

func josepRing(pt *node) {
	//fmt.Print("指向的结点")
	//fmt.Println(pt.index)

	pre := pt
	pt = pt.next

	for i := 1; i < key-1; i++ {
		pre = pt
		pt = pt.next
	}

	fmt.Print("本次出列结点")
	fmt.Println(pt.index)

	//fmt.Print("前一个节点")
	//fmt.Println(pre.index)

	pre.next = pt.next
	num--

	fmt.Print("出列后剩余节点")
	printAllNode()
	head = pre.next

}
