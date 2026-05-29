package main

import (
	"errors"
	"fmt"
)

func operer(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("division par zero")
		}
		return a / b, nil
	default:
		return 0, errors.New("operation inconnue")
	}
}

func creerOperation(op string) func(float64, float64) float64 {
	switch op {
	case "+":
		return func(a, b float64) float64 { return a + b }
	case "-":
		return func(a, b float64) float64 { return a - b }
	case "*":
		return func(a, b float64) float64 { return a * b }
	case "/":
		return func(a, b float64) float64 { return a / b }
	default:
		return func(a, b float64) float64 { return 0 }
	}
}

func main() {
	for {
		var a, b float64
		var op string

		fmt.Print("Entrez: nombre nombre operation (ou quit): ")
		_, err := fmt.Scan(&a, &b, &op)
		if err != nil {
			fmt.Println("Erreur de saisie:", err)
			return
		}

		if op == "quit" {
			fmt.Println("Fin.")
			break
		}

		resultat, err := operer(a, b, op)
		if err != nil {
			fmt.Println("Erreur:", err)
			continue
		}

		operation := creerOperation(op)
		resultatClosure := operation(a, b)

		fmt.Printf("Resultat: %.2f\n", resultat)
		fmt.Printf("Resultat (closure): %.2f\n", resultatClosure)
	}
}
