package main

import (
	"calculadora/calculadora"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Uso: calculadora <numero1> <operacao> <numero2>")
		os.Exit(1)
	}

	num1, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		log.Fatal("Erro ao converter o primeiro número:", err)
	}

	num2, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		log.Fatal("Erro ao converter o segundo número:", err)
	}

	resultado, err := calculadora.ExecutarOperacao(num1, os.Args[2], num2)
	if err != nil {
		log.Fatal("Erro ao executar a operação: ", err)
	}

	fmt.Printf("Resultado: %.2f\n", resultado)
}
