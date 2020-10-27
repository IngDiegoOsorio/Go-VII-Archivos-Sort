package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type Ascendente []string

func (a Ascendente) Len() int           { return len(a) }
func (a Ascendente) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Ascendente) Less(i, j int) bool { return a[i] < a[j] }

type Descendente []string

func (a Descendente) Len() int           { return len(a) }
func (a Descendente) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Descendente) Less(i, j int) bool { return a[i] > a[j] }

func main() {
	var cadenas []string
	var cadena string
	var inputs int

	fmt.Printf("Cuantas cadenas vas a ingresar? ")
	fmt.Scanln(&inputs)

	for i := 0; i < inputs; i++ {
		fmt.Print("Cadena", i, " ")
		fmt.Scanln(&cadena)
		cadenas = append(cadenas, cadena)
	}

	sort.Sort(Ascendente(cadenas))

	cadena = strings.Join(cadenas, "\n")

	file, err := os.Create("asecendente.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	file.WriteString(cadena)

	sort.Sort(Descendente(cadenas))

	cadena = strings.Join(cadenas, "\n")

	file2, err := os.Create("descendente.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file2.Close()

	file2.WriteString(cadena)
}
