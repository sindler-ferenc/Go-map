package main

import "fmt"

func main() {
	var revenue int
	var expenses int
	var taxRate float64

	fmt.Print("Jövedelem: ")
	fmt.Scan(&revenue)

	fmt.Print("Kiadások: ")
	fmt.Scan(&expenses)

	fmt.Print("Adókulcs: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := float64(ebt) * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
}
