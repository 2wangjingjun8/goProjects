package cal
import (
	_"fmt"
	"testing"
)

func TestSum(t *testing.T){
	res := addSum(10)
	if res != 55 {
		t.Fatalf("addSum(10)执行错误, 期望值:%v,实际值：%v", 55, res)
	}
	// fmt.Println(res)
}

func TestSub (t *testing.T){
	res := sub(10, 3)
	if res != 7 {
		t.Fatalf("sub(10, 3)执行错误, 期望值:%v,实际值：%v", 7, res)
	}
}
