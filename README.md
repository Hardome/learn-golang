# Урок 027. Создание разделяемого модуля

Разделяемый модуль объединяет пакеты, которые могут импортировать разные приложения. У библиотеки
свой `go.mod`. Приложение объявляет её зависимостью, а для локальной разработки может указать каталог
с исходниками через `replace`.

В примере два модуля. Корневой модуль содержит приложение, а вложенный модуль `example.com/messages`
предоставляет пакет `messages`.

```text
.                     # модуль приложения
├── go.mod
├── hello.go           # пакет main
└── greetings/        # модуль example.com/messages
    ├── go.mod
    └── messages.go    # пакет messages
```

Файлы уже лежат в этих каталогах. Для воспроизведения шагов с нуля используйте отдельный пустой каталог.

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

`replace` задаёт источник кода, а `require` добавляет зависимость в граф модулей.
`go mod tidy` добавляет требование, когда находит импорт `example.com/messages` в `hello.go`.
По завершении в [go.mod](go.mod) приложения появляются две записи.

```text
require example.com/messages v0.0.0-00010101000000-000000000000

replace example.com/messages => ./greetings
```

В этом примере `go mod tidy` записывает служебную псевдоверсию. Чтобы собрать приложение, Go берёт данные
по локальной замене. При обычной сборке замена в `go.mod` действует для главного модуля приложения.
Если другое приложение использует ту же библиотеку локально, оно задаёт свою замену. Например,
из соседнего каталога путь к `greetings` будет `../greetings`.

## Запуск готового проекта

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
