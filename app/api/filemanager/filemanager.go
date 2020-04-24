package filemanager

import (
	"gf-app/app/model/content_list"
	content_list2 "gf-app/app/service/content_list"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/gtime"
	"io"
)

var imp = content_list2.StuContentListServiceImp{}

// Upload uploads files to /tmp .
func Upload(r *ghttp.Request) {
	saveDir := "./tmp/"
	rtId := r.GetString("RtId")
	content := r.GetString("Content")
	for _, item := range r.GetMultipartFiles("file") {
		file, err := item.Open()
		if err != nil {
			r.Response.Write(err)
			return
		}
		defer func() {
			if file!= nil{
				file.Close()
			}

		}()
		f, err := gfile.Create(saveDir + gfile.Basename(item.Filename))
		if err != nil {
			r.Response.Write(err)
			return
		}
		if rtId == ""{
			rtId = "1"
		}
		if content == ""{
			content = "默认标题"
		}
		entity := content_list.Entity{}
		entity.RtId = rtId
		entity.VideoUrl = f.Name()
		entity.CreateTime = gtime.Now()
		entity.Content = content
		defer func() {
			if f!= nil{
				 f.Close()
			}
		}()

		if _, err := io.Copy(f, file); err != nil {
			r.Response.Write(err)
			return
		}
		err = imp.InsertLabel(&entity)
	}
	r.Response.Write("upload successfully")
}