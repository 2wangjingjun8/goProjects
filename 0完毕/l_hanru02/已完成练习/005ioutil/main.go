package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// filePath := "aa.txt"
	// data, err := ioutil.ReadFile(filePath)
	// if err != nil {
	// 	fmt.Println("err = ", err)
	// }
	// fmt.Println("dataByte = ", data)
	// fmt.Println("data = ", string(data))

	// str := "hello world 啊啊啊"
	// err = ioutil.WriteFile(filePath, []byte(str), 0777)
	// if err != nil {
	// 	fmt.Println("err = ", err)
	// }

	// ioutil.ReadAll
	// s1 := "eee，取消那个夏天博，豹猫浮绿水"
	// d := strings.NewReader(s1)
	// data, err := ioutil.ReadAll(d)
	// if err != nil {
	// 	fmt.Println("err = ", err)
	// }
	// fmt.Println(data)
	// fmt.Println(string(data))

	srcPath := "E://009-Go/goProjects/src/l_hanru"
	// fileInfos, err := ioutil.ReadDir(srcPath)
	// if err != nil {
	// 	fmt.Println("err = ", err)
	// 	return
	// }
	// for _, val := range fileInfos {
	// 	fmt.Printf("name = %v \t\t\t isDir = %v\n", val.Name(), val.IsDir())
	// }
	loopDir(srcPath, 0)

}

func loopDir(fileDir string, level int) {
	fileInfos, err := ioutil.ReadDir(fileDir)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	s := "|--"
	for i := 0; i < level; i++ {
		s = "|--" + s
	}
	for _, val := range fileInfos {
		filePath := fileDir + "/" + val.Name()
		fmt.Printf("%s %v \n", s, filePath)
		if val.IsDir() {
			loopDir(filePath, level+1)
		}
	}
}
