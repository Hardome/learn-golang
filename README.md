# Урок 014. Замыкания в Go

Замыкание — функция, которая сохраняет доступ к переменным из окружающей лексической области видимости.
В Go литерал функции может читать и менять переменные окружающей функции. Если вернуть такой литерал,
захваченные переменные остаются доступными даже после возврата из функции, которая их объявила.

Для практики запусти [hello.go](hello.go) командой `go run hello.go` с Go 1.22 или новее.
Синтаксис литералов разобран в [уроке 013][lesson-013].

## Захват и время жизни переменных

Замыкание обращается к той же переменной, что и окружающий код. Область видимости определяется местом записи
литерала, а значение захваченной переменной читается при исполнении тела. Одноимённая переменная, объявленная
в более узком блоке после создания литерала, не меняет его привязку.

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

Разные литералы, созданные в одном вызове фабрики, могут захватывать одну переменную. Например,
`newAccumulator` возвращает две функции с общей суммой.

```go
func newAccumulator() (add func(int), total func() int) {
    sum := 0
    return func(value int) { sum += value }, func() int { return sum }
}
```

После `add, total := newAccumulator()` вызов `add(5)` меняет переменную, которую читает `total()`.
Если вызвать фабрику ещё раз, новая пара будет работать с собственной суммой.

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
Тело отложенной функции выполняется при выходе из окружающей функции, в обратном порядке регистрации,
в том числе при раскрутке стека из-за паники.
Чтобы остановить панику, вызови
`recover()` непосредственно в теле отложенной функции в той же горутине. Тело функции, запущенной через
`go`, выполняется в новой горутине.

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

При конкурентном доступе к захваченной переменной нужна синхронизация, если хотя бы одна операция — запись.
Например, параллельные вызовы одного счётчика без синхронизации приводят к гонке данных.
Состояние можно защитить мьютексом или поручить одной горутине, которая обрабатывает запросы через канал.

## Рекурсия и захват функционального значения

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
к новому значению. Если присвоить переменной `nil`, следующий рекурсивный вызов приведёт к панике.
Для рекурсии без локального состояния обычно достаточно именованной функции.

## Источники

- [Go Language Specification — Function literals](https://go.dev/ref/spec#Function_literals)
- [Go Language Specification — Declarations and scope](https://go.dev/ref/spec#Declarations_and_scope)
- [Go Language Specification — Calls](https://go.dev/ref/spec#Calls)
- [Go Language Specification — Order of evaluation](https://go.dev/ref/spec#Order_of_evaluation)
- [Go Language Specification — For statements](https://go.dev/ref/spec#For_statements)
- [Go 1.22 Release Notes — Changes to the language](https://go.dev/doc/go1.22#language)
- [The Go Blog — Fixing For Loops in Go 1.22](https://go.dev/blog/loopvar-preview#the-fix)
- [Go Language Specification — Defer statements](https://go.dev/ref/spec#Defer_statements)
- [Go Language Specification — Handling panics](https://go.dev/ref/spec#Handling_panics)
- [Go Language Specification — Go statements](https://go.dev/ref/spec#Go_statements)
- [Effective Go — Defer](https://go.dev/doc/effective_go#defer)
- [Go Memory Model — Introduction](https://go.dev/ref/mem#introduction)
- [Go FAQ — размещение переменных](https://go.dev/doc/faq#stack_or_heap)

[lesson-013]: https://github.com/Hardome/learn-golang/tree/lesson-013-anonymous-functions
