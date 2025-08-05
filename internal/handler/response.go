package handler

import (
	"encoding/json"
	"net/http"
)

func (r DefaultResponse) Response(w http.ResponseWriter, header int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(header)
	json.NewEncoder(w).Encode(DefaultResponse{Type: r.Type, Message: r.Message})
}

func (r RtpResponse) Response(w http.ResponseWriter, header int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(header)
	json.NewEncoder(w).Encode(RtpResponse{RTP: r.RTP})
}

type DefaultResponse struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type RtpResponse struct {
	RTP float64 `json:"result"`
}
