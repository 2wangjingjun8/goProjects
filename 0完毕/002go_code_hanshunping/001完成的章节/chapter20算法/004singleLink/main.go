package main

import(
	"fmt"
)

type HerosNode struct{
	No int
	Name string
	Next *HerosNode
}

// 末尾插入链表
func InsertNode(headNode *HerosNode, newHerosNode * HerosNode)  {
	temp := headNode
	for{
		if temp.Next == nil {
			// 该链表是最后的链表
			temp.Next = newHerosNode
			break
		}
		temp = temp.Next
	}
}
// 有序插入链表
func InsertNode2(headNode *HerosNode, newHerosNode * HerosNode)  {
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
	herosNode4 := & HerosNode{
		No:3,
		Name :"毛阿莫",
	}
	showLink(headNode)
	InsertNode2(headNode, herosNode1)
	InsertNode2(headNode, herosNode4)
	InsertNode2(headNode, herosNode2)
	InsertNode2(headNode, herosNode3)
	showLink(headNode)
	DeleteNode(headNode,2)
	showLink(headNode)
	DeleteNode(headNode,5)
	showLink(headNode)
	DeleteNode(headNode,3)
	showLink(headNode)
	DeleteNode(headNode,1)
	showLink(headNode)

}