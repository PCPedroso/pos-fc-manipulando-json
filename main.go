package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Numero int
	Saldo  float64
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100.00}

	res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		fmt.Println(err)
	}

	jsonPuro := []byte(`{"Numero": 2, "Saldo": 110}`)
	var novaConta Conta
	err = json.Unmarshal(jsonPuro, &novaConta)
	if err != nil {
		panic(err)
	}
	fmt.Println(novaConta)
}
