package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/bb9leko/api-cep-tempo/internal/service"
)

type tempResponse struct {
	Cep        string  `json:"cep,omitempty"`
	Localidade string  `json:"localidade,omitempty"`
	TempC      float64 `json:"temp_C,omitempty"`
	TempF      float64 `json:"temp_F,omitempty"`
	TempK      float64 `json:"temp_K,omitempty"`
	Code       int     `json:"code"`
	Error      string  `json:"error"`
}

func CEPHandler(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	data, err := service.GetCEPAndTempoInfo(cep)
	if err != nil {
		var code int
		var msg string
		switch {
		case strings.Contains(err.Error(), "CEP inválido"):
			code = http.StatusUnprocessableEntity
			msg = "invalid zipcode"
		case strings.Contains(err.Error(), "CEP não encontrado"):
			code = http.StatusNotFound
			msg = "can not find zipcode"
		default:
			code = http.StatusInternalServerError
			msg = "internal error"
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(tempResponse{
			Cep:        cep,
			Localidade: "",
			Code:       code,
			Error:      msg,
		})
		return
	}

	resp := tempResponse{
		Cep:        data.Cep,
		Localidade: data.Localidade,
		TempC:      data.Temperatura.Celsius,
		TempF:      data.Temperatura.Fahrenheit,
		TempK:      data.Temperatura.Kelvin,
		Code:       http.StatusOK,
		Error:      "",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
