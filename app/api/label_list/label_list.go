package label_list

import (
	"gf-app/app/service/label_list"
	"gf-app/utils"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
)


var imp = label_list.StuLabelServiceImp{}

func All(r *ghttp.Request) {
	entities, e := imp.All()
	if e!=nil {
		r.Response.Write(g.Map{"code":500})
		return
	}
	if entities == nil{
		r.Response.Write(g.Map{"code":200,"result":"null"})
		return
	}
	result, e := utils.EntityToJson(entities)
	if e!=nil {
		r.Response.Write(g.Map{"code":500})
		glog.Error("entity to string Error occurred")
		return
	}
	r.Response.Write(g.Map{"code":200,"result":result})
}
