# Урок 018. Структуры в Go

Структура в Go — составной тип данных, который задаёт упорядоченный набор полей. У каждого поля
есть имя и тип. Поля описывают свойства объекта.

Ключевое слово `struct` задаёт набор полей, а `type` позволяет определить именованный тип на его основе.

Примеры для запуска собраны в [hello.go](hello.go). Команда запуска — `go run hello.go`.

## Объявление структуры

```go
type person struct {
    name string // Имя человека.
    age  int    // Возраст в годах.
}
```

`struct { ... }` описывает поля. `type person` даёт описанной структуре имя. После этого
`person` можно использовать в объявлениях переменных, полей и параметров функций.

Имена полей в одной структуре уникальны, кроме `_`: такое поле нельзя выбрать через точку.
Поле доступно из другого пакета, если его имя начинается с заглавной буквы. Объявление типа на уровне пакета
видно во всех файлах пакета, а объявление внутри функции действует в своём блоке.

## Создание значения и изменение полей

### Нулевое значение

Каждое поле получает нулевое значение своего типа. Если поле само является структурой или массивом,
это правило применяется к его полям или элементам.

```go
var tom person
fmt.Printf("%q %d\n", tom.name, tom.age) // "" 0
fmt.Println(tom == (person{}))          // true, пустой литерал даёт то же значение.
```

### Позиционный литерал и литерал с именами

```go
tom := person{"Tom", 23}                 // Первый элемент идёт в name, второй — в age.
alice := person{age: 23, name: "Alice"}   // По именам поля можно перечислить в любом порядке.
bob := person{name: "Bob"}                // Пропущенное поле age получает 0.
fmt.Println(tom, alice, bob)              // {Tom 23} {Alice 23} {Bob 0}

// _ = person{"Tom"}          // Ошибка компиляции: в позиционном литерале нужны все поля.
// _ = person{"Tom", age: 23}  // Ошибка компиляции: нельзя смешивать позиции и имена.
```

Имена в литерале помогают различить поля одного типа. Неэкспортируемое поле структуры из другого пакета нельзя
явно заполнить в литерале.

### Чтение и запись через точку

```go
tom := person{name: "Tom", age: 24}
fmt.Println(tom.name, tom.age) // Tom 24

tom.age = 38                  // Записываем новое значение в поле age переменной tom.
fmt.Println(tom.name, tom.age) // Tom 38
```

## Неименованная структура

Тип можно записать прямо в объявлении или литерале. К полям такой структуры обращаются через точку.

### Когда типы совместимы

```go
value := struct{ name string; age int }{name: "Tom", age: 23}
var tom person = value // Базовые типы совпадают, а тип value неименованный.

type visitor struct{ name string; age int }
// var guest visitor = tom // Ошибка компиляции: оба типа определены отдельно.
guest := visitor(tom) // Явное преобразование допустимо.
fmt.Println(guest.name, guest.age) // Tom 23
```

Порядок полей, имена, типы, теги и признак встраивания участвуют в определении идентичности
типов структур. Неэкспортируемые имена из разных пакетов считаются разными.

## Указатели и адресуемость

### Указатель на структуру и указатель на поле

```go
tom := person{name: "Tom", age: 22}
var pointer *person = &tom // В pointer записан адрес переменной tom.

fmt.Println(pointer.name, (*pointer).age) // Tom 22
pointer.age = 23                          // То же, что (*pointer).age = 23.
fmt.Println(tom.age)                      // 23

age := &tom.age // У age тип *int. Он указывает только на поле age.
*age = 35
fmt.Println(tom.age) // 35

// var missing *person
// fmt.Println(missing.age) // Паника: чтение поля через nil.
```

В выражении `pointer.age` Go разыменовывает указатель автоматически. Чтение и запись через него
обращаются к полю той же переменной `tom`. Указатель `age` позволяет обращаться к полю напрямую.

### Поле элемента карты и поле элемента среза

```go
people := map[string]person{"tom": {name: "Tom", age: 22}}
// people["tom"].age++  // Ошибка компиляции: элемент карты неадресуем.
// _ = &people["tom"].age // По той же причине адрес его поля взять нельзя.
tom := people["tom"] // Получаем копию значения из карты.
tom.age++
people["tom"] = tom  // Записываем изменённое значение целиком.
fmt.Println(people["tom"].age) // 23

team := []person{tom}
team[0].age++ // Элемент среза адресуем. Поле можно обновить на месте.
fmt.Println(team[0].age, tom.age) // 24 23. В team скопировали значение tom.
```

```go
tom := person{name: "Tom", age: 22}
people := map[string]*person{"tom": &tom}
people["tom"].age++ // Из карты читаем адрес и меняем поле переменной tom.
fmt.Println(tom.age) // 23

// people["missing"].age++ // Паника: для отсутствующего ключа карта вернёт nil.
```

## `new` и адрес литерала

`new(person)` создаёт нулевую структуру и возвращает указатель. Выражение `&person{...}` сразу задаёт поля.
С Go 1.26 `new(value)` создаёт отдельную переменную и копирует туда переданное значение.

```go
zero := new(person)
fmt.Println(zero.age) // 0, поле получило нулевое значение.

tom := &person{name: "Tom", age: 23}
fmt.Println(tom.name, tom.age) // Tom 23

copied := new(*tom) // Go 1.26+. Копируется структура по адресу tom.
copied.age++
fmt.Println(tom.age, copied.age) // 23 24
```

Подробнее о том, где хранится переменная и как работают указатели, см. в уроке 016.

## Копирование и параметры функций

### Изменение простого поля в копии

```go
tom := person{name: "Tom", age: 41}
tomas := tom   // В tomas копируются оба поля.
tomas.age = 18 // Записываем только в поле переменной tomas.
fmt.Println(tom.age, tomas.age) // 41 18
```

Массивы и вложенные структуры также копируются по значению вместе со своими элементами и полями.

### Изменение данных, на которые ссылается поле

```go
original := struct{ tags []string }{tags: []string{"go"}}
copied := original    // Скопировано описание среза. Массив у срезов общий.
copied.tags[0] = "struct"
fmt.Println(original.tags[0], copied.tags[0]) // struct struct

copied.tags = []string{"copy"} // В поле copied записываем уже другой срез.
fmt.Println(original.tags[0], copied.tags[0]) // struct copy
```

При копировании поля-указателя оба поля указывают на одну переменную. Скопированное поле-карта
обращается к той же карте. Если нужны независимые данные, отдельно копируйте данные
за такими полями. Структуру с полем `sync.Mutex` нельзя копировать после первого использования
мьютекса.

### Передача структуры и указателя в функцию

```go
func nextBirthday(value person) person {
    value.age++  // Меняем локальную копию. Чтобы получить её снаружи, возвращаем значение.
    return value
}

func incrementAge(value *person) {
    value.age++ // Через копию указателя меняем исходного человека.
}

// Внутри вызывающей функции.
tom := person{name: "Tom", age: 41}
older := nextBirthday(tom)
fmt.Println(tom.age, older.age) // 41 42

incrementAge(&tom)   // Передаём адрес переменной tom.
fmt.Println(tom.age) // 42
```

Параметр-значение подходит для небольших самостоятельных данных. Параметр-указатель подходит,
если функция меняет исходный объект или если значение нельзя копировать.

## Сравнение структур

### Поля, адреса и ключи карты

```go
tom := person{name: "Tom", age: 41}
bob := person{name: "Bob", age: 41}
tomas := person{name: "Tom", age: 41}

fmt.Println(tom == bob)     // false. Поля name различаются, хотя возраст совпадает.
fmt.Println(tom == tomas)   // true. Совпадают оба поля.
fmt.Println(&tom == &tomas) // false. Адреса принадлежат разным переменным.

roles := map[person]string{tom: "admin"}
fmt.Println(roles[tomas]) // admin. Для поиска сравниваются значения ключей.
```

Операторы `==` и `!=` доступны, если типы всех полей сравнимы. Сравнение идёт в порядке объявления
полей до первого различия. Поля `_` пропускаются при сравнении значений, но их типы участвуют
в проверке сравнимости. Для структур нет операций порядка `<` и `>`. Указатели на разные
переменные нулевого размера могут сравниваться как равные.

### Ошибка компиляции и паника при сравнении

```go
type profile struct{ tags []string }
left, right := profile{}, profile{}
// _ = left == right // Ошибка компиляции: тип поля []string несравним, даже если оба среза nil.
fmt.Println(left.tags == nil, right.tags == nil) // true true. Срез можно сравнить с nil.
```

Поле-срез, поле-карта или поле-функция делает структуру несравнимой. Это правило действует
и для таких полей внутри вложенной структуры или массива. Если сравнение нужно, определите
его явно, например поэлементно для срезов.

```go
type entry struct {
    id    int
    value interface{}
}
first := entry{id: 1, value: []int{1}}
second := entry{id: 2, value: []int{1}}
fmt.Println(first == second) // false. Разный id, поэтому до value сравнение не дошло.

second.id = 1
// fmt.Println(first == second) // Паника: оба поля value содержат несравнимый тип []int.
```

Интерфейсное поле формально сравнимо. Паника возникает, если до него дошло сравнение, динамический
тип совпал и он несравним. По той же причине хеширование структурного ключа с таким полем может
вызвать панику при обращении к карте.

## Именованное вложение и встраивание

### Именованное поле требует явного пути

```go
type account struct {
    login   string
    profile person // Имя поля — profile, его тип — person.
}

// Внутри функции.
tom := account{
    login:   "tom",
    profile: person{name: "Tom", age: 41}, // Заполняем вложенную структуру целиком.
}
fmt.Println(tom.profile.name) // Tom. Сначала выбираем profile, затем его поле name.
// fmt.Println(tom.name) // Ошибка компиляции: поле profile именованное, его поля не продвигаются.
```

### Встроенное поле даёт короткий доступ

```go
type member struct {
    login string
    person // Встроенное поле. Его имя берётся из имени типа person.
}

// Внутри функции.
alice := member{login: "alice", person: person{name: "Alice", age: 23}}
fmt.Println(alice.person.name) // Alice. Полный путь к полю.
fmt.Println(alice.name)        // Alice. Продвинутый селектор ведёт к тому же полю.
alice.age++
fmt.Println(alice.person.age) // 24. Короткая запись изменила возраст во вложенном person.
```

Встраивать можно имя типа `T` или `*T`. Имя поля берётся без звёздочки и имени пакета.
`T` не может быть указательным типом или параметром типа. Интерфейс можно встроить по имени,
а указатель на интерфейс встроить нельзя.

```go
type record struct {
    string         // Встроенное поле с именем string.
    company string // Второе поле того же типа получило явное имя.
    int            // Встроенное поле с именем int.
}

// Внутри функции.
tom := record{string: "Tom", company: "Acme", int: 41}
fmt.Println(tom.string, tom.company, tom.int) // Tom Acme 41

// type duplicate struct{ string; string } // Ошибка: два поля с именем string.
// type duplicate struct{ person; *person } // Ошибка: у обоих полей имя person.
```

### Как Go выбирает поле при совпадении имён

```go
type nickname struct{ name string }

type badge struct {
    person
    nickname
}

// Внутри функции.
tom := badge{person: person{name: "Tom"}, nickname: nickname{name: "T"}}
// fmt.Println(tom.name) // Ошибка компиляции: person.name и nickname.name на одной глубине.
fmt.Println(tom.person.name, tom.nickname.name) // Tom T. Полные пути однозначны.
```

Go выбирает единственный подходящий селектор на наименьшей глубине. Собственное поле внешней
структуры имеет глубину ноль, поэтому оно скрывает одноимённое продвинутое поле.

```go
tom := struct {
    name string // Это поле находится прямо во внешней структуре.
    person
}{name: "Team Tom", person: person{name: "Tom"}}

fmt.Println(tom.name, tom.person.name) // Team Tom Tom
```

### Продвинутый селектор как ключ литерала

```go
alice := member{person: person{name: "Alice"}} // Явно задаём встроенное поле целиком.
fmt.Println(alice.name) // Alice

bob := member{name: "Bob"} // Go 1.27+. Ключ name задаёт вложенное поле person.name.
fmt.Println(bob.person.name) // Bob

// _ = member{person: person{}, name: "Bob"} // Ошибка: person уже задан целиком.
```

До Go 1.27 ключом литерала могло быть только непосредственно объявленное поле. С Go 1.27
продвинутый селектор допустим, если путь к нему не проходит через встроенный указатель.

### Встроенный указатель и нулевое значение

```go
type linkedMember struct{ *person }

// Внутри функции.
var linked linkedMember
fmt.Println(linked.person == nil) // true. Встроенное поле-указатель пока нулевое.
// fmt.Println(linked.name) // Паника: для чтения name пришлось бы разыменовать nil.
// _ = linkedMember{name: "Tom"} // Ошибка даже в Go 1.27: путь к name проходит через *person.

linked.person = &person{name: "Tom", age: 41}
fmt.Println(linked.name) // Tom. Теперь селектор идёт через действующий указатель.
```

### Продвижение методов

Методы встроенного поля тоже могут быть доступны через внешнюю структуру. Какие методы входят в её набор,
зависит от того, встроено значение `T` или указатель `*T`. Правила и примеры вызовов см. в
[уроке 019 о методах][methods-lesson].

## Рекурсивные структуры

### Почему в узле хранится указатель

```go
// type badNode struct{ next badNode } // Ошибка: badNode пришлось бы вложить в него самого.
// type badArray struct{ next [0]badArray } // Ошибка даже с массивом длины ноль.

type node struct {
    value int
    next  *node // Поле хранит указатель. Сам следующий узел может жить отдельно.
}
```

Go запрещает рекурсию через цепочку, состоящую только из структур и массивов. Указатель
или срез разрывает такую цепочку. Например, поле `children []node` допустимо.

## Теги, размер и выравнивание

### Тег как метаданные для библиотеки

```go
type publicPerson struct {
    Name string `json:"name"` // encoding/json использует ключ "name" вместо "Name".
}

// Внутри функции.
encoded, err := json.Marshal(publicPerson{Name: "Tom"})
fmt.Println(string(encoded), err) // {"name":"Tom"} <nil>
```

Тег — часть описания поля и участвует в идентичности типа. Смысл тега определяет читающая
его библиотека. В примере `json` импортируется из `encoding/json`.

### Размер и выравнивание

`unsafe.Sizeof` учитывает промежутки между полями для выравнивания. Размер зависит от архитектуры и реализации.
Размер `struct{}{}` равен нулю.

См. также [урок 016 об указателях](https://github.com/Hardome/learn-golang/tree/lesson-016-pointers)
и [урок 017 об определениях типов](https://github.com/Hardome/learn-golang/tree/lesson-017-type-aliases).

## Источники

- [Спецификация Go — структуры и встроенные поля](https://go.dev/ref/spec#Struct_types)
- [Спецификация Go — литералы структур](https://go.dev/ref/spec#Composite_literals)
- [Спецификация Go — область видимости](https://go.dev/ref/spec#Declarations_and_scope)
- [Спецификация Go — экспорт имён](https://go.dev/ref/spec#Exported_identifiers)
- [Спецификация Go — идентичность типов](https://go.dev/ref/spec#Type_identity)
- [Спецификация Go — присваиваемость](https://go.dev/ref/spec#Assignability)
- [Спецификация Go — преобразования](https://go.dev/ref/spec#Conversions)
- [Спецификация Go — нулевые значения](https://go.dev/ref/spec#The_zero_value)
- [Спецификация Go — встроенная функция new](https://go.dev/ref/spec#Allocation)
- [Спецификация Go — селекторы](https://go.dev/ref/spec#Selectors)
- [Спецификация Go — адресуемость](https://go.dev/ref/spec#Address_operators)
- [Спецификация Go — присваивание](https://go.dev/ref/spec#Assignment_statements)
- [Спецификация Go — передача аргументов](https://go.dev/ref/spec#Calls)
- [Спецификация Go — сравнение](https://go.dev/ref/spec#Comparison_operators)
- [Спецификация Go — типы карт и их ключей](https://go.dev/ref/spec#Map_types)
- [Спецификация Go — размеры и выравнивание](https://go.dev/ref/spec#Size_and_alignment_guarantees)
- [Effective Go — составные литералы](https://go.dev/doc/effective_go#composite_literals)
- [Go 1.26 — выражение как аргумент new](https://go.dev/doc/go1.26#language) и
  [Go 1.27 — селекторы в литералах структур](https://go.dev/doc/go1.27#language).
- [Стандартная библиотека — sync.Mutex](https://pkg.go.dev/sync#Mutex)
- [Стандартная библиотека — encoding/json.Marshal](https://pkg.go.dev/encoding/json#Marshal)
- [Стандартная библиотека — unsafe.Sizeof](https://pkg.go.dev/unsafe#Sizeof)

[methods-lesson]: https://github.com/Hardome/learn-golang/tree/lesson-019-methods
