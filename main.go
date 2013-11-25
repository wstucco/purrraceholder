package main

import (
	"github.com/pilu/traffic"
)

var router *traffic.Router

func main() {
	router = traffic.New()

	router.Get("/", RootHandler)

	router.Get(`/:width/?(:height/?)?`, ImageHandler).
		AddBeforeFilter(RequireValidImageParameters)

	// Custom not found handler
	router.NotFoundHandler = NotFoundHandler

	router.Run()
}
