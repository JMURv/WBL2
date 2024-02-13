package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int) // Создали канал

	go func() { // Анон. горутина для записи в новый канал
		for _, v := range vs {
			c <- v                                                        // Записываем переданные в ф. значения в канал
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) // С задержкой от 1 до 1000 милисек
		}
		close(c) // Закрываем канал
	}()
	return c // Возвращаем канал
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int) // Создали канал
	go func() {
		for { // Ждём значений из 2-х каналов и записываем в 3 канал
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
			// Вариант решения: проверять каналы на закрытость и выходить из цикла после закрытия обоих
		} // Цикл не завершается никогда + читаем из закрытых каналов нулевые значения
	}()
	return c
}

func main() {
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Println(v)
	}
}

// Вывод: 0000000...
