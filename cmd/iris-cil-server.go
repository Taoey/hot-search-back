package main

import (
	"github.com/Taoey/iris-cli/pkg/api"
	"github.com/Taoey/iris-cli/pkg/sysinit"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"log"
)

var App *iris.Application

//程序入口
func main() {
	// 初始化
	sysinit.InitConf()
	//myinit.InitMongo()
	//myinit.InitQuartz()

	// 初始化App
	App = iris.New()
	SetRoutes()

	// 启动
	run := App.Run(iris.Addr(sysinit.GCF.UString("server.url")), iris.WithCharset("UTF-8"))
	log.Fatal(run)
}

// 设置路由
func SetRoutes() {

	// 跨域设置
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
	})

	//主页
	App.Get("/", api.ApiIndex)

	//根API
	RootApi := App.Party("api/v1", crs).AllowMethods(iris.MethodOptions)

	RootApi.Get("/test/index", api.ApiIndexJsonTest)
	// upload
	RootApi.Post("/upload/ali_bill", iris.LimitRequestBodySize(5<<20), api.UploadAliBill)

}
