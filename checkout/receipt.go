package checkout

import "fmt"

// Receipt возвращает строку с итоговой стоимостью покупки в рублях.
func Receipt(unitPrice, quantity int) string {
	return fmt.Sprintf("Итого: %d ₽", total(unitPrice, quantity)) // total объявлена в другом файле checkout.
}
