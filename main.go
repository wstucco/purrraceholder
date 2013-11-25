package main

import (
	"github.com/pilu/traffic"
)

var router *traffic.Router

func main() {
	router = traffic.New()

	router.Get("/", RootHandler)

	// match (width)x(height) format
	// we cannot choose which character set the named routes can take
	// they only stop matching on / # ? ( ) . \
	// we should be able to do something like
	// (:width=[\d+])(x(:height=[\d+]))?
	// or something even simplier like
	// (:width=:digits)(x(:height=:digits))(.(:format=json|xml|atom))
	router.Get(`/(?P<width>\d+)x(?P<height>\d+)?/?`, ImageHandler).
		AddBeforeFilter(RequireValidImageParameters).
		AddBeforeFilter(GenerateImageCache)

	router.Get(`/:width/?(:height)/?`, ImageHandler).
		AddBeforeFilter(RequireValidImageParameters).
		AddBeforeFilter(GenerateImageCache)

	// Executed before all handlers
	router.AddBeforeFilter(PoweredByHandler)

	// Custom not found handler
	router.NotFoundHandler = NotFoundHandler

	router.Run()
}
