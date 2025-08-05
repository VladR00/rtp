package handler

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"test/internal/calculation"
)

type HandlerStorage struct {
	RTP float64
}

func NewHandlerStorage(rtp float64) *HandlerStorage {
	return &HandlerStorage{RTP: rtp}
}

func (storage *HandlerStorage) GetFunc(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		DefaultResponse{Type: "Error", Message: "Only GET method allowed"}.Response(w, http.StatusMethodNotAllowed)
		return
	}

	length, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/get/"))
	if err != nil {
		log.Fatal(err)
	}
	str := calculation.NewStorage(storage.RTP)
	rtp := str.Calculation(length)
	RtpResponse{RTP: rtp}.Response(w, http.StatusOK)
}
