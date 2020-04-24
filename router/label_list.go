package router

import (

	"gf-app/app/api/label_list"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/labelList", func(group *ghttp.RouterGroup) {
		group.ALL("/all", label_list.All)
	})


}