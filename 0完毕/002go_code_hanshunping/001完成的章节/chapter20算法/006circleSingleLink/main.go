package main

import(
	"fmt"
)

type CatNode struct{
	Id int
	Name string
	Next *CatNode
}

func Insert(head *CatNode, newCatNode *CatNode)  {
	if head.Next == nil {
		// 添加第一个猫猫
		head.Id = newCatNode.Id
		head.Name = newCatNode.Name
		head.Next = head
		return
	}
	// 找到最后的结点
	temp := head
	for{
		if temp.Next == head {
			break
		}
		temp = temp.Next
	}
	// 加入到链表中
	temp.Next = newCatNode
	newCatNode.Next = head
}

func List(head *CatNode)  {
	fmt.Println("环形链表情况如下：")
	if head.Next == nil {
		fmt.Println("环形链表为空链表")
	}
	temp := head
	for{
		fmt.Printf("%d %s -->%d %s\n",temp.Id, temp.Name,temp.Next.Id, temp.Next.Name)
		if temp.Next == head{
			break
		}
		temp = temp.Next
	}
}

func main()  {
	head := &CatNode{}
	cat1 := &CatNode{
		Id:1,
		Name:"tom",
	}
	cat2 := &CatNode{
		Id:2,
		Name:"Jomn",
	}
	cat3 := &CatNode{
		Id:3,
		Name:"dad",
	}
	Insert(head, cat1)
	Insert(head, cat2)
	Insert(head, cat3)
	List(head)

}