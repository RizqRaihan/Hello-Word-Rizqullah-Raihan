package main

import "fmt"

func main() {
	var a, b, c, hasil int
	fmt.Scanln(&a, &b, &c)

	i := 0
	for i <= a {
		hasil = b * i * c
		fmt.Println("+", hasil)
		i++
	}
}
