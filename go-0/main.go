package main

import (
	"fmt"
)

func main() {
	estados, err := os10maioresEstadosDoBrasil()
	if err == nil {
		fmt.Println(estados)
	}
}

func os10maioresEstadosDoBrasil() ([]string, error) {
	var data []string
	data = append(data, "Amazonas",
		"Pará",
		"Mato Grosso",
		"Minas Gerais",
		"Bahia",
		"Mato Grosso do Sul",
		"Goiás",
		"Maranhão",
		"Rio Grande do Sul",
		"Tocantins")
	return data, nil
}
