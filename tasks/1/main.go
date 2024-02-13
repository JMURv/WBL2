package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	ntpNow, err := ntp.Time("pool.ntp.org")
	if err != nil {
		log.Printf("Ошибка получения времени: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Текущее локальное время:", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("Точное время по NTP:", ntpNow.Format("2006-01-02 15:04:05"))
}
