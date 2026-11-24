package main

import "fmt"

func main() {
	a, b, c := 5, 6, 5

	// Если обе проверки дали false, a меньше либо равно и b, и c.
	var relation string

	if a > b {
		relation = "a > b"
	} else if a > c {
		relation = "a > c"
	} else {
		relation = "a <= b и a <= c"
	}
	fmt.Println("Сравнение:", relation) // Сравнение: a <= b и a <= c

	// remainder доступна в условии и внутри обеих веток.
	var divisionResult string

	if remainder := 17 % 3; remainder == 0 {
		divisionResult = "делится на 3"
	} else {
		divisionResult = fmt.Sprintf("остаток %d", remainder)
	}
	fmt.Println("17 / 3:", divisionResult) // 17 / 3: остаток 2

	// Три значения показывают обычное совпадение, fallthrough и default.
	for _, value := range [3]int{4, c, 7} {
		switch value {
		case 3, 4:
			fmt.Printf("switch(%d), case 3 или 4\n", value) // switch(4), case 3 или 4
		case 5:
			fmt.Printf("switch(%d), case 5\n", value) // switch(5), case 5
			fallthrough                               // Выполняет следующий блок, даже если value != 6.
		case 6:
			fmt.Printf("switch(%d), case 6\n", value) // switch(5), case 6
		default:
			fmt.Printf("switch(%d), default\n", value) // switch(7), default
		}
	}

	// Без выражения после switch каждая ветка задаёт условие.
	temperature := 18

	var weather string

	switch {
	case temperature < 0:
		weather = "морозно"
	case temperature >= 25:
		weather = "жарко"
	default:
		weather = "умеренно"
	}
	fmt.Println("На улице:", weather) // На улице: умеренно

	// Type switch проверяет динамический тип интерфейса.
	var value any = 42
	var kind string
	switch typed := value.(type) {
	case int:
		kind = fmt.Sprintf("int %d", typed) // Здесь typed имеет тип int.
	case nil:
		kind = "nil"
	default:
		kind = "другой тип"
	}
	fmt.Println("Динамический тип:", kind) // Динамический тип: int 42
}
