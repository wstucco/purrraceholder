package main

import (
	"github.com/pilu/traffic"
	"io/ioutil"
)

type ResponseData struct {
	ImageUrl string
}

func RootHandler(w traffic.ResponseWriter, r *traffic.Request) {
	lastImageGenerates, err := ioutil.ReadFile("latest")
	if err != nil {
		panic(err)
	}

	responseData := &ResponseData{string(lastImageGenerates)}
	w.RenderTemplate("index", responseData)
}
