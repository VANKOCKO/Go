package main

import "fmt"

/**
   Intro,
**/
func main() {

	fmt.Printf("Hello, world \n")

	/**
	   Valeurs
		**/

	fmt.Println("Valeurs")
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0=", 7.0/3.0)
	fmt.Println("true && false :", true && false)
	fmt.Println("true || false :", true || false)
	fmt.Println("!true :", !true)

	/**
	     Variables
			**/

	fmt.Println("Variables")
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "short"
	fmt.Println(f)
}
