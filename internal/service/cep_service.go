package service

import (
	"errors"
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

	return &model.CEPTempoResponse{
		Cep:        data.Cep,
		Localidade: data.Localidade,
		Temperatura: model.TempoResponse{
			Celsius: tempC,
		},
	}, nil
}
