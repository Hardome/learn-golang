# Урок 020. Срезы в Go

**Срез** (`slice`) — составной тип данных, который описывает последовательный участок массива.
Все элементы имеют один тип. Запись `[]T` задаёт тип элементов `T`, а длина и ёмкость относятся
к конкретному значению среза. Поэтому `[3]string` и `[4]string` — разные типы массивов, а срезы
из трёх и четырёх строк могут иметь один тип `[]string`.

Примеры для запуска собраны в [hello.go](hello.go). Команда — `go run hello.go`. Для всего урока
нужен Go 1.22 или новее, чтобы пример с `slices.Delete` обнулял хвост. Сам пакет `slices` появился в Go 1.21.

## Объявление и создание срезов

### Нулевой и пустой срез

```go
var users []string                        // Нулевое значение среза равно nil.
empty := []string{}                       // Пустой литерал даёт ненулевой срез.
named := []string{"Tom", "Alice", "Kate"} // Литерал создаёт массив для трёх строк.

fmt.Println(users == nil, len(users), cap(users)) // true 0 0
fmt.Println(empty == nil, len(empty), cap(empty)) // false 0 0
fmt.Println(named)                                // [Tom Alice Kate]
```

По `nil`-срезу можно пройти циклом `range`. Итераций будет ноль. С ним также работают `append`,
`copy` и выражение `users[:0]`. Чтение `users[0]` вызовет панику: нет элемента с таким индексом.

### `make` задаёт длину и ёмкость

```go
users := make([]string, 3)                 // Уже есть три элемента со значением "".
fmt.Printf("%q\n", users)                  // ["" "" ""]
users[0], users[1], users[2] = "Tom", "Alice", "Bob"
fmt.Println(users, len(users), cap(users)) // [Tom Alice Bob] 3 3

planned := make([]string, 0, 3)            // Длина 0, место под три будущих элемента.
planned = append(planned, "Tom")
fmt.Println(planned, len(planned), cap(planned)) // [Tom] 1 3
// planned[1] = "Alice" // Паника: индекс за пределами длины, хотя место в массиве есть.
```

Если ёмкость пропущена, она равна длине. Должно выполняться `0 <= len <= cap`.
Противоречащие этому константы дают ошибку компиляции, а недопустимые значения переменных —
панику при вызове. Выбирайте положительную длину, когда нужно сразу заполнять элементы
по индексам, и нулевую длину с запасом ёмкости, когда будете добавлять через `append`.

## Выражения среза

### Срез массива, указателя на массив и другого среза

```go
initialUsers := [8]string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}

middle := initialUsers[2:6] // Включает индексы 2, 3, 4 и 5. Индекс 6 исключён.
first := initialUsers[:4]  // Пропущенная левая граница равна 0.
last := initialUsers[3:]   // Пропущенная правая граница равна len(initialUsers).

fmt.Println(middle) // [Kate Sam Tom Paul]
fmt.Println(first)  // [Bob Alice Kate Sam]
fmt.Println(last)   // [Sam Tom Paul Mike Robert]

fmt.Println(middle[:2])     // [Kate Sam]. Индексы отсчитываются от начала middle.
pointer := &initialUsers
fmt.Println(pointer[2:4])   // [Kate Sam]. Сокращение для (*pointer)[2:4].
```

У результата `s[i:j]` длина `j-i`. Для массива и указателя на массив границы
удовлетворяют `0 <= i <= j <= len(s)`, для среза — `0 <= i <= j <= cap(s)`. Массив должен
быть адресуемым. Например, переменную-массив можно нарезать напрямую, а для литерала
можно взять адрес и нарезать через указатель. Границы проверяются при компиляции там,
где это позволяет тип и константы. Остальные нарушения вызывают панику при выполнении.

### Длина ограничивает индекс, ёмкость — расширение

```go
array := [6]int{1, 2, 3, 4, 5, 6}
window := array[1:5] // От индекса 1 до 4, но исходный массив продолжается до индекса 5.
fmt.Println(window, len(window), cap(window)) // [2 3 4 5] 4 5

wider := window[:5]              // Правая граница допустима, потому что cap(window) == 5.
fmt.Println(wider[4], len(window)) // 6 4. Длина window прежняя.
// fmt.Println(window[4]) // Паника: индекс проверяется по len(window), а не по cap(window).
```

`cap` показывает, до какой длины можно расширить срез, сохраняя его начало и текущий массив.
Ёмкость по умолчанию ограничена концом массива или ёмкостью исходного среза.
У пустого среза длина может быть нулевой, а ёмкость — положительной.

### Третий индекс ограничивает ёмкость

```go
array := [6]int{1, 2, 3, 4, 5, 6}
limited := array[1:3:3] // len = 3 - 1, cap = 3 - 1. Элементы по-прежнему общие.
limited[0] = 22
fmt.Println(limited, len(limited), cap(limited)) // [22 3] 2 2
fmt.Println(array)                               // [1 22 3 4 5 6]

// _ = limited[:3] // Паника: новый предел превышает cap(limited).
```

Для `s[low:high:max]` действует `0 <= low <= high <= max <= cap(s)`.
Длина равна `high-low`, ёмкость — `max-low`. Пропустить можно только `low`.
Выражение `s[:len(s):len(s)]` оставляет все текущие элементы и ограничивает ёмкость длиной.
Если после этого добавить хотя бы один элемент, `append` выделит новый массив.

Для строки `text[i:j]` индексы считаются в байтах. Подстрока ограничена длиной строки и не поддерживает
третий индекс.

## Общий массив, присваивание и перебор

### Запись видна через оба представления

```go
array := [6]int{1, 2, 3, 4, 5, 6}
window := array[1:5]
array[1] = 22                       // array[1] — это window[0].
fmt.Println(window)                // [22 3 4 5]
window[1] = 33                      // window[1] — это array[2].
fmt.Println(array)                 // [1 22 33 4 5 6]

alias := window                     // Копируем описание, элементы остаются общими.
alias = alias[:1]                   // Длина меняется только у alias.
fmt.Println(len(window), len(alias)) // 4 1
```

Рабочая модель среза — указатель на участок массива, длина и ёмкость. Спецификация
описывает эти свойства и совместное использование элементов без требования к конкретной
раскладке значения в памяти. В аргумент функции описание передаётся по значению. Функция
может поменять элементы общего массива. Чтобы вызывающий получил новую длину, функция
обычно возвращает получившийся срез.

### Изменение элемента через индекс в `range`

```go
users := []string{"Tom", "Alice", "Kate"}
fmt.Println(users[2]) // Kate
users[2] = "Katherine"

for _, name := range users {
    fmt.Println(name) // Tom, Alice, Katherine на отдельных строках.
}

scores := []int{1, 2, 3}
for index, score := range scores {
    scores[index] = score * 10 // score — копия, а scores[index] — элемент массива.
}
fmt.Println(scores) // [10 20 30]
```

При `range` по срезу срез для обхода определяется перед циклом. Увеличение длины переменной
через `append` внутри цикла не увеличит число итераций. Запись в ещё не прочитанные элементы
общего массива может изменить значения, которые цикл получит дальше.

```go
numbers := []int{1, 2}
for _, value := range numbers {
    numbers = append(numbers, value*10) // Добавляем элементы, но обход по-прежнему из двух шагов.
}
fmt.Println(numbers) // [1 2 10 20]
```

### Срезы срезов

```go
rows := [][]int{{1, 2}, {3, 4}, {5, 6, 7}} // Длины строк могут отличаться.
rows[1][0] = 30                            // Выбираем строку с индексом 1, затем её элемент 0.
fmt.Println(rows)                         // [[1 2] [30 4] [5 6 7]]

grid := make([][]int, 2)
fmt.Println(grid[0] == nil) // true. make создал внешние элементы, а строки ещё nil.
grid[0] = make([]int, 3)
fmt.Println(grid)          // [[0 0 0] []]
```

Каждый элемент внешнего среза — отдельное описание внутреннего среза. Можно выделить массив
для каждой строки или нарезать один общий буфер. Если строки созданы из одного буфера,
учитывайте общую память и ёмкость каждой строки.

## Добавление через `append`

### Сохраните возвращённый срез

```go
users := []string{"Tom", "Alice", "Kate"}
users = append(users, "Bob")
users = append(users, []string{"Sam", "Paul"}...) // ... передаёт элементы второго среза.
fmt.Println(users) // [Tom Alice Kate Bob Sam Paul]

var numbers []int
numbers = append(numbers, 1)
fmt.Println(numbers) // [1]. Нулевой срез подходит для первого append.
```

Если места хватает, `append` использует прежний массив. Если места не хватает, создаёт новый
массив достаточной ёмкости и копирует в него прежние элементы. Формула роста ёмкости
зависит от реализации. В обоих случаях длина результата видна только после присваивания
или передачи результата дальше. Допустимы перекрывающиеся исходный и целевой участки.

### Когда добавление перезаписывает соседний элемент

```go
array := [4]int{10, 20, 30, 40}
prefix := array[:2]                       // len = 2, cap = 4.
extended := append(prefix, 99)            // Места хватает. Пишем в array[2].
fmt.Println(prefix, extended, array)     // [10 20] [10 20 99] [10 20 99 40]

limited := array[:2:2]                    // Те же первые элементы, но cap == len.
detached := append(limited, 77)           // Для добавления нужен новый массив.
detached[0] = 100
fmt.Println(detached, array)              // [100 20 77] [10 20 99 40]
```

Ограничение ёмкости защищает соседнюю область при добавлении, хотя до `append` запись
в видимый элемент всё ещё меняет общий массив. Когда нужна независимая копия уже сейчас,
используйте `copy` или `slices.Clone`.

## Удаление элементов

### Удаление с сохранением порядка

```go
users := []string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
before := users // Второй срез с исходной длиной, но на том же массиве.
index := 3
users = append(users[:index], users[index+1:]...) // Сдвигаем хвост на место Sam.
fmt.Println(users)                              // [Bob Alice Kate Tom Paul Mike Robert]
fmt.Println(before)                             // [Bob Alice Kate Tom Paul Mike Robert Robert]
clear(before[len(users):])                      // Обнуляем освободившиеся элементы.
fmt.Printf("%q\n", before[len(users)])           // ""
```

Операция стоит `O(len(s)-index)`, потому что сдвигает элементы после удаляемого.
Индекс должен быть в пределах длины. У других срезов того же массива длина не меняется,
но сдвиг и очистку они видят. Обнуление особенно важно для указателей и других значений,
удерживающих данные в памяти.

### `slices.Delete` и удаление без сохранения порядка

```go
numbers := []int{10, 20, 30, 40}
original := numbers
numbers = slices.Delete(numbers, 1, 3) // Удаляет индексы 1 и 2. Граница 3 исключена.
fmt.Println(numbers, original)        // [10 40] [10 40 0 0]. Хвост обнулён с Go 1.22.
```

`slices.Delete(s, i, j)` меняет содержимое исходного массива и возвращает укороченный срез.
Если порядок не важен, перенесите последний элемент на место удаляемого. Тогда сдвиг не нужен.

```go
numbers := []int{10, 20, 30, 40}
index := 1
last := len(numbers) - 1
numbers[index] = numbers[last] // На место 20 записываем 40.
numbers[last] = 0              // Для указателей здесь записали бы nil.
numbers = numbers[:last]
fmt.Println(numbers) // [10 40 30]
```

## Копирование и время жизни массива

### `copy(dst, src)` использует длину обоих срезов

```go
source := []int{1, 2, 3, 4, 5, 6}
empty := make([]int, 0, 6)
short := make([]int, 3)
long := make([]int, 8)

fmt.Println(copy(empty, source), empty) // 0 []. Ёмкость назначения не увеличивает число копий.
fmt.Println(copy(short, source), short) // 3 [1 2 3]
fmt.Println(copy(long, source), long)   // 6 [1 2 3 4 5 6 0 0]. Последние элементы прежние.
```

`copy` возвращает `min(len(dst), len(src))` и сохраняет длины обоих срезов.
Перекрытие допустимо. Есть специальная форма `copy(dst, text)`, где для назначения
типа `[]byte` из строки копируются байты.

```go
values := []int{1, 2, 3, 4}
count := copy(values[1:], values[:3]) // Оба аргумента ссылаются на один массив.
fmt.Println(count, values)           // 3 [1 1 2 3]
```

### Независимые элементы и поверхностная копия

```go
original := []int{1, 2, 3}
cloned := slices.Clone(original)
cloned[0] = 99
fmt.Println(original, cloned) // [1 2 3] [99 2 3]

rows := [][]int{{1, 2}}
clonedRows := slices.Clone(rows)
clonedRows[0][0] = 10          // Скопированы только описания строк. Их массивы общие.
fmt.Println(rows, clonedRows) // [[10 2]] [[10 2]]
```

Без `slices.Clone` можно создать `make([]T, len(source))` и вызвать `copy`.
Чтобы клонировать вложенные срезы, нужно копировать каждую строку отдельно. Маленький
подсрез может удерживать весь большой массив. Если нужен только этот участок, скопируйте
его элементы и отпустите ссылки на прежний массив. Укорочение или ограничение ёмкости
само по себе не отсоединяет срез от массива.

## Сортировка и поиск

`slices.Sort` меняет порядок элементов в исходном массиве. Если его нужно сохранить, сортируйте копию.
`slices.Index` ищет линейно, а `slices.BinarySearch` требует уже отсортированный срез.

```go
scores := []int{33, 11, 22}
fmt.Println(slices.Index(scores, 22)) // 2, поиск в исходном порядке.

slices.Sort(scores)
fmt.Println(scores)                        // [11 22 33]
fmt.Println(slices.BinarySearch(scores, 20)) // 1 false, место для вставки без совпадения.
```

`slices.Reverse` разворачивает элементы в текущем порядке. Чтобы получить порядок по убыванию таким способом,
сначала отсортируйте срез.

## Сравнение срезов

```go
first := []int{1, 2, 3, 4}
same := []int{1, 2, 3, 4}
short := []int{1, 2, 3}
var missing []int

// _ = first == same // Ошибка компиляции: срез можно сравнить только с nil.
fmt.Println(slices.Equal(first, same), slices.Equal(first, short)) // true false
fmt.Println(slices.Equal(missing, []int{}))                        // true
fmt.Println(reflect.DeepEqual(missing, []int{}))                    // false
fmt.Println(reflect.DeepEqual(first, []string{"Tom", "Bob"}))        // false
```

`slices.Equal` сравнивает длину и элементы по порядку. Тип элементов должен допускать
оператор `==`. Для сравнения по своему правилу, в том числе для `[][]int`, подходит
`slices.EqualFunc`: в качестве правила для строки можно передать функцию, вызывающую `slices.Equal`.
Если сравниваются элементы интерфейсного типа, одинаковый несравнимый динамический тип
может вызвать панику. Например, это произойдёт при сравнении двух `[]any{[]int{1}}`.

`reflect.DeepEqual` проверяет идентичность типов, различает `nil` и ненулевой пустой срез,
а для срезов одинаковой длины может считать общую начальную позицию достаточной для равенства.
Например, срез с `NaN` может быть глубоко равен самому себе, хотя поэлементное сравнение `NaN`
через `==` даёт `false`.

## Источники

- [Спецификация Go — типы срезов и общий массив](https://go.dev/ref/spec#Slice_types)
- [Спецификация Go — индексы](https://go.dev/ref/spec#Index_expressions)
- [Спецификация Go — выражения среза](https://go.dev/ref/spec#Slice_expressions)
- [Спецификация Go — циклы range](https://go.dev/ref/spec#For_range)
- [Спецификация Go — make](https://go.dev/ref/spec#Making_slices_maps_and_channels)
- [Спецификация Go — append и copy](https://go.dev/ref/spec#Appending_and_copying_slices)
- [Спецификация Go — clear](https://go.dev/ref/spec#Clear)
- [Спецификация Go — сравнение](https://go.dev/ref/spec#Comparison_operators)
- [Блог Go — устройство срезов и удержание массива](https://go.dev/blog/slices-intro)
- [Блог Go — обнуление хвоста и общая память](https://go.dev/blog/generic-slice-functions)
- [Заметки Go 1.21 — стандартный пакет slices](https://go.dev/doc/go1.21#slices)
- [Go 1.22 — изменения стандартной библиотеки](https://go.dev/doc/go1.22#minor_library_changes)
- [Стандартная библиотека — slices.Delete](https://pkg.go.dev/slices#Delete)
- [Стандартная библиотека — slices.Clone](https://pkg.go.dev/slices#Clone)
- [Стандартная библиотека — slices.Sort](https://pkg.go.dev/slices#Sort)
- [Стандартная библиотека — slices.SortStableFunc](https://pkg.go.dev/slices#SortStableFunc)
- [Стандартная библиотека — sort.SearchStrings](https://pkg.go.dev/sort#SearchStrings)
- [Стандартная библиотека — slices.BinarySearch](https://pkg.go.dev/slices#BinarySearch)
- [Стандартная библиотека — slices.Reverse](https://pkg.go.dev/slices#Reverse)
- [Стандартная библиотека — sort.Reverse](https://pkg.go.dev/sort#Reverse)
- [Стандартная библиотека — slices.Equal и EqualFunc](https://pkg.go.dev/slices#Equal)
- [Стандартная библиотека — reflect.DeepEqual](https://pkg.go.dev/reflect#DeepEqual)
