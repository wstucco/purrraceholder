package main

import (
	"github.com/wstucco/traffic"
)

func NotFoundHandler(w traffic.ResponseWriter, r *traffic.Request) {
	w.Render("404")
}
