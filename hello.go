package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"image"
	_ "image/gif" // Подключает поддержку GIF для пакета image.
	"math"
	texttemplate "text/template" // Даём второму пакету template своё имя.
)

func sayHello(name string) {
	fmt.Printf("Привет, %s!\n", name) // Подставляет имя в приветствие.
}

func main() {
	sayHello("Go")              // Привет, Go!
	fmt.Println(math.Pow(2, 3)) // 8

	// Оба пакета называются template. Вызовы различаются по имени перед точкой.
	fmt.Println(template.New("page").Name())     // page
	fmt.Println(texttemplate.New("note").Name()) // note

	showGIFConfig() // 1 1 gif
}

func showGIFConfig() {
	// Строка хранит GIF размером в один пиксель. DecodeString восстанавливает его байты.
	const pixelGIF = "R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7"
	data, err := base64.StdEncoding.DecodeString(pixelGIF)
	if err != nil {
		panic(err)
	}

	// Если убрать импорт image/gif, image не распознает формат и вернёт image.ErrFormat.
	config, format, err := image.DecodeConfig(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}
	fmt.Println(config.Width, config.Height, format) // 1 1 gif
}
