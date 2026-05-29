package main

import (
	"errors"
	"fmt"
	"strings"
)

type Produit struct {
	ID        int
	Nom       string
	Marque    string
	Prix      float64
	Stock     int
	Categorie string
	Actif     bool
}

type Catalogue struct {
	Produits []Produit
}

func (c *Catalogue) AjouterProduit(p Produit) error {
	for _, prod := range c.Produits {
		if prod.ID == p.ID {
			return errors.New("ID duplique")
		}
	}
	c.Produits = append(c.Produits, p)
	return nil
}

func (c Catalogue) TrouverParID(id int) (Produit, error) {
	for _, prod := range c.Produits {
		if prod.ID == id {
			return prod, nil
		}
	}
	return Produit{}, errors.New("produit introuvable")
}

func (c Catalogue) TrouverParCategorie(cat string) []Produit {
	resultats := []Produit{}
	for _, prod := range c.Produits {
		if strings.EqualFold(prod.Categorie, cat) {
			resultats = append(resultats, prod)
		}
	}
	return resultats
}

func (c *Catalogue) AppliquerReduction(categorie string, pct float64) int {
	modifies := 0
	for i, prod := range c.Produits {
		if strings.EqualFold(prod.Categorie, categorie) {
			reduction := prod.Prix * (pct / 100)
			c.Produits[i].Prix = prod.Prix - reduction
			modifies++
		}
	}
	return modifies
}

func (c *Catalogue) Vendre(id int, qte int) error {
	for i, prod := range c.Produits {
		if prod.ID == id {
			if qte <= 0 {
				return errors.New("quantite invalide")
			}
			if prod.Stock < qte {
				return errors.New("stock insuffisant")
			}
			c.Produits[i].Stock = prod.Stock - qte
			return nil
		}
	}
	return errors.New("produit introuvable")
}

func (c Catalogue) Rapport() string {
	total := 0.0
	for _, prod := range c.Produits {
		total += prod.Prix * float64(prod.Stock)
	}
	return fmt.Sprintf("Produits: %d | Valeur stock: %.2f", len(c.Produits), total)
}

func afficherProduit(p Produit) {
	fmt.Printf("ID: %d | %s %s | %.2f | Stock: %d | Categorie: %s | Actif: %t\n",
		p.ID, p.Marque, p.Nom, p.Prix, p.Stock, p.Categorie, p.Actif)
}

func main() {
	catalogue := Catalogue{}

	_ = catalogue.AjouterProduit(Produit{ID: 1, Nom: "iPhone 15", Marque: "Apple", Prix: 1199.00, Stock: 8, Categorie: "Smartphone", Actif: true})
	_ = catalogue.AjouterProduit(Produit{ID: 2, Nom: "MacBook Air M3", Marque: "Apple", Prix: 1499.00, Stock: 5, Categorie: "Laptop", Actif: true})
	_ = catalogue.AjouterProduit(Produit{ID: 3, Nom: "Galaxy S24", Marque: "Samsung", Prix: 999.00, Stock: 10, Categorie: "Smartphone", Actif: true})
	_ = catalogue.AjouterProduit(Produit{ID: 4, Nom: "ThinkPad X1", Marque: "Lenovo", Prix: 1799.00, Stock: 3, Categorie: "Laptop", Actif: true})
	_ = catalogue.AjouterProduit(Produit{ID: 5, Nom: "PlayStation 5", Marque: "Sony", Prix: 549.00, Stock: 6, Categorie: "Console", Actif: true})

	for {
		fmt.Println("\n[1] Ajouter [2] Chercher [3] Soldes [4] Vendre [5] Rapport [0] Quitter")
		fmt.Print("Choix: ")

		var choix int
		if _, err := fmt.Scan(&choix); err != nil {
			fmt.Println("Erreur de saisie:", err)
			return
		}

		switch choix {
		case 1:
			var p Produit
			fmt.Print("ID: ")
			if _, err := fmt.Scan(&p.ID); err != nil {
				fmt.Println("Erreur:", err)
				continue
			}
			fmt.Print("Nom: ")
			if _, err := fmt.Scan(&p.Nom); err != nil {
				fmt.Println("Erreur:", err)
				continue
			}
			fmt.Print("Marque: ")
			if _, err := fmt.Scan(&p.Marque); err != nil {
				fmt.Println("Erreur:", err)
				continue
			}
			fmt.Print("Prix: ")
			if _, err := fmt.Scan(&p.Prix); err != nil {
				fmt.Println("Erreur:", err)
				continue
			}
			fmt.Print("Stock: ")
			if _, err := fmt.Scan(&p.Stock); err != nil {
				fmt.Println("Erreur:", err)
				continue
			}
			fmt.Print("Categorie: ")
			if _, err := fmt.Scan(&p.Categorie); err != nil {
				fmt.Println("Erreur:", err)
				continue
			}
			p.Actif = true
			if err := catalogue.AjouterProduit(p); err != nil {
				fmt.Println("Erreur:", err)
				continue
			}
			fmt.Println("Produit ajoute.")
		case 2:
			fmt.Print("Chercher par [1] ID ou [2] Categorie: ")
			var sousChoix int
			if _, err := fmt.Scan(&sousChoix); err != nil {
				fmt.Println("Erreur:", err)
				continue
			}
			if sousChoix == 1 {
				var id int
				fmt.Print("ID: ")
				if _, err := fmt.Scan(&id); err != nil {
					fmt.Println("Erreur:", err)
					continue
				}
				prod, err := catalogue.TrouverParID(id)
				if err != nil {
					fmt.Println("Erreur:", err)
					continue
				}
				afficherProduit(prod)
			} else if sousChoix == 2 {
				var cat string
				fmt.Print("Categorie: ")
				if _, err := fmt.Scan(&cat); err != nil {
					fmt.Println("Erreur:", err)
					continue
				}
				produits := catalogue.TrouverParCategorie(cat)
				if len(produits) == 0 {
					fmt.Println("Aucun produit trouve.")
					continue
				}
				for _, prod := range produits {
					afficherProduit(prod)
				}
			} else {
				fmt.Println("Choix invalide.")
			}
		case 3:
			var cat string
			var pct float64
			fmt.Print("Categorie: ")
			if _, err := fmt.Scan(&cat); err != nil {
				fmt.Println("Erreur:", err)
				continue
			}
			fmt.Print("Reduction (%): ")
			if _, err := fmt.Scan(&pct); err != nil {
				fmt.Println("Erreur:", err)
				continue
			}
			modifies := catalogue.AppliquerReduction(cat, pct)
			fmt.Printf("Produits modifies: %d\n", modifies)
		case 4:
			var id int
			var qte int
			fmt.Print("ID: ")
			if _, err := fmt.Scan(&id); err != nil {
				fmt.Println("Erreur:", err)
				continue
			}
			fmt.Print("Quantite: ")
			if _, err := fmt.Scan(&qte); err != nil {
				fmt.Println("Erreur:", err)
				continue
			}
			if err := catalogue.Vendre(id, qte); err != nil {
				fmt.Println("Erreur:", err)
				continue
			}
			fmt.Println("Vente OK.")
		case 5:
			fmt.Println(catalogue.Rapport())
		case 0:
			fmt.Println("Au revoir.")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
