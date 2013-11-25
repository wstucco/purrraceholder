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
const cache_folder string = "cache"

func ImageHandler(w traffic.ResponseWriter, r *traffic.Request) {
	width := w.GetVar("width").(int)
	height := w.GetVar("height").(int)

	src_image := loadImageFromFile(image_file)
	pattern := resizeImage(src_image, width, height)

	var dst_image image.Image
	if width > height {
		dst_image = tileImageHorizontally(pattern, width, height)
	} else {
		dst_image = tileImageVertically(pattern, width, height)
	}

	// output the image with the correct content-type
	w.Header().Set("Content-Type", "image/jpeg")
	jpeg.Encode(w, dst_image, &jpeg.Options{jpeg.DefaultQuality})
}

func RequireValidImageParameters(w traffic.ResponseWriter, r *traffic.Request) {

	width, err := strconv.Atoi(r.Param("width"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	height, err := strconv.Atoi(r.Param("height"))
	if err != nil {
		height = width
	}

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
		// bad request
		w.WriteHeader(http.StatusBadRequest)
		w.RenderTemplate("400", nil)
	}

}

func GenerateImageCache(w traffic.ResponseWriter, r *traffic.Request) {

	filename := fmt.Sprintf("%s/%dx%d.jpg", cache_folder, w.GetVar("width"), w.GetVar("height"))
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// file does not exists, generate a cached version
		// src_image := loadImageFromFile(image_file)

	}
}

func newImage(width int, height int) draw.Image {
	return image.NewRGBA(image.Rect(0, 0, width, height))
}

func loadImageFromFile(filename string) image.Image {
	file, err := os.Open(filename)
	if err != nil {
		// panic is trapped by Traffic and show us a nice stack trace in the browser
		// a proper error handling should be provided, but in this simple example
		// it's used to remind you to always check for errors
		panic(err)
	}
	defer file.Close()

	// decode jpeg into image.Image
	image, err := jpeg.Decode(file)
	if err != nil {
		panic(err)
	}

	return image
}

func resizeImage(im image.Image, width int, height int) image.Image {
	if width > height {
		return resize.Resize(0, uint(height), im, resize.Lanczos3)
	}

	return resize.Resize(uint(width), 0, im, resize.Lanczos3)
}

func tileImageHorizontally(im image.Image, width int, height int) image.Image {
	canvas := newImage(width, height)

	b := im.Bounds()
	dx := b.Dx()

	for i := 0; i < width; i = i + dx {
		r := image.Rect(b.Min.X+i, b.Min.Y, b.Max.X+i, b.Max.Y)
		draw.Draw(canvas, r, im, image.ZP, draw.Src)
	}

	return canvas
}

func tileImageVertically(im image.Image, width int, height int) image.Image {
	canvas := newImage(width, height)

	b := im.Bounds()
	dy := b.Dy()

	for i := 0; i < height; i = i + dy {
		r := image.Rect(b.Min.X, b.Min.Y+i, b.Max.X, b.Max.Y+i)
		draw.Draw(canvas, r, im, image.ZP, draw.Src)
	}

	return canvas
}
