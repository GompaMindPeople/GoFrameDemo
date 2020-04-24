package main

import (
	_ "gf-app/boot"
	_ "gf-app/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	s := g.Server()
	//s.EnableHTTPS("/home/https/Nginx/1_xzmtest.xyz_bundle.crt", "/home/https/Nginx/2_xzmtest.xyz.key")
	//s.SetHTTPSPort(443)
	//s.SetPort(80)
	s.Run()
}
