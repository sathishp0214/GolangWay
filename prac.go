package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// func main() {
// 	// start := time.Now()

// 	// wg.Add(4)

// 	// firstchan := make(chan string)
// 	// secondchan := make(chan string)
// 	// thirdchan := make(chan string)
// 	// fourthchan := make(chan string)

// 	// go FourthFunc(fourthchan)
// 	// go SecondFunc(secondchan)
// 	// go FirstFunc(firstchan)
// 	// go ThirdFunc(thirdchan)

// 	// fmt.Println(<-firstchan)
// 	// fmt.Println(<-secondchan)
// 	// fmt.Println(<-thirdchan)
// 	// fmt.Println(<-fourthchan)

// 	// wg.Wait()

// 	// fmt.Printf("Total time to finish : %s \n", time.Since(start).String())

// 	f := func() {
// 		n := 0
// 		fmt.Println("hello", n)
// 		n++
// 	}

// 	fmt.Println(f)
// 	f()
// 	f()
// 	f()

// }

type DivisionError struct {
	IntA int
	IntB int
	Msg  string
}

func (e *DivisionError) Error() string {
	return e.Msg
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, &DivisionError{
			Msg:  fmt.Sprintf("cannot divide '%d' by zero", a),
			IntA: a, IntB: b,
		}
	}
	return a / b, nil
}

func main() {
	a, b := 10, 0
	result, err := Divide(a, b)
	if err != nil {
		var divErr *DivisionError
		switch {
		case errors.As(err, &divErr):
			fmt.Printf("%d / %d is not mathematically valid: %s\n",
				divErr.IntA, divErr.IntB, divErr.Error())
		default:
			fmt.Printf("unexpected division error: %s\n", err)
		}
		return
	}

	fmt.Printf("%d / %d = %d\n", a, b, result)
}

func FirstFunc(ch chan<- string) {
	fmt.Println("-- Executing first function --")
	time.Sleep(7 * time.Second)
	defer wg.Done()

	ch <- "-- First Function finished --"
}

func SecondFunc(ch chan<- string) {
	fmt.Println("-- Executing second function --")
	// time.Sleep(5 * time.Second)
	defer wg.Done()

	ch <- "-- Second Function finished --"
}

func ThirdFunc(ch chan<- string) {
	fmt.Println("-- Executing third function --")
	time.Sleep(25 * time.Second)
	defer wg.Done()

	ch <- "-- Third Function finished --"
}

func FourthFunc(ch chan<- string) {
	fmt.Println("-- Executing fourth function --")
	// time.Sleep(10 * time.Second)
	defer wg.Done()

	ch <- "-- Fourth Function finished --"
}
