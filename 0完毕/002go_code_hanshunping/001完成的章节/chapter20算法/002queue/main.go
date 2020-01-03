package main

import(
	"fmt"
	"os"
	"errors"
)

type Queue struct{
	MaxSize int
	Array [5]int
	Front int
	Rear int
}


func (this *Queue) SetQueue(num int) (err error) {
	if this.Rear == this.MaxSize-1 {
		// fmt.Println("queue full")
		return errors.New("queue full")
	}
	this.Rear++
	this.Array[this.Rear] = num
	return
}

func (this *Queue) GetQueue() (val int, err error) {
	if this.Rear == this.Front {
		// fmt.Println("queue null")
		return -1,errors.New("queue null")
	}

	this.Front++
	val = this.Array[this.Front]
	return val,err

}
func (this *Queue) ShowQueue()  {
	fmt.Println("队列当前的情况是：")
	if this.Rear == this.Front {
		fmt.Println("queue null")
	}
	for i := this.Front+1; i <= this.Rear; i++ {
		fmt.Printf("Array[%v] = %v\t",i,this.Array[i])
	}
	fmt.Println()
}

func main()  {
	queue := &Queue{
		MaxSize:5,
		Front:-1,
		Rear:-1,
	}

	var key string
	var num int
	for{
		fmt.Println("---1. 查看当前队列---")
		fmt.Println("---2. 加入一个数字到队列---")
		fmt.Println("---3. 从队列中取走一个数字---")
		fmt.Println("---4. 退出---")
		fmt.Println("---请选择（1-4）---")
		fmt.Scanln(&key)
		switch key {
			case "1":
				queue.ShowQueue()
			case "2":
				fmt.Println("---请输入数字：---")
				fmt.Scanln(&num)
				err := queue.SetQueue(num)
				if err != nil {
					fmt.Println("err = ",err)
				}else{
					fmt.Println("加入队列成功")
				}
			case "3":
				val,err := queue.GetQueue()
				if err != nil {
					fmt.Println("err = ",err)
				}
				fmt.Println("当前取出的数字是：",val)
			case "4":
				os.Exit(0)
		}
	}
}