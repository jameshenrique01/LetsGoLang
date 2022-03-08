package main

import "fmt"

func main() {

	idade := 20
	nome := "James"

	// Print
	fmt.Print("Olá, ")
	fmt.Print("mundo! \n")
	fmt.Print("nova linha \n ")

	// Println
	fmt.Println("Olá pessoas!")
	fmt.Println("Tchau pessoas!")
	fmt.Println("Minha idade é", idade, "e meu nome é", nome)

	// Printf (formatted strings)%_ = especifica o formato
	fmt.Printf("Minha idade é %v e meu nome é %v \n", idade, nome)
	fmt.Printf("Minha idade é %q e meu nome é %q \n", idade, nome)
	fmt.Printf("O tipo da minha idade é %T", idade)
	fmt.Printf("Seu placar é %f pontos! \n", 255.55)
	fmt.Printf("Seu placar é %0.1f pontos! \n", 255.55)

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("Minha idade é %v e meu nome é %v \n", idade, nome)
	fmt.Println("a string salve é:", str)

}
