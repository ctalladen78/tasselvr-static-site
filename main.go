package main

import (
	"fmt"
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func main() {
	router := fasthttprouter.New()
	router.GET("/", Index)
	router.GET("/test/testvar", Subpage)

	log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}

// show index page
func Index(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Welcome!\n")

}

// show test params
func Subpage(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "testing params, %s \n", ctx.UserValue("name"))

}
