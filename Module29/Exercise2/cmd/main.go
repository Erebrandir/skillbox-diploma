/*
Задание 2. Graceful shutdown
Что нужно сделать
В работе часто возникает потребность правильно останавливать приложения. Например, когда наш сервер обслуживает
соединения, а нам хочется, чтобы все текущие соединения были обработаны и лишь потом произошло выключение сервиса.
Для этого существует паттерн graceful shutdown.
Напишите приложение, которое выводит квадраты натуральных чисел на экран, а после получения сигнала ^С обрабатывает
этот сигнал, пишет «выхожу из программы» и выходит.

Советы и рекомендации
Для реализации данного паттерна воспользуйтесь каналами и оператором select с default-кейсом.
*/

package main

import (
	"context"
	"errors"
	"fmt"
	"math"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func initGracefulShutdown(cancelFunc context.CancelFunc, wg *sync.WaitGroup) {
	defer wg.Done()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGILL)
	<-sigs
	fmt.Println("Graceful Shutdown")
	cancelFunc()
}

func action(cancelCtx context.Context) error {
	time.Sleep(100 * time.Millisecond)
	return errors.New("failed")
}

func main() {
	cancelCtx, cancelFunc := context.WithCancel(context.Background())

	go func() {
		err := action(cancelCtx)

		if err != nil {
			cancelFunc()
		}
	}()

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	shutdownChan := make(chan string)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func(shutdown chan string, wg *sync.WaitGroup) {
		i := 1

		for {
			time.Sleep(time.Second)
			i++
			select {
			case <-c:
				fmt.Println("Выхожу из программы... ")
				wg.Done()
				return
			default:
				fmt.Println(math.Pow(float64(i), 2))
			}
		}
	}(shutdownChan, wg)

	go initGracefulShutdown(cancelFunc, wg)
	wg.Wait()
}
