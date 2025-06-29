package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bb9leko/api-cep-tempo/configs"
	"github.com/bb9leko/api-cep-tempo/internal/handler"
)

func main() {
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Erro ao carregar configuração: %v", err)
	}

	// Disponibiliza a chave para o restante do sistema via variável de ambiente
	os.Setenv("WEATHERAPI_KEY", cfg.WeatherAPIKey)
	fmt.Println("WEATHERAPI_KEY carregada:", cfg.WeatherAPIKey)

	//http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/cep", handler.CEPHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Servidor rodando na porta %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
