# Урок 001. Hello, world!

Минимальная программа на Go начинается с объявления пакета. Запуск исполняемой программы передаёт управление функции
`main`.

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, world!") // Hello, world!
}
```

- `package main` объявляет пакет исполняемой программы.
- `import "fmt"` подключает стандартный пакет форматированного ввода-вывода.
- `func main()` объявляет точку входа без параметров и результатов.
- `fmt.Println("Hello, world!")` печатает приветствие и добавляет перевод строки.

## Запуск и сборка

Установите Go по официальной инструкции. Из каталога урока можно запустить или собрать [пример](hello.go).

| Команда | Что делает |
| --- | --- |
| `go version` | Показывает версию инструментария, ОС и архитектуру. |
| `go run hello.go` | Собирает и сразу запускает программу без сохранения бинарника в каталоге урока. |
| `go build hello.go` | Собирает `hello.exe` на Windows и `hello` на Linux и macOS. |

Собранный файл запускается через `.\hello.exe` в PowerShell или `./hello` в Linux и macOS. При передаче имён файлов
команда `go` собирает указанные Go-файлы из одного каталога как один пакет. Это позволяет запускать отдельный урок без
`go.mod`.

## Когда запускается main

Go сначала инициализирует пакеты, затем вызывает `main`. Когда `main` возвращается, программа завершается.
Подробнее о запуске и импортах см. в [уроке 022][packages].

[packages]: https://github.com/Hardome/learn-golang/tree/lesson-022-packages-and-imports

## Источники

- [Точка входа и завершение программы](https://go.dev/ref/spec#Program_execution)
- [Правила импорта](https://go.dev/ref/spec#Import_declarations)
- [go run](https://pkg.go.dev/cmd/go#hdr-Compile_and_run_Go_program) и
  [go build](https://pkg.go.dev/cmd/go#hdr-Compile_packages_and_dependencies)
- [fmt.Println](https://pkg.go.dev/fmt#Println)
- [Установка Go](https://go.dev/doc/install)
