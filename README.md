# Урок 029. Пакет strings. Операции со строками

`strings` — пакет стандартной библиотеки Go для поиска, разбиения, объединения и преобразования строк.
Go хранит строку как последовательность байтов. **Поиск подстроки сравнивает байты**, а операции вроде
`ContainsAny` читают текст как последовательность рун Unicode.

Строковые значения неизменяемы. Чтобы сохранить результат преобразования, присвойте его переменной.
Фрагменты используют импорты `"fmt"` и `"strings"`. Команда `go run hello.go` запускает рабочие примеры.
[Код примеров в hello.go](hello.go)

## Верхний и нижний регистр

### Верхний регистр

`strings.ToUpper` переводит буквы Unicode в верхний регистр. Так можно вывести заголовок прописными буквами.
Цифры и знаки препинания без регистрового варианта сохраняются.

```go
fmt.Println(strings.ToUpper("Привет, Go! 42")) // ПРИВЕТ, GO! 42
```

### Нижний регистр

`strings.ToLower` переводит буквы в нижний регистр. Аргумент и результат имеют тип `string`.

```go
fmt.Println(strings.ToLower("Привет, Go! 42")) // привет, go! 42
```

### Отличие от JavaScript при смене регистра

`strings.ToUpper` использует простое преобразование регистра Unicode, где результат для каждой руны —
одна кодовая точка. В этой таблице для `ß` нет прописного соответствия, поэтому она сохраняется.
В JavaScript `"straße".toUpperCase()` возвращает `"STRASSE"`, поскольку там допустима замена одной
кодовой точки несколькими.

```go
fmt.Println(strings.ToUpper("straße")) // STRAßE
```

## Проверка префикса и суффикса

### Префикс

`strings.HasPrefix(s, prefix)` сравнивает префикс с началом `s`. Например, по `"report-"` можно отобрать
имена файлов с отчётами.

```go
filename := "report-2026.csv"
fmt.Println(strings.HasPrefix(filename, "report-")) // true
```

### Суффикс

`strings.HasSuffix(s, suffix)` сравнивает суффикс с концом `s`. Обе проверки требуют точного совпадения
байтов на нужном краю. Поэтому у файла с расширением `.csv` проверка `.CSV` возвращает `false`.

```go
filename := "report-2026.csv"
fmt.Println(strings.HasSuffix(filename, ".csv")) // true
fmt.Println(strings.HasSuffix(filename, ".CSV")) // false
```

## Поиск текста

### Поиск целой подстроки

`strings.Contains(s, substr)` ищет байты подстроки подряд в заданном порядке. Функция подходит для проверки
фрагмента текста. Поиск отдельного слова потребовал бы ещё проверить его границы.

```go
text := "Hello World"
fmt.Println(strings.Contains(text, "World")) // true
fmt.Println(strings.Contains(text, "Work"))  // false
```

### Поиск любой руны из набора

`strings.ContainsAny(s, chars)` рассматривает `chars` как набор рун. Достаточно **любой одной руны** из набора.
Так можно проверить наличие запрещённых знаков. Для `"Work"` в примере достаточно совпадения по букве `W`.

```go
text := "Hello World"
fmt.Println(strings.ContainsAny(text, "Work")) // true — в World есть W.
fmt.Println(strings.ContainsAny(text, "xyz"))  // false
```

## Количество вхождений в строку

`strings.Count(s, substr)` считает вхождения точной подстроки. Регистр имеет значение, поэтому `World`
и `world` считаются отдельно. При отсутствии совпадений функция возвращает `0`.

```go
message := "Hello World. Good bye World"
fmt.Println(strings.Count(message, "World")) // 2
fmt.Println(strings.Count(message, "world")) // 0
```

### Перекрывающиеся вхождения

`Count` считает **неперекрывающиеся вхождения**. В `"aaa"` первое совпадение с `"aa"` занимает байты `0` и `1`.
Следующий поиск начинается после него. Одного оставшегося байта для второго `"aa"` мало.

```go
fmt.Println(strings.Count("aaa", "aa")) // 1
```

### Пустая подстрока

Пустая подстрока совпадает в начале строки и после каждой руны. В `"Пр"` две руны, поэтому `Count` находит
три позиции, хотя байтов четыре.

```go
fmt.Println(strings.Count("Пр", "")) // 3 — начало, место между буквами и конец.
```

## Объединение среза строк

`strings.Join(elems, sep)` собирает `[]string` в одну строку. Разделитель вставляется только **между**
соседними элементами. Для предложения подойдёт пробел, для объединения без разделителей — `""`.

```go
words := []string{"Привет", "прекрасный", "мир"}
fmt.Println(strings.Join(words, " ")) // Привет прекрасный мир
```

Когда все части уже лежат в срезе, `Join` удобен и для больших наборов. Текущая реализация при двух и более
элементах сначала вычисляет длину результата и резервирует буфер, затем записывает в него части и разделители.
Так она обходится без конкатенации растущей строки на каждом шаге. Число выделений памяти зависит от реализации.

## Замена подстроки

### Замена заданного числа вхождений

`strings.Replace(s, old, new, n)` заменяет первые `n` неперекрывающихся вхождений. Чтобы поменять первое
`World`, передадим `1` и сохраним результат в `message`.

```go
message := "Hello World. Good bye World"
message = strings.Replace(message, "World", "Go", 1)
fmt.Println(message) // Hello Go. Good bye World
```

- `n > 0` — заменить до `n` совпадений. Если их меньше, заменяются все найденные.
- `n == 0` — вернуть исходную строку.
- `n < 0` — снять ограничение по количеству замен.

Если `old` отсутствует, результат совпадает с исходной строкой.

### Замена всех вхождений

Для замены всех вхождений используйте `strings.ReplaceAll(s, old, new)`. Это то же поведение, что у `Replace`
с `n = -1`, без отдельного аргумента для лимита.

```go
message := "Hello World. Good bye World"
fmt.Println(strings.ReplaceAll(message, "World", "Go")) // Hello Go. Good bye Go
```

### Замена пустой подстроки

Пустой `old` совпадает в тех же позициях, которые считает `Count(s, "")`. У `"Пр"` их три. Поэтому `ReplaceAll`
вставит дефис перед `П`, между `П` и `р` и после `р`.

```go
fmt.Println(strings.ReplaceAll("Пр", "", "-")) // -П-р-
```

## Разбиение строки на части

### Разбиение по разделителю

`strings.Split(s, sep)` возвращает `[]string` из частей между совпадениями с разделителем `sep`.
Например, номер можно разделить по дефисам. Формат `%q` показывает границы отдельных элементов.

```go
phone := "+7-987-654-32-10"
fmt.Printf("%q\n", strings.Split(phone, "-")) // ["+7" "987" "654" "32" "10"]
```

Если непустого разделителя в строке нет, вся строка становится одним элементом. Между двумя соседними
запятыми и после последней запятой текст пустой, поэтому **`Split` сохраняет пустые поля**.

```go
fmt.Printf("%q\n", strings.Split("7999", "-"))  // ["7999"]
fmt.Printf("%q\n", strings.Split("a,,b,", ",")) // ["a" "" "b" ""]
```

### Пустой разделитель

При `sep == ""` функция выделяет по одному элементу на последовательность UTF-8. В корректном UTF-8
это одна руна на элемент. Если оба аргумента пусты, результат — пустой срез.

```go
fmt.Printf("%q\n", strings.Split("П!", "")) // ["П" "!"]
```

### Разбиение по пробельным символам

`strings.Fields(s)` выделяет непустые части, разделённые пробельными символами Unicode. Подряд идущие пробелы,
табуляции и переводы строк служат общей границей. Пробелы по краям строки пропускаются.

```go
text := " \tHello  World\nGood bye World "
fmt.Printf("%q\n", strings.Fields(text)) // ["Hello" "World" "Good" "bye" "World"]
```

## Удаление с краёв строки

### По набору рун

`Trim` удаляет с обоих краёв руны, перечисленные во втором аргументе. `TrimLeft` работает с начала,
`TrimRight` — с конца. На каждом краю обработка останавливается на первой руне вне набора.

```go
text := "01.2300"
fmt.Println(strings.Trim(text, "0"))      // 1.23
fmt.Println(strings.TrimLeft(text, "0"))  // 1.2300
fmt.Println(strings.TrimRight(text, "0")) // 01.23
fmt.Println(strings.Trim("?!Go!?", "!?")) // Go — порядок рун в наборе произвольный.
```

### По условию

`TrimFunc`, `TrimLeftFunc` и `TrimRightFunc` работают с обоими краями, началом и концом соответственно.
Вместо набора рун они принимают `func(rune) bool`. Значение `true` говорит, что руну с текущего края нужно убрать.
В примере отсечение с каждого края заканчивается на первой цифре ASCII, поэтому внутренние дефисы остаются.

```go
phone := "+7-987-654-32-10!"
outsideDigit := func(symbol rune) bool {
	return symbol < '0' || symbol > '9'
}
fmt.Println(strings.TrimFunc(phone, outsideDigit))      // 7-987-654-32-10
fmt.Println(strings.TrimLeftFunc(phone, outsideDigit))  // 7-987-654-32-10!
fmt.Println(strings.TrimRightFunc(phone, outsideDigit)) // +7-987-654-32-10
```

### Пробельные символы

`TrimSpace` убирает пробельные символы Unicode с обоих краёв. Это удобно для очистки пользовательского
ввода с табуляцией или переводом строки в конце. Пробелы внутри значения сохраняются.

```go
fmt.Printf("%q\n", strings.TrimSpace(" \t7 987 654 32 10\n")) // "7 987 654 32 10"
```

### Точный префикс и суффикс

`TrimPrefix` удаляет указанную последовательность байтов с начала, `TrimSuffix` — с конца. Каждая функция
удаляет совпадение **один раз**. Если на нужном краю совпадения нет, возвращается исходное значение.

```go
numberLine := "+7 987 654 32 10;"
fmt.Printf("%q\n", strings.TrimPrefix(numberLine, "+7")) // " 987 654 32 10;" — пробел остаётся.
fmt.Printf("%q\n", strings.TrimSuffix(numberLine, ";"))  // "+7 987 654 32 10"
```

На строке `"--Go"` разница видна по количеству удалённых дефисов.

```go
fmt.Println(strings.TrimPrefix("--Go", "-")) // -Go — снят один префикс.
fmt.Println(strings.TrimLeft("--Go", "-"))   // Go — сняты все дефисы с начала.
```

## Индекс подстроки

`Index`, `LastIndex` и варианты с `Any` возвращают **байтовое смещение**. В строке `"П: Go, Go!"`
буква `П` занимает два байта, поэтому первое `Go` начинается с байта `4`.

```text
Байт    0  1 | 2 | 3 | 4  5 | 6 | 7 | 8  9 | 10
Текст      П | : | ␠ | G  o | , | ␠ | G  o |  !
```

Знаком `␠` обозначен пробел.

### Первое вхождение

`strings.Index(s, substr)` находит начало первого вхождения. `-1` означает отсутствие совпадения.
**Перед взятием подстроки проверьте индекс.** Ноль означает совпадение в начале, а отрицательный индекс
при вычислении границы вызовет панику.

```go
text := "П: Go, Go!"
offset := strings.Index(text, "Go")
fmt.Println(offset) // 4
if offset >= 0 {
	fmt.Println(text[offset:]) // Go, Go!
}
fmt.Println(strings.Index(text, "go")) // -1
```

### Последнее вхождение

`strings.LastIndex(s, substr)` находит начало последнего вхождения. При отсутствии подстроки ответ тоже `-1`.

```go
text := "П: Go, Go!"
fmt.Println(strings.LastIndex(text, "Go")) // 8 — начало второго Go.
```

## Индекс любой руны из набора

### Первое совпадение

`strings.IndexAny(s, chars)` находит первую по ходу строки руну из `chars`. В `"П: Go, Go!"` руна `o`
встречается раньше `!`, хотя в наборе `"!o"` она записана второй.

```go
text := "П: Go, Go!"
fmt.Println(strings.IndexAny(text, "!o")) // 5 — байтовое смещение первой o.
fmt.Println(strings.IndexAny(text, "xy")) // -1
```

### Последнее совпадение

`strings.LastIndexAny(s, chars)` возвращает байтовое смещение последней подходящей руны или `-1`.

```go
text := "П: Go, Go!"
fmt.Println(strings.LastIndexAny(text, "!o")) // 10 — байтовое смещение !.
```

## Какую функцию выбрать

| Нужно | Функция |
| --- | --- |
| Проверить точную подстроку или любую руну из набора | `Contains` или `ContainsAny` |
| Ограничить число замен или заменить все вхождения | `Replace` или `ReplaceAll` |
| Сохранить пустые поля или выделить непустые части среди пробелов | `Split` или `Fields` |
| Снять набор рун с краёв или один точный фрагмент | `Trim` или `TrimPrefix`/`TrimSuffix` |
| Найти позицию подстроки или позицию любой подходящей руны | `Index` или `IndexAny` |

## Источники

- [Go Specification · String types](https://go.dev/ref/spec#String_types)
- [Go Specification · Slice expressions](https://go.dev/ref/spec#Slice_expressions)
- `strings` · [ToUpper](https://pkg.go.dev/strings#ToUpper), [ToLower](https://pkg.go.dev/strings#ToLower)
- [Реализация strings.ToUpper](https://go.dev/src/strings/strings.go#L668)
- [ECMAScript · String.prototype.toUpperCase](https://tc39.es/ecma262/#sec-string.prototype.touppercase)
- `strings` · [HasPrefix](https://pkg.go.dev/strings#HasPrefix), [HasSuffix](https://pkg.go.dev/strings#HasSuffix)
- `strings` · [Contains](https://pkg.go.dev/strings#Contains), [ContainsAny](https://pkg.go.dev/strings#ContainsAny)
- `strings` · [Count](https://pkg.go.dev/strings#Count), [Join](https://pkg.go.dev/strings#Join)
- [Реализация strings.Join](https://go.dev/src/strings/strings.go#L468)
- `strings` · [Replace](https://pkg.go.dev/strings#Replace), [ReplaceAll](https://pkg.go.dev/strings#ReplaceAll)
- `strings` · [Split](https://pkg.go.dev/strings#Split), [Fields](https://pkg.go.dev/strings#Fields)
- `strings` · [Trim](https://pkg.go.dev/strings#Trim), [TrimLeft](https://pkg.go.dev/strings#TrimLeft),
  [TrimRight](https://pkg.go.dev/strings#TrimRight)
- `strings` · [TrimFunc](https://pkg.go.dev/strings#TrimFunc), [TrimLeftFunc](https://pkg.go.dev/strings#TrimLeftFunc),
  [TrimRightFunc](https://pkg.go.dev/strings#TrimRightFunc)
- `strings` · [TrimSpace](https://pkg.go.dev/strings#TrimSpace), [TrimPrefix](https://pkg.go.dev/strings#TrimPrefix),
  [TrimSuffix](https://pkg.go.dev/strings#TrimSuffix)
- `strings` · [Index](https://pkg.go.dev/strings#Index), [LastIndex](https://pkg.go.dev/strings#LastIndex)
- `strings` · [IndexAny](https://pkg.go.dev/strings#IndexAny), [LastIndexAny](https://pkg.go.dev/strings#LastIndexAny)
- [unicode.IsSpace](https://pkg.go.dev/unicode#IsSpace) · [fmt · Printing](https://pkg.go.dev/fmt#hdr-Printing)
