package main

import (
	"github.com/pilu/traffic"
)

func PoweredByHandler(w traffic.ResponseWriter, r *traffic.Request) {
	w.Header().Set("X-Powered-By", "Grumpy cat")
}
