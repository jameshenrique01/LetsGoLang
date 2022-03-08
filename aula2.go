//arrays e slices

package main

import "fmt"

func main() {

	// var idades [3]int = [3]int{20, 25, 30}
	var idades = [3]int{20, 25, 30}

	nomes := [4]string{"yoshi", "mario", "peach", "bowser"}
	nomes[1] = "luigi"

	fmt.Println(idades, len(idades))
	fmt.Println(nomes, len(nomes))

	//slices (use arrays under the hood)
	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 85)

	fmt.Println(scores, len(scores))

	//slices ranges

	rangeOne := nomes[1:3]
	rangeTwo := nomes[2:]
	rangeThree := nomes[:3]

	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "koopa")

	fmt.Println(rangeOne)

}
