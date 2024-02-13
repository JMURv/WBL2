package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

// go run main.go -host=google.com -port=80 --timeout=10s

func main() {
	// Определение флагов командной строки
	host := flag.String("host", "", "Хост для подключения")
	port := flag.String("port", "", "Порт для подключения")
	timeout := flag.Duration("timeout", 10*time.Second, "Таймаут подключения")
	flag.Parse()

	if *host == "" || *port == "" {
		fmt.Println("Необходимо указать хост и порт")
		return
	}

	conn, err := net.DialTimeout("tcp", *host+":"+*port, *timeout)
	if err != nil {
		fmt.Println("Ошибка при подключении к серверу:", err)
		return
	}
	defer conn.Close()
	fmt.Println("Connected to ", *host)

	done := make(chan struct{})
	defer close(done)

	// Горутина для копирования данных из сокета и вывода в STDOUT
	go func() {
		io.Copy(os.Stdout, conn)
		done <- struct{}{}
	}()

	// Горутина для чтения данных из STDIN и записи в сокет
	go func() {
		io.Copy(conn, os.Stdin)
		done <- struct{}{}
	}()

	select {
	case <-done:
	case <-time.After(*timeout):
		fmt.Println("Таймаут при ожидании завершения работы программы")
	}
}
