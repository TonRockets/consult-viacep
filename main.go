package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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
	for _, cep := range os.Args[1:] {
		req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")

		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao consultar a API ViaCEP! %v\n", err)
		}

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler a resposta da API ViaCEP! %v\n", err)
		}

		defer req.Body.Close()

		var data ViaCEP
		parseErr := json.Unmarshal(res, &data)
		if parseErr != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer a conversao dos dados trazidos da API ViaCEP! %v\n", parseErr)
		}

		fmt.Println("Seu logradouro e:", data.Logradouro)
		fmt.Println("Sua cidade e:", data.Localidade)
		fmt.Println("Seu estado e:", data.Uf)

		file, err := os.Create("resultado-viacep.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar o arquivo com os resultados da API ViaCEP! %v\n", err)
		}

		defer file.Close()

		_, writeFileErr := file.WriteString(fmt.Sprintf("O seu resultado e:\n CEP: %s,\n Logradouro: %s,\n Cidade: %s,\n UF: %s\n", data.Cep, data.Logradouro, data.Localidade, data.Uf))
		if writeFileErr != nil {
			fmt.Printf("Erro ao registrar resultados no arquivo! %v\n", writeFileErr)
		}
		println("Arquivo criado com sucesso!")
	}
}
