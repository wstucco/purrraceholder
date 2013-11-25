package main

import (
	"github.com/pilu/traffic"
)

func NotFoundHandler(w traffic.ResponseWriter, r *traffic.Request) {
	w.RenderTemplate("404", nil)
}
