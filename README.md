# Урок 027. Создание разделяемого модуля

Когда один и тот же пакет нужен нескольким приложениям, его можно вынести в отдельный модуль с собственным
`go.mod`. В этом уроке пакет `messages` формирует приветствие по имени. Приложение импортирует его из модуля
`example.com/messages` и берёт исходники из локального каталога `greetings`.

## Создание модуля библиотеки

Чтобы повторить пример с нуля, создайте в проекте каталог `greetings` и выполните в нём
`go mod init example.com/messages`. Команда создаст [greetings/go.mod](greetings/go.mod) и запишет в него путь
модуля. Этот файл определяет корень модуля библиотеки.

Добавьте в тот же каталог [messages.go](greetings/messages.go).

```go
package messages

// Greeting возвращает приветствие для переданного имени.
func Greeting(name string) string {
	return "Привет, " + name + "!"
}
```

Функция `Greeting` объявлена на уровне пакета, и её имя начинается с заглавной буквы. Поэтому она доступна
из пакета `main`.

## Импорт пакета в приложении

В корне проекта выполните `go mod init github.com/Hardome/learn-golang`. Этот модуль будет содержать приложение.
Рядом с его `go.mod` добавьте [hello.go](hello.go).

```go
package main

import (
	"fmt"

	"example.com/messages" // Пакет находится в корне модуля библиотеки.
)

func main() {
	fmt.Println(messages.Greeting("Анна")) // Привет, Анна!
}
```

Путь импорта `example.com/messages` совпадает с путём модуля, потому что пакет лежит в его корне. В вызове
`messages.Greeting` имя `messages` взято из строки `package messages`.

## Подключение локального модуля

Чтобы Go взял исходники `example.com/messages` из каталога `greetings`, выполните из корня приложения:

```text
go mod edit -replace="example.com/messages=./greetings"
go mod tidy
```

Первая команда записывает `replace` в [go.mod](go.mod) приложения. Путь `./greetings` отсчитывается от каталога,
где лежит этот `go.mod`. В каталоге замены должен быть свой `go.mod` с путём `example.com/messages`.

`replace` задаёт источник кода. Чтобы модуль стал зависимостью приложения, нужно ещё требование `require`.
`go mod tidy` добавляет его, когда находит импорт `example.com/messages` в `hello.go`. В `go.mod` приложения
появляются две записи.

```text
require example.com/messages v0.0.0-00010101000000-000000000000

replace example.com/messages => ./greetings
```

Для новой зависимости с такой заменой `go mod tidy` подставляет псевдоверсию. Если другое приложение использует
ту же локальную библиотеку, оно добавляет `replace` в свой `go.mod`. Например, из соседнего с `greetings`
каталога путь к библиотеке будет `../greetings`.

## Структура проекта и запуск

В этой ветке исходники и оба файла `go.mod` уже на месте.

```text
.                     # модуль приложения
├── go.mod
├── hello.go           # пакет main
└── greetings/        # модуль example.com/messages
    ├── go.mod
    └── messages.go    # пакет messages
```

Из корня проекта выполните `go run .`. В этом примере `go run hello.go` даёт тот же результат.

```text
Привет, Анна!
```

## Источники

- [Go Modules Reference · Modules, packages, and versions](https://go.dev/ref/mod#modules-overview)
- [Go Modules Reference · go mod init](https://go.dev/ref/mod#go-mod-init)
- [Go Specification · Exported identifiers](https://go.dev/ref/spec#Exported_identifiers)
- [Go Specification · Import declarations](https://go.dev/ref/spec#Import_declarations)
- [Go Modules Reference · go mod edit](https://go.dev/ref/mod#go-mod-edit)
- [Go Modules Reference · replace](https://go.dev/ref/mod#go-mod-file-replace)
- [Go Modules Reference · require](https://go.dev/ref/mod#go-mod-file-require)
- [Go Modules Reference · go mod tidy](https://go.dev/ref/mod#go-mod-tidy)
- [Исходный код Go · относительный путь replace](https://go.dev/src/cmd/go/internal/modload/modfile.go#L350)
- [Исходный код Go · псевдоверсия локальной замены](https://go.dev/src/cmd/go/internal/modload/import.go#L533)
- [cmd/go · go run](https://pkg.go.dev/cmd/go#hdr-Compile_and_run_Go_program)
- [fmt.Println](https://pkg.go.dev/fmt#Println)
