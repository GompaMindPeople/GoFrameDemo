package user

import (
	"gf-app/app/service/user"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
)


var service = user.StructUserServer{}

func Index(r *ghttp.Request) {

	error:= r.Response.WriteTpl("index.html")
	//r.Response.WriteTpl("main.html")
	if error != nil{
		glog.Error("输出html的时候发生错误，",error)
		return
	}
}

func Login(r *ghttp.Request) {
	account := r.GetString("account")
	password := r.GetString("password")

	entity, e := service.Login(account, password)

	if e != nil{
		glog.Error(e.Error())
		r.Response.Write(g.Map{"code":500})
		return
	}
	e = r.Session.Set("loginState", true)
	if e != nil{
		glog.Error(e.Error())
		r.Response.Write(g.Map{"code":500})
		return
	}
	e = r.Session.Set("userModel", entity)
	if e != nil{
		glog.Error(e.Error())
		r.Response.Write(g.Map{"code":500})
		return
	}
	r.Response.Write(g.Map{"code":200})
}




func Register(r *ghttp.Request) {

	r.Response.Write(g.Map{"code":200})
}