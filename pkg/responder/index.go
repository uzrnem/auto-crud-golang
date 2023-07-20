package responder

import (
	"encoding/json"
	"net/http"
)

type Reponder struct {
}

var (
	Resp *Reponder
)

func Load() error {
	resp := &Reponder{}
	Resp = resp
	return nil
}

func (resp *Reponder) SendJson(w http.ResponseWriter, statusCode int, body any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
	return nil
}

func (resp *Reponder) SendError(w http.ResponseWriter, err error) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)
	body := map[string]string{"message": err.Error()}
	json.NewEncoder(w).Encode(body)
	return nil
}
