package main

import "fmt"

type library []string

func (books library) print() {
	for _, title := range books {
		fmt.Println(title)
	}
}

func (books library) replaceFirst(title string) {
	books[0] = title // Вызывающая сторона передаёт непустой срез. Массив остаётся общим.
}

func (books library) keepFirst() {
	books = books[:1] // Новая длина хранится только в локальном описании среза.
	fmt.Println(len(books))
}

func (books *library) add(title string) {
	*books = append(*books, title) // Сохраняем новое описание в исходной переменной.
}

type person struct {
	name string
	age  int
}

func (p person) print() {
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
}

func (p person) eat(meal string) {
	fmt.Println(p.name, "eats", meal)
}

func (p person) currentAge() int {
	return p.age
}

func (p person) updateAgeOnCopy(newAge int) {
	p.age = newAge // Меняется только поле получателя, то есть копии person.
	fmt.Println("inside", p.age)
}

func (p person) withAge(newAge int) person {
	p.age = newAge
	return p
}

func (p *person) updateAge(newAge int) {
	p.age = newAge // Сокращение для (*p).age = newAge.
}

func (p *person) nameOrUnknown() string {
	if p == nil {
		return "unknown"
	}
	return p.name
}

type ageReader interface{ currentAge() int }
type ageUpdater interface{ updateAge(int) }

// Эти присваивания проверяют наборы методов на этапе компиляции.
var _ ageReader = person{}
var _ ageReader = (*person)(nil)
var _ ageUpdater = (*person)(nil)

// var _ ageUpdater = person{} // Ошибка компиляции: набор person не содержит updateAge.

func main() {
	showPerson()
	showReceivers()
	showLibrary()
	showAddressabilityAndInterfaces()
	showMethodFunctions()
	showNilReceiver()
	showGenericMethod()
}

func showLibrary() {
	lib := library{"Book1", "Book2"}
	lib.print() // Book1, Book2 на отдельных строках

	lib.replaceFirst("Go")
	fmt.Println(lib)      // [Go Book2]
	lib.keepFirst()       // 1
	fmt.Println(len(lib)) // 2

	lib.add("Book3") // lib адресуемая. Метод получает &lib и обновляет описание среза.
	fmt.Println(lib) // [Go Book2 Book3]
}

func showPerson() {
	tom := person{name: "Tom", age: 24}
	tom.print()                   // Name: Tom, затем Age: 24 на новой строке
	tom.eat("apples")             // Tom eats apples
	fmt.Println(tom.currentAge()) // 24
}

func showReceivers() {
	tom := person{name: "Tom", age: 24}
	fmt.Println("before", tom.age) // before 24
	tom.updateAgeOnCopy(33)        // inside 33
	fmt.Println("after", tom.age)  // after 24

	older := tom.withAge(33)
	fmt.Println(tom.age, older.age) // 24 33

	var pointer *person = &tom
	fmt.Println("before", tom.age) // before 24
	pointer.updateAge(33)
	fmt.Println("after", tom.age) // after 33
	tom.updateAge(34)             // Сокращённая форма вызова (&tom).updateAge(34).
	fmt.Println(tom.age)          // 34
}

func showAddressabilityAndInterfaces() {
	tom := person{name: "Tom", age: 33}
	team := []person{tom}
	team[0].updateAge(35)
	fmt.Println(team[0].age, tom.age) // 35 33

	people := map[string]person{"tom": tom}
	// people["tom"].updateAge(36) // Ошибка компиляции: элемент карты неадресуем.
	fmt.Println(people["tom"].currentAge()) // 33
	// (person{name: "Bob"}).updateAge(40) // Ошибка компиляции: литерал неадресуем.

	// var updater ageUpdater = tom // Ошибка компиляции: присваивание интерфейсу не подставляет &.
	var updater ageUpdater = &tom
	updater.updateAge(36)
	fmt.Println(tom.age) // 36
}

func showMethodFunctions() {
	tom := person{age: 24}
	readAge := tom.currentAge // Сохраняет копию tom при создании значения метода.
	update := tom.updateAge   // Сохраняет адрес tom при создании значения метода.
	update(33)
	fmt.Println(readAge(), tom.age) // 24 33

	readFrom := person.currentAge
	updateFor := (*person).updateAge
	updateFor(&tom, 40)                     // Выражение метода получает указатель первым аргументом.
	fmt.Println(readFrom(tom))              // 40
	fmt.Println((*person).currentAge(&tom)) // 40
	// _ = person.updateAge // Ошибка компиляции: в наборе person этого метода нет.
}

func showNilReceiver() {
	var missing *person
	fmt.Println(missing.nameOrUnknown()) // unknown
	// missing.updateAge(33) // Паника при записи в поле через nil.
	// missing.currentAge() // Паника при разыменовании nil до входа в метод.

	var namer interface{ nameOrUnknown() string } = missing
	fmt.Println(namer == nil, namer.nameOrUnknown()) // false unknown
}

type shelf[T any] []T

func (s shelf[T]) collect[R any](convert func(T) R) []R {
	result := make([]R, 0, len(s))
	for _, item := range s {
		result = append(result, convert(item))
	}
	return result
}

func showGenericMethod() {
	books := shelf[string]{"Go", "Methods"}
	lengths := books.collect[int](func(title string) int { return len(title) })
	fmt.Println(lengths) // [2 7]
}
