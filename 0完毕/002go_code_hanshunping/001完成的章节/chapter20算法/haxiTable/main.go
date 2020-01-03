package main

import(
	"fmt"
	"os"
)
type Emp struct{
	Id int
	Name string
	Next *Emp
}
type EmpLink struct{
	Head *Emp
}
//添加员工的方法， 保证添加的编号是从小到大
func (this *EmpLink) Insert(emp *Emp)  {
	cur := this.Head //这是辅助指针
	var pre *Emp = nil // 这是一个辅助指针 pre在cur前面
	// 如果当前的EmpLink是一个空链表
	if cur == nil{
		this.Head = emp
		return
	}
	//如果不是一个空链表，给emp找到合适的位置并插入
	for{
		if cur != nil {
			// 比较
			if cur.Id >emp.Id {
				//找到位置
				break
			}
			pre = cur // 保证同步
			cur = cur.Next
		}else{
			break
		}
	}
	// 退出时，我们将看下是不是将emp添加到链表最后
	pre.Next = emp
	emp.Next = cur
}

func (this * EmpLink) ShowAll(No int)  {
	if this.Head == nil {
		fmt.Printf("链表%d为空\n",No)
	}
	cur := this.Head //辅助指针
	for{
		if cur != nil {
			fmt.Printf("链表%d 雇员id=%d 名字=%s -->",No,cur.Id,cur.Name)
			cur = cur.Next
		}else{
			break
		}
	}
	fmt.Println()
}
func (this * EmpLink) FindById(id int)  {
	cur := this.Head
	for{
		if cur != nil && cur.Id == id {
			// 找到了
			fmt.Printf("该雇员信息为:%d %s\n",cur.Id,cur.Name)
			break
		}else if cur == nil{
			// 找不到
			fmt.Println("找不到该雇员")
			break
		}
		cur = cur.Next
	}
}

type HashTable struct{
	LinkArr [7]EmpLink
}

// 给HashTable编写Insert雇员的方法
func (this * HashTable) Insert(emp *Emp)  {
	linkNo := this.HashFun(emp.Id)
	this.LinkArr[linkNo].Insert(emp)
}

// 获取要添加雇员的对应链表的下标
func (this *HashTable) HashFun(id int) int  {
	return id % 7 //得到的是对应的链表的下标
}

func (this *HashTable) ShowAll()  {
	for i := 0; i < len(this.LinkArr); i++ {
		this.LinkArr[i].ShowAll(i)
	}
}
func (this *HashTable) FindById(id int)  {
	linkNo := this.HashFun(id)
	this.LinkArr[linkNo].FindById(id)
}

func main()  {
	var key int
	var id int
	var name string
	var hashTable HashTable
	for {
		fmt.Println("------雇员系统菜单------")
		fmt.Println("------1 添加雇员------")
		fmt.Println("------2 显示雇员------")
		fmt.Println("------3 查找雇员------")
		fmt.Println("------4 退出系统------")
		fmt.Println("------请输入（1-4）------")
		fmt.Scanln(&key)

		switch key {
			case 1:
				fmt.Println("请输入Id:")
				fmt.Scanln(&id)
				fmt.Println("请输入Name:")
				fmt.Scanln(&name)
				emp := &Emp{
					Id :id,
					Name:name,
				}
				hashTable.Insert(emp)
			case 2:
				hashTable.ShowAll()
			case 3:
				fmt.Println("请输入Id:")
				fmt.Scanln(&id)
				hashTable.FindById(id)
			case 4:
				os.Exit(0)
			
		}
	}
	
}