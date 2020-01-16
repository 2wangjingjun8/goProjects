package main
import(
	"fmt"
	_"net/http"
	_"io/ioutil"
)

func main()  {
	// resp,err := http.Get("http://www.ice20.top")
	// if err != nil {
	// 	fmt.Println("http.Get err = ",err)
	// 	return
	// }
	// defer resp.Body.Close()
	// // fmt.Println("resp=",resp)
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("ioutil.ReadAll err = ",err)
	// 	return
	// }
	// fmt.Println("body=",string(body))
	aa := "\/pages\/goodsDetails\/goodsDetails?id=1111002449&d=miallgz.miallnb&s=misall.top&b=Mide12&c=\u5e7f\u5dde\u79d8\u5965\u670d\u88c5\u6279\u53d1"
	fmt.Println(aa)
}