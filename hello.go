package main

import (
	"fmt"
	"maps"
	"math"
	"reflect"
	"slices"
	"sync"
)

func main() {
	demonstrateInitialization()
	demonstrateOperations()
	demonstrateKeys()
	demonstrateIteration()
	demonstrateCopying()
	demonstrateComparison()
	demonstrateCollections()
	demonstrateConcurrentUpdates()
}

func demonstrateInitialization() {
	var zero map[string]int
	empty := map[string]int{}
	planned := make(map[string]int, 64)                  // Подсказка о размере, записей пока нет.
	fmt.Println(zero, empty, planned)                    // map[] map[] map[]
	fmt.Println(zero == nil, empty == nil, len(planned)) // true false 0

	names := map[string]int{"Tom": 1, "tom": 2, "TOM": 3, "Sam": 4}
	fmt.Println(names) // map[Sam:4 TOM:3 Tom:1 tom:2]
}

func demonstrateOperations() {
	var scores map[string]int
	score, present := scores["Tom"]
	delete(scores, "Tom") // Безопасно для nil-карты.
	clear(scores)
	fmt.Println(score, present) // 0 false
	// scores["Tom"] = 1 // Паника: запись в nil-карту.

	scores = map[string]int{"Tom": 0}
	score, present = scores["Tom"]
	fmt.Println(score, present) // 0 true
	score, present = scores["Kate"]
	fmt.Println(score, present) // 0 false

	scores["Tom"] = 1
	scores["Tom"]++   // Заменяет значение существующей записи.
	scores["Bob"] = 2 // Добавляет запись.
	scores["Alice"]++ // Начинает с нулевого значения и записывает 1.
	delete(scores, "Bob")
	delete(scores, "Kate") // Для отсутствующего ключа ничего не меняется.
	fmt.Println(scores)    // map[Alice:1 Tom:2]
	clear(scores)
	fmt.Println(len(scores), scores == nil) // 0 false
}

func demonstrateKeys() {
	type route struct {
		method string
		path   string
	}
	hits := map[route]int{{method: "GET", path: "/"}: 1}
	fmt.Println(hits[route{method: "GET", path: "/"}])  // 1
	fmt.Println(hits[route{method: "POST", path: "/"}]) // 0
	// _ = map[[]byte]int{} // Ошибка компиляции: срезы несравнимы.

	labels := map[any]string{int(1): "int", int64(1): "int64", nil: "nil"}
	fmt.Println(labels[int(1)], labels[int64(1)], labels[nil]) // int int64 nil
	// labels[[]int{1}] = "slice" // Паника: динамический тип []int несравним.

	nan := math.NaN()
	values := map[float64]string{nan: "первый"}
	values[nan] = "второй" // NaN не равен самому себе, появляется ещё одна запись.
	_, present := values[nan]
	delete(values, nan)               // Ни один ключ не совпадает.
	fmt.Println(len(values), present) // 2 false
	clear(values)
	fmt.Println(len(values)) // 0
}

func demonstrateIteration() {
	scores := map[string]int{"Tom": 1, "Bob": 2, "Alice": 3}
	total := 0
	for name, score := range scores {
		total += score // Сумма одинакова при любом порядке обхода.
		if score < 2 {
			delete(scores, name) // Удаляет только текущую запись Tom.
		}
	}
	fmt.Println(total) // 6
	for _, name := range slices.Sorted(maps.Keys(scores)) {
		fmt.Println(name, scores[name]) // Сначала Alice 3, затем Bob 2.
	}
}

func demonstrateCopying() {
	scores := map[string]int{"Tom": 1}
	shared := scores
	shared["Tom"] = 2          // Обе переменные ссылаются на одну карту.
	fmt.Println(scores["Tom"]) // 2
	shared = map[string]int{"Tom": 9}
	fmt.Println(scores["Tom"], shared["Tom"]) // 2 9

	groups := map[string][]string{"dev": {"Tom"}}
	cloned := maps.Clone(groups)
	cloned["dev"][0] = "Bob"                             // Изменение видно через оба среза.
	cloned["qa"] = []string{"Alice"}                     // Новый ключ появляется только в копии.
	fmt.Println(groups["dev"], len(groups), len(cloned)) // [Bob] 1 2

	type player struct{ score int }
	players := map[string]player{"Tom": {score: 1}}
	// players["Tom"].score++ // Ошибка компиляции. Элемент карты неадресуем.
	current := players["Tom"]
	current.score++
	players["Tom"] = current
	fmt.Println(players["Tom"].score) // 2
}

func demonstrateComparison() {
	first := map[string]int{"Tom": 1}
	second := map[string]int{"Tom": 1}
	changed := map[string]int{"Tom": 2}
	fmt.Println(maps.Equal(first, second), maps.Equal(first, changed)) // true false

	var zero map[string]int
	empty := map[string]int{}
	fmt.Println(maps.Equal(zero, empty), reflect.DeepEqual(zero, empty)) // true false

	reviewed := map[string][]int{"Tom": {1, 2}}
	submitted := map[string][]int{"Tom": {1, 2}}
	fmt.Println(maps.EqualFunc(reviewed, submitted, slices.Equal[[]int, int])) // true
	fmt.Println(reflect.DeepEqual(reviewed, submitted))                        // true

	withNaN := map[float64]int{math.NaN(): 1}
	// maps.Equal ищет ключ через индекс, DeepEqual распознаёт общую таблицу.
	fmt.Println(maps.Equal(withNaN, withNaN), reflect.DeepEqual(withNaN, withNaN)) // false true
}

func demonstrateCollections() {
	seen := map[string]struct{}{"Tom": {}}
	_, present := seen["Tom"]
	fmt.Println(present) // true

	groups := make(map[string][]string)
	groups["dev"] = append(groups["dev"], "Tom") // append принимает nil-срез.
	fmt.Println(groups["dev"])                   // [Tom]

	scores := make(map[string]map[string]int)
	if scores["Go"] == nil {
		scores["Go"] = make(map[string]int) // Внутренняя карта нужна до записи.
	}
	scores["Go"]["Tom"]++
	fmt.Println(scores["Go"]["Tom"]) // 1
}

func demonstrateConcurrentUpdates() {
	counts := make(map[string]int)
	var mutex sync.Mutex
	var workers sync.WaitGroup

	for range 2 {
		workers.Add(1)
		go func() {
			defer workers.Done()
			mutex.Lock()
			counts["Go"]++ // Один замок защищает чтение и запись при инкременте.
			mutex.Unlock()
		}()
	}

	workers.Wait()            // После завершения горутин писателей больше нет.
	fmt.Println(counts["Go"]) // 2
}
