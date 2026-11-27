# Урок 023. Модули

Модуль — это набор пакетов Go с общим путём и версией. В его корне лежит файл `go.mod`. Один проект может
использовать одну версию библиотеки, другой — другую. Go выбирает версии для каждого проекта и может хранить
их в общем кеше. Библиотечный модуль можно опубликовать в репозитории, чтобы его могли подключить другие проекты.

## go.mod и go.sum

### go.mod

`go.mod` хранит путь модуля, минимальную версию Go и требования к зависимостям. Например, такой файл описывает модуль
`example.org/hello`, который использует библиотеку `rsc.io/quote`.

```text
module example.org/hello

go 1.24.4

require rsc.io/quote v1.5.2
```

При `go mod init` строка `go` зависит от версии запущенного инструмента. Начиная с Go 1.21 команда `go`
строго проверяет этот минимум. `require` задаёт минимальную версию зависимости.
Пометка `// indirect` обычно означает, что этот модуль нужен через другую зависимость.

Если один модуль требует библиотеку версии `v1.2.0`, а другой — `v1.4.0` по тому же пути, Go выберет `v1.4.0`.
Он берёт наибольшую из требуемых версий. Этот алгоритм называется **Minimal Version Selection** (MVS).
Список выбранных версий показывает `go list -m all`.

### go.sum

`go.sum` хранит контрольные суммы скачанных модулей. Строка с `/go.mod` проверяет файл `go.mod` зависимости,
а строка без этого суффикса — файлы модуля. При несовпадении суммы Go выдаст ошибку. Выбор версий при
этом определяется требованиями модулей. В проекте только со стандартной библиотекой `go.sum` может отсутствовать.
Если файл появился, его сохраняют в Git вместе с `go.mod`.

## Создание и запуск модуля

1. Создайте отдельную папку и выполните `go mod init`. Команда создаст `go.mod` в текущем каталоге.

   ```text
   mkdir hello
   cd hello
   go mod init example.org/hello
   ```

   Для проекта на GitHub обычно указывают адрес репозитория без `https://` и `.git`. Например,
   `go mod init github.com/your-username/your-project`. Для локального примера подойдёт и `go mod init hello`.

2. Создайте `main.go` в той же папке.

   ```go
   package main

   import "fmt"

   func main() {
       fmt.Println("Привет, Go!") // Привет, Go!
   }
   ```

3. Запустите и соберите программу.

   ```text
   go run main.go
   go run .
   go build .
   ```

   Обе команды `go run` выведут `Привет, Go!`. Точка выбирает весь пакет из текущего каталога. `go build .`
   создаст `hello.exe` на Windows или `hello` на Linux и macOS. Имя взято из конца пути `example.org/hello`.

В этом уроке уже есть [go.mod](go.mod) и два файла одного пакета: [hello.go](hello.go) и [greeting.go](greeting.go).
Если запустить `go run .` из корня репозитория, сначала появится `Привет, читатель!`, затем `Привет, Go!`.
При `go run hello.go` выводится только `Привет, Go!`, потому что в команду передан один файл.
Работа с несколькими файлами разобрана в [уроке о пакетах][packages-lesson].

## Команды для зависимостей

1. В пробном `main.go` добавьте импорт `"rsc.io/quote"` и замените вызов `fmt.Println`.

   ```go
   fmt.Println(quote.Go()) // Don't communicate by sharing memory, share memory by communicating.
   ```

2. Подключите версию `v1.5.2` и запустите программу.

   ```text
   go get rsc.io/quote@v1.5.2
   go run .
   ```

Обычные `go build` и `go test` с Go 1.16 по умолчанию сообщают об ошибке, если нужно добавить зависимость.
Выберите версию через `go get` или выполните `go mod tidy`, чтобы Go подобрал недостающие зависимости.

- `go get rsc.io/quote@latest` — выбирает последнюю версию по этому пути, предпочитая стабильные релизы.
- `go get -u` — обновляет прямые и косвенные зависимости пакета в текущей папке.
- `go get -u=patch` — обновляет те же зависимости в рамках патч-версий.
- `go mod tidy` — добавляет недостающее, удаляет лишнее и обновляет `go.sum`. Учитывает импорты в тестах.

Например, если удалить из программы импорт `rsc.io/quote` и вызов `quote.Go()`, после `go mod tidy`
эта зависимость исчезнет, если её больше ничего не использует.

Для мажорных версий начиная с `v2` меняется путь модуля, например с `example.org/lib` на `example.org/lib/v2`.
Чтобы перейти на такую версию, нужно поменять импорт. Если запросить коммит без тега версии, Go запишет
псевдоверсию с датой и частью хеша коммита.

## Локальная копия зависимости

Если библиотека `example.org/lib` лежит рядом с проектом в папке `lib`, её можно временно подключить через
`replace` в `go.mod` проекта.

```text
require example.org/lib v1.2.0
replace example.org/lib => ../lib
```

Импорты остаются прежними, но Go берёт исходники из `../lib`. В этой папке должен быть свой `go.mod`.
`replace` сам по себе не добавляет зависимость, поэтому в примере нужен и `require`.

## Источники

- [Модули и версии](https://go.dev/ref/mod#modules-overview) · [go mod init](https://go.dev/ref/mod#go-mod-init) ·
  [go run](https://pkg.go.dev/cmd/go#hdr-Compile_and_run_Go_program) ·
  [go build](https://pkg.go.dev/cmd/go#hdr-Compile_packages_and_dependencies).
- [Директивы go.mod](https://go.dev/ref/mod#go-mod-file) · [файл go.sum](https://go.dev/ref/mod#go-sum-files) ·
  [выбор версий MVS](https://go.dev/ref/mod#minimal-version-selection).
- [Минимальная версия Go](https://go.dev/ref/mod#go-mod-file-go) ·
  [команды сборки с Go 1.16](https://go.dev/doc/go1.16#go-command).
- [go get](https://go.dev/ref/mod#go-get) · [go mod tidy](https://go.dev/ref/mod#go-mod-tidy) ·
  [запросы версий](https://go.dev/ref/mod#version-queries) · [replace](https://go.dev/ref/mod#go-mod-file-replace).
- [Мажорные версии](https://go.dev/ref/mod#major-version-suffixes) ·
  [псевдоверсии](https://go.dev/ref/mod#pseudo-versions).
- [Код для примера с rsc.io/quote](https://go.dev/doc/tutorial/getting-started#call).

[packages-lesson]: https://github.com/Hardome/learn-golang/tree/lesson-022-packages-and-imports
