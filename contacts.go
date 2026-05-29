package main

import "fmt"

type Personne struct {
	Prenom string `json:"prenom"`
	Nom    string `json:"nom"`
	Age    int    `json:"age"`
	Email  string `json:"email"`
}

func (p Personne) NomComplet() string {
	return fmt.Sprintf("%s %s", p.Prenom, p.Nom)
}

func (p Personne) Presentation() string {
	return fmt.Sprintf("%s, %d ans, email: %s", p.NomComplet(), p.Age, p.Email)
}

type Adresse struct {
	Rue        string `json:"rue"`
	Ville      string `json:"ville"`
	CodePostal string `json:"code_postal"`
}

func (a Adresse) Format() string {
	return fmt.Sprintf("%s, %s %s", a.Rue, a.CodePostal, a.Ville)
}

type Employe struct {
	Personne
	Adresse
	Poste   string  `json:"poste"`
	Salaire float64 `json:"salaire"`
}

func (e Employe) FicheEmploye() string {
	return fmt.Sprintf(
		"Employe: %s\nPresentation: %s\nAdresse: %s\nPoste: %s\nSalaire: %.2f",
		e.NomComplet(),
		e.Presentation(),
		e.Adresse.Format(),
		e.Poste,
		e.Salaire,
	)
}

func (e *Employe) AugmenterSalaire(pct float64) {
	e.Salaire = e.Salaire * (1 + pct/100)
}

type Etudiant struct {
	Personne
	Promo   string  `json:"promo"`
	Moyenne float64 `json:"moyenne"`
}

func (e Etudiant) MentionObtenue() string {
	switch {
	case e.Moyenne >= 16:
		return "TB"
	case e.Moyenne >= 14:
		return "B"
	case e.Moyenne >= 12:
		return "AB"
	default:
		return "P"
	}
}

func main() {
	employe1 := Employe{
		Personne: Personne{Prenom: "Sara", Nom: "Dupont", Age: 32, Email: "sara.dupont@exemple.com"},
		Adresse:  Adresse{Rue: "12 rue des Lilas", Ville: "Lyon", CodePostal: "69003"},
		Poste:    "Developpeuse",
		Salaire:  3200.00,
	}
	employe2 := Employe{
		Personne: Personne{Prenom: "Amine", Nom: "Belaid", Age: 41, Email: "amine.belaid@exemple.com"},
		Adresse:  Adresse{Rue: "8 avenue Victor Hugo", Ville: "Paris", CodePostal: "75016"},
		Poste:    "Chef de projet",
		Salaire:  4200.00,
	}

	employe1.AugmenterSalaire(5)

	employes := map[string]Employe{
		"E001": employe1,
		"E002": employe2,
	}

	etudiant1 := Etudiant{
		Personne: Personne{Prenom: "Nina", Nom: "Martin", Age: 21, Email: "nina.martin@exemple.com"},
		Promo:    "2026",
		Moyenne:  15.2,
	}
	etudiant2 := Etudiant{
		Personne: Personne{Prenom: "Yassine", Nom: "Kader", Age: 22, Email: "yassine.kader@exemple.com"},
		Promo:    "2025",
		Moyenne:  11.4,
	}

	etudiants := map[string]Etudiant{
		"S101": etudiant1,
		"S102": etudiant2,
	}

	fmt.Println(employe1.FicheEmploye())
	fmt.Println("---")
	fmt.Println(employe2.FicheEmploye())
	fmt.Println("---")
	fmt.Printf("Etudiant: %s\nPresentation: %s\nPromo: %s\nMoyenne: %.2f\nMention: %s\n",
		etudiant1.NomComplet(), etudiant1.Presentation(), etudiant1.Promo, etudiant1.Moyenne, etudiant1.MentionObtenue())
	fmt.Println("---")
	fmt.Printf("Etudiant: %s\nPresentation: %s\nPromo: %s\nMoyenne: %.2f\nMention: %s\n",
		etudiant2.NomComplet(), etudiant2.Presentation(), etudiant2.Promo, etudiant2.Moyenne, etudiant2.MentionObtenue())
	fmt.Printf("Total employes: %d | Total etudiants: %d\n", len(employes), len(etudiants))
}
