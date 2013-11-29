package main

import (
	"github.com/pilu/traffic"
	"io/ioutil"
)

type ResponseData struct {
	ImageUrl string
}

func RootHandler(w traffic.ResponseWriter, r *traffic.Request) {
	last_image_generated, _ := ioutil.ReadFile("cache/latest")

	responseData := &ResponseData{string(last_image_generated)}
	w.Render("index", responseData)
}
