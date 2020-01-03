package main

import(
	"fmt"
	"reflect"
)

type Cal struct{
	Num1 int
	Num2 int
}

func (c *Cal) GetSub(name string) string {
	res := fmt.Sprintf("%v完成了减法运行：%v",name,c.Num1 - c.Num2)
	return res

}

func main()  {
	cal := Cal{8, 3}
	// cal.GetSub("xiaoxiao")
	// rType := reflect.TypeOf(cal)
	rVal := reflect.ValueOf(&cal)
	num := rVal.Elem().NumField()
	fmt.Printf("字段个数：%v",num)

	// 设置num1和num2值
	
	// 调用方法
	var params []reflect.Value
	params = append(params,reflect.ValueOf("tom"))
	res := rVal.Method(0).Call(params)
	fmt.Printf("res = %v",res)



}