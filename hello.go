package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(7 == 2, 2 < 7)             // false true
	fmt.Println("Go" < "Go!", "Go" < "go") // true true

	nan := math.NaN()
	fmt.Println(nan == nan, math.IsNaN(nan)) // false true

	// Интерфейс хранит динамический тип *int, даже если указатель nil.
	var pointer *int
	var boxed any = pointer
	fmt.Println(pointer == nil, boxed == nil) // true false

	divisor := 0
	// Левая часть уже определяет оба результата.
	fmt.Println(divisor != 0 && 7/divisor > 2) // false
	fmt.Println(divisor == 0 || 7/divisor > 2) // true
}
