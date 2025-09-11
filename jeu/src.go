package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

)

type Inventaire struct {
    Potions int
	Items [] string 
}

type Character struct {
    Nom        string
    Classe     string
    Niveau     float64
    Pdv        int
    PdvMax     int
    Inventaire Inventaire
}

func InitCharacter(nom, classe string, niveau float64, pdv, pdvmax int) Character {
    return Character{
        Nom:        nom,
        Classe:     classe,
        Niveau:     niveau,
        Pdv:        pdv,
        PdvMax:     pdvmax,
        Inventaire: Inventaire{
			Potions : 3,
			Items : [] string{},
		},
			
    }
}

func DisplayInfo(c Character) {
    fmt.Println("Nom :", c.Nom)
    fmt.Println("Classe :", c.Classe)
    fmt.Println("Niveau :", c.Niveau)
    fmt.Println("PV :", c.Pdv, "/", c.PdvMax)
    fmt.Println("Potions dans l'inventaire :", c.Inventaire.Potions)
}

func accessInventory(c Character) {
    fmt.Println("=== Inventaire de", c.Nom, "===")
    fmt.Println("Potions :", c.Inventaire.Potions)
    if len(c.Inventaire.Items) == 0 {
        fmt.Println("Aucun objets dans l'inventaire.")
    } else {
        fmt.Println("Objets :")
        for _, Items := range c.Inventaire.Items {
            fmt.Println("-", Items)
        }
    }
}

func main() {
	fmt.Println("Bienvenue dans Milousque G@ming")
	fmt.Println("Voulez vous créer un personnage ou reprendre un personnage existant ?")
	fmt.Println("Tapez CREER pour créer un nouveau personnage ou REPRENDRE pour selectionner un personnage existant")
	reader := bufio.Newreader(os.Stdin)
}
