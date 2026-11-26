package main

import "fmt"

type person struct {
	name string
	age  int
}

type account struct {
	login   string
	profile person
}

type member struct {
	login string
	person
}

type node struct {
	value int
	next  *node
}

// type badNode struct{ next badNode } // Ошибка компиляции: рекурсия через поле по значению.

func main() {
	showInitialization()
	showUnnamedStructs()
	showPointers()
	showCopies()
	showComparison()
	showEmbedding()
	showLinkedList()
}

func showInitialization() {
	// Каждое поле получает нулевое значение своего типа.
	var zero person
	fmt.Printf("%q %d\n", zero.name, zero.age) // "" 0

	tom := person{"Tom", 23}
	alice := person{age: 23, name: "Alice"}
	bob := person{name: "Bob"}
	var empty = person{}
	fmt.Println(tom, alice, bob, empty) // {Tom 23} {Alice 23} {Bob 0} { 0}
	// _ = person{"Tom"} // Ошибка компиляции: позиционный литерал должен задать все поля.
	// _ = person{"Tom", age: 23} // Ошибка компиляции: нельзя смешивать формы литерала.

	tom.age = 38
	fmt.Println(tom.name, tom.age) // Tom 38

	// Этот тип объявлен в блоке функции. Два поля одного типа записаны вместе.
	type employee struct {
		name, company string
		age           int
	}
	var employeeTom employee = employee{name: "Tom", company: "Acme", age: 38}
	fmt.Println(employeeTom.company) // Acme
}

func showUnnamedStructs() {
	// Неименованный тип можно записать при объявлении переменной.
	var tom struct {
		name string
		age  int
	}
	tom.name = "Tom"
	tom.age = 23

	bob := struct {
		name string
		age  int
	}{
		name: "Bob",
		age:  41,
	}
	// Базовые типы совпадают, а тип bob неименованный.
	var named person = bob
	fmt.Println(tom.name, tom.age, named.name, named.age) // Tom 23 Bob 41
}

func showPointers() {
	tom := person{name: "Tom", age: 22}
	var pointer *person = &tom
	// Селектор автоматически разыменовывает указатель при обращении к полю.
	fmt.Println(pointer.name, (*pointer).age) // Tom 22
	pointer.age = 23
	fmt.Println(tom.age) // 23
	age := &tom.age
	*age = 35
	fmt.Println(tom.age) // 35

	// new(person) возвращает указатель на новую переменную с нулевым значением.
	zero := new(person)
	fmt.Println(*zero) // { 0}
	zero.name = "Tom"
	zero.age = 41
	fmt.Println(*zero) // {Tom 41}

	alice := &person{name: "Alice", age: 23}
	fmt.Println(*alice) // {Alice 23}

	anonymous := new(struct {
		name, company string
		age           int
	})
	anonymous.name = "Bob"
	anonymous.company = "Acme"
	anonymous.age = 46
	fmt.Println(*anonymous) // {Bob Acme 46}

	// var missing *person
	// fmt.Println(missing.age) // Паника: чтение поля через nil.

	// Запись в карту обновляет целое значение. Поле элемента карты неадресуемо.
	people := map[string]person{"tom": tom}
	// people["tom"].age++ // Ошибка компиляции: поле элемента карты неадресуемо.
	updated := people["tom"]
	updated.age++
	people["tom"] = updated
	fmt.Println(people["tom"].age, tom.age) // 36 35
}

func showCopies() {
	tom := person{name: "Tom", age: 41}
	copyOfTom := tom
	copyOfTom.age = 18
	fmt.Println(tom.age, copyOfTom.age) // 41 18

	older := nextBirthday(tom)
	fmt.Println(tom.age, older.age) // 41 42
	incrementAge(&tom)
	fmt.Println(tom.age) // 42

	type profile struct {
		owner person
		tags  []string
	}
	original := profile{owner: tom, tags: []string{"go"}}
	copied := original
	copied.owner.age = 18
	// Поле owner скопировано по значению. Оба среза по-прежнему используют общий массив.
	copied.tags[0] = "struct"
	fmt.Println(original.owner.age, copied.owner.age, original.tags[0]) // 42 18 struct

	copied.tags = []string{"copy"}
	fmt.Println(original.tags[0], copied.tags[0]) // struct copy
}

func nextBirthday(value person) person {
	value.age++
	return value
}

func incrementAge(value *person) {
	value.age++
}

func showComparison() {
	tom := person{name: "Tom", age: 41}
	bob := person{name: "Bob", age: 41}
	tomas := person{name: "Tom", age: 41}
	fmt.Println(tom == bob, tom == tomas, &tom == &tomas) // false true false

	// Ключи сравниваются по значению, поэтому tomas находит запись для tom.
	roleByPerson := map[person]string{tom: "admin"}
	fmt.Println(roleByPerson[tomas]) // admin

	// _ = struct{ tags []string }{} == struct{ tags []string }{} // Ошибка компиляции: поле-срез.
	// entry := struct{ value interface{} }{value: []int{1}}
	// fmt.Println(entry == entry) // Паника: динамический тип []int несравним.
}

func showEmbedding() {
	tom := account{login: "tom", profile: person{name: "Tom", age: 41}}
	alice := member{login: "alice", person: person{name: "Alice", age: 23}}
	fmt.Println(tom.profile.name)              // Tom
	fmt.Println(alice.name, alice.person.name) // Alice Alice

	// Имена встроенных полей здесь совпадают с именами их типов.
	record := struct {
		string
		company string
		int
	}{string: "Tom", company: "Acme", int: 41}
	fmt.Println(record.string, record.company, record.int) // Tom Acme 41
	// type duplicate struct{ string; string } // Ошибка компиляции: два поля string.

	type nickname struct{ name string }
	colliding := struct {
		person
		nickname
	}{person: alice.person, nickname: nickname{name: "Al"}}
	// fmt.Println(colliding.name) // Ошибка компиляции: два поля name на одной глубине.
	fmt.Println(colliding.person.name, colliding.nickname.name) // Alice Al

	var linked struct{ *person }
	fmt.Println(linked.person == nil) // true
	// fmt.Println(linked.name) // Паника: продвинутое поле читается через nil.
}

func showLinkedList() {
	first := node{value: 4}
	second := node{value: 5}
	third := node{value: 6}
	first.next = &second
	second.next = &third

	// У последнего узла next равен nil. Цикл выводит значения по одному на строку.
	for current := &first; current != nil; current = current.next {
		fmt.Println(current.value) // 4, 5, 6 на отдельных строках
	}
}
