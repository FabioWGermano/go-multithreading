package response

import (
	"encoding/json"
	"net/http"
)

type Sucess struct {
	statusCode int
	result     interface{}
}

func NewSucess(result interface{}, status int) Sucess {
	return Sucess{
		statusCode: status,
		result:     result,
	}
}

func (s Sucess) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(s.statusCode)
	_ = json.NewEncoder(w).Encode(s.result)
}
