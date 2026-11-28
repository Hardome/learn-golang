# Урок 025. Загрузка внешних зависимостей

В записи `import` указывают путь пакета. Для внешнего пакета команда `go` берёт исходники из выбранной версии
модуля. Требования к версиям хранятся в `go.mod`, а проверочные суммы скачанных файлов — в `go.sum`.

## Добавление зависимости

В этой ветке зависимости уже записаны в [go.mod](go.mod), поэтому готовый пример запускается из корня
командой `go run .`. Чтобы проследить изменения файлов, начните в отдельном пустом каталоге.

```text
go mod init example.org/quotes
go get rsc.io/quote@v1.5.2
```

После `go mod init` файл содержит директивы `module` и `go`. Требований `require` пока нет.
`go get` добавляет требование с указанной версией `v1.5.2` и при необходимости загружает модули.
В `go.sum` появляются их проверочные суммы.

Поскольку в каталоге пока нет файла с импортом `rsc.io/quote`, требование входит в блок с пометкой `// indirect`.

```text
require (
    golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
    rsc.io/quote v1.5.2 // indirect
    rsc.io/sampler v1.3.0 // indirect
)
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

Пакет лежит в корне модуля `rsc.io/quote`, поэтому путь пакета совпадает с путём модуля.
Версия указана в `require`. Из корня пробного модуля выполните `go mod tidy`, затем `go run .`.
Программа напечатает две строки из комментариев.

### Прямые и косвенные зависимости

`go mod tidy` сопоставляет требования с импортами исходных файлов. Поскольку `hello.go` импортирует
`rsc.io/quote`, его модуль становится прямой зависимостью. Пакет `quote` использует `rsc.io/sampler`,
а тот использует `golang.org/x/text`. Эти два модуля остаются косвенными зависимостями.
Так выглядит итоговый блок требований в `go.mod` этой ветки.

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

[modules-lesson]: https://github.com/Hardome/learn-golang/tree/lesson-023-modules

## Источники

- [Go Modules Reference · go mod init](https://go.dev/ref/mod#go-mod-init)
- [Go Modules Reference · go get](https://go.dev/ref/mod#go-get)
- [Go Modules Reference · require](https://go.dev/ref/mod#go-mod-file-require)
- [Go Modules Reference · go mod tidy](https://go.dev/ref/mod#go-mod-tidy)
- [Go Modules Reference · go.sum files](https://go.dev/ref/mod#go-sum-files)
- [Go Modules Reference · Module cache](https://go.dev/ref/mod#module-cache)
- [cmd/go · go env](https://pkg.go.dev/cmd/go#hdr-Print_Go_environment_information)
- [cmd/go · go run](https://pkg.go.dev/cmd/go#hdr-Compile_and_run_Go_program)
- [quote.Go](https://pkg.go.dev/rsc.io/quote@v1.5.2#Go) · [quote.Opt](https://pkg.go.dev/rsc.io/quote@v1.5.2#Opt)
- [rsc.io/quote · go.mod](https://github.com/rsc/quote/blob/v1.5.2/go.mod)
- [rsc.io/sampler · go.mod](https://github.com/rsc/sampler/blob/v1.3.0/go.mod)
