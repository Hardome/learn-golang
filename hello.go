package main

import "fmt"

// Все пять элементов получают нулевое значение типа int.
var numbers [5]int
var values = [5]int{1, 2, 3, 4, 5}

func main() {
	letters := [5]string{"a", "b", "c", "d", "e"}
	inferredLength := [...]int{1} // Go выводит тип [1]int.
	partial := [5]int{1}          // Остальные четыре элемента равны 0.
	sparse := [...]int{3: 7}      // Последний индекс 3, поэтому длина равна 4.

	fmt.Println("Нулевые значения:", numbers)                             // Нулевые значения: [0 0 0 0 0]
	fmt.Println("Заданные значения:", values)                             // Заданные значения: [1 2 3 4 5]
	fmt.Println("Строки:", letters)                                       // Строки: [a b c d e]
	fmt.Printf("Выведенный тип: %T %v\n", inferredLength, inferredLength) // Выведенный тип: [1]int [1]
	fmt.Println("Частичное заполнение:", partial)                         // Частичное заполнение: [1 0 0 0 0]
	fmt.Println("Элемент по индексу 3:", sparse)                          // Элемент по индексу 3: [0 0 0 7]

	// Индексы начинаются с нуля. Последний индекс равен len(values) - 1.
	fmt.Println("Длины:", len(letters), len(inferredLength))             // Длины: 5 1
	fmt.Println("Первый и последний:", values[0], values[len(values)-1]) // Первый и последний: 1 5

	// Массивы одного типа сравниваются по элементам на одинаковых позициях.
	nums1 := [4]int{3, 4, 5, 6}
	nums2 := [4]int{3, 4, 5}
	nums3 := [4]int{3, 4, 5, 0}
	fmt.Println("nums1 == nums2:", nums1 == nums2) // nums1 == nums2: false
	fmt.Println("nums2 == nums3:", nums2 == nums3) // nums2 == nums3: true

	// Присваивание копирует массив целиком. Изменяем первый элемент копии.
	copied := values
	copied[0] = 99
	fmt.Println("Оригинал:", values, "Копия:", copied) // Оригинал: [1 2 3 4 5] Копия: [99 2 3 4 5]

	// range возвращает индекс и значение элемента. Здесь индекс пропускаем через _.
	total := 0
	for _, value := range values {
		total += value
	}
	fmt.Println("Сумма элементов:", total) // Сумма элементов: 15

	// Перебор значения массива читает элементы снимка, сделанного до цикла.
	sample := [3]int{1, 2, 3}
	snapshotSum := 0
	for index, value := range sample {
		if index == 0 {
			sample[1] = 20
		}
		snapshotSum += value
	}
	fmt.Println("Сумма снимка и массив:", snapshotSum, sample) // Сумма снимка и массив: 6 [1 20 3]
}
