package main

import (
		"github.com/pilu/traffic"
		"os"
)

var router *traffic.Router

func main() {
	traffic.Logger().Printf("Port: %s\n", os.Getenv("PORT"))

	router = traffic.New()

	router.Get("/", RootHandler)

	// match (width)x(height) format
	// we cannot choose which character set the named routes can take
	// they only stop matching on / # ? ( ) . \
	// we should be able to do something like
	// (:width=[\d+])(x(:height=[\d+]))?
	// or something even simplier like
	// (:width=:digits)(x(:height=:digits))(.(:format=json|xml|atom))
	router.Get(`/(?P<width>\d+)(x(?P<height>\d+)?)?/?`, ImageHandler).
		AddBeforeFilter(RequireValidImageParameters).
		AddBeforeFilter(GenerateImageCache)

	router.Get(`/:width/?(:height)?/?`, ImageHandler).
		AddBeforeFilter(RequireValidImageParameters).
		AddBeforeFilter(GenerateImageCache)

	// Executed before all handlers
	router.AddBeforeFilter(PoweredByHandler)

	// Custom not found handler
	router.NotFoundHandler = NotFoundHandler

	// if not in development, add the static handler
	if traffic.Env() == "production" {
	  router.Use(traffic.NewStaticMiddleware(traffic.PublicPath()))
	}	

	router.Run()

}
