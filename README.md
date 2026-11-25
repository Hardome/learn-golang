# Изучаем Go

Учебный репозиторий для изучения Go и повторения перед техническими собеседованиями. Уроки начинаются с базовых
конструкций и разбирают их до уровня senior: гарантии языка, пограничные случаи, особенности версий и поведение кода, на
которое можно опираться.

Каждый урок хранится в отдельной ветке `lesson-NNN-topic`. В её `README.md` находятся теория, пояснения и ссылки на
официальные источники. В `hello.go` собраны запускаемые примеры с ожидаемым выводом в комментариях.

## Оглавление

| № | Урок |
| --- | --- |
| 001 | [Hello, world!](https://github.com/Hardome/learn-golang/tree/lesson-001-hello-world) |
| 002 | [Переменные](https://github.com/Hardome/learn-golang/tree/lesson-002-variables) |
| 003 | [Типы](https://github.com/Hardome/learn-golang/tree/lesson-003-types) |
| 004 | [Константы](https://github.com/Hardome/learn-golang/tree/lesson-004-const) |
| 005 | [Арифметические операции](https://github.com/Hardome/learn-golang/tree/lesson-005-arithmetic-operations) |
| 006 | [Условные выражения](https://github.com/Hardome/learn-golang/tree/lesson-006-conditional-expressions) |
| 007 | [Поразрядные операции](https://github.com/Hardome/learn-golang/tree/lesson-007-bitwise-operations) |
| 008 | [Операторы присваивания](https://github.com/Hardome/learn-golang/tree/lesson-008-assignment-operators) |
| 009 | [Массивы](https://github.com/Hardome/learn-golang/tree/lesson-009-arrays) |
| 010 | [Условные конструкции](https://github.com/Hardome/learn-golang/tree/lesson-010-conditional-statements) |
| 011 | [Циклы](https://github.com/Hardome/learn-golang/tree/lesson-011-loops) |
| 012 | [Функции и их параметры](https://github.com/Hardome/learn-golang/tree/lesson-012-functions-and-parameters) |

## Запуск урока

Установите Go 1.22 или новее, чтобы запустить все текущие уроки.
[Инструкция по установке Go](https://go.dev/doc/install)

```sh
git clone https://github.com/Hardome/learn-golang.git
cd learn-golang
git switch lesson-001-hello-world
go run hello.go
go vet hello.go
```

Для другого урока подставьте имя его ветки. Ветки содержат самостоятельные примеры, поэтому команда запуска указывает
файл напрямую. Сверяйте результат с комментариями к примерам в `hello.go`.

## Как дополнить урок

Откройте ветку по теме, обновите её теорию и примеры, добавьте ссылки на конкретные разделы официальной документации.
Правила оформления, проверки и коммитов описаны в [AGENTS.md](AGENTS.md).

## Документация Go

- [Спецификация языка Go](https://go.dev/ref/spec)
- [Effective Go](https://go.dev/doc/effective_go)
- [Стандартная библиотека](https://pkg.go.dev/std)
- [Примечания к выпускам](https://go.dev/doc/devel/release)
