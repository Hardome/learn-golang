package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"image"
	_ "image/gif" // Инициализация пакета регистрирует формат GIF в image.
	"math"
	texttemplate "text/template" // Псевдоним отличает этот пакет от html/template.
)

// Порядок инициализации учитывает ссылку на base в теле функции.
var doubled = doubleBase()
var base = 21

func doubleBase() int {
	return base * 2
}

// init выполняется после инициализации переменных и до main.
func init() {
	fmt.Println(doubled) // 42
	doubled++
}

func main() {
	fmt.Println(doubled) // 43

	// Обычные импорты вводят имена fmt и math в область видимости файла.
	fmt.Println(math.Pow(2, 3)) // 8

	// Оба пакета объявлены как template. В этом файле для одного задан псевдоним.
	fmt.Println(template.New("page").Name(), texttemplate.New("note").Name()) // page note

	showGIFConfig() // 1 1 gif
}

func showGIFConfig() {
	// Однопиксельный GIF хранится в Base64. Имя base64 берётся из объявления пакета.
	const pixelGIF = "R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7"
	data, err := base64.StdEncoding.DecodeString(pixelGIF)
	if err != nil {
		panic(err)
	}

	// Если убрать пустой импорт image/gif, DecodeConfig вернёт image.ErrFormat.
	config, format, err := image.DecodeConfig(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}
	fmt.Println(config.Width, config.Height, format) // 1 1 gif
}
