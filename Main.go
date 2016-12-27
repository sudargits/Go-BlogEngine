package main

import (
	"github.com/kataras/iris"
	"github.com/valyala/fasthttp"
)

func main() {
	api := iris.New()
	api.Get("/", func(ctx *iris.Context) {
		ctx.Write("Hello from the server")
	})

	//route for post
	blog := api.Party("/blog")
	blog.Get("/all",allPost)
	blog.Get("/detail/:idblog",detailPost)



	api.Build()
	fsrv := &fasthttp.Server{Handler: api.Router}
	fsrv.ListenAndServe(":8080")
}
