package main

import (
	"bytes"
	"encoding/json"
	"github.com/hoisie/web"
	"io"
)

type Response struct {
	Hello string `json:"hello"`
}

func hello(ctx *web.Context, val string) {
	var buf bytes.Buffer

	result, _ := json.Marshal(&Response{Hello: val})

	buf.Write(result)
	//copy buf directly into the HTTP response
	io.Copy(ctx, &buf)
}

func main() {
	web.Get("/(.*)", hello)
	web.Run("0.0.0.0:9999")
}
