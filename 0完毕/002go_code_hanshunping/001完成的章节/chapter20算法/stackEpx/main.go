package main

import(
	"fmt"
	"errors"
	"strconv"
)

type Stack struct{
	MaxTop int // 栈最大可以存放数的个数
	Top int // 表示栈顶，因为栈顶固定，因此直接用Top
	arr [20]int // 数组模拟栈
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
func (this *Stack) List()  {
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

// 判断一个字符是不是一个运算符【=-*/】
func (this *Stack) IsOPer(val int) bool {
	if val == 42 || val == 43 || val == 45|| val == 47{
		return true
	}else{
		return false
	}
}
func (this *Stack) Cal(num1 int,num2 int, oper int ) int {
	res := 0
	switch oper {
		case 42:
			res = num2 * num1
		case 43:
			res = num2 + num1
		case 45:
			res = num2 - num1
		case 47:
			res = num2 / num1
		default:
			fmt.Println("运算符错误")
	}
	return res
}

// 编写一个方法，返回某个运算符的优先级
// [*/ => 1  +- => 0 ]
func (this *Stack) Priority(oper int) int {
	res := 0
	if oper == 42 || oper == 47 {
		res = 1
	}else if oper == 43 || oper == 47{
		res = 0
	}
	return res

}

 
func main()  {
	// 符号栈
	operStack := &Stack{
		MaxTop:20,
		Top:-1, //当栈顶为-1时，表示栈为空
	}
	// 数栈
	numStack := &Stack{
		MaxTop:20,
		Top:-1, //当栈顶为-1时，表示栈为空
	}
	exp := "3+20*5-200"
	//定义一个index,帮助扫描exp
	index := 0
	num1 := 0
	num2 := 0
	oper := 0
	result := 0
	keepNum := ""
	for{
		ch := exp[index:index+1]
		temp := int([]byte(ch)[0])
		if operStack.IsOPer(temp) {
			// 如果是个空栈，直接入栈
			if operStack.Top == -1 {
				operStack.Push(temp)
			}else{
				// 判断当前要入栈的符号，和栈顶符号优先级；如果栈顶优先级较高，则先运算
				if operStack.Priority(operStack.arr[operStack.Top]) >= operStack.Priority(temp) {
					num1,_ = numStack.Pop()
					num2,_ = numStack.Pop()
					oper,_ = operStack.Pop()
					result = operStack.Cal(num1,num2,oper)
					//计算完，就把当前结果入栈，当前符号入栈
					numStack.Push(result)
					operStack.Push(temp)
				}else{
					operStack.Push(temp)
				}
			}
			
		}else{//说明是数
			keepNum += ch
			if index == len(exp) - 1 {
				val,_ := strconv.ParseInt(keepNum,10,64)
				numStack.Push(int(val))
			}else{
				// 向index后面测试看看是不是运算符
				if operStack.IsOPer(int([]byte(exp[index+1:index+2])[0])) {
					val,_ := strconv.ParseInt(keepNum,10,64)
					numStack.Push(int(val))
					keepNum = ""
				}
			}
			// val,_ := strconv.ParseInt(ch,10,64)
			// numStack.Push(int(val))
		}

		// 继续扫描
		// 先判断index是否扫描到当前计算表达式的最后
		if index + 1 == len(exp){
			break
		}
		index++

	}
	// 如果扫描完毕，依次从数栈符号栈取出数据进行运算
	for{
		if operStack.Top != -1 {
			num1,_ = numStack.Pop()
			num2,_ = numStack.Pop()
			oper,_ = operStack.Pop()
			result = operStack.Cal(num1,num2,oper)
			//计算完，就把当前结果入栈，当前符号入栈
			numStack.Push(result)
		}else{
			break
		}
	}

	// 如果我们的运算没有问题，表达式也是正确，其结果就是numStack中的最后数
	res,_ := numStack.Pop()
	fmt.Println("表达式结果res = ",res)
}