package calculadora

import (
	"errors"
	"fmt"
	"math"
)

var ErrDivisaoPorZero = errors.New("divisão por zero")

// Soma retorna a soma de dois números
func somar(a, b float64) float64 {
	return a + b
}

// Subtracao retorna a diferença entre dois números
func subtrair(a, b float64) float64 {
	return a - b
}

// Multiplicacao retorna o produto de dois números
func multiplicar(a, b float64) float64 {
	return a * b
}

// Divisao retorna o resultado da divisão entre dois números
func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivisaoPorZero
	}
	return a / b, nil
}

// Elevar retorna o resultado do exponencial
func elevar(a, b float64) float64 {
	return math.Pow(a, b)
}

func somarPorcentagem(a, b float64) float64 {
	porcent := (a * b) / 100

	return a + porcent
}

func subtrairPorcentagem(a, b float64) float64 {
	porcent := (a * b) / 100

	return a - porcent
}

func ExecutarOperacao(a float64, operacao string, b float64) (float64, error) {
	switch operacao {
	case "+":
		return somar(a, b), nil
	case "-":
		return subtrair(a, b), nil
	case "*":
		return multiplicar(a, b), nil
	case "/":
		return dividir(a, b)
	case "^":
		return elevar(a, b), nil
	case "%+":
		return somarPorcentagem(a, b), nil
	case "%-":
		return subtrairPorcentagem(a, b), nil
	default:
		return 0, fmt.Errorf("operação desconhecida: %s", operacao)
	}
}
