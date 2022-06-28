package main

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"iHome/models"
	_ "iHome/models"
	_ "iHome/routers"
	"log"
	"net/http"
	"strings"
)

func main() {
	name, id, err := models.FdfsUploadByFileName("testDemo.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name, id)
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
