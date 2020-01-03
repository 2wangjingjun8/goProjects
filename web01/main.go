package main

import (
	"fmt"
	_ "web01/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	driverName := beego.AppConfig.String("driverName")
	orm.RegisterDriver(driverName, orm.DRMySQL)

	user := beego.AppConfig.String("user")
	pwd := beego.AppConfig.String("pwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("dbname")
	// root:root@/orm_test?charset=utf8
	dbConn := "" + user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"
	err := orm.RegisterDataBase("default", driverName, dbConn)
	if err != nil {
		fmt.Println("连接数据库出错")
		return
	}
	fmt.Println("连接数据库成功")
}
func main() {
	beego.Run()
}
