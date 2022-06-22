package controllers

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"iHome/models"
)

type AreaController struct {
	beego.Controller
}

func (this *AreaController) RetData(resp map[string]interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}
func (c *AreaController) GetArea() {
	logs.Info("connect success")
	resp := make(map[string]interface{})
	defer c.RetData(resp)
	//从session获取数据
	//从数据库获取area数据
	var areas []models.Area
	o := orm.NewOrm()
	qs := o.QueryTable("area")
	num, err := qs.All(&areas)
	if err != nil {
		logs.Info("数据错误")
		resp["errno"] = 4001
		resp["errmsg"] = "查询数据错误"
		return
	}
	if num == 0 {
		resp["errno"] = 4002
		resp["errmsg"] = "没有查到数据"
		return
	}
	resp["errno"] = 0
	resp["errmsg"] = "OK"
	resp["data"] = areas
	//封装成json返回前端
	//logs.Info("query data success,resp=", resp)

}
