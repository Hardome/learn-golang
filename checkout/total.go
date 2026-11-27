package checkout

// total вычисляет стоимость покупки и доступна во всех файлах пакета checkout.
func total(unitPrice, quantity int) int {
	return unitPrice * quantity
}
