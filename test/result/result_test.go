package test

import (
	"skillbox-diploma/internal/result"
	"skillbox-diploma/internal/simulator"
	"sync"
	"testing"
	"time"
)

func TestResultMultithreading(t *testing.T) {
	go simulator.StartSimulatorServer()
	time.Sleep(5 * time.Second)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			_ = result.GetResultData()
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			_ = result.GetResultData()
		}
	}()
	wg.Wait()
}
