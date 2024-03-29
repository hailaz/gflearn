package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func main() {
	s := g.Server()
	s.AddStaticPath("/static", "static")
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Write("helloworld")
	})
	// s.SetAddr("127.0.0.1:8080") // 看config.toml配置文件
	s.Run()
}

// http://127.0.0.1:8080/static/1.txt
