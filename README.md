# Урок 001. Hello, world!

Go собирает программу из пакетов. Пакет объединяет объявления из одного или нескольких файлов.
Для исполняемой программы Go начинает сборку с пакета `main`. После инициализации пакетов программа вызывает
функцию `main` из этого пакета.

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, world!") // Hello, world!
}
```

- `package main` относит файл к пакету исполняемой программы.
- `import "fmt"` делает имя пакета `fmt` доступным в этом файле. Через него программа обращается к `fmt.Println`.
- `func main()` объявляет точку входа. Такая функция принимает ноль параметров и возвращает ноль результатов.
- `fmt.Println("Hello, world!")` пишет приветствие в стандартный вывод. `Println` добавляет в конце перевод строки.

Когда управление возвращается из `main`, программа завершается.
Порядок инициализации пакетов описан в [уроке 022][packages].

## Запуск и сборка

Установите Go по [официальной инструкции][install]. Выполняйте команды из каталога с [hello.go](hello.go).

| Команда | Что делает |
| --- | --- |
| `go version` | Показывает версию инструментария, ОС и архитектуру. |
| `go run hello.go` | Собирает и сразу запускает программу без сохранения бинарника в каталоге урока. |
| `go build hello.go` | Собирает `hello.exe` на Windows и `hello` на Linux и macOS. |

После `go build hello.go` запустите готовый файл через `.\hello.exe` в PowerShell или `./hello` в Linux и macOS.
В обоих случаях вывод совпадает с результатом `go run hello.go`.

Когда команде `go` передают имена файлов из одного каталога, она собирает из перечисленных файлов один пакет.
В этом примере список состоит только из `hello.go`, а `fmt` входит в стандартную библиотеку. Поэтому для запуска
и сборки достаточно одного исходного файла, даже если в каталоге нет `go.mod`.

[packages]: https://github.com/Hardome/learn-golang/tree/lesson-022-packages-and-imports
[install]: https://go.dev/doc/install

## Источники

- [Точка входа и завершение программы](https://go.dev/ref/spec#Program_execution)
- [Правила импорта](https://go.dev/ref/spec#Import_declarations)
- [Явный список Go-файлов](https://pkg.go.dev/cmd/go#hdr-List_of__go_files)
- [go run](https://pkg.go.dev/cmd/go#hdr-Compile_and_run_Go_program) и
  [go build](https://pkg.go.dev/cmd/go#hdr-Compile_packages_and_dependencies)
- [fmt.Println](https://pkg.go.dev/fmt#Println)
- [Установка Go](https://go.dev/doc/install)
