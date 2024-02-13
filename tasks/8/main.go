package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("UNIX> ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка чтения ввода:", err)
			continue
		}

		// Удаляем символ новой строки из ввода
		input = strings.TrimSpace(input)

		// Выходим из цикла, если пользователь ввел \quit
		if input == "\\q" {
			fmt.Println("Выход из программы.")
			os.Exit(0)
		}

		args := strings.Fields(input)
		if len(args) == 0 {
			continue
		}

		switch args[0] {
		case "cd":
			if len(args) < 2 {
				fmt.Println("Не указана директория для смены")
				return
			}
			err := os.Chdir(args[1])
			if err != nil {
				fmt.Println("Ошибка при смене директории:", err)
			}
		case "pwd":
			dir, err := os.Getwd()
			if err != nil {
				fmt.Println("Ошибка при получении текущей директории:", err)
				return
			}
			fmt.Println(dir)
		case "echo":
			fmt.Println(strings.Join(args[1:], " "))
		case "kill":
			if len(args) < 2 {
				fmt.Println("Не указан процесс для завершения")
				return
			}
			pid := args[1]
			err := exec.Command("kill", pid).Run()
			if err != nil {
				fmt.Println("Ошибка при завершении процесса:", err)
			}
		case "ps":
			cmd := exec.Command("ps", "aux")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println("Ошибка при выполнении команды ps:", err)
			}
		default:
			// Попытка выполнить пользовательскую команду
			cmd := exec.Command(args[0], args[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println("Ошибка при выполнении команды:", err)
			}
		}
	}
}
