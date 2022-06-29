package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"iHome/models"
	"path"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) RetData(resp interface{}) {
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

func (c *UserController) Avatar() {
	resp := SessionResp{Errno: models.RECODE_OK, Errmsg: models.RecodeText(models.RECODE_OK)}
	defer c.RetData(&resp)
	fileData, fileHandler, err := c.GetFile("avatar")
	if err != nil {
		resp.Errno = models.RECODE_REQERR
		resp.Errmsg = models.RecodeText(models.RECODE_REQERR)
		return
	}
	fmt.Println(fileData)
	//2. 得到文件后缀
	suffix := path.Ext(fileHandler.Filename)
	fmt.Println(suffix)

	//3. 存储文件到fastdfs
}
