package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"howett.net/plist"
)

func kek(ipaPath string) {
	// Открываем .ipa (как zip-файл)
	zipReader, err := zip.OpenReader(ipaPath)
	if err != nil {
		log.Fatalf("Ошибка при открытии .ipa: %v", err)
	}
	defer zipReader.Close()

	var plistData []byte

	// Ищем Info.plist в распакованном архиве
	for _, file := range zipReader.File {
		if bytes.Contains([]byte(file.Name), []byte("Info.plist")) {
			// Открываем файл внутри архива
			f, err := file.Open()
			if err != nil {
				log.Fatalf("Ошибка при открытии Info.plist: %v", err)
			}
			defer f.Close()

			// Читаем данные из Info.plist
			plistData, err = io.ReadAll(f)
			if err != nil {
				log.Fatalf("Ошибка при чтении Info.plist: %v", err)
			}
			break
		}
	}

	if plistData == nil {
		log.Fatalln("Info.plist не найден")
	}

	// Парсим данные Info.plist с использованием go-plist
	var parsedPlist map[string]any
	decoder := plist.NewDecoder(bytes.NewReader(plistData))
	err = decoder.Decode(&parsedPlist)
	if err != nil {
		log.Fatalf("Ошибка при разборе Info.plist: %v", err)
	}

	fmt.Println(parsedPlist["Cndljhhhjjhhjhje"])

	// Выводим распарсенные данные
	fmt.Println("Распарсенные данные Info.plist:")
	for key, value := range parsedPlist {
		fmt.Printf("%s: %v\n", key, value)
	}

}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Использование: go run main.go <путь к .ipa> <папка для сохранения>")
		return
	}

	ipaPath := os.Args[1]
	// outputDir := os.Args[2]

	kek(ipaPath)

	// err := extractIPAIcon(ipaPath, outputDir)
	// if err != nil {
	// 	fmt.Println("Ошибка:", err)
	// }
}
