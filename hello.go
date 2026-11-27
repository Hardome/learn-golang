package main

import "fmt"

func sayHello(name string) {
	fmt.Printf("Привет, %s!\n", name) // Подставляет имя в приветствие.
}

func main() {
	sayHello("Go") // Привет, Go!
}
