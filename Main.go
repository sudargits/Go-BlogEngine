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
	api.Get("/mypath", func(ctx *iris.Context) {
		ctx.Write("Hello from the server on path /mypath")
	})
	api.Get("/blog",allPost)

	api.Build()
	fsrv := &fasthttp.Server{Handler: api.Router}
	fsrv.ListenAndServe(":8080")
}
