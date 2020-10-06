package main

import (
	. "github.com/Taoey/hot-search-back/pkg"
	"github.com/Taoey/hot-search-back/pkg/api"
	. "github.com/Taoey/hot-search-back/pkg/sysinit"
	_ "github.com/go-sql-driver/mysql"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/pprof"
	"log"
)

var App *iris.Application

//程序入口
func main() {
	// 初始化
	InitConf()

	InitLogger()
	//InitMongo()
	InitQuartz()
	InitMysql()

	// 初始化App
	App = iris.New()
	SetRoutes()
	SetWebStatic()

	// 启动
	run := App.Run(iris.Addr(GCF.UString("server.url")), iris.WithCharset("UTF-8"))
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
	//App.Get("/", api.ApiIndex)

	App.Get("/aa", func(ctx iris.Context) {
		ctx.HTML("<h1> Please click <a href='/debug/pprof'>here</a>")
	})

	p := pprof.New()
	App.Any("/debug/pprof", p)
	App.Any("/debug/pprof/{action:path}", p)

	//根API
	RootApi := App.Party("api/v1", crs).AllowMethods(iris.MethodOptions)

	RootApi.Get("/test/index", api.ApiIndexJsonTest)
	// upload
	RootApi.Post("/upload/ali_bill", iris.LimitRequestBodySize(5<<20), api.UploadAliBill)

	// 知乎
	RootApi.Post("/zhihu/query", api.ApiZhihuQuery) // 查询
}

// 设置web资源
func SetWebStatic() {
	App.Favicon("./web/favicon.ico")
	App.HandleDir("/css", "./web/css")
	App.HandleDir("/js", "./web/js")

	App.HandleDir("/", "./web")
}
