package main

import (
	"fmt"
	"math/rand"
	"time"
)

// iota = création automatique de constantes
const (
	Faible = iota
	Moyen
	Fort
	Super
)

func main() {

	rand.Seed(time.Now().UnixNano())

	// Slice des héros
	heros := []string{
		"🦸 Chat Ninja",
		"🤖 Robot Danseur",
		"🦄 Licorne Magique",
		"🏴 Pirate Gourmand",
	}

	// Slice des missions
	missions := []string{
		"Trouver le trésor perdu",
		"Sauver le royaume",
		"Battre le dragon WiFi",
		"Manger 100 tacos",
	}

	fmt.Println("================================")
	fmt.Println("🎮 MINI JEU GO - AVENTURE")
	fmt.Println("================================")

	// Choix aléatoire d’un héros
	hero := heros[rand.Intn(len(heros))]
	fmt.Println("Votre héros :", hero)

	// Choix aléatoire d’une mission
	mission := missions[rand.Intn(len(missions))]
	fmt.Println("Mission :", mission)

	// Énergie aléatoire
	energie := rand.Intn(4)

	fmt.Println("\n⚡ Niveau d'énergie :")

	// Switch + fallthrough
	switch energie {

	case Faible:
		fmt.Println("Énergie faible...")
		fallthrough

	case Moyen:
		fmt.Println("Le héros commence à se réveiller...")
		fallthrough

	case Fort:
		fmt.Println("Le héros devient puissant !")
		fallthrough

	case Super:
		fmt.Println("🔥 MODE SUPER ACTIVÉ 🔥")
	}

	fmt.Println("\n📜 Progression de la mission :")

	// Boucle for unique
	for i, etape := range []string{
		"Préparation",
		"Combat",
		"Victoire",
	} {

		fmt.Printf("%d -> %s\n", i+1, etape)

		time.Sleep(1 * time.Second)
	}

	fmt.Println("\n🏆 Mission terminée avec succès !")
}
