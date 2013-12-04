package main

import (
	"github.com/wstucco/traffic"
	"io/ioutil"
)

type ResponseData struct {
	ImageUrl string
}

func RootHandler(w traffic.ResponseWriter, r *traffic.Request) {
	last_image_generated, err := ioutil.ReadFile("cache/latest")

	if err != nil {
		responseData := &ResponseData{string(last_image_generated)}
		w.Render("index", responseData)
	}
}
