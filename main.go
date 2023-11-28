package main

import (
	"fmt"
	"net/http"
)

func main() {

	func() {
		println("Servidor rodando na porta 5051!")
		err := http.ListenAndServe(":5051", nil)
		if err != nil {
			fmt.Printf("Erro ao iniciar o servidor!\n %v", err)
		}
	}()
}
