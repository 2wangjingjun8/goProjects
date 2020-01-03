package main

import(
	"fmt"
)

type Boy struct{
	No int
	Next *Boy
}

// 编写一个函数，构建单向的环形链表
// num 表示小孩的个数
// *Boy 返回该环形链表的第一个小孩的指针
func AddBoy(num int) *Boy {
	first := &Boy{} // 空结点
	curBoy := &Boy{} //辅助
	// 判断
	if num < 1 {
		fmt.Println("num的值不对")
		return first
	}
	//循环地构建这个环形链表
	for i := 1; i <= num; i++ {
		boy := &Boy{
			No: i,
		}
		// 分析构成训话链表，需要一个辅助指针
		// 1.因为第一个小孩比较特殊
		if i == 1 {
			first = boy
			curBoy = boy
			curBoy.Next = first
		}else{
			curBoy.Next = boy
			curBoy = boy //移动指针到下一个小孩
			curBoy.Next = first // 构成环形链表
		}
	}
	return first
}

func ShowBoy(first *Boy)  {
	// 处理一下如果环形链表为空
	if first.Next == nil {
		fmt.Println("链表为空，没有小孩")
	}
	// 创建一个指针，帮助遍历
	curBoy := first
	for{
		fmt.Printf("小孩编号是%d-->",curBoy.No)
		if curBoy.Next == first {
			break
		}
		curBoy = curBoy.Next
	}
}

func PlayGame(first *Boy, startNo int, countNum int)  {

	// 空的链表我们单独处理
	if first.Next == nil{
		fmt.Println("空的链表，没有小孩")
		return
	}

	// 判断startNo <= 小孩的总数
	// 需要定义辅助的指针，帮助我们删除小孩
	tail := first
	// 让tail指向环形链表的最后一个小孩
	// 因为tail在删除小孩需要使用到
	for{
		if tail.Next == first { // 说明tail到了最后一个小孩
			break
		}
		tail = tail.Next
	}
	// 让first移动到stratNo[后面我们删除小孩，就以first为准]
	for i := 1; i <= startNo - 1; i++ {
		first = first.Next
		tail = tail.Next
	}

	fmt.Println()
	// 开始数countNum，然后删除first指向的小孩
	for{
		for i := 1; i <= countNum - 1; i++ {
			first = first.Next
			tail = tail.Next
		}
		fmt.Printf("小孩编号为%d 出列\n",first.No)
		// 删除first指向的小孩
		first = first.Next
		tail.Next = first
		// 判断如果tail =first，圈子只有一个小孩
		if first == tail {
			break
		}
	}
	fmt.Printf("最后小孩编号为%d",first.No)

	
}

func main()  {
	first := AddBoy(5)
	ShowBoy(first)
	PlayGame(first,2,3)
}