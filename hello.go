package main

import (
	"fmt"
	"reflect"
	"slices"
	"sort"
)

func main() {
	showCreation()
	showSliceExpressions()
	showSharingAndRange()
	showRows()
	showAppend()
	showDelete()
	showCopy()
	showSortAndSearch()
	showComparison()
}

func showCreation() {
	var missing []string
	empty := []string{}
	fmt.Println(missing == nil, len(missing), cap(missing)) // true 0 0
	fmt.Println(empty == nil, len(empty), cap(empty))       // false 0 0

	users := make([]string, 3)
	fmt.Printf("%q\n", users) // ["" "" ""]
	users[0], users[1], users[2] = "Tom", "Alice", "Bob"
	fmt.Println(users, len(users), cap(users)) // [Tom Alice Bob] 3 3

	planned := make([]string, 0, 3)
	// Ёмкость позволяет добавить элемент, но обращаться по индексу пока нельзя.
	// planned[0] = "Tom" // Паника: длина равна нулю.
	planned = append(planned, "Tom")
	fmt.Println(planned, len(planned), cap(planned)) // [Tom] 1 3
}

func showSliceExpressions() {
	initialUsers := [8]string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	middle := initialUsers[2:6]   // Индекс 6 не входит в результат.
	fmt.Println(middle)           // [Kate Sam Tom Paul]
	fmt.Println(initialUsers[:4]) // [Bob Alice Kate Sam]
	fmt.Println(initialUsers[3:]) // [Sam Tom Paul Mike Robert]
	fmt.Println(middle[:2])       // [Kate Sam]

	pointer := &initialUsers
	fmt.Println(pointer[2:4]) // [Kate Sam]

	array := [4]int{10, 20, 30, 40}
	window := array[:2]
	fmt.Println(window, len(window), cap(window)) // [10 20] 2 4
	wider := window[:4]
	fmt.Println(wider[3], len(window)) // 40 2
	// fmt.Println(window[2]) // Паника. Индекс 2 выходит за текущую длину.
	limited := array[:2:2]
	limited[0] = 11
	fmt.Println(limited, len(limited), cap(limited)) // [11 20] 2 2
	fmt.Println(array)                               // [11 20 30 40]

	message := "Hello World"
	fmt.Println(message[6:]) // World
	word := "ёж"
	fmt.Println(len(word), word[:2]) // 4 ё
	fmt.Printf("% x\n", word[:1])    // d1
}

func showSharingAndRange() {
	array := [4]int{10, 20, 30, 40}
	window := array[:2]
	alias := window
	alias[1] = 22
	fmt.Println(window, array) // [10 22] [10 22 30 40]

	alias = alias[:1]
	fmt.Println(len(window), len(alias)) // 2 1

	users := []string{"Tom", "Alice", "Kate"}
	fmt.Println(users[2]) // Kate
	users[2] = "Katherine"
	for _, name := range users {
		fmt.Println(name) // Tom, Alice, Katherine на отдельных строках
	}

	scores := []int{1, 2, 3}
	for index, score := range scores {
		scores[index] = score * 10 // Получаем копию score, а пишем в элемент массива.
	}
	fmt.Println(scores) // [10 20 30]
}

func showRows() {
	rows := [][]int{{1, 2}, {3, 4}, {5, 6, 7}}
	rows[1][0] = 30
	fmt.Println(rows) // [[1 2] [30 4] [5 6 7]]

	grid := make([][]int, 2)
	fmt.Println(grid[0] == nil) // true
	grid[0] = make([]int, 3)
	fmt.Println(grid) // [[0 0 0] []]
}

func showAppend() {
	users := []string{"Tom", "Alice", "Kate"}
	users = append(users, "Bob")
	fmt.Println(users) // [Tom Alice Kate Bob]

	var missing []int
	missing = append(missing, 1, 2)
	fmt.Println(missing) // [1 2]

	array := [4]int{10, 20, 30, 40}
	prefix := array[:2]
	extended := append(prefix, 99)       // Свободная ёмкость позволяет перезаписать array[2].
	fmt.Println(prefix, extended, array) // [10 20] [10 20 99] [10 20 99 40]

	other := [4]int{10, 20, 30, 40}
	detached := append(other[:2:2], 99) // При len == cap выделяется новый массив.
	detached[0] = 100
	fmt.Println(detached, other) // [100 20 99] [10 20 30 40]
}

func showDelete() {
	users := []string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	original := users
	index := 3
	users = append(users[:index], users[index+1:]...)
	fmt.Println(users)    // [Bob Alice Kate Tom Paul Mike Robert]
	fmt.Println(original) // [Bob Alice Kate Tom Paul Mike Robert Robert]
	clear(original[len(users):])
	fmt.Printf("%q\n", original[len(users)]) // ""

	values := []int{10, 20, 30, 40}
	shortened := slices.Delete(values, 1, 3)
	// Начиная с Go 1.22 хвост массива обнуляется.
	fmt.Println(shortened, values) // [10 40] [10 40 0 0]
}

func showCopy() {
	source := []int{1, 2, 3, 4, 5, 6}
	empty := make([]int, 0, 6)
	short := make([]int, 3)
	long := make([]int, 8)
	fmt.Println(copy(empty, source), empty) // 0 []
	fmt.Println(copy(short, source), short) // 3 [1 2 3]
	fmt.Println(copy(long, source), long)   // 6 [1 2 3 4 5 6 0 0]

	overlapping := []int{1, 2, 3, 4}
	copied := copy(overlapping[1:], overlapping[:3])
	fmt.Println(copied, overlapping) // 3 [1 1 2 3]

	cloned := slices.Clone(source)
	cloned[0] = 99
	fmt.Println(source[0], cloned[0]) // 1 99

	rows := [][]int{{1, 2}}
	clonedRows := slices.Clone(rows)
	clonedRows[0][0] = 10         // Описания строк скопированы, но массивы строк общие.
	fmt.Println(rows, clonedRows) // [[10 2]] [[10 2]]
}

func showSortAndSearch() {
	users := []string{"Tom", "Bob", "Sam"}
	fmt.Println(slices.Index(users, "Sam")) // 2
	sort.Strings(users)
	fmt.Println(users)                                                                // [Bob Sam Tom]
	fmt.Println(sort.SearchStrings(users, "Sam"), sort.SearchStrings(users, "Alice")) // 1 0

	values := []int{11, 22, 33, 44, 55, 66}
	fmt.Println(sort.SearchInts(values, 33), sort.SearchInts(values, 45)) // 2 4
	fmt.Println(slices.BinarySearch(values, 45))                          // 4 false

	numbers := []int{3, 1, 2}
	slices.Reverse(numbers)
	fmt.Println(numbers) // [2 1 3]
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	fmt.Println(numbers) // [3 2 1]
}

func showComparison() {
	first := []int{1, 2, 3, 4}
	same := []int{1, 2, 3, 4}
	short := []int{1, 2, 3}
	// _ = first == same // Ошибка компиляции: два среза нельзя сравнить через ==.
	fmt.Println(slices.Equal(first, same), slices.Equal(first, short)) // true false

	var missing []int
	fmt.Println(slices.Equal(missing, []int{}))                   // true
	fmt.Println(reflect.DeepEqual(missing, []int{}))              // false
	fmt.Println(reflect.DeepEqual(first, []string{"Tom", "Bob"})) // false
}
