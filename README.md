# Урок 010. Условные конструкции в Go

Условные конструкции выбирают блок кода по результату проверки.

В `if` ветки задают условия. `switch` выбирает ветку по значению, условию или динамическому типу.

## `if`, `else if` и `else`

Go проверяет условия сверху вниз и выполняет первую ветку, условие которой истинно.

Блок `else` выполняется, если все проверки дали `false`.

Ветки `else if` и `else` необязательны.

```go
a, b, c := 5, 6, 5
var relation string

if a > b {
    relation = "a > b"
} else if a > c {
    relation = "a > c"
} else {
    relation = "a <= b и a <= c"
}
fmt.Println(relation) // a <= b и a <= c
```

Обе строгие проверки дают `false`, поэтому `relation` получает значение из `else`. Здесь `a` и `c` равны.

### Короткая инструкция в `if`

Перед условием в `if` можно выполнить простую инструкцию. Здесь она объявляет `remainder`, доступную
в условии и в обеих ветках. Остаток равен `2`, поэтому срабатывает `else`.

```go
if remainder := 17 % 3; remainder == 0 {
    fmt.Println("делится на 3")
} else {
    fmt.Printf("остаток %d\n", remainder) // остаток 2
}
// fmt.Println(remainder) // Ошибка компиляции. Область видимости закончилась вместе с if.
```

## `switch`

### `switch` по значению

В `switch value` Go вычисляет значение один раз и сравнивает его с вариантами в `case`.

Несколько вариантов с общим телом перечисляют через запятую. Например, `case 3, 4:` подходит для `3` и для `4`.

Перед выражением допускается короткая инструкция, как в `switch value := 5; value {`.

### `switch` без выражения

В `switch { ... }` выражения в `case` задают логические условия. Подходит первое истинное условие.

```go
temperature := 18
var weather string

switch {
case temperature < 0:
    weather = "морозно"
case temperature >= 25:
    weather = "жарко"
default:
    weather = "умеренно"
}
fmt.Println(weather) // умеренно
```

### Порядок проверок и `default`

Go проверяет варианты по порядку до первого совпадения. Выражения после него пропускаются вместе с их побочными
эффектами.

Если совпадений нет и указана ветка `default`, Go выполняет её.

В `switch` без выражения температура `18` не подходит под условия `temperature < 0` и `temperature >= 25`, поэтому
выполняется `default`.

Ветку `default` можно разместить в любом месте конструкции.

### Переход `fallthrough`

По умолчанию после выполнения тела выбранной ветки `switch` завершается.

Инструкция `fallthrough` передаёт управление телу следующей ветки без проверки её условия.

```go
for _, value := range [3]int{4, 5, 7} {
    switch value {
    case 3, 4:
        fmt.Println(value, "case 3 или 4") // 4 case 3 или 4
    case 5:
        fmt.Println(value, "case 5")      // 5 case 5
        fallthrough
    case 6:
        fmt.Println(value, "case 6")      // 5 case 6
    default:
        fmt.Println(value, "default")     // 7 default
    }
}
```

Для `4` подходит первый `case`. Для `5` срабатывает `case 5`, затем переход сразу в тело `case 6`.
У `7` совпадений нет, поэтому выполняется `default`.

`fallthrough` допустим как последняя непустая инструкция ветки обычного `switch`, если за ней есть следующая ветка.

## `switch` по типу

Конструкция `switch typed := value.(type)` проверяет динамический тип интерфейса, то есть тип хранящегося в нём
значения.

```go
func describeType(value any) string {
    switch typed := value.(type) {
    case int:
        return fmt.Sprintf("int, следующее число %d", typed+1)
    case string, bool:
        _, isString := typed.(string)
        return fmt.Sprintf("строка: %t", isString)
    case nil:
        return "nil"
    default:
        return "другой тип"
    }
}

// Внутри функции.
fmt.Println(describeType(42))   // int, следующее число 43
fmt.Println(describeType("Go")) // строка: true
fmt.Println(describeType(nil))  // nil
```

- **Один тип в `case`.** В `case int` переменная `typed` имеет тип `int`, поэтому можно вычислить `typed+1`.
- **Несколько типов в `case` или ветка `default`.** У `typed` сохраняется тип интерфейсного выражения,
  в примере это `any`. Выражение `typed.(string)` проверяет, содержит ли интерфейс строку.
- **`case nil`.** Ветка срабатывает для интерфейса без динамического типа.

В `switch` по типу `fallthrough` запрещён.

## Запуск

Выполните `go run hello.go` из каталога урока, чтобы запустить [примеры](hello.go).

## Источники

- [Конструкция if](https://go.dev/ref/spec#If_statements)
- [Две формы switch](https://go.dev/ref/spec#Switch_statements)
- [Переход fallthrough](https://go.dev/ref/spec#Fallthrough_statements)
- [Идиоматичные switch](https://go.dev/doc/effective_go#switch)
