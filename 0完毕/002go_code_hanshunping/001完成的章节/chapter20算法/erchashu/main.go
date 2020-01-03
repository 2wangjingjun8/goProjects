package main

import(
	"fmt"
)

type Heros struct{
	No int
	Name string
	Left *Heros
	Right *Heros
}

func Beforeheros(node *Heros)  {
	if node != nil {
		fmt.Printf("%d %s\n",node.No,node.Name)
		Beforeheros(node.Left)
		Beforeheros(node.Right)
	}
}
func Middleheros(node *Heros)  {
	if node != nil {
		Middleheros(node.Left)
		fmt.Printf("%d %s\n",node.No,node.Name)
		Middleheros(node.Right)
	}
}
func Afterheros(node *Heros)  {
	if node != nil {
		Afterheros(node.Left)
		Afterheros(node.Right)
		fmt.Printf("%d %s\n",node.No,node.Name)
	}
}

func main()  {
	// 创建一个二叉树
	root := &Heros{
		No:1,
		Name:"老大",
	}
	left1 := &Heros{
		No:2,
		Name:"老2",
	}
	right1 := &Heros{
		No:3,
		Name:"老3",
	}
	root.Left = left1
	root.Right = right1

	left21 := &Heros{
		No:4,
		Name:"老4",
	}
	right22 := &Heros{
		No:5,
		Name:"老5",
	}
	left1.Left = left21
	left1.Right = right22

	left23 := &Heros{
		No:6,
		Name:"老6",
	}
	right24 := &Heros{
		No:7,
		Name:"老7",
	}
	right1.Left = left23
	right1.Right = right24

	fmt.Println("二叉树的前序遍历情况是：")
	Beforeheros(root)
	fmt.Println("二叉树的中序遍历情况是：")
	Middleheros(root)
	fmt.Println("二叉树的后序遍历情况是：")
	Afterheros(root)


}