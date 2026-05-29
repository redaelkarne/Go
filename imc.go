package main

import "fmt"

func main() {
	// Variables
	poids := 70.5
	taille := 1.75

	// Constantes
	const (
		IMCMaigreur = 18.5
		IMCNormal   = 25.0
		IMCSurpoids = 30.0
		Nom         = "Reda"
	)

	imc := poids / (taille * taille)

	fmt.Printf("IMC : %.2f\n", imc)

	categorie := ""
	switch {
	case imc < IMCMaigreur:
		categorie = "Maigreur"
	case imc < IMCNormal:
		categorie = "Normal"
	case imc < IMCSurpoids:
		categorie = "Surpoids"
	default:
		categorie = "Obesite"
	}

	fmt.Printf("Categorie : %s\n", categorie)
	fmt.Printf("Nom : %s\n", Nom)
}
