
package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	file, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	fmt.Println("Введите строки для записи в файл (для завершения введите 'exit'):")
	scanner := bufio.NewScanner(os.Stdin)
	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text()
		if line == "exit" {
			break
		}

		currentTime := time.Now().Format("2006-01-02 15:04:05")

		_, err := fmt.Fprintf(file, "%d. %s %s\n", lineNumber, currentTime, line)
		if err != nil {
			fmt.Println("Ошибка записи в файл:", err)
			return
		}

		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения ввода:", err)
	}
	fmt.Println("Программа завершила работу.")
}
