package main

import (
	"fmt"

	"github.com/Hardome/learn-golang/checkout"
	"github.com/Hardome/learn-golang/checkout/card"
	"github.com/Hardome/learn-golang/checkout/cash"
)

func main() {
	fmt.Println(checkout.Receipt(250, 3)) // Итого: 750 ₽
	fmt.Println(card.Label())             // Карта
	fmt.Println(cash.Label())             // Наличные

	// checkout.total(250, 3) // Ошибка компиляции: total доступна только внутри checkout.
}
