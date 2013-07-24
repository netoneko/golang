package main

import (
	"bytes"
	"github.com/hoisie/web"
	//	"html/template"
	"io"
)

func hello(val string) string { return "hello " + val }
func index(ctx *web.Context) {
	var buf bytes.Buffer
	buf.WriteString("hello")
	//copy buf directly into the HTTP response
	io.Copy(ctx, &buf)
}

func main() {
	web.Get("/", index)
	web.Get("/(.*)", hello)

	web.Run("0.0.0.0:9999")
}
