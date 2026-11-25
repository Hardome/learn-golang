# Урок 016. Указатели в Go

Указатель хранит адрес переменной. Например, значение типа `*int` может указывать на переменную типа `int`.
Оператор `&` берёт адрес, а оператор `*` позволяет прочитать или изменить переменную по этому адресу.

Примеры можно запустить из [hello.go](hello.go) командой `go run hello.go`.

## Адрес и разыменование

```go
score := 10
pointer := &score          // тип *int
fmt.Println(*pointer)     // 10

alias := pointer
*alias = 15
fmt.Println(score, *pointer, pointer == alias) // 15 15 true
```

В `*int` звёздочка входит в запись типа. В `*pointer` она обозначает разыменование. Копия указателя
содержит тот же адрес, поэтому изменение через `alias` видно через `pointer` и через имя `score`.

Два указателя одного типа равны, если указывают на одну переменную или оба равны `nil`. Указатели на разные
переменные нулевого размера, например `struct{}`, могут быть равны. Для обычных указателей доступны
`==` и `!=`. Арифметика вроде `pointer + 1` не поддерживается. Число типа `uintptr` имеет целочисленную
семантику и само по себе не удерживает объект от сборки мусора.

## Нулевое значение и создание указателя

Нулевое значение типа `*T` — `nil`. Оно означает, что указатель никуда не указывает. Разыменование `nil`
вызывает панику. Это касается и выражения `&*pointer`, поскольку оно сначала вычисляет `*pointer`.

```go
var pointer *int
fmt.Println(pointer == nil) // true
// fmt.Println(*pointer)    // паника

pointer = new(int)
fmt.Println(*pointer) // 0
```

`new(T)` создаёт переменную с нулевым значением и возвращает `*T`. Например, `new([]int)` возвращает указатель
на нулевой срез, а `make([]int, 2)` — сам срез с двумя элементами. Для указателя на структуру с заданными
полями удобна запись `&Point{X: 3}`. Начиная с Go 1.26, `new(42)` также допустимо и возвращает `*int`,
указывающий на новую переменную со значением 42.

Поле типа `*int` удобно, когда нужно различать отсутствие значения (`nil`) и явный ноль (указатель на 0).
Если положить `nil` типа `*int` в интерфейс `any`, интерфейс будет содержать динамический тип `*int`
и сам уже не будет равен `nil`.

## Указатель как параметр

Go передаёт аргументы по значению. Параметр типа `*int` получает копию адреса, поэтому через него можно изменить
переменную вызывающего кода. Присваивание самому параметру меняет только локальную копию указателя.

```go
func increment(target *int) {
    *target = *target + 1
}

func redirect(target, replacement *int) *int {
    target = replacement
    return target
}
```

```go
score, backup := 10, 99
pointer := &score
increment(pointer)
fmt.Println(score) // 11

redirect(pointer, &backup)
fmt.Println(*pointer) // 11
pointer = redirect(pointer, &backup)
fmt.Println(*pointer) // 99
```

Если функция должна заменить сам указатель вызывающего кода, обычно проще вернуть новый указатель.
Можно также передать адрес переменной-указателя. Например, при `slot := &pointer` переменная `slot`
имеет тип `**int`, и присваивание `*slot = &score` меняет `pointer`.

Для небольших значений удобно принимать значение и возвращать результат. Указатель полезен для изменения
существующего объекта, представления отсутствия значения или уменьшения стоимости копирования больших структур.
При совместном использовании изменяемого объекта из нескольких горутин нужен согласованный доступ к нему.

## У каких выражений можно взять адрес

Можно взять адрес переменной, адресуемого поля структуры, элемента адресуемого массива или элемента среза.
Литерал составного типа — отдельное исключение. Например, `&Point{X: 3}` создаёт новую переменную сразу.

```go
type Point struct{ X int }

point := &Point{X: 3}
point.X = 4 // то же, что (*point).X = 4
fmt.Println(point.X) // 4

numbers := []int{1, 2}
element := &numbers[0]
*element = 7
fmt.Println(numbers) // [7 2]
```

У константы, строкового элемента, результата вызова функции и элемента карты адрес взять нельзя. Например,
`&42`, `&len(numbers)` и `&scores["go"]` не компилируются. Чтобы хранить указатели в карте, можно объявить
`map[string]*int` и явно положить туда адрес переменной.

В `for _, value := range numbers` выражение `&value` даёт адрес переменной цикла, в которую скопировали элемент.
Для указателя на сам элемент используй `&numbers[index]`. Если `append` выделит новый базовый массив,
старый указатель останется привязан к элементу прежнего массива.

## Время жизни

Можно вернуть адрес локальной переменной. Она остаётся доступной, пока её можно достичь через указатель.

```go
func newScore(initial int) *int {
    score := initial
    return &score
}
```

Выбор стека или кучи делает компилятор, в том числе с помощью анализа убегания. Само использование `&` или `new`
не гарантирует размещение в куче. Диагностику текущего компилятора можно посмотреть командой
`go build -gcflags=-m hello.go`.

## Источники

- [Go Language Specification — Pointer types](https://go.dev/ref/spec#Pointer_types)
- [Go Language Specification — Address operators](https://go.dev/ref/spec#Address_operators)
- [Go Language Specification — Comparison operators](https://go.dev/ref/spec#Comparison_operators)
- [Go Language Specification — Allocation](https://go.dev/ref/spec#Allocation)
- [Go 1.26 Release Notes — Changes to the language](https://go.dev/doc/go1.26#language)
- [Go Language Specification — Calls](https://go.dev/ref/spec#Calls)
- [Go Language Specification — Selectors](https://go.dev/ref/spec#Selectors)
- [Go Language Specification — For statements](https://go.dev/ref/spec#For_statements)
- [Go Language Specification — Appending and copying slices](https://go.dev/ref/spec#Appending_and_copying_slices)
- [Go FAQ — nil внутри интерфейса](https://go.dev/doc/faq#nil_error)
- [Go FAQ — размещение переменных](https://go.dev/doc/faq#stack_or_heap)
- [Go Memory Model — Introduction](https://go.dev/ref/mem#introduction)
- [unsafe.Pointer — отличие от uintptr](https://pkg.go.dev/unsafe#Pointer)
