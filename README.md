# Урок 013. Анонимные функции в Go

Литерал функции `func(параметры) результаты { тело }` создаёт функциональное значение.
Его можно вызвать сразу, сохранить в переменной, передать другой функции или вернуть из неё.
Тело выполняется при вызове.

Примеры из [hello.go](hello.go) запускаются командой `go run hello.go` на Go 1.21 или новее.
Параметры и результаты описаны в [уроке 012][lesson-012].

## Литерал функции и вызов

```go
double := func(value int) int {
    return value * 2
}
fmt.Println(double(4)) // 8

fmt.Println(func(value int) int {
    return value * 2
}(5)) // 10
```

`double` — переменная типа `func(int) int`. В первом примере она получает функциональное значение,
которое затем вызывается через `double(4)`. Во втором скобки после тела сразу вызывают функцию с аргументом `5`.

Литерал можно написать внутри другой функции и обратиться из него к переменным окружающей области видимости.
Это создаёт замыкание. Правила захвата переменных разобраны в [уроке 014][lesson-014].

## Функция как аргумент

Короткое правило, которое нужно только для одного вызова, удобно передать литералом. Например,
`slices.SortFunc` сортирует срез на месте по переданному компаратору.

```go
type job struct {
    name     string
    priority int
}

jobs := []job{{name: "backup", priority: 2}, {name: "deploy", priority: 1}}
slices.SortFunc(jobs, func(left, right job) int {
    return cmp.Compare(left.priority, right.priority)
})
fmt.Println(jobs) // [{deploy 1} {backup 2}]
```

Компаратор возвращает отрицательное число, ноль или положительное число в зависимости от порядка элементов.
`cmp.Compare` избегает переполнения, возможного при вычитании двух приоритетов типа `int`.
`slices.SortFunc` требует строгого слабого порядка. Например, отношение «меньше» должно быть транзитивным.
Для сохранения относительного порядка элементов с одинаковым приоритетом есть `slices.SortStableFunc`.

Если у правила есть имя в предметной области, несколько мест использования или сложные ветвления, отдельную
именованную функцию проще читать и проверять.

## Локальная инициализация

Немедленный вызов полезен, когда для вычисления одного значения нужны локальные переменные и несколько шагов.
Например, на уровне пакета его можно использовать как выражение инициализации без отдельного объявления функции.

```go
var defaultPorts = func() []int {
    ports := []int{8080, 443}
    slices.Sort(ports)
    return ports
}()
```

При инициализации пакета `defaultPorts` получает возвращённый срез, а временное имя `ports` видно
только внутри литерала.

## Хранение функциональных значений

Функциональные значения можно хранить в полях структуры, срезах и значениях карты. Сигнатура литерала должна
подходить для типа назначения. Например, присваивание литерала переменной определённого типа
`transform` допустимо, потому что у них одинаковые базовые типы, а тип литерала неименованный.

```go
type transform func(int) int

operations := map[string]transform{
    "square": func(value int) int { return value * value },
}
if square, ok := operations["square"]; ok && square != nil {
    fmt.Println(square(5)) // 25
}
```

Имена параметров и результатов не меняют тип функции. Количество, порядок, типы параметров и результатов,
а также вариативность входят в сигнатуру. Функциональный тип `func(...int)` отличается от `func([]int)`.
Для двух разных определённых функциональных типов с одинаковым базовым типом нужно явное приведение
при присваивании. Встроенные функции вроде `len` нельзя сохранять как функциональные значения, но можно
вызвать их внутри литерала.

## Нулевое значение и сравнение функций

Нулевое значение функционального типа — `nil`. Оно не указывает на функцию, поэтому попытка вызова
приводит к панике. Для функционального значения доступны сравнения `== nil` и `!= nil`.

```go
var callback func()
var boxed any = callback
fmt.Println(callback == nil, boxed == nil) // true false
// callback() // Паника при вызове nil.
// fmt.Println(callback == callback) // Ошибка компиляции. Два значения функции несравнимы.
// fmt.Println(boxed == boxed) // Паника. Оба интерфейса содержат несравнимый тип func().
```

Интерфейс `boxed` содержит динамический тип `func()` и значение `nil` этого типа. Сравнение `boxed == boxed`
компилируется и вызывает панику, поскольку обе стороны содержат один и тот же несравнимый динамический тип.

[lesson-012]: https://github.com/Hardome/learn-golang/tree/lesson-012-functions-and-parameters
[lesson-014]: https://github.com/Hardome/learn-golang/tree/lesson-014-closures

## Источники

- [Go Language Specification — Function literals](https://go.dev/ref/spec#Function_literals)
- [Go Language Specification — Assignability](https://go.dev/ref/spec#Assignability)
- [Go Language Specification — Built-in functions](https://go.dev/ref/spec#Built-in_functions)
- [Go Language Specification — Function types](https://go.dev/ref/spec#Function_types)
- [Go Language Specification — Declarations and scope](https://go.dev/ref/spec#Declarations_and_scope)
- [Go Language Specification — Calls](https://go.dev/ref/spec#Calls)
- [Go Language Specification — Comparison operators](https://go.dev/ref/spec#Comparison_operators)
- [slices.SortFunc](https://pkg.go.dev/slices#SortFunc)
- [slices.SortStableFunc](https://pkg.go.dev/slices#SortStableFunc)
- [cmp.Compare](https://pkg.go.dev/cmp#Compare)
