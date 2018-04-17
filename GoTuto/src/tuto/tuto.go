package main

import "fmt"
import "math"

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

	/**
	     Constantes
			**/

	const s string = "constant"
	fmt.Println(s)
	const n = 5000000
	fmt.Println(n)
	//const d = 3e20 / n
	//fmt.Println(d)

	fmt.Println("sinus =", math.Sin(n))

	/**
	     For
			**/

	i := 1
	for i <= 3 {
		fmt.Println("i = ", i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println("j = ", j)
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println("n = ", n)
	}

	/**
	      if else
		**/

	if 7%2 == 0 {
		fmt.Println(" 7 est paire")
	} else {
		fmt.Println(" 7 est impaire ")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "est negatif")
	} else if num < 10 {
		fmt.Println(num, "est bit")
	} else {
		fmt.Println(num, "est un multiple de 10")
	}

	/**
	      switch
		**/

	l := 2
	fmt.Print("Write", l, "as")
	switch l {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
}
