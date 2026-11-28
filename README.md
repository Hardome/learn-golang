# Урок 019. Методы в Go

**Метод** — функция, связанная с определённым типом. У метода есть получатель (receiver) —
один дополнительный параметр перед именем метода. При вызове `value.method(arg)` выражение
`value` передаётся в получатель, а `arg` — в обычный параметр. Тип получателя определяет,
с чем работает метод внутри: с копией значения или с копией указателя.

Примеры из [hello.go](hello.go) запускаются командой `go run hello.go` на Go 1.27 или новее.
Эта версия нужна для метода с собственным параметром типа.

## Объявление и вызов метода

Получатель указан перед именем метода. В следующем примере `p` — получатель, а `meal` — обычный параметр.

### Методы структуры, параметры и результат

```go
type person struct {
    name string
    age  int
}

func (p person) print() {
    fmt.Println("Name:", p.name) // Читаем поля конкретного получателя p.
    fmt.Println("Age:", p.age)
}

func (p person) eat(meal string) {
    fmt.Println(p.name, "eats", meal) // Имя берём из получателя, еду — из параметра.
}

func (p person) currentAge() int {
    return p.age // Метод может возвращать результат как обычная функция.
}

// Внутри функции.
tom := person{name: "Tom", age: 24}
tom.print()                   // Name: Tom, затем Age: 24 на новой строке.
tom.eat("apples")             // Tom eats apples. Получатель — tom, аргумент — "apples".
fmt.Println(tom.currentAge()) // 24
```

### Для каких типов можно объявлять методы

Метод объявляется на уровне пакета. Базовый тип получателя должен быть определён в том же
пакете и не может быть указателем или интерфейсом. В объявлении получателя можно записать
сам тип `T` или указатель `*T`. У одного базового типа имена методов уникальны. У структуры
имя метода также не может совпасть с именем её поля.

```go
type counter int

func (c counter) doubled() int { return int(c) * 2 }

// func (c *counter) doubled() int { return int(*c) * 2 } // Ошибка: имя doubled уже занято.
// func (n int) doubled() int { return n * 2 }           // Ошибка: int определён вне пакета.
// func (values []string) print() {}                     // Ошибка: []string неименованный тип.

// Внутри функции.
fmt.Println(counter(5).doubled()) // 10
```

`library.print` и `person.print` могут сосуществовать, поскольку принадлежат разным базовым типам.
Интерфейс перечисляет методы в своём объявлении, а реализацию предоставляет конкретный тип.
Псевдоним чужого типа не позволяет добавить к нему методы. Обычный псевдоним собственного
необобщённого типа может использоваться в получателе. Метод тогда связывается с исходным типом.

## Получатель-значение и получатель-указатель

### Получатель `person`: изменение остаётся в копии

```go
func (p person) updateAgeOnCopy(newAge int) {
    p.age = newAge // Меняем поле локальной копии.
    fmt.Println("inside", p.age)
}

// Внутри функции.
tom := person{name: "Tom", age: 24}
fmt.Println("before", tom.age) // before 24
tom.updateAgeOnCopy(33)        // inside 33
fmt.Println("after", tom.age)  // after 24. Исходное поле осталось прежним.
```

Если нужна изменённая копия, метод с получателем-значением может её вернуть. Вызывающая
сторона затем решает, сохранить ли её в другой переменной или присвоить обратно.

```go
func (p person) withAge(newAge int) person {
    p.age = newAge
    return p // Возвращаем изменённую копию.
}

// Внутри функции.
tom := person{name: "Tom", age: 24}
older := tom.withAge(33)
fmt.Println(tom.age, older.age) // 24 33
```

### Получатель `*person`: изменение видно в исходной структуре

```go
func (p *person) updateAge(newAge int) {
    p.age = newAge // Сокращение для (*p).age = newAge. Пишем по адресу.
}

// Внутри функции.
tom := person{name: "Tom", age: 24}
var pointer *person = &tom
fmt.Println("before", tom.age) // before 24
pointer.updateAge(33)          // В получатель p копируется адрес tom.
fmt.Println("after", tom.age)  // after 33

tom.updateAge(34)     // tom адресуемая. Go вызывает (&tom).updateAge(34).
fmt.Println(tom.age) // 34
```

Оба получателя передаются по значению. В `p person` копируется структура, а в `p *person` —
адрес. Присваивание нового адреса самой переменной `p` внутри метода меняет только локальный
параметр. Присваивание в `p.age` меняет объект по сохранённому адресу.

### Получатель-значение может менять общие данные

```go
type library []string

func (books library) replaceFirst(title string) {
    books[0] = title // Запись в общий массив. У среза должен быть хотя бы один элемент.
}

func (books library) keepFirst() {
    books = books[:1] // Меняем длину только локального описания среза.
    fmt.Println(len(books))
}

func (books *library) add(title string) {
    *books = append(*books, title) // Записываем новое описание среза через указатель.
}

// Внутри функции.
lib := library{"Book1", "Book2"}
lib.replaceFirst("Go")
fmt.Println(lib)      // [Go Book2]
lib.keepFirst()       // 1
fmt.Println(len(lib)) // 2
lib.add("Book3")
fmt.Println(lib) // [Go Book2 Book3]
```

При копировании среза копируется его описание, и оба среза продолжают обращаться к одному
массиву. При замене элемента меняется массив. Чтобы сохранить для вызывающего новую длину
после `append`, верните новый срез или запишите его через указатель, как в `add`.

### Как выбрать получателя

- Получатель-значение подходит, если значение легко копировать и метод должен работать с копией.
- Получатель-указатель подходит, если метод меняет исходный объект или объект нельзя копировать.

## Адресуемость и набор методов

### Где Go автоматически берёт адрес

```go
tom := person{name: "Tom", age: 24}
tom.updateAge(33) // Вызов через адресуемую переменную.
(&person{name: "Bob"}).updateAge(40) // Явно берём адрес составного литерала.

team := []person{tom}
team[0].updateAge(35)          // Элемент среза адресуем.
fmt.Println(team[0].age)       // 35
fmt.Println(tom.age)           // 33. В срез положили копию значения tom.

people := map[string]person{"tom": tom}
// people["tom"].updateAge(36) // Ошибка: элемент карты неадресуем.
fmt.Println(people["tom"].currentAge()) // 33. Получателю-значению адрес не нужен.

// (person{name: "Bob"}).updateAge(40) // Ошибка: литерал сам по себе неадресуем.
```

Результат функции типа `person` тоже неадресуем. Сохраните его в переменную или возвращайте
`*person`, если функция должна выдать указатель. Для карты с элементами `*person` значение указателя
можно передать в метод напрямую. Сам элемент карты от этого адресуемым не становится.

### Набор методов при присваивании интерфейсу

**Набор методов** определяет, какие методы принадлежат типу. Для определённого типа `T` в набор
входят методы с получателем `T`. Для `*T` в набор входят методы с получателем `T` и `*T`.
Продвинутые методы встроенных полей тоже могут входить в набор.

```go
type ageReader interface{ currentAge() int }
type ageUpdater interface{ updateAge(int) }

var _ ageReader = person{}  // currentAge входит в набор person.
var _ ageReader = &person{} // currentAge входит и в набор *person.
var _ ageUpdater = &person{}
// var _ ageUpdater = person{} // Ошибка: updateAge отсутствует в наборе person.

// Внутри функции.
tom := person{age: 24}
tom.updateAge(33) // Вызов разрешён: Go использует &tom.

// var updater ageUpdater = tom // Ошибка: при присваивании интерфейсу & не подставляется.
var updater ageUpdater = &tom
updater.updateAge(34)
fmt.Println(tom.age) // 34
```

Встраивание `T` продвигает его методы с получателем `T` в наборы `S` и `*S`, а методы
с получателем `*T` — в набор `*S`. Встраивание `*T` может продвинуть обе группы в оба набора.
Пример разрешения конфликтов имён есть в [уроке 018 о структурах][structs].

## Значение метода и выражение метода

### Значение метода сохраняет получатель

```go
tom := person{age: 24}
readAge := tom.currentAge  // func() int. Сейчас сохраняется копия person.
update := tom.updateAge    // func(int). Сейчас сохраняется указатель &tom.
update(33)

fmt.Println(readAge(), tom.age) // 24 33
```

Получатель вычисляется при создании значения метода, ещё до вызова сохранённой функции.
При получателе-указателе сохраняется конкретный адрес. Если позже записать другой адрес
в переменную-указатель, сохранённая функция продолжит обращаться по прежнему адресу.

```go
tom, bob := person{age: 24}, person{age: 30}
pointer := &tom
update := pointer.updateAge // Сохраняем текущий адрес, то есть &tom.
pointer = &bob              // Меняем только переменную pointer.
update(33)
fmt.Println(tom.age, bob.age) // 33 30
```

### Выражение метода принимает получатель явно

```go
tom := person{age: 24}
readAge := person.currentAge  // func(person) int.
update := (*person).updateAge // func(*person, int).
fmt.Println(readAge(tom))     // 24. Получатель передан первым аргументом.

update(&tom, 33)
readFromPointer := (*person).currentAge // func(*person) int. Разыменует переданный адрес.
fmt.Println(readFromPointer(&tom))      // 33

// _ = person.updateAge // Ошибка: в наборе person метода updateAge нет.
```

## Вызовы через `nil`

```go
func (p *person) nameOrUnknown() string {
    if p == nil {
        return "unknown" // Метод сам обрабатывает nil до чтения полей.
    }
    return p.name
}

// Внутри функции.
var missing *person
fmt.Println(missing.nameOrUnknown()) // unknown
// missing.updateAge(33)            // Паника внутри метода при записи в поле через nil.
// missing.currentAge()             // Паника до входа в метод. Для получателя нужен person.

var namer interface{ nameOrUnknown() string } = missing
fmt.Println(namer == nil, namer.nameOrUnknown()) // false unknown
```

У интерфейса `namer` есть динамический тип `*person`, поэтому интерфейс отличен от `nil`,
даже когда сам указатель равен `nil`. Вызов метода у интерфейса без динамического типа
вызывает панику.

## Методы и параметры типов

С Go 1.18 метод обобщённого типа может использовать параметры получателя. Например, у типа
`stack[T]` метод `func (s *stack[T]) push(value T)` принимает значение того же типа `T`.

### Собственные параметры типа метода с Go 1.27

```go
type shelf[T any] []T

func (s shelf[T]) collect[R any](convert func(T) R) []R {
    result := make([]R, 0, len(s))
    for _, item := range s {
        result = append(result, convert(item)) // T приходит от shelf, R задаёт метод.
    }
    return result
}

// Внутри функции.
books := shelf[string]{"Go", "Methods"}
lengths := books.collect[int](func(title string) int { return len(title) })
fmt.Println(lengths) // [2 7]
```

У `books` тип `shelf[string]`, поэтому `T` в получателе обозначает `string`. Метод отдельно объявляет
`R`, который в вызове `books.collect[int](...)` равен `int`. Функция `convert` принимает элементы среза
и возвращает значения нового типа. В примере это длины строк в байтах.

В Go 1.26 и раньше такую операцию объявляют отдельной обобщённой функцией с параметрами
`T` и `R`. Интерфейсные методы не объявляют собственные параметры типов. Обобщённый метод
с собственными параметрами не реализует метод интерфейса.

## Источники

- [Спецификация Go — объявления методов и ограничения получателя](https://go.dev/ref/spec#Method_declarations)
- [Спецификация Go — вызовы и автоматическое взятие адреса](https://go.dev/ref/spec#Calls)
- [Спецификация Go — селекторы и nil](https://go.dev/ref/spec#Selectors)
- [Спецификация Go — наборы методов](https://go.dev/ref/spec#Method_sets)
- [Спецификация Go — реализация интерфейсов](https://go.dev/ref/spec#Implementing_an_interface)
- [Спецификация Go — встраивание и продвижение методов](https://go.dev/ref/spec#Struct_types)
- [Спецификация Go — операции с адресом](https://go.dev/ref/spec#Address_operators)
- [Спецификация Go — срезы и общий массив](https://go.dev/ref/spec#Slice_types)
- [Спецификация Go — значения методов](https://go.dev/ref/spec#Method_values)
- [Спецификация Go — выражения методов](https://go.dev/ref/spec#Method_expressions)
- [Go Code Review Comments — выбор типа получателя](https://go.dev/wiki/CodeReviewComments#receiver-type)
- [Заметки о выпуске Go 1.18 — обобщённые типы](https://go.dev/doc/go1.18#language)
- [Заметки о выпуске Go 1.27 — обобщённые методы](https://go.dev/doc/go1.27#language)

[structs]: https://github.com/Hardome/learn-golang/tree/lesson-018-structs
