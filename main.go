package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
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
	http.HandleFunc("/", BuscaCEPHandler)
	func() {
		println("Servidor rodando na porta 5051!")
		err := http.ListenAndServe(":5051", nil)
		if err != nil {
			fmt.Printf("Erro ao iniciar o servidor!\n %v", err)
		}
	}()
}

func BuscaCEPHandler(resWriter http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/" {
		resWriter.WriteHeader(http.StatusNotFound)
		return
	}

	cepParam := req.URL.Query().Get("cep")
	if cepParam == "" {
		resWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	validCep := isValidCEP(cepParam)

	if !validCep {
		resWriter.WriteHeader(http.StatusBadRequest)
		resWriter.Write([]byte("Cep invalido!"))
		return
	}

	result, err := BuscaCEP(cepParam)
	if err != nil {
		resWriter.WriteHeader(http.StatusInternalServerError)
		resWriter.Write([]byte(err.Error()))
		return
	}

	resWriter.Header().Set("Content-Type", "application/json")
	resWriter.WriteHeader(http.StatusOK)
	json.NewEncoder(resWriter).Encode(result)
	return
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

func isValidCEP(cep string) bool {
	cepPattern := regexp.MustCompile(`^\d{5}\d{3}$`)

	if !cepPattern.MatchString(cep) {
		return false
	}

	return true
}
