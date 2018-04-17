package main

import (
	"fmt"
	"math"
)

/**
   Intro,
**/
func main() {

	fmt.Printf("Hello, world \n")

	/**
	   Valeurs


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

	/**
	      array
		**/

	var tab [10]int
	fmt.Println("vide", tab)

	tab[4] = 100
	fmt.Println("set", tab)
	fmt.Println("get", tab[4])
	fmt.Println("longueur de tab", len(tab))

	tab1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("declaration de tab ", tab1)

	var tabDouble [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			tabDouble[i][j] = i + j
		}
	}
	fmt.Println("DeuxD", tabDouble)

	/**
			  Slice
			  make() permet de creer un tableau
	 		**/

	slice := make([]string, 3)
	fmt.Println("slice", slice)

	slice[0] = "a"
	slice[1] = "b"
	slice[2] = "c"

	fmt.Println("set", slice)
	fmt.Println("get", slice[2])
	fmt.Println("longueur de slice ", len(slice))

	slice = append(slice, "d")
	slice = append(slice, "e", "f")
	fmt.Println("append", slice)

	slice1 := make([]string, len(slice))
	copy(slice1, slice)
	fmt.Println("cppie", slice1)

	l1 := slice[2:5]
	fmt.Println("sl1", l1)

	l1 = slice[:5]
	fmt.Println("sl2", l1)

	l1 = slice[2:]
	fmt.Println("sl3", l1)

	t1 := []string{"g", "h", "i"}
	fmt.Println("declaration de t1", t1)

	/**
			  map

	 		**/

	m1 := make(map[string]int)
	m1["k1"] = 7
	m1["k2"] = 13

	fmt.Println("map:", m1)

	v1 := m1["k1"]
	fmt.Println("v1 = ", v1)

	/**
	  range
	 **/

	nums := []int{2, 3, 4}
	sum := 0

	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum", sum)

	for i, num := range nums {
		if num == 4 {
			fmt.Println("index", i)
		}
	}

	/**
	   appel des fonctions
	**/

	res := plus(1, 2)
	fmt.Println("res", res)
	res2 := plusplus(1, 2, 3)
	fmt.Println("res2", res2)
	a1, a2 := vals()
	fmt.Println("a1", a1)
	fmt.Println("a2", a2)
}

/**
 functions
**/

func plus(a, b int) int {
	return a + b

}
func plusplus(a, b, c int) int {
	return a + b + c
}
func vals() (int, int) {
	return 2, 4
}
