package main

import (
	_ "testgo/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"testgo/models"
	_ "github.com/go-sql-driver/mysql"
)

func init(){
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname :=beego.AppConfig.String("dbname")
	dsn := dbuser + ":" +dbpassword +"@tcp("+dbhost+":"+dbport+")/"+dbname+"?charset=utf8&loc=Asia%2FShanghai"
	orm.RegisterDataBase("default","mysql",dsn)
	// 注册model模型
	orm.RegisterModel(new(models.User))
	//调用 RunCommand 执行 orm 命令。
	orm.RunCommand()

}




func main() {
	beego.Run()
}

