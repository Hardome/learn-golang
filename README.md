# Урок 025. Загрузка внешних зависимостей

Внешний пакет можно импортировать в свой код так же, как пакет стандартной библиотеки. Команда `go` берёт
внешний пакет из версии модуля, выбранной для проекта. В этом уроке подключим демонстрационный пакет
`rsc.io/quote`, вызовем две функции и посмотрим, как меняются `go.mod` и `go.sum`.

## Добавление зависимости

В новом каталоге сначала выполните `go mod init example.org/quotes`. Команда создаст в нём `go.mod`.
У этого урока уже есть [go.mod](go.mod), поэтому готовый пример можно сразу запустить из корня репозитория.

Чтобы добавить зависимость в новый модуль, выполните:

```text
go get rsc.io/quote@v1.5.2
```

`rsc.io/quote` здесь одновременно путь импорта пакета и путь модуля, потому что пакет находится в корне
модуля. Суффикс `@v1.5.2` запрашивает конкретную версию. В режиме модулей начиная с Go 1.18 команда
`go get` управляет зависимостями проекта. Она обновляет требования в `go.mod` и при необходимости
скачивает модули. После первого скачивания в примере появляется и `go.sum`.

Пока исходники проекта ещё не импортируют `rsc.io/quote`, команда добавляет запись с пометкой `// indirect`:

```text
rsc.io/quote v1.5.2 // indirect
```

## Импорт пакета и запуск

Создайте `hello.go` в корне нового модуля. Готовый пример лежит в [hello.go](hello.go).

```go
package main

import (
	"fmt"

	"rsc.io/quote" // Пакет из внешнего модуля.
)

func main() {
	// Обе функции возвращают фиксированные строки.
	fmt.Println(quote.Go())  // Don't communicate by sharing memory, share memory by communicating.
	fmt.Println(quote.Opt()) // If a program is too slow, it must have a loop.
}
```

В `import` указывают путь пакета. Версия хранится в требованиях модуля. Из его корня выполните
`go mod tidy`, затем `go run .`. Программа напечатает две строки, указанные в комментариях.

### Прямые и косвенные зависимости

После добавления импорта `go mod tidy` убирает у `rsc.io/quote` пометку `// indirect`, потому что теперь
его пакет импортируется в `hello.go`. Пакет `quote` использует `rsc.io/sampler`, а тот использует
`golang.org/x/text`. Для нашего модуля это косвенные зависимости. В `go.mod` урока получились такие требования:

```text
require rsc.io/quote v1.5.2

require (
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/sampler v1.3.0 // indirect
)
```

## Контрольные суммы и кеш модулей

В [go.sum](go.sum) Go хранит контрольные суммы, чтобы при загрузке проверить целостность модулей. Для
`rsc.io/quote` в нём есть две строки:

```text
rsc.io/quote v1.5.2 h1:w5fcysjrx7yqtD/aO+QwRjYZOKnaM9Uh2b40tElTs3Y=
rsc.io/quote v1.5.2/go.mod h1:LzX7hefJvL54yjefDEDHNONDjII0t9xZLPXsUe+TKr0=
```

Первая сумма относится к содержимому архива модуля, вторая проверяет только его `go.mod`.
Сохраняйте `go.mod` и `go.sum` в Git вместе с кодом проекта.

Загруженные модули лежат в общем для проектов кеше `GOMODCACHE`. По умолчанию это каталог `pkg/mod`
в первом каталоге `GOPATH`. Фактические пути можно посмотреть командами:

```text
go env GOPATH
go env GOMODCACHE
```

## Что меняет go mod tidy

`go mod tidy` приводит `go.mod` и `go.sum` в соответствие с импортами проекта и их зависимостями,
включая импорты из тестов.

- Если в новом модуле сначала написать импорт `rsc.io/quote` и выполнить `go mod tidy`, Go сам найдёт
  модуль для этого пакета и добавит требование. `go get rsc.io/quote@v1.5.2` позволяет явно указать версию.
- Если в пробном проекте убрать импорт `rsc.io/quote` и заменить оба вызова одной строкой
  `fmt.Println("Готово")`, после `go mod tidy` блоки `require` исчезнут, а `go.sum` опустеет или будет
  удалён. Все три модуля больше не нужны этому проекту. Их файлы сохранятся в кеше.

О пути модуля и структуре `go.mod` подробнее в [уроке 023 о модулях][modules-lesson].

## Источники

- [Go Modules Reference · go mod init](https://go.dev/ref/mod#go-mod-init)
- [Go Modules Reference · go get](https://go.dev/ref/mod#go-get)
- [Go 1.18 Release Notes · go get](https://go.dev/doc/go1.18#go-get)
- [Go Modules Reference · require](https://go.dev/ref/mod#go-mod-file-require)
- [Go Modules Reference · go mod tidy](https://go.dev/ref/mod#go-mod-tidy)
- [Go Modules Reference · go.sum files](https://go.dev/ref/mod#go-sum-files)
- [Go Modules Reference · Module cache](https://go.dev/ref/mod#module-cache)
- [cmd/go · go env](https://pkg.go.dev/cmd/go#hdr-Print_Go_environment_information)
- [cmd/go · go run](https://pkg.go.dev/cmd/go#hdr-Compile_and_run_Go_program)
- [quote.Go](https://pkg.go.dev/rsc.io/quote@v1.5.2#Go) · [quote.Opt](https://pkg.go.dev/rsc.io/quote@v1.5.2#Opt)
- [rsc.io/quote · go.mod](https://github.com/rsc/quote/blob/v1.5.2/go.mod)
- [rsc.io/sampler · go.mod](https://github.com/rsc/sampler/blob/v1.3.0/go.mod)

[modules-lesson]: https://github.com/Hardome/learn-golang/tree/lesson-023-modules
