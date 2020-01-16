package main
import(
	"fmt"
	"os"
	_"bufio"
	"io"
)
func copyFile(dstFileName string,srcFileName string) (written int64, err error) {
	srcFile,err := os.Open(srcFileName)
	defer srcFile.Close()
	if err != nil {
		fmt.Println("open srcFile err = ",err)
	}
	// reader := bufio.NewReader(srcFile)
	
	dstFile,err := os.OpenFile(dstFileName,os.O_WRONLY | os.O_CREATE,0666)
	defer dstFile.Close()
	if err != nil {
		fmt.Println("open dstFile err = ",err)
	}
	// writer := bufio.NewWriter(dstFile)
	return io.Copy(dstFile, srcFile)
}

func main()  {
	dstFile := "g:/bb.png"
	srcFile := "g:/mm.jpg"
	_,err := copyFile(dstFile,srcFile)
	
	if err != nil {
		fmt.Println("拷贝失败")
	}else{
		fmt.Println("拷贝完成")
	}
}