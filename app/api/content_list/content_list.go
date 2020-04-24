package content_list

import (
	"gf-app/app/service/content_list"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var imp = content_list.StuContentListServiceImp{}

func GetByLabelId(r *ghttp.Request) {
	id := r.GetString("LabelId")
	pageNumber := r.GetString("pageNumber")
	if id == ""{
		r.Response.Write(g.Map{"code":500,"result":"LabelId cannot be empty"})
		return
	}
	entities, e := imp.GetByLabelId(id,pageNumber)
	if e!=nil {
		r.Response.Write(g.Map{"code":500})
		return
	}
	if entities == nil{
		r.Response.Write(g.Map{"code":200,"result":"null"})
		return
	}
	decode, e := gjson.Encode(entities)
	if e!=nil {
		r.Response.Write(g.Map{"code":500,"result":"json Encode Error occurred"})
		return
	}
	result, e := gjson.DecodeToJson(decode)
	if e!=nil {
		r.Response.Write(g.Map{"code":500,"result":"json to string Error occurred"})
		return
	}
	r.Response.Write(g.Map{"code":200,"result":result})
}
func GetByContentId(r *ghttp.Request) {
	id := r.GetString("ContentId")

	if id == ""{
		r.Response.Write(g.Map{"code":500,"result":"ContentId cannot be empty"})
		return
	}
	entities, e := imp.GetByContentId(id)
	if e!=nil {
		r.Response.Write(g.Map{"code":500})
		return
	}
	if entities == nil{
		r.Response.Write(g.Map{"code":200,"result":"null"})
		return
	}
	decode, e := gjson.Encode(entities)
	if e!=nil {
		r.Response.Write(g.Map{"code":500,"result":"json Encode Error occurred"})
		return
	}
	result, e := gjson.DecodeToJson(decode)
	if e!=nil {
		r.Response.Write(g.Map{"code":500,"result":"json to string Error occurred"})
		return
	}
	r.Response.Write(g.Map{"code":200,"result":result})
}