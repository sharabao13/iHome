package main

import (
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	_ "iHome/models"
	_ "iHome/routers"
	"net/http"
	"strings"
)

func main() {
	//fileId := models.UploadFile("testDemo.txt")
	//fmt.Println(fileId)
	//models.FdfsUploadByFileName("testDemo.txt")
	ignoreStaticPath()
	beego.Run()
}

func ignoreStaticPath() {
	//透明static
	beego.InsertFilter("/", beego.BeforeRouter, TransparentStatic)
	beego.InsertFilter("/*", beego.BeforeRouter, TransparentStatic)
}

func TransparentStatic(ctx *context.Context) {
	orpath := ctx.Request.URL.Path
	logs.Debug("request url: ", orpath)
	//如果请求url还有api字段，说明是指令应该取消静态资源路径重定向
	if strings.Index(orpath, "api") >= 0 {
		return
	}
	http.ServeFile(ctx.ResponseWriter, ctx.Request, "static/html/"+ctx.Request.URL.Path)
}
