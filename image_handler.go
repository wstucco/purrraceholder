package main

import (
	"fmt"
	"github.com/pilu/traffic"
	"io/ioutil"
	"net/http"
	"strconv"
)

func ImageHandler(w traffic.ResponseWriter, r *traffic.Request) {

	w.Header().Set("Content-Type", "image/jpeg")
	image, err := ioutil.ReadFile("assets/images/grumpy.jpg")
	if err != nil {
		// panic is trapped by Traffic and show us a nice stack trace in the browser
		// a proper error handling should be provided, but in this simple example
		// it's used to remind you to always check for errors
		panic(err)
	}

	w.Write(image)
}

func RequireValidImageParameters(w traffic.ResponseWriter, r *traffic.Request) {

	width, err := strconv.Atoi(r.Param("width"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	height, err := strconv.Atoi(r.Param("height"))
	if err != nil {
		height = width
	}

	if width > 2560 || height > 2560 {
		w.WriteHeader(http.StatusBadRequest)
		w.RenderTemplate("400", nil)
	} else {

		w.SetVar("width", width)
		w.SetVar("height", height)

		// log latest greatest creation
		err = ioutil.WriteFile("latest", []byte(fmt.Sprintf("%d/%d", width, height)), 0644)
		if err != nil {
			// panic is trapped by Traffic and show us a nice stack trace in the browser
			// a proper error handling should be provided, but in this simple example
			// it's used to remind you to always check for errors
			panic(err)
		}

	}

}
