# Урок 023. Модули в Go

Пакет объединяет Go-файлы, которые компилируются вместе. **Модуль** объединяет пакеты, которые
распространяются и версионируются вместе. Его корень определяется файлом `go.mod`. По этому файлу
команда `go` узнаёт путь модуля, требования к версии Go и версии зависимостей.

## Путь модуля и версия Go

В корне репозитория уже есть [go.mod](go.mod).

```text
module github.com/Hardome/learn-golang

go 1.27.4
```

Директива `module` задаёт путь модуля. Для пакета в корне модуля это также его путь импорта.
Для пакетов в подкаталогах к пути модуля добавляется путь каталога, как показано в [уроке 024][layout].
Пакет `main` в этом уроке служит точкой входа, а его файлы лежат в корне модуля.

Директива `go` задаёт минимальную версию Go для модуля и версию языка для его пакетов.
В примере минимальная версия — `1.27.4`. Директива `go` отдельно от версий библиотек, которые
указываются в `require`.

## Зависимости и выбор версий

У модуля урока нет внешних зависимостей. Если бы приложение импортировало пакет из модуля
`rsc.io/quote`, в `go.mod` можно было бы увидеть такое требование.

```text
require rsc.io/quote v1.5.2
```

`require` задаёт **минимальную версию** по указанному пути модуля. Если приложение требует
`example.org/lib v1.2.0`, а одна из его зависимостей требует `example.org/lib v1.4.0`, Go выберет
`v1.4.0`. Это наибольшая из требуемых версий в графе зависимостей для одного пути. Такой алгоритм
называется **Minimal Version Selection** (MVS). Команда `go list -m all` показывает выбранные версии.

Добавление зависимости через `go get` и согласование файлов модулей через `go mod tidy` разобраны
в [уроке 025][dependencies]. Подключение локальной копии через `replace` разобрано в [уроке 027][shared].

### Контрольные суммы в `go.sum`

В `go.sum` записываются контрольные суммы загруженных зависимостей. Запись с суффиксом `/go.mod`
относится к файлу модуля, а запись без суффикса — к содержимому модуля. Если при проверке сумма
не совпадёт, Go сообщит об ошибке. Выбранные версии можно посмотреть через `go list -m all`, поскольку
`go.sum` может содержать суммы нескольких версий одного модуля.

В этой ветке используются только пакеты стандартной библиотеки, поэтому `go.sum` отсутствует.
Если в проекте с внешними зависимостями он появился, сохраняйте его в Git вместе с `go.mod`.

## Запуск готового модуля

В корне лежат [hello.go](hello.go) с функцией `main` и [greeting.go](greeting.go) с функцией `init`.
Оба файла относятся к одному пакету. Разница в выводе зависит от того, какие файлы входят в сборку.

| Команда из корня | Вывод по порядку |
| --- | --- |
| `go run hello.go` | `Привет, Go!` |
| `go run hello.go greeting.go` | `Привет, читатель!`, затем `Привет, Go!` |
| `go run .` | `Привет, читатель!`, затем `Привет, Go!` |

При явном перечислении Go собирает перечисленные файлы. Точка обозначает пакет из текущего
каталога, поэтому Go сам находит оба файла. Когда `greeting.go` входит в пакет, его `init` выполняется
перед `main`. Правила работы пакета из нескольких файлов разобраны в [уроке 022][packages].

## Создание отдельного модуля

В новом каталоге выполните `go mod init example.org/hello`. Команда создаст `go.mod` в этом каталоге.
Добавьте рядом файл `main.go`.

```go
package main

import "fmt"

func main() {
    fmt.Println("Привет, Go!") // Привет, Go!
}
```

Команда `go run .` запустит пакет и выведет `Привет, Go!`. Команда `go build .` соберёт `hello.exe`
на Windows и `hello` на Linux или macOS. Имя исполняемого файла берётся из конца пути `example.org/hello`.

[packages]: https://github.com/Hardome/learn-golang/tree/lesson-022-packages-and-imports
[layout]: https://github.com/Hardome/learn-golang/tree/lesson-024-project-packages
[dependencies]: https://github.com/Hardome/learn-golang/tree/lesson-025-external-dependencies
[shared]: https://github.com/Hardome/learn-golang/tree/lesson-027-shared-modules

## Источники

- [Модули, пакеты и версии](https://go.dev/ref/mod#modules-overview)
- [Директива go](https://go.dev/ref/mod#go-mod-file-go)
- [Директива require](https://go.dev/ref/mod#go-mod-file-require)
- [Выбор версий MVS](https://go.dev/ref/mod#minimal-version-selection)
- [Файл go.sum](https://go.dev/ref/mod#go-sum-files)
- [go mod init](https://go.dev/ref/mod#go-mod-init)
- [Явный список файлов](https://pkg.go.dev/cmd/go#hdr-List_of__go_files)
- [go run](https://pkg.go.dev/cmd/go#hdr-Compile_and_run_Go_program)
- [go build](https://pkg.go.dev/cmd/go#hdr-Compile_packages_and_dependencies)
- [Выполнение init](https://go.dev/ref/spec#Package_initialization)
