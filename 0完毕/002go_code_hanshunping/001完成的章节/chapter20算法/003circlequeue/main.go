package main

import(
	"fmt"
	"os"
	"errors"
)

type CircleQueue struct{
	MaxSize int
	Array [5] int
	Head int
	Tail int
}

func (this * CircleQueue) Push(num int) (err error)  {
	if this.IsFull(){
		return errors.New("Queue is full!")
	}
	this.Array[this.Tail] = num
	this.Tail = (this.Tail + 1) % this.MaxSize
	return
}

func (this * CircleQueue) Pop() (val int, err error)  {
	if this.IsEmpty(){
		return 0,errors.New("Queue is empty!")
	}
	val = this.Array[this.Head]
	this.Head = (this.Head + 1) % this.MaxSize
	return 
}

func (this * CircleQueue) List()  {
	fmt.Println("环形队列如下：")
	size := this.Size()
	if size == 0{
		fmt.Println("队列为空")
	}

	// 设置一个辅助的变量，指向Head
	tempHead := this.Head
	for i := 0; i < size; i++ {
		fmt.Printf("Array[%d] = %d\t", tempHead, this.Array[tempHead])
		tempHead = (tempHead + 1) % this.MaxSize
	}
	fmt.Println()

}

func (this * CircleQueue) IsFull() bool {
	return (this.Tail + 1) % this.MaxSize == this.Head
}
func (this * CircleQueue) IsEmpty() bool {
	return this.Tail == this.Head
}
func (this * CircleQueue) Size() int {
	return (this.Tail - this.Head +this.MaxSize) % this.MaxSize
}
func main()  {
	queue := &CircleQueue{
		MaxSize:5,
		Head:0,
		Tail:0,
	}

	var key string
	var num int
	for{
		fmt.Println("---1. 查看队列---")
		fmt.Println("---2. 加入队列---")
		fmt.Println("---3. 弹出队列---")
		fmt.Println("---4. 退出---")
		fmt.Println("---请选择（1-4）---")
		fmt.Scanln(&key)
		switch key {
			case "1":
				queue.List()
			case "2":
				fmt.Println("---请输入数字：---")
				fmt.Scanln(&num)
				err := queue.Push(num)
				if err != nil {
					fmt.Println("err = ",err)
				}else{
					fmt.Println("加入队列成功")
				}
			case "3":
				val,err := queue.Pop()
				if err != nil {
					fmt.Println("err = ",err)
				}
				fmt.Println("当前取出的数字是：",val)
			case "4":
				os.Exit(0)
		}
	}
}