package main

import "fmt"

// var a string

func main() {
	// a = "Danilo"

	// b := "Raquel"
	// b := "Miguel"

	// c := "Pedro"
	// c = "Miguel"

	e := 10
	f := "World"
	g := 3.144
	h := false
	i := `ouuuuuu

	teste
	`

	fmt.Printf("%T \n", e)
	fmt.Printf("%T \n", f)
	fmt.Printf("%T \n", g)
	fmt.Printf("%T \n", h)
	fmt.Printf("%T \n", i)

	fmt.Printf(" ------------------------- \n")

	result := Soma(1, 1)

	fmt.Printf("%T", result)
	fmt.Printf("%v", result)

}

func Soma(a int, b int) int {
	return a + b
}
