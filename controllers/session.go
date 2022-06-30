package controllers

import (
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"iHome/models"
)

type SessionController struct {
	beego.Controller
}

type Name struct {
	Name string `json:"name"`
}

type SessionResp struct {
	Errno  string      `json:"errno"`
	Errmsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}

func (this *SessionController) RetData(resp interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}
func (c *SessionController) GetSessionData() {
	resp := SessionResp{Errno: models.RECODE_OK, Errmsg: models.RecodeText(models.RECODE_OK)}
	defer c.RetData(&resp)

	name := c.GetSession("name")
	if name == nil {
		resp.Errno = models.RECODE_SESSIONERR
		resp.Errno = models.RecodeText(resp.Errno)
		return
	}

	NameData := Name{Name: name.(string)}
	resp.Data = NameData

	return

}

func (c *SessionController) DeleteSessionData() {
	resp := SessionResp{Errno: models.RECODE_OK, Errmsg: models.RecodeText(models.RECODE_OK)}
	defer c.RetData(&resp)

	err := c.DelSession("name")
	if err != nil {
		resp.Errno = models.RECODE_SERVERERR
		resp.Errmsg = models.RecodeText(resp.Errno)
	}

}

func (c *SessionController) Login() {
	//1. 获取用户信息
	resp := make(map[string]interface{})
	defer c.RetData(resp)
	//获取前端输入的jsonshju
	json.Unmarshal(c.Ctx.Input.RequestBody, &resp)
	logs.Info("==============name==========", resp["mobile"], "===============password==========", resp["password"])

	//2. 判断是否合法
	if resp["mobile"] == nil || resp["password"] == nil {
		resp["errno"] = models.RECODE_DATAERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DATAERR)
		logs.Info("==============name==========合法验证")
		return
	}
	//3. 数据库匹配
	o := orm.NewOrm()
	user := models.User{Name: resp["mobile"].(string)}
	qs := o.QueryTable("user")
	_, err := qs.Filter("mobile", resp["mobile"]).All(&user)

	if err != nil {
		resp["errno"] = models.RECODE_DATAERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DATAERR)
		logs.Info("==============name==========数据库匹配")

		return
	}
	if user.Password_hash != resp["password"] {
		resp["errno"] = models.RECODE_DATAERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DATAERR)
		logs.Info("==============name==========密码匹配")
		return
	}
	//4. 添加session
	c.SetSession("name", user.Name)
	//c.SetSession("mobile", resp["mobile"])
	c.SetSession("user_id", user.Id)
	//5. 返回json数据给前端

	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)

}
