package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fmt"
	"reflect"
	)

type TestController struct {
	beego.Controller
}


func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}

func (c *TestController) Get() {
	orm.Debug=true
	o :=orm.NewOrm()

	var maps []orm.Params
	fmt.Printf("%p", &maps)
	num, err := o.QueryTable("user").Values(&maps)
	fmt.Println(num)
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num)
		for _, m := range maps {
			fmt.Println(m["Id"], m["Name"])
		}
	}

	fmt.Printf("%p", &maps)

	c.Data["json"] =  maps
	c.ServeJSON()



}
