# Изучаем Go

Учебный репозиторий для изучения Go и повторения перед собеседованиями. У каждого урока своя тема, короткие
примеры и пограничные случаи, которые помогают понять поведение кода.

Каждый урок хранится в отдельной ветке `lesson-NNN-topic`.

`README.md` объясняет тему целиком: вводит понятия, разбирает примеры и даёт ссылки на официальные источники.

В `hello.go` собраны запускаемые примеры для практики. Ожидаемый вывод указан в комментариях.

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
| 013 | [Анонимные функции][lesson-013] |
| 014 | [Замыкания][lesson-014] |
| 015 | [Рекурсивные функции][lesson-015] |
| 016 | [Указатели][lesson-016] |
| 017 | [Псевдонимы и определения типов][lesson-017] |
| 018 | [Структуры][lesson-018] |
| 019 | [Методы][lesson-019] |
| 020 | [Срезы][lesson-020] |
| 021 | [Карты][lesson-021] |
| 022 | [Пакеты и их импорт][lesson-022] |
| 023 | [Модули][lesson-023] |

## Запуск урока

Установите Go 1.27.4 или новее, чтобы запустить все уроки. Версия указана в модуле урока 023.
[Инструкция по установке Go](https://go.dev/doc/install)

```sh
git clone https://github.com/Hardome/learn-golang.git
cd learn-golang
git switch lesson-001-hello-world
go run hello.go
go vet hello.go
```

Для другого урока подставьте имя его ветки. Если в уроке несколько файлов, используйте команду из его README.
Сверяйте результат с комментариями к примерам.

## Как дополнить урок

Откройте ветку по теме. Дайте объяснения и короткие примеры в `README.md`, чтобы урок можно было понять при чтении.

Запускаемые примеры храните в `hello.go`.

Добавьте ссылки на конкретные разделы официальной документации. Правила оформления, проверки и коммитов описаны
в [AGENTS.md](AGENTS.md).

## Документация Go

- [Спецификация языка Go](https://go.dev/ref/spec)
- [Effective Go](https://go.dev/doc/effective_go)
- [Стандартная библиотека](https://pkg.go.dev/std)
- [Примечания к выпускам](https://go.dev/doc/devel/release)

[lesson-013]: https://github.com/Hardome/learn-golang/tree/lesson-013-anonymous-functions
[lesson-014]: https://github.com/Hardome/learn-golang/tree/lesson-014-closures
[lesson-015]: https://github.com/Hardome/learn-golang/tree/lesson-015-recursive-functions
[lesson-016]: https://github.com/Hardome/learn-golang/tree/lesson-016-pointers
[lesson-017]: https://github.com/Hardome/learn-golang/tree/lesson-017-type-aliases
[lesson-018]: https://github.com/Hardome/learn-golang/tree/lesson-018-structs
[lesson-019]: https://github.com/Hardome/learn-golang/tree/lesson-019-methods
[lesson-020]: https://github.com/Hardome/learn-golang/tree/lesson-020-slices
[lesson-021]: https://github.com/Hardome/learn-golang/tree/lesson-021-maps
[lesson-022]: https://github.com/Hardome/learn-golang/tree/lesson-022-packages-and-imports
[lesson-023]: https://github.com/Hardome/learn-golang/tree/lesson-023-modules
