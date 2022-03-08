package main

import "fmt"

func main() {

	fmt.Println("Hello")

	//strings

	var nomeUm string = "James"
	var nomeDois = "Henrique"
	var nomeTres string

	fmt.Println(nomeUm, nomeDois, nomeTres)

	nomeUm = "peach"
	nomeDois = "bowser"

	fmt.Println(nomeUm, nomeDois, nomeTres)

	nomeQuatro := "yoshi"
	fmt.Println(nomeQuatro)

	// inteiros

	var anoUm int = 20
	var anoDois = 30
	anoTres := 40

	fmt.Println(anoUm, anoDois, anoTres)

	// bits & memoria

	// var numUm int8 = 12
	// var numTwo int8 = -128
	// var numThree uint16 = 256

	var scoreOne float32 = 25.88
	var scoreTwo float64 = 737892173298137.5
	scoreThree := 1.5
}
