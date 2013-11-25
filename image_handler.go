package main

import (
	"fmt"
	"github.com/nfnt/resize"
	"github.com/pilu/traffic"
	"image"
	"image/draw"
	"image/jpeg"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

const image_file string = "assets/images/grumpy.jpg"

func ImageHandler(w traffic.ResponseWriter, r *traffic.Request) {

	file, err := os.Open(image_file)
	if err != nil {
		// panic is trapped by Traffic and show us a nice stack trace in the browser
		// a proper error handling should be provided, but in this simple example
		// it's used to remind you to always check for errors
		panic(err)
	}
	defer file.Close()

	// decode jpeg into image.Image
	src_image, err := jpeg.Decode(file)
	if err != nil {
		panic(err)
	}

	width := w.GetVar("width").(int)
	height := w.GetVar("height").(int)

	dst_image := image.NewRGBA(image.Rect(0, 0, width, height))

	if width > height {
		grumpy := resize.Resize(0, uint(height), src_image, resize.Lanczos3)
		b := grumpy.Bounds()

		for i := 0; i < width; i = i + grumpy.Bounds().Dx() {
			r := image.Rect(b.Min.X+i, b.Min.Y, b.Max.X+i, b.Max.Y)
			draw.Draw(dst_image, r, grumpy, image.ZP, draw.Src)
		}
	} else {
		grumpy := resize.Resize(uint(width), 0, src_image, resize.Lanczos3)
		b := grumpy.Bounds()

		for i := 0; i < height; i = i + grumpy.Bounds().Dy() {
			r := image.Rect(b.Min.X, b.Min.Y+i, b.Max.X, b.Max.Y+i)
			draw.Draw(dst_image, r, grumpy, image.ZP, draw.Src)
		}
	}

	w.Header().Set("Content-Type", "image/jpeg")

	jpeg.Encode(w, dst_image, &jpeg.Options{jpeg.DefaultQuality})
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

	traffic.Logger().Printf("%dx%d\n", width, height)
	if (width <= 2560 && width > 0) && (height <= 2560 && height > 0) {

		w.SetVar("width", width)
		w.SetVar("height", height)

		// log latest greatest creation
		err = ioutil.WriteFile("cache/latest", []byte(fmt.Sprintf("%d/%d", width, height)), 0644)
		if err != nil {
			// panic is trapped by Traffic and show us a nice stack trace in the browser
			// a proper error handling should be provided, but in this simple example
			// it's used to remind you to always check for errors
			panic(err)
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.RenderTemplate("400", nil)
	}

}
