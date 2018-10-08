package main

import (
	"fmt"
)

func printNumbers() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
}

func printLetters1() {
	for i := 'A'; i < 'A'+10; i++ {
		fmt.Printf("%c", i)
	}
}

func print1() {
	printNumbers()
	printLetters1()
}

func goPrint1() {
	go printNumbers()
	go printLetters1()
}

func main() {
	
}