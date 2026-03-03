package main

import "fmt"

func printlnE() {
	for i := 1; i <= 7; i++ {
		if i == 1 || i == 4 || i == 7 {
			fmt.Println("*******")
		} else {
			fmt.Println("*")
		}
	}
}

func printlnC() {
	for i := 1; i <= 7; i++ {
		if i == 1 || i == 7 {
			fmt.Println("  ****")
		} else if i == 2 || i == 6 {
			fmt.Println(" *    *")
		} else {
			fmt.Println("*")
		}
	}
}
func main() {
	fmt.Println("課題1-1 : \n Hello World!")
	fmt.Print("課題1-2 : \n Hello")
	fmt.Print("World\n")
	fmt.Println("課題1-3 :")
	printlnE()
	fmt.Println("課題1-4 :")
	printlnE()
	printlnC()
	printlnC()
}
