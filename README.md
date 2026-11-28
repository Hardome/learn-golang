# Урок 024. Разделение проекта на пакеты

Проект на Go можно разделить на пакеты по назначению. В этом уроке пакет `checkout` формирует строку с итоговой
стоимостью покупки, а пакеты `card` и `cash` содержат названия способов оплаты. Все пакеты входят в один модуль.

## Каталоги, пакеты и пути импорта

[go.mod](go.mod) задаёт путь модуля `github.com/Hardome/learn-golang`. Из корня [hello.go](hello.go) использует
три пакета, исходники которых лежат в отдельных каталогах.

```text
.                 # пакет main
├── go.mod
├── hello.go
└── checkout/     # пакет checkout
    ├── receipt.go
    ├── total.go
    ├── card/     # пакет card
    │   └── label.go
    └── cash/     # пакет cash
        └── label.go
```

Для пакета в подкаталоге путь импорта состоит из пути модуля и пути к каталогу от корня.
У пакета в `checkout/` получается `github.com/Hardome/learn-golang/checkout`. Его имя берётся
из `package checkout`. Подробнее о директиве `module` в [уроке 023][modules-lesson].

## Один пакет в нескольких файлах

Go собирает файлы одного пакета из одного каталога. В файлах `checkout/receipt.go` и `checkout/total.go`
указано одно имя пакета `checkout`. Функция `Receipt` вызывает `total` напрямую, потому что обе объявлены
в одном пакете.

`checkout/receipt.go`

```go
package checkout

import "fmt"

func Receipt(unitPrice, quantity int) string {
	return fmt.Sprintf("Итого: %d ₽", total(unitPrice, quantity)) // total объявлена в другом файле пакета.
}
```

`checkout/total.go`

```go
package checkout

func total(unitPrice, quantity int) int {
	return unitPrice * quantity // Цена одной единицы умножается на количество.
}
```

### Доступ из другого пакета

Функция, объявленная на уровне пакета, экспортируется, если её имя начинается с заглавной буквы Unicode.
Поэтому из пакета `main` можно вызвать `checkout.Receipt`, а функция `total` доступна только внутри
`checkout`. В импорте указывают полный путь к пакету. В вызове без псевдонима используют имя из его
объявления `package checkout`.

```go
package main

import (
	"fmt"

	"github.com/Hardome/learn-golang/checkout"
)

func main() {
	fmt.Println(checkout.Receipt(250, 3)) // Итого: 750 ₽
	// checkout.total(250, 3) // Ошибка компиляции: total доступна только внутри checkout.
}
```

## Пакеты во вложенных каталогах

Каталоги `checkout/card/` и `checkout/cash/` содержат отдельные пакеты со своими именами и путями импорта.
Вложенность каталогов группирует код. Каждый пакет импортируется отдельно. Правила экспорта действуют
и между `checkout` и пакетами в его подкаталогах.

`checkout/card/label.go`

```go
package card

func Label() string {
	return "Карта" // Название оплаты картой.
}
```

`checkout/cash/label.go`

```go
package cash

func Label() string {
	return "Наличные" // Название оплаты наличными.
}
```

Файлы можно назвать одинаково, потому что они лежат в разных каталогах. Функции `Label` тоже могут иметь
одинаковые имена. Вызовы различаются по имени пакета.

```go
package main

import (
	"fmt"

	"github.com/Hardome/learn-golang/checkout/card"
	"github.com/Hardome/learn-golang/checkout/cash"
)

func main() {
	fmt.Println(card.Label()) // Карта
	fmt.Println(cash.Label()) // Наличные
}
```

## Запуск примера

Из корня проекта выполните `go run .`. Точка выбирает пакет `main` из текущего каталога. Команда `go`
сама загружает импортируемые пакеты из их каталогов. В этом примере `go run hello.go` даёт тот же результат.

```text
Итого: 750 ₽
Карта
Наличные
```

[modules-lesson]: https://github.com/Hardome/learn-golang/tree/lesson-023-modules

## Источники

- [Organizing a Go module · Multiple packages](https://go.dev/doc/modules/layout#multiple-packages)
- [Go Specification · Package clause](https://go.dev/ref/spec#Package_clause)
- [Go Specification · Declarations and scope](https://go.dev/ref/spec#Declarations_and_scope)
- [Go Specification · Exported identifiers](https://go.dev/ref/spec#Exported_identifiers)
- [Go Specification · Import declarations](https://go.dev/ref/spec#Import_declarations)
- [Package names · Package paths](https://go.dev/blog/package-names#package-paths)
- [cmd/go · go run](https://pkg.go.dev/cmd/go#hdr-Compile_and_run_Go_program)
- [fmt.Sprintf](https://pkg.go.dev/fmt#Sprintf) · [fmt.Println](https://pkg.go.dev/fmt#Println)
