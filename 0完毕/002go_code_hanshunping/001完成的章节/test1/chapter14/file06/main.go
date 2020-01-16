package main

import(
	"fmt"
	"io/ioutil"
	"os"
)

func pathExist(path string) (bool,error){
	_,err := os.Stat(path)
	if err == nil {
		return true,nil
	}
	if os.IsNotExist(err) {
		return false,nil
	}
	return false,err

}
func main() {
	file,err := ioutil.ReadFile("./test1.txt")
	if err != nil {
		fmt.Println("err = ",err)
	}
	path := "./test3.txt"
	b,_ := pathExist(path) 
	if false == b {
		fmt.Println("文件不存在")
		os.Create(path)
		fmt.Println("创建文件成功")
		
	}
	err = ioutil.WriteFile(path,file,0666)
	if err != nil {
		fmt.Println("err = ",err)
	}

	fmt.Println(string(file))
}