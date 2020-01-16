package main

import (
	"fmt"
	_"time"
	"errors"
)

func test() {
	defer func () {
		err := recover()
		if err != nil {
			fmt.Printf("向管理员发送错误信息：%v,%T\n",err,err)
		}

	}()
	num1 := 10
	num2 := 0
	res := num1/num2
	fmt.Println(res)
}

func readConf(name string ) (err error) {
	if name == "config.ini" {
		return nil
	}else{
		return errors.New("读取文件失败。。。")
	}
}

func test02() {
	err := readConf("config2.ini")
	if err != nil {
		panic(err)
	}
	fmt.Println("test02程序继续执行...")
}

func main() {
	test()
	// for{
	// 	fmt.Println("程序继续往下执行。。。")
	// 	time.Sleep(time.Second)
	// }
	test02()
	fmt.Println("main程序继续执行...")
	
}
