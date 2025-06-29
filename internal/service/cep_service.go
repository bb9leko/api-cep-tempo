package service

import (
	"errors"
	"math"
	"regexp"

	"github.com/bb9leko/api-cep-tempo/internal/client"
	"github.com/bb9leko/api-cep-tempo/internal/model"
)

func IsValidCEP(cep string) bool {
	re := regexp.MustCompile(`^\d{8}$`)
	return re.MatchString(cep)
}

func GetCEPAndTempoInfo(cep string) (*model.CEPTempoResponse, error) {
	if !IsValidCEP(cep) {
		return nil, errors.New("CEP inválido. Deve conter 8 números.")
	}
	data, err := client.FetchCEP(cep)
	if err != nil {
		return nil, err
	}
	if data.Erro {
		return nil, errors.New("CEP não encontrado.")
	}

	// Consulta a WeatherAPI usando a localidade retornada
	tempC, err := client.FetchTemperatura(data.Localidade)
	if err != nil {
		return nil, err
	}

	tempF := tempC*1.8 + 32 // Fahrenheit
	tempK := tempC + 273    // Kelvin (conforme solicitado, sem casas decimais)

	// Arredonda para 1 casa decimal
	tempC = math.Round(tempC*10) / 10
	tempF = math.Round(tempF*10) / 10
	tempK = math.Round(tempK*10) / 10

	return &model.CEPTempoResponse{
		Cep:        data.Cep,
		Localidade: data.Localidade,
		Temperatura: model.TempoResponse{
			Celsius:    tempC,
			Fahrenheit: tempF,
			Kelvin:     tempK,
		},
	}, nil
}
