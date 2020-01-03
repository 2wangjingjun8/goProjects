package main

import(
	"fmt"
	"errors"
)

type Stack struct{
	MaxTop int // 栈最大可以存放数的个数
	Top int // 表示栈顶，因为栈顶固定，因此直接用Top
	arr [5]int // 数组模拟栈
}

func (this *Stack) Push(val int) (err error) {
	// 先判断栈是否满了
	if this.Top == this.MaxTop -1 {
		fmt.Println("stack full")
		return errors.New("stack full")
	}
	this.Top ++
	// 放入数据
	this.arr[this.Top] = val

	return
}
func (this *Stack) Pop() (val int, err error) {
	if this.Top == -1 {
		fmt.Println("stack empty")
		return 0,errors.New("stack empty")
	}
	val = this.arr[this.Top]
	this.Top--
	return val, nil
}

// 遍历栈，注意需要从栈顶开始遍历
func (this *Stack) list()  {
	// 先判断栈是否为空
	if this.Top == -1 {
		fmt.Println("stack empty")
		return
	}
	fmt.Println("栈的情况如下：")
	for i := this.Top; i >=0; i-- {
		fmt.Printf("arr[%d] = %d\n", i, this.arr[i])
	}
}

func main()  {
	stack := &Stack{
		MaxTop:5,
		Top:-1, //当栈顶为-1时，表示栈为空
	}
	//入栈
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.list()
	val,_ := stack.Pop()
	fmt.Println("出栈：",val)
	
	stack.list()
}