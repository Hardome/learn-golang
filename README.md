# Урок 013. Анонимные функции в Go

Анонимная функция записывается литералом `func(параметры) результаты { тело }`. Литерал создаёт значение
функционального типа. Его можно сразу вызвать, сохранить в переменной, передать другой функции или вернуть из неё.
Замыкание может обращаться к переменным окружающей функции и сохраняет доступ к ним после её завершения.

Для практики запусти [hello.go](hello.go) командой `go run hello.go` с Go 1.22 или новее.
Синтаксис параметров и результатов разобран в [уроке 012][lesson-012].

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

## Замыкания и время жизни переменных

Замыкание обращается к той же переменной, что и окружающий код. Область видимости определяется местом записи
литерала, а значение захваченной переменной читается при исполнении тела.

```go
label := "до"
readLabel := func() string { return label }
label = "после"
fmt.Println(readLabel()) // после
```

Чтобы сохранить текущее значение, передай его параметром в функцию-фабрику или сделай отдельную переменную
до создания литерала и захвати её. Присваивание копирует значение по обычным правилам Go. Например, копия
среза может ссылаться на тот же массив, поэтому изменение элемента видно через обе копии.

```go
func newCounter(start int) func() int {
    return func() int {
        start++
        return start
    }
}
```

```go
first := newCounter(0)
sameCounter := first
second := newCounter(10)
fmt.Println(first(), sameCounter(), second()) // 1 2 11
```

У каждого вызова `newCounter` своя переменная `start`. Копия функционального значения `sameCounter`
работает с тем же `start`, что и `first`. Вызовы в аргументах `fmt.Println` выполняются слева направо.
Захваченная переменная остаётся доступной, пока до неё можно добраться через замыкание.

Долгоживущее замыкание может удерживать данные в памяти. Например, функция, которая читает один элемент
захваченного среза, может удерживать весь его базовый массив. Если нужно только значение элемента,
его можно извлечь до создания замыкания.

Язык гарантирует доступность захваченной переменной. Размещение на стеке или в куче зависит от реализации,
анализа убегания и оптимизаций компилятора. Диагностику компилятора Go можно посмотреть командой
`go build -gcflags=-m hello.go`.

## Захват переменных цикла

С версии языка Go 1.22 каждая итерация получает отдельную переменную, если она объявлена через `:=`
в заголовке `for`. Это касается и `for _, value := range values`, и трёхчастного `for index := 0; ...`.
При старой семантике замыкания из разных итераций разделяли одну переменную. В модульном проекте учитывай
директиву `go` в `go.mod`, даже когда установлен более новый компилятор.

```go
var readers []func() int
for _, value := range []int{1, 2, 3} {
    readers = append(readers, func() int { return value })
}
fmt.Println(readers[0](), readers[1](), readers[2]()) // 1 2 3 с Go 1.22+, 3 3 3 раньше
```

Присваивание через `=` использует ранее объявленную переменную и в новых версиях. В следующем примере
все функции вызываются после цикла и читают значение из последней итерации.

```go
var value int
var readers []func() int
for _, value = range []int{1, 2, 3} {
    readers = append(readers, func() int { return value })
}
fmt.Println(readers[0](), readers[1](), readers[2]()) // 3 3 3
```

Чтобы получить отдельную переменную при `=` или при старой семантике, добавь `iterationValue := value` в тело
цикла и захвати `iterationValue`. При немедленном вызове внутри той же итерации замыкание читает текущее значение.
Отдельная переменная итерации также не копирует данные за указателем или базовый массив среза.

## Аргументы и захват в defer и go

В инструкциях `defer` и `go` функциональное значение и аргументы вычисляются сразу в вызывающей горутине.
Тело отложенной функции выполняется при выходе из окружающей функции, в обратном порядке регистрации.
Отложенные вызовы также выполняются при раскрутке стека из-за паники. Чтобы остановить панику, вызови
`recover()` непосредственно в теле отложенной функции в той же горутине. Тело функции, запущенной через
`go`, выполняется в новой горутине. Результаты вызовов через `defer` и `go` отбрасываются.

```go
func showDeferredValues() {
    state := "начало"
    defer func(saved string) { fmt.Println(saved) }(state)
    defer func() { fmt.Println(state) }()
    state = "конец"
}
```

`showDeferredValues()` печатает `конец` и `начало` на отдельных строках. Первый зарегистрированный
вызов сохранил значение аргумента. Второй читает захваченную переменную при выполнении.

Отложенное замыкание может изменить именованный результат после вычисления `return` и до возврата
вызывающему коду.

```go
func adjusted() (result int) {
    defer func() { result++ }()
    return 4
}
// adjusted() возвращает 5
```

Для короткой задачи в горутине можно передать входное значение параметром. Канал позволяет дождаться результата.

```go
result := make(chan int)
value := 7
go func(input int) { result <- input * 2 }(value)
value = 9
fmt.Println(<-result) // 14
```

Захваченное изменяемое состояние требует синхронизации при одновременном доступе. Например, параллельные
вызовы одного счётчика без синхронизации приводят к гонке данных. Состояние можно защитить мьютексом
или поручить одной горутине, которая обрабатывает запросы через канал.

## Рекурсия, nil и сравнение функций

Рекурсивный литерал обращается к заранее объявленной переменной. В `factorial := func(...) { ... }` область
видимости новой переменной начинается после всего короткого объявления. Внутри литерала это новое имя ещё
недоступно.

```go
var factorial func(int) int
factorial = func(value int) int {
    if value < 2 {
        return 1
    }
    return value * factorial(value-1)
}
fmt.Println(factorial(5)) // 120
```

Пример рассчитан на малые неотрицательные числа. При рекурсивном вызове тело заново читает захваченную
переменную `factorial`. Если её переназначить, даже сохранённая ранее копия функции будет обращаться
к новому значению. Для рекурсии без локального состояния обычно достаточно именованной функции.

Нулевое значение функционального типа — `nil`. Вызов такого значения приводит к панике.
При `defer callback()` с `nil` паника возникнет при отложенном вызове.
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
- [Go Language Specification — Function types](https://go.dev/ref/spec#Function_types)
- [Go Language Specification — Declarations and scope](https://go.dev/ref/spec#Declarations_and_scope)
- [Go Language Specification — Calls](https://go.dev/ref/spec#Calls)
- [Go Language Specification — Order of evaluation](https://go.dev/ref/spec#Order_of_evaluation)
- [Go Language Specification — Comparison operators](https://go.dev/ref/spec#Comparison_operators)
- [Go Language Specification — For statements](https://go.dev/ref/spec#For_statements)
- [Go 1.22 Release Notes — Changes to the language](https://go.dev/doc/go1.22#language)
- [The Go Blog — Fixing For Loops in Go 1.22](https://go.dev/blog/loopvar-preview#the-fix)
- [Go Language Specification — Defer statements](https://go.dev/ref/spec#Defer_statements)
- [Go Language Specification — Handling panics](https://go.dev/ref/spec#Handling_panics)
- [Go Language Specification — Go statements](https://go.dev/ref/spec#Go_statements)
- [Effective Go — Defer](https://go.dev/doc/effective_go#defer)
- [Go Memory Model — Introduction](https://go.dev/ref/mem#introduction)
- [Go FAQ — размещение переменных](https://go.dev/doc/faq#stack_or_heap)
- [slices.SortFunc](https://pkg.go.dev/slices#SortFunc) · [slices.SortStableFunc](https://pkg.go.dev/slices#SortStableFunc)
- [cmp.Compare](https://pkg.go.dev/cmp#Compare)

[lesson-012]: https://github.com/Hardome/learn-golang/tree/lesson-012-functions-and-parameters
