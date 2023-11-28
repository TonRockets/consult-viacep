package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {

	func() {
		println("Servidor rodando na porta 5051!")
		err := http.ListenAndServe(":5051", nil)
		if err != nil {
			fmt.Printf("Erro ao iniciar o servidor!\n %v", err)
		}
	}()
}

func BuscaCEP(cep string) (*ViaCEP, error) {
	res, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var viaCEP ViaCEP
	err = json.Unmarshal(body, &viaCEP)
	if err != nil {
		return nil, err
	}
	return &viaCEP, nil
}
