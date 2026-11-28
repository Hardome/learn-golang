package main

import "fmt"

type mile uint
type kilometer uint
type mileAlias = uint
type routeMile = mile

type BinaryOp = func(int, int) int
type NamedBinaryOp func(int, int) int

// Обобщённый псевдоним доступен с Go 1.24.
type set[K comparable] = map[K]bool

func main() {
	showDistanceTypes()
	showOperations()
	showAliasIdentity()
	showGenericAlias()
}

func showDistanceTypes() {
	// Нетипизированная константа 5 представима в mile, а арифметика работает как у uint.
	var distance mile = 5
	fmt.Println(distance) // 5
	distance += 5
	fmt.Println(distance)     // 10
	distanceToEnemy(distance) // 10 миль до противника

	var routeDistance routeMile = distance
	distanceToEnemy(routeDistance) // 10 миль до противника

	// Обе переменные имеют тип uint, поэтому присваивание возможно в обе стороны.
	var raw uint = 10
	var aliasDistance mileAlias = raw
	raw = aliasDistance
	fmt.Println(raw, aliasDistance) // 10 10

	km := kilometer(16)
	// distanceToEnemy(km) // Ошибка компиляции: kilometer и mile — разные типы.
	// distanceToEnemy(raw) // Ошибка компиляции: uint и mile — разные типы.
	// distanceToEnemy(-1) // Ошибка компиляции: -1 не представимо в mile.
	distanceToEnemy(mile(raw)) // 10 миль до противника
	// Преобразование типа не переводит километры в мили.
	fmt.Println(mile(km)) // 16
}

func distanceToEnemy(distance mile) {
	fmt.Println(distance, "миль до противника")
}

func (distance mile) Double() mile {
	return distance * 2
}

func showAliasIdentity() {
	// У псевдонима тот же набор методов и тот же динамический тип.
	var distance routeMile = 5
	fmt.Println(distance.Double()) // 10

	value := any(distance)
	_, isMile := value.(mile)
	_, isUint := value.(uint)
	fmt.Println(isMile, isUint) // true false

	// Срезы идентичны, так как идентичны их типы элементов.
	var route []routeMile = []mile{5, 10}
	fmt.Println(route) // [5 10]
	// _ = []uint(route) // Ошибка компиляции: mile и uint — разные типы элементов.
}

func showOperations() {
	// Псевдоним и исходная сигнатура обозначают один тип функции.
	var rawOperation func(int, int) int = add
	action(10, 25, rawOperation) // 35
	var operation BinaryOp = rawOperation
	action(10, 35, operation) // 45

	type anotherOperation func(int, int) int
	var namedOperation NamedBinaryOp = rawOperation // Тип rawOperation неименованный.
	// var invalid anotherOperation = namedOperation // Ошибка компиляции. Оба типа именованные.
	other := anotherOperation(namedOperation)
	action(2, 3, other) // 5

	var unset BinaryOp
	fmt.Println(unset == nil) // true
	// action(10, 35, unset) // Паника: вызов функции со значением nil.
}

func action(left, right int, op BinaryOp) {
	fmt.Println(op(left, right))
}

func add(left, right int) int {
	return left + right
}

func showGenericAlias() {
	// set[string] идентичен map[string]bool, значение отсутствующего ключа равно false.
	var active set[string] = map[string]bool{"go": true}
	fmt.Println(active["go"], active["rust"]) // true false
}
