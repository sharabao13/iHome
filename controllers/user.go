package controllers

import (
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"iHome/models"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) RetData(resp map[string]interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}
func (c *UserController) Reg() {
	resp := make(map[string]interface{})
	defer c.RetData(resp)
	//获取前端输入的jsonshju
	json.Unmarshal(c.Ctx.Input.RequestBody, &resp)

	/* 调试查询数据
	logs.Info(`resp["mobile"]=`, resp["mobile"])
	logs.Info(`resp["password"]=`, resp["password"])
	logs.Info(`resp["sms_code"]=`, resp["sms_code"])
	*/
	//插入数据库
	o := orm.NewOrm()
	user := models.User{}
	user.Password_hash = resp["password"].(string) //断言
	user.Name = resp["mobile"].(string)
	user.Mobile = resp["mobile"].(string)

	id, err := o.Insert(&user)
	if err != nil {
		resp["errno"] = models.RECODE_NODATA
		resp["errmsg"] = models.RecodeText(models.RECODE_NODATA)
		return
	}
	logs.Info("注册成功,id: ", id)
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	c.SetSession("name", user.Name)

}
