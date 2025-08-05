package main

import (
	"flag"
	"log"
	"net/http"
	"test/internal/handler"
)

func main() {
	rtp := flag.Float64("rtp", 0.0, "Значение для rtp") // По умолчанию 0.0
	flag.Parse()

	storage := handler.NewHandlerStorage(*rtp)

	http.HandleFunc("/get/", storage.GetFunc)
	log.Println("Server start at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
