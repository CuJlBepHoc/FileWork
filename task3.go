package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.WriteFile("read-only.txt", []byte("Это тестовая строка"), 0444)
	if err != nil {
		fmt.Println("Ошибка при создании файла и записи данных:", err)
		return
	}

	file, err := os.OpenFile("read-only.txt", os.O_RDONLY, 0444)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("Новая строка")
	if err != nil {
		fmt.Println("Ошибка при записи данных:", err)
	} else {
		fmt.Println("Данные успешно записаны")
	}
}
