package controllers

import (
	"encoding/json"
	"github.com/beego/beego/v2/client/cache"
	_ "github.com/beego/beego/v2/client/cache/redis"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"golang.org/x/net/context"
	"iHome/models"
	"time"
)

type AreaController struct {
	beego.Controller
}

type AreaResp struct {
	Errno  string      `json:"errno"`
	Errmsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}

func (this *AreaController) RetData(resp interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (c *AreaController) GetArea() {
	logs.Info("connect success")
	resp := AreaResp{Errno: models.RECODE_OK, Errmsg: models.RecodeText(models.RECODE_OK)}
	//resp := make(map[string]interface{})
	defer c.RetData(&resp)
	//从session获取数据
	cache_conn, err := cache.NewCache("redis", `{"key":"iHome","conn":"120.48.15.88:16379","dbNum":"0","password":"Passw0rd@2022"}`)
	if err != nil {
		logs.Debug("connect redis server error")
		resp.Errno = models.RECODE_DATAERR
		resp.Errmsg = models.RecodeText(models.RECODE_DBERR)
		return
	}

	if areaSession, _ := cache_conn.Get(context.TODO(), "area"); areaSession != nil {
		logs.Info("======================= Get Session From Cache ==========================")
		//返回前端
		resp.Data = areaSession
		return
	} else {
		logs.Info("获取area缓存失败")
	}

	//从数据库获取area数据
	var areas []models.Area
	o := orm.NewOrm()
	qs := o.QueryTable("area")
	num, err := qs.All(&areas)
	if err != nil {
		logs.Info("数据错误")
		resp.Errno = models.RECODE_DBERR
		resp.Errmsg = models.RecodeText(models.RECODE_DBERR)
		return
	}
	if num == 0 {
		resp.Errno = models.RECODE_NODATA
		resp.Errmsg = models.RecodeText(models.RECODE_NODATA)
		return
	}
	resp.Errno = models.RECODE_OK
	resp.Errmsg = models.RecodeText(models.RECODE_OK)
	logs.Info("======================= Get Session From DataBase ==========================")
	resp.Data = areas
	//把数据转换成json存入redis
	json_str, err := json.Marshal(areas)
	if err != nil {
		logs.Info("封装json失败")
		return
	}
	cache_conn.Put(context.TODO(), "area", json_str, time.Second*1200)
	//封装成json返回前端
	//logs.Info("query data success,resp=", resp)
}
