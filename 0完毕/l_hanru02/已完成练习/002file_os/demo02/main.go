package main

func main() {
	// 1. 路径
	// filePath1 := "./demo01/aa.txt"
	// filePath2 := "E://009-Go/goProjects/src/l_hanru/002file_os/demo01/aa.txt"
	// fmt.Println(filepath.IsAbs(filePath1))
	// fmt.Println(filepath.IsAbs(filePath2))

	// fmt.Println(filepath.Abs(filePath1))

	// fmt.Println(path.Join(filePath2, ".."))

	// 2.创建目录
	// err := os.Mkdir("bb", 0777)
	// if err != nil {
	// 	fmt.Println("err = ", err)
	// 	return
	// }
	// err := os.MkdirAll("bb/aa/cc", 0777)
	// if err != nil {
	// 	fmt.Println("err = ", err)
	// 	return
	// }
	// fmt.Println("文件夹创建成功")

	// 2.创建文件
	// fp, err := os.Create("aa.txt")
	// if err != nil {
	// 	fmt.Println("err = ", err)
	// 	return
	// }
	// fmt.Println("文件创建成功 fp = ", fp)

	// 3. 打开文件
	// fp, err := os.Open("aa.txt")
	// if err != nil {
	// 	fmt.Println("err = ", err)
	// 	return
	// }
	// fmt.Println(fp)
	// // fmt.Println(fp.Read())
	// fp.Close()

	// 删除文件
	// err = os.Remove("aa.txt")
	// if err != nil {
	// 	fmt.Println("err = ", err)
	// 	return
	// }

	// err := os.RemoveAll("bb")
	// if err != nil {
	// 	fmt.Println("err = ", err)
	// 	return
	// }
}
