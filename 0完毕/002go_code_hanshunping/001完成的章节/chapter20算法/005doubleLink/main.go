package main

import(
	"fmt"
)

type HerosNode struct{
	No int
	Name string
	Prev *HerosNode
	Next *HerosNode
}

// 末尾插入双向链表
func InsertNode(headNode *HerosNode, newHerosNode *HerosNode)  {
	temp := headNode
	for{
		if temp.Next == nil {
			// 该链表是最后的链表
			temp.Next = newHerosNode
			newHerosNode.Prev = temp
			break
		}
		temp = temp.Next
	}
}
// 有序插入双向链表
func InsertNode2(headNode *HerosNode, newHerosNode *HerosNode)  {
	temp := headNode
	flag := true
	for{
		if temp.Next == nil {
			// 到了最后的链表了
			break
		}else if temp.Next.No >= newHerosNode.No {
			// 找到插入的位置了
			break
		}else if temp.Next.No == newHerosNode.No {
			// 该No的链表已经存在该
			flag = false
			break
		}
		temp = temp.Next
	}
	if !flag {
		fmt.Println("对不起，该链表已存在")
		return
	}else{
		newHerosNode.Next = temp.Next
		newHerosNode.Prev = temp
		if temp.Next != nil {
			temp.Next.Prev = newHerosNode
		}
		temp.Next = newHerosNode
	}
}

// 删除链表结点
func DeleteNode(headNode *HerosNode, id int)  {
	temp := headNode
	flag := false
	for{
		if temp.Next == nil {
			// 到了最后的链表了
			break
		}else if temp.Next.No == id {
			// 该No的链表已经找到
			flag = true
			break
		}
		temp = temp.Next
	}
	if !flag {
		fmt.Println("对不起，该链表不存在")
		return
	}else{

		temp.Next = temp.Next.Next
		if temp.Next != nil {
			temp.Next.Prev = temp
		}


	}
}

// 反向显示双向链表
func RevshowLink(headNode *HerosNode){
	temp := headNode
	if temp.Next == nil {
		// 空链表
		fmt.Println("空链表")
	}
	// 让头结点定位到最后一个结点
	for{
		if temp.Next == nil {
			// 找到最后的链表
			break
		}
		temp = temp.Next
	}
	// 开始输入
	for{
		if temp.Prev == nil {
			// 到头了
			fmt.Println("到头了")
			break
		}
		fmt.Printf("%d %s --> ",temp.No,temp.Name)
		temp = temp.Prev
	}
}
// 显示链表
func showLink(headNode *HerosNode){
	temp := headNode
	for{
		if temp.Next == nil {
			// 空链表
			fmt.Println("空链表")
			break
		}
		fmt.Printf("%d %s --> ",temp.Next.No,temp.Next.Name)
		temp = temp.Next
	}
}

func main()  {
	headNode := &HerosNode{}
	herosNode1 := & HerosNode{
		No:1,
		Name :"小小",
	}
	herosNode2 := & HerosNode{
		No:2,
		Name :"毛毛",
	}
	herosNode3 := & HerosNode{
		No:3,
		Name :"哒哒",
	}
	showLink(headNode)
	InsertNode(headNode, herosNode1)
	InsertNode(headNode, herosNode2)
	InsertNode(headNode, herosNode3)
	showLink(headNode)
	RevshowLink(headNode)
}