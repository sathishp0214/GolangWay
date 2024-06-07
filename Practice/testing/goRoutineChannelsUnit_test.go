package main

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func Process(data int, resultChan chan<- int) {
	go func() {
		// Simulate some processing
		time.Sleep(1 * time.Second)
		resultChan <- data * 2
	}()
}

func TestProcess(t *testing.T) {
	resultChan := make(chan int, 1)
	Process(5, resultChan)

	select {
	case result := <-resultChan:
		if result != 10 {
			t.Errorf("Expected 10, but got %d", result)
		}
	case <-time.After(2 * time.Second):
		t.Error("Test timed out")
	}
}

func ProcessWithContext(ctx context.Context, data int, resultChan chan<- int) { //Instead of creating channel inside function, Having channel in function parameter makes unit testing possible.
	go func() {
		select {
		case <-time.After(1 * time.Second):
			resultChan <- data * 2
		case <-ctx.Done():
			fmt.Println("Operation cancelled")
			return
		}
	}()
}

func TestProcessWithContext(t *testing.T) {
	resultChan := make(chan int, 1)
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	ProcessWithContext(ctx, 5, resultChan)

	select {
	case <-ctx.Done():
		t.Error("Operation timed out before completion")
	case result := <-resultChan:
		if result != 10 {
			t.Errorf("Expected 10, but got %d", result)
		}
	}
}

func IncrementCounter(wg *sync.WaitGroup, counter *int) {
	defer wg.Done()
	*counter++
}

func TestIncrementCounter(t *testing.T) {
	var counter int
	var wg sync.WaitGroup

	wg.Add(1)
	go IncrementCounter(&wg, &counter)

	wg.Wait()

	if counter != 1 {
		t.Errorf("Expected counter to be 1, but got %d", counter)
	}
}
