// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego/context"
	"hello/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	//基本GET路由
	beego.Get("/get", func(ctx *context.Context) {
		ctx.Output.Body([]byte("hello world get"))
	})

	//基本POST路由
	beego.Post("/post", func(ctx *context.Context) {
		ctx.Output.Body([]byte("hello beego post"))
	})

	//Any
	beego.Any("/any", func(ctx *context.Context) {
		ctx.Output.Body([]byte("hello beego any"))
	})

	//ALL METHOD
	//beego.Get(router, beego.FilterFunc)
	//beego.Post(router, beego.FilterFunc)
	//beego.Put(router, beego.FilterFunc)
	//beego.Patch(router, beego.FilterFunc)
	//beego.Head(router, beego.FilterFunc) //head请求得不到body ...
	//beego.Options(router, beego.FilterFunc)
	//beego.Delete(router, beego.FilterFunc)
	//beego.Any(router, beego.FilterFunc

	//注解路由
	// 1 cms.go mapping
	// 2 方法注解
	// 3 路由include
	beego.Include(&controllers.CMSController{})

	//RESTful Controller

	//自动路由
	beego.AutoRouter(&controllers.AdminController{})

}
