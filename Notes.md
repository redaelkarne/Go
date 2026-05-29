# Notes Go — Révision perso

---

## 1. Types & Variables

- Types de base déjà maîtrisés (int, string, bool, float64, etc.) ✓
- **Déclaration courte** avec `:=` — Go infère le type automatiquement :
  ```go
  nom := "REDA"
  x, y := 10, 30        // multi-déclaration, même avec types différents
  a, b := 10, "hello"   // ça marche aussi
  ```
- **Déclaration classique** avec `var` (utile pour les variables globales ou sans valeur initiale) :
  ```go
  var age int
  var prenom string = "REDA"
  ```
- **Zéro value** : toute variable non initialisée a une valeur par défaut (`0`, `""`, `false`, `nil`).

---

## 2. Slices

Un slice est un tableau de taille **modifiable**, contrairement à un array dont la taille est fixe.

```go
s := []int{1, 2, 3}
s = append(s, 4)        // ajoute un élément
fmt.Println(len(s))     // 4
```

- Un slice est en fait une *vue* sur un array sous-jacent.
- `len()` = nombre d'éléments, `cap()` = capacité allouée.

---

## 3. Arrays & Maps

**Array** : taille fixe, rien de spécial par rapport aux autres langages.
```go
arr := [3]int{1, 2, 3}
```

**Map** : dictionnaire clé → valeur, comme en Python.
```go
ages := map[string]int{
    "Reda": 25,
    "Ali":  30,
}
ages["Sara"] = 22
delete(ages, "Ali")

val, ok := ages["Reda"] // ok = false si la clé n'existe pas
```

---

## 4. Fonctions

- **Pas de surcharge** en Go — chaque fonction a un nom unique. Pas de variantes selon les types de paramètres.
- Une fonction peut **retourner plusieurs valeurs** (très courant en Go) :
  ```go
  func divide(a, b float64) (float64, error) {
      if b == 0 {
          return 0, errors.New("division par zéro")
      }
      return a / b, nil
  }
  ```
- **Fonction variadique** : accepte un nombre variable d'arguments avec `...` :
  ```go
  func sum(nums ...int) int {
      total := 0
      for _, n := range nums {
          total += n
      }
      return total
  }

  sum(1, 2, 3)       // 6
  sum(1, 2, 3, 4, 5) // 15
  ```

---

## 5. Switch & Fallthrough

Par défaut, un `switch` en Go **ne tombe pas** dans le cas suivant automatiquement (contrairement à C/Java). Pour forcer ce comportement, on utilise `fallthrough` :

```go
switch x {
case 1:
    fmt.Println("un")
    fallthrough         // exécute aussi le case 2
case 2:
    fmt.Println("deux")
case 3:
    fmt.Println("trois")
}
```

Un `switch` peut aussi s'écrire sans expression, comme un `if/else` :
```go
switch {
case x > 10:
    fmt.Println("grand")
case x > 0:
    fmt.Println("petit")
}
```

---

## 6. Pointeurs

Un pointeur stocke l'**adresse mémoire** d'une variable, pas sa valeur directement.

```go
x := 42
p := &x          // p pointe vers x
fmt.Println(*p)  // 42  → déréférencement
*p = 100
fmt.Println(x)   // 100 → x a été modifié
```

- `&variable` → donne l'adresse
- `*pointeur` → accède à la valeur (déréférencement)
- Utile pour modifier une variable dans une fonction sans la copier.

---

## 7. Visibilité (Exported vs Unexported)

Pas de mots-clés `public` / `private` en Go. La visibilité est déterminée par la **casse** :

| Casse | Visibilité |
|-------|-----------|
| `MaFonction` | Exportée (accessible hors du package) |
| `maFonction` | Non exportée (interne au package) |

Valable pour les fonctions, types, variables, constantes, et champs de struct.

---

## 8. OOP à la Go — Structs, Méthodes & Interfaces

Go n'a **pas d'héritage**. À la place : composition + interfaces.

### Struct
```go
type Personne struct {
    Nom string
    Age int
}
```

### Méthode sur un type
```go
func (p Personne) Saluer() string {
    return "Bonjour, je suis " + p.Nom
}

p := Personne{Nom: "Reda", Age: 25}
fmt.Println(p.Saluer())
```

### Interface
Une interface définit un **comportement** (un ensemble de méthodes). Tout type qui implémente ces méthodes satisfait l'interface — implicitement, sans `implements`.

```go
type Animal interface {
    Parler() string
}

type Chien struct{}
func (c Chien) Parler() string { return "Woof" }

type Chat struct{}
func (c Chat) Parler() string { return "Miaou" }

func faireParler(a Animal) {
    fmt.Println(a.Parler())
}
```

### Composition (à la place de l'héritage)
```go
type Employe struct {
    Personne           // embedding — Employe "hérite" des champs de Personne
    Entreprise string
}

e := Employe{Personne: Personne{Nom: "Reda"}, Entreprise: "Anthropic"}
fmt.Println(e.Nom)    // accessible directement
```

---

## Rappel général — Ce qui distingue Go

| Concept | Go |
|---------|----|
| Héritage | ❌ — on utilise la composition |
| Surcharge de fonctions | ❌ — pas supportée |
| Visibilité | Par la casse (Majuscule / minuscule) |
| Gestion des erreurs | Retour de valeur, pas d'exceptions |
| Typage | Statique, inférence avec `:=` |