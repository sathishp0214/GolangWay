package main

import "fmt"

func main() {
	ch := make(chan int)
	ch1 := make(chan int)
	ch3 := make(chan int)

	go func() {
		defer close(ch)
		for i := 0; i <= 10; i++ {
			ch <- i
		}
	}()

	go func() {
		defer close(ch1)
		for i := 11; i <= 20; i++ {
			ch1 <- i
		}
	}()

	go func() {
		// go func() {
		// 	for i := range ch {
		// 		fmt.Println("Listening from first Goroutine", i)
		// 		ch3 <- i
		// 	}
		// }()

		// for i := range ch1 {
		// 	fmt.Println("Listening from Second Goroutine", i)
		// 	ch3 <- i
		// }
		flag := false
		flag1 := false

		for {
			if flag && flag1 {
				return
			}
			select {
			case v1, ok := <-ch:
				if !ok {
					flag = true
				}
				ch3 <- v1
			case v2, ok := <-ch1:
				if !ok {
					flag1 = true
				}
				ch3 <- v2

			default:

			}
		}

	}()

	for {
		select {
		case val := <-ch3:
			fmt.Println("From main function", val)
		default:
		}
	}
}
