# Урок 013. Анонимные функции в Go

Анонимная функция записывается литералом `func(параметры) результаты { тело }`. Литерал создаёт значение
функционального типа. Его можно сразу вызвать, сохранить в переменной, передать другой функции или вернуть из неё.

Для практики запусти [hello.go](hello.go) командой `go run hello.go` с Go 1.21 или новее.
Синтаксис параметров и результатов разобран в [уроке 012][lesson-012]. Захват переменных разобран
в [уроке 014 о замыканиях][lesson-014].

## Литерал функции и вызов

```go
double := func(value int) int {
    return value * 2
}
fmt.Println(double(4)) // 8

fmt.Println(func(left, right int) int {
    return left + right
}(3, 4)) // 7
```

`double` — переменная типа `func(int) int`. Скобки с аргументами после тела во втором примере сразу вызывают
функцию. При вычислении самого литерала тело ещё не выполняется.

Литерал можно написать внутри другой функции. Объявление именованной функции допустимо только на уровне пакета.
Литерал не может объявлять собственные параметры типа, но может использовать параметр окружающей обобщённой
функции. Например, `func identityFor[T any]() func(T) T` может вернуть `func(value T) T { return value }`.

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

Переменная получает результат вызова, а область видимости временной переменной ограничена телом функции.
Для простого выражения обычно достаточно обычного присваивания.

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

Нулевое значение функционального типа — `nil`. Вызов такого значения приводит к панике.
Функциональное значение можно сравнить только с `nil`. Выражение `callback == callback` не компилируется,
и функциональный тип нельзя использовать в качестве типа ключа карты.

```go
var callback func()
var boxed any = callback
fmt.Println(callback == nil, boxed == nil) // true false
```

Интерфейс `boxed` содержит динамический тип `func()` и значение `nil` этого типа. Сравнение `boxed == boxed`
компилируется и вызывает панику, поскольку обе стороны содержат один и тот же несравнимый динамический тип.

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

[lesson-012]: https://github.com/Hardome/learn-golang/tree/lesson-012-functions-and-parameters
[lesson-014]: https://github.com/Hardome/learn-golang/tree/lesson-014-closures
