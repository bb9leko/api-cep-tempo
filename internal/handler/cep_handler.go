package handler

import (
	"encoding/json"
	"net/http"

	"github.com/bb9leko/api-cep-tempo/internal/service"
)

func CEPHandler(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	data, err := service.GetCEPAndTempoInfo(cep)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
