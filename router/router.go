package router

import (
	"gf-app/app/api/content_list"
	"gf-app/app/api/filemanager"
	"gf-app/app/api/user"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(Middleware)
		group.ALL("/",user.Index)
		group.ALL("/fileManager/batch", filemanager.Upload)
		group.POST("/user/login",user.Login)
	})
	//s.Group("/template", func(group *ghttp.RouterGroup) {
	//
	//	group.ALL("/", filemanager.Upload)
	//})


	s.Group("/video", func(group *ghttp.RouterGroup) {
		group.ALL("/getByLabelId", content_list.GetByLabelId)
		group.ALL("/GetByContentId", content_list.GetByContentId)
	})
}
func Middleware(r *ghttp.Request) {
	// 中间件处理逻辑
	loginState := r.Session.GetBool("loginState")
	if !loginState{
		if  r.RequestURI != "/user/login"{
				error:= r.Response.WriteTpl("login.html")
				//r.Response.WriteTpl("main.html")
				if error != nil{
					glog.Error("输出html的时候发生错误，",error)
				}
				return
		}

	}

	r.Middleware.Next()
}