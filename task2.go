package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("log.txt")
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	buffer := make([]byte, 1024)

	for {
		n, err := file.Read(buffer)
		if err != nil && err.Error() != "EOF" {
			fmt.Println("Ошибка чтения файла:", err)
			return
		}

		if n == 0 {
			break
		}

		fmt.Print(string(buffer[:n]))
	}
}
