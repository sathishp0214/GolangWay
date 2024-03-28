package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

var p = fmt.Println

func main() {
	// ch := make(chan int, 2)
	// // go func() {
	// // 	ch <- 10
	// // }()'
	// ch <- 10
	// ch <- 10
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println("completed")

	// SendingReadingbufferedChannelOneafterOne()

	// dummy()

	// GoRoutinesExecutionOrderWithMutexLockUnlock()

	// input := "abcccccddefggg"
	// i := 0
	// for i = 0; i < len(input)-1; i++ {
	// 	if input[i] == input[i+1] {
	// 		continue
	// 	}
	// 	fmt.Print(string(input[i]))
	// }
	// fmt.Print(string(input[i]))

	// fmt.Println(4 % 7)

	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	// output := input[0]
	// for i:=1;i<len(input);i++ {
	// 	for j:=
	// }

	fmt.Println(maxSubArray(input))

	df := []string{"dd", "a", "aaaa", "cccccccccc"}
	sort.SliceStable(df, func(i, j int) bool {
		return len(df[i]) > len(df[j])
	})
	fmt.Println(df)

	df = []string{"123", "90", "5000", "4"}
	sort.SliceStable(df, func(i, j int) bool {
		return df[i] > df[j]
	})
	fmt.Println(df)

	//longest increasing subarray values
	a := []int{10, 20, 30, 0, -1, 50, 55, 60, 1}
	output := [][]int{}
	tmp := []int{}
	for i := 0; i < len(a)-1; i++ {

		if a[i] < a[i+1] {
			tmp = append(tmp, a[i])
		} else {
			if len(tmp) > 0 {
				tmp = append(tmp, a[i])
				output = append(output, tmp)
				tmp = []int{}
			}

		}
	}
	if len(tmp) > 0 {
		output = append(output, tmp)
	}
	fmt.Println(output)

	sort.SliceStable(output, func(i, j int) bool {
		return len(output[i]) > len(output[j])
	})
	fmt.Println(output, "==========", output[0])

}

func maxSubArray(nums []int) int {
	lenNums := len(nums)

	currentMax := 0
	max := -1000
	for i := 0; i < lenNums; i++ {
		currentMax = currentMax + nums[i]
		if currentMax > max {
			max = currentMax
		}

		if currentMax < 0 {
			currentMax = 0
		}

	}
	return max
}

func dummy() {
	ch := make(chan int, 2)
	go func() {
		ch <- 10
		ch <- 100
		ch <- 101
	}()
	// ch <- 10
	// ch <- 10
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Println("completed")
	for i := 0; i < 10; i++ {
		_ = 100
		time.Sleep(1 * time.Second)
		fmt.Println("ppppppppppp")
	}
	fmt.Println(<-ch)

}

func SendingReadingbufferedChannelOneafterOne() {

	wg := sync.WaitGroup{}
	length := 15
	ch := make(chan int, length)

	wg.Add(2)

	go func() {
		defer wg.Done()
		defer close(ch) //Closing the channel is crucial here, After we sent all the data
		for i := 1; i <= length+50; i++ {
			fmt.Println("sending value into channel", i)
			ch <- i
		}
	}()
	// wg.Wait()

	// wg.Add(1)
	go func() {
		defer wg.Done()
		for value := range ch {
			fmt.Println("printing square", value*value)
		}

		//can use the below method as more reliable
		// for {
		// 	select {
		// 	case value, ok := <-ch:
		// 		if !ok {
		// 			return
		// 		}
		// 		fmt.Println("printing square", value*value)
		// 	default:

		// 	}
		// }

	}()
	wg.Wait()
}
