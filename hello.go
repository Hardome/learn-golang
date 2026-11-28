package main

import "fmt"

func main() {
	next := newCounter(5)
	fmt.Println(next(), next(), next()) // 6 7 8

	first := newCounter(0)
	sameCounter := first
	second := newCounter(10)
	// Копия функции использует тот же start. У второго счётчика свой start.
	fmt.Println(first(), sameCounter(), second()) // 1 2 11

	showMultiplier()
	showCapture()
	showLoopCapture()
	showDeferredValues()    // конец, затем начало (две строки)
	fmt.Println(adjusted()) // 5
	showGoroutineArgument()
}

func showMultiplier() {
	// Первый вызов задаёт множитель. Последующие передают числа для умножения.
	timesFive := multiply(5)
	fmt.Println(timesFive(6))   // 30
	fmt.Println(timesFive(5))   // 25
	fmt.Println(multiply(5)(6)) // 30
}

func multiply(factor int) func(int) int {
	// Вложенная функция использует параметр внешней функции.
	return func(value int) int {
		return factor * value
	}
}

func showCapture() {
	label := "до"
	readLabel := func() string { return label }
	// Замыкание читает ту же переменную после присваивания.
	label = "после"
	fmt.Println(readLabel()) // после

	// Два разных литерала захватывают одну переменную sum.
	add, total := newAccumulator()
	fmt.Println(total()) // 0
	add(5)
	fmt.Println(total()) // 5
}

func newCounter(start int) func() int {
	// Параметр start остаётся доступным после возврата из newCounter.
	return func() int {
		start++
		return start
	}
}

func newAccumulator() (add func(int), total func() int) {
	sum := 0
	return func(value int) { sum += value }, func() int { return sum }
}

func showLoopCapture() {
	values := []int{1, 2, 3}

	// Начиная с Go 1.22, := объявляет свою переменную на каждой итерации.
	var perIteration []func() int
	for _, value := range values {
		perIteration = append(perIteration, func() int { return value })
	}
	fmt.Println(perIteration[0](), perIteration[1](), perIteration[2]()) // 1 2 3

	// При = все итерации присваивают значения одной внешней переменной.
	var shared []func() int
	var snapshots []func() int
	var value int
	for _, value = range values {
		shared = append(shared, func() int { return value })

		// Копия в теле даёт отдельную переменную на каждой итерации.
		iterationValue := value
		snapshots = append(snapshots, func() int { return iterationValue })
	}
	fmt.Println(shared[0](), shared[1](), shared[2]())          // 3 3 3
	fmt.Println(snapshots[0](), snapshots[1](), snapshots[2]()) // 1 2 3
}

func showDeferredValues() {
	state := "начало"
	// Аргумент вычисляется сразу. Отложенные вызовы выполняются в обратном порядке.
	defer func(saved string) { fmt.Println(saved) }(state) // начало (выводится вторым)
	defer func() { fmt.Println(state) }()                  // конец (выводится первым)
	state = "конец"
}

func adjusted() (result int) {
	// return сначала присваивает результату 4, затем замыкание увеличивает его.
	defer func() { result++ }()
	return 4
}

func showGoroutineArgument() {
	result := make(chan int)
	value := 7
	// Параметр получает значение до запуска горутины. Канал позволяет дождаться ответа.
	go func(input int) { result <- input * 2 }(value)
	value = 9
	fmt.Println(<-result) // 14
}
