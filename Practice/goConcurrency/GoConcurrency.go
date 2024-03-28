package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var p = fmt.Println

var globalFlagValue bool

var (
	globalWaitGroup = sync.WaitGroup{}
)

/*

Python threading vs GoRoutine:
GoRoutine has very light weight GoROutines compared to Python threads.
Python thread is preferred for I/O bound tasks only, Where python multiprocessing is preferred for CPU-bound tasks, Where GoROutine is fine for both I/O bound and CPU bound tasks, So Routines don't need multiprocessing seperately.


Parallelism in general(multiple things are executing actually at same time) - Multiple core processors running at same time. Each core processor can run multiple threads.
Concurrency in general(multiple things are managing at same time) - Multiple threads/Go Routines run in one single/multi core processor.

Go runtime can create 100s or 1000s logical GoRoutines in one single system thread. My personal laptop 1215U has 8 system threads. We don't know, In too depth internal in GO Runtime of how multiple goroutines are running at same time? always running GoROutines? How switches back between goRoutines?,

Go runtime can have access to run all threads in system, So that automtically have access to all cores Since each core has two threads. we can't say Go routine whether uses concurrency or both concurrency and parallism.

single goRoutine take 8kb stack size defaultly in Go Runtime.
*/

/*channels vs WaitGroup use cases:
WaitGroup alone-

In parent GoROutine, Waits expected single or multiple child GoRoutines to complete.
Can be used when multiple GoRoutines are running, Without need to passing any data to other goRoutines.

Channel alone -

In parent GoROutine, Waits expected child GoRoutines to complete by sending dummy values from child GoRoutine to Parent GoRoutine. (But ensure you use proper for loops with range or for loop with select cases)
Can be used to pass the data between multiple goRoutines.


We can use both channels and waitgroup together for scenarios and also for more reliability in production
*/

/*
Buffered Channel vs UnBuffered channel use case:

UnBuffered channel: (synchronous communication)
Without any capacity in a channel.
Both send and receive channel, Must be available at the time for data passing otherwise deadlock panic occurs. So it blocks for the next value processing until the current value is processed in both send/receive ways channel values.

Unbuffered channel use case: Should use when there is direct exchange of data EX: Checking authentication

Buffered channel: (Asynchronous communication)
Have a buffer capacity in a channel.

Within the capacity in a buffered channel - Stores the data until the buffer capacity, Once buffer capacity reached, Then buffered channel sends data to the channel's reciever side.

Once capacity is exceed in a buffered channel - Then it acts and works like unbuffered channel EX: send/receive at same time and blocking as mentioned above.

Buffered channel use case: Doing things one after one EX: writing and reading one after one. Overall its used to improve the performance among the goRoutines and also reduces the chances of deadlock errors due to the Asynchronous communication.

EX: A goRoutine received any data from Unbuffered channel, Then that data needs to do high multi level processing flow. So once the current data is proceesed then only next data can be readed from the unbuffered channel due to its blocking synchronous communication. But if used buffered channel, GoRoutine keep on reading the next data and added into queue upto the given buffer channel length. So overall improves the performance.

So we can buffered channel as well for all the below examples in this file.
*/

/*
Closing a channel:
Channel should be closed only in the value sending side.

If you are sure, All the expected number of values is sent. Then We can close the channel in that sending function itself.

If we sending data on closed channel, Will get a panic.

Refer this function -- ChannelWithClose()
https://go.dev/tour/concurrency/4#:~:text=Note%3A%20Only%20the%20sender%20should,to%20terminate%20a%20range%20loop.



ii)If One Sender channel sends value in multiple ways or multiple sender channels sends value in multiple ways:

Then we should wg.Wait() and closing channels in below way:

go func() {
	wg.Wait()
	close(ch)
}()

Refer this functions:
multipleSenderCloseBasic()
multipleSenderCloseV1()

*/

/*
Deadlock - When two or more goRoutines are involved and none of goRoutine can't proceed, Because one GoRoutine holding the lock for more time/forever and not unlocks, So other GoRoutines can't proceed further. Go runtime able to sense this, So it issues deadlock error.

We can't recover deadlock error like panic error.

Deadlock scenarios:

Unbuffered channel:
While read data from channel - Not declared/Take more time/Related read code not hits.
Both read and write channel declared in same goROutine

waitGroup:
wg.Add(1) -- Puts more than expected goRoutine number, Waitgroup waits for Extra goRoutine, That is not present at all.

wg.Done() -- Forgets to add the wg.Done(), So waitGroup thinks that goRoutine still running.

<-ch Reading channel value after wg.Wait()

Mutex lock:
Mutex locked and but not unlocked.


Racing Conditions:
sync.Mutex and sync.RWMutex - Has lock() and unlock() for prevent non-current goROutine to read/write the common values EX:Database, all global variables, values in struct like flag,map Variables.

sync.RWMutex has exclusive read Rlock() and RUnlock() we can use this for read operations alone. This gives better performance than normal mutex's and rwmutex's lock() and Unlock() methods.
Refer - CheckRaceConditionDBSample()


sync/atomic package:
It prevents racing conditions on simple integer increment and decrement values.
Refer - checkRaceConditionWithAtomic()


Concurrency design patterns:
for-Select-Done pattern
Waitgroup pattern
Mutex pattern

Fanout and fanin pattern:
Fanout - Launch multiple goRoutines for completing task. EX:In each for loop iteration, launching a goROutine - go func(){}

FanIn - Aggregating the results from the multiple GoRoutines.
Refer this - FanOut_FanIn_ConcurrencyPattern()





Check this after sometime:
https://www.sobyte.net/post/2022-07/go-sync-cond/
sync.NewCond()


*/

/*
i)Empty return is possible from GoROutine for killing/stopping goRoutine

ii)Returning valuesf rom GoRoutine is not possible causes compilation error and contradictory: EX: a := go function().
GoRoutines are running at same time, Assume returing value from goRoutine make parent goRoutine to wait and process returned value and totally opposite for the GoRoutine processing. So should use "channels for returning like values"
*/

func main() {

	// CheckRaceConditionDBSample()
	// checkRaceConditionWithAtomic()

	// FanOut_FanIn_ConcurrencyPattern()

	ReadsSingleValueFromMultipleChannelValuesV1()

}

func SendingReadingbufferedChannelOneafterOne() {

	wg := sync.WaitGroup{}
	length := 15
	ch := make(chan int, length)

	wg.Add(2)

	go func() {
		defer wg.Done()
		defer close(ch) //Closing the channel is crucial here, After we sent all the data
		for i := 1; i <= length; i++ {
			fmt.Println("sending value into channel", i)
			ch <- i
		}
	}()

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

func ChannelWithClose() {

	ch := make(chan int, 5)
	ch <- 1
	ch <- 10
	ch <- 100
	ch <- 1000
	close(ch) //Channel should be closed in sender side
	//When channel is closed, Sending data into channel will not allowed, Will cause panic.
	//Even channel is closed, Still we can receive the datas from channel

	a, ok := <-ch
	fmt.Println(a, ok)
	a, ok = <-ch
	fmt.Println(a, ok)
	a, ok = <-ch
	fmt.Println(a, ok)
	a, ok = <-ch
	fmt.Println(a, ok)
	a, ok = <-ch //4 values are sent into channel, Then channel closed. Now reading 5th value, This "ok" turns false,
	fmt.Println(a, ok)

	// 1 true
	// 10 true
	// 100 true
	// 1000 true
	// 0 false

}

func multipleChannelsReceivingValuesWithCloseExample() {
	ch := make(chan string)
	ch1 := make(chan string)
	sum := 0
	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		defer close(ch) //Channels Should be closed in sender side. Here closing the channel, After sent all values in this Routine.
		defer wg.Done()
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				// p("GO ROutine number", i)
				sum++
				ch <- fmt.Sprintf("Go Routine 1:%d ------ Value: %d", i, j)
			}
		}
	}()

	go func() {
		defer close(ch1)
		defer wg.Done()
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				// p("GO ROutine number", i)
				sum++
				ch1 <- fmt.Sprintf("Go Routine 2:%d ------ Value: %d", i, j)
			}
		}
	}()

	/*Below ways to receive the values from channels using for loop with range and for loop with select  */

	//For loop with range for channel -- Its very easy and simple way to implement and avoid any deadlocks
	// for i := range ch { //Here one variable is enough "i" in for range loop.
	// 	p("receiving values from channel 1 ", i, sum)
	// }

	// for i := range ch1 {
	// 	p("receiving values from channel 2 ", i, sum)
	// }

	//Another method using for loop with select statement
	channelFlag := true
	channelFlag1 := true
	for channelFlag || channelFlag1 { //using flag values to exit this loop.
		select {
		case i, ok := <-ch:
			if !ok { //Once channel closed, This condition passes
				channelFlag = false
			}
			p("receiving values from channel 1 ------------", i, sum)
		case i, ok := <-ch1:
			if !ok {
				channelFlag1 = false
			}
			p("receiving values from channel 2 -------------", i, sum)
		default: //Better to have default case always in select statement
		}
	}

	wg.Wait()

	p("multipleChannelsReceivingValuesWithCloseExample() completed-----------")
}

func UniDirectionalChannels() {

	//unidirectional channels declaration with make() - But not possible to use only one way send/Recieve channel at all. One Channel should do both send and recieve, So channel sending in one GoRoutine, Should receive in same GoROutine or other GoROutine, So declaring unidirectional channels with make(), Will not have any use mostly.

	//So we have our default bidirectional channels, Then Child GoRoutine called function differentUniDirectionalChannels(), We can control the send only or receive only channel to avoid the different channel send/receive operations by mistake and reduce complexity to avoid unexpected runtime errors.

	// SendOnlychannel := make(chan<- int)
	// ReceiveOnlyChannel := make(<-chan int)

	DefaultBiDirectionChannel := make(chan int)
	DefaultBiDirectionChannel1 := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go differentUniDirectionalChannels(wg, DefaultBiDirectionChannel, DefaultBiDirectionChannel1)

	p("receiving from channel in UniDirectionalChannels()", <-DefaultBiDirectionChannel)

	p("Sending into channel in UniDirectionalChannels()")
	DefaultBiDirectionChannel1 <- 110
	wg.Wait()

}

// channel are defaulty bidirectional channels, But we can make channel as unidirectional channels by below ways in function arguments. Even in above UniDirectionalChannels(), Sends only default bidirectional channels,In differentUniDirectionalChannels() function arguments We can control as unidirectional channels

// SendIntochannel chan<- int
// ReceiveFromChannel <-chan int
func differentUniDirectionalChannels(wg *sync.WaitGroup, SendIntochannel chan<- int, ReceiveFromChannel <-chan int) {
	defer wg.Done()
	p("Sending into channel in differentUniDirectionalChannels()")
	SendIntochannel <- 100

	// <-SendIntochannel   //This is not allowed from compile time, Receiving values from send only channel

	// ReceiveFromChannel<-100   //This is not allowed from compile time, Sending values into receive only channel

	p("receiving from channel in differentUniDirectionalChannels()", <-ReceiveFromChannel)
}

func ChannelreturnValue() {
	//scenario use case example where a function returns channel and reads,
	ch := checking()
	p("Returned channel value", <-ch)
}

func checking() <-chan int {
	ch := make(chan int)
	go func() { //Since checking() is normal function, We should assign value to channel inside goRoutine only, So passing value in a anonmyous function.
		ch <- 10
	}()
	return ch
}

func NormalFunction() {
	p("before", runtime.NumGoroutine())
	go task1()
	go task2()
	for {
		p(runtime.NumGoroutine())
		time.Sleep(1 * time.Second)
	}
}

func task1() {
	for i := 0; i < 20; i++ {
		p("task 1")
	}
}

func task2() {
	for i := 0; i < 20; i++ {
		p("task 2")
	}
}
func MainFunctionIsGoRoutine() {

	//runtime.NumGoroutine() - Prints the number of current goRoutines running
	p("MainFunctionIsGoRoutine()", runtime.NumGoroutine()) //Prints 1 Go ROutine that is main() function, Assume this MainFunctionIsGoRoutine() only called from main(),
}

func comparePerformanceNormalVSGoRoutineUse() {

	//GoRoutines improved the speed almost 3 or 300 % times

	before1 := time.Now()
	for i := 0; i < 10000; i++ {
		sampleTask1()
	}
	fmt.Println("Time difference without GoRoutines", time.Since(before1)) //Took 23 - 25 seconds to complete

	wg := &sync.WaitGroup{}
	before2 := time.Now()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go sampleTask(wg)
	}
	wg.Wait()
	fmt.Println("Time difference withGoRoutines", time.Since(before2)) //Took 7 - 9 seconds to complete
}

func sampleTask1() {
	for i := 0; i < 10000000; i++ {
		_ = 10
	}

}

func sampleTask(wg *sync.WaitGroup) {
	for i := 0; i < 10000000; i++ {
		_ = 10
	}
	wg.Done()
}

func WriteReadMap() {
	wg := sync.WaitGroup{}
	ch := make(chan int)

	mapValue := map[int]int{}

	wg.Add(2)

	go func() {
		defer wg.Done()
		defer close(ch) //This is crucial
		for i := 1; i <= 5; i++ {
			fmt.Println("Adding value into map", i)
			mapValue[i] = i
			ch <- i
		}
	}()

	go func() {
		defer wg.Done()

		//This alternate for -range method for reading channel value
		// for value := range ch {
		// 	fmt.Println("printing square", value*value)
		// }

		for {
			select {
			case value, ok := <-ch:
				if !ok {
					return
				}
				fmt.Println("Reading value into map", value, mapValue)
			default:

			}
		}

	}()

	wg.Wait()
}

func WriteReadMapWithChannelDummySignal() {
	ch := make(chan int, 1)
	ch1 := make(chan int, 1)

	mapValue := map[int]int{}

	go func() {
		defer close(ch)
		for i := 1; i <= 15; i++ {
			fmt.Println("Adding value into map", i)
			mapValue[i] = i
		}
		ch <- 0
	}()

	go func() {
		defer close(ch1)
		for i := 16; i <= 30; i++ {
			fmt.Println("Adding value into map version 1", i)
			mapValue[i] = i
		}
		ch1 <- 0 //After added all the values into map, Sending this dummy value into channel
	}()

	// time.Sleep(1 * time.Second)   //If we are not using any channels in this example, Then this time sleep is required to wait in main thread for child goRoutines to complete

	//Instead of WaitGroup, Here we using dummy value channels for waiting it.
	<-ch //Even buffered channels, It waits for this line gets completed
	<-ch1

	fmt.Println("Reading value into map", mapValue)
}

func ChannelwithWaitGroupBasic() {
	wg := sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(1)
	go func() {
		ch <- 12
		wg.Done()
	}()
	<-ch //Read value from channel should declare above wg.Wait()
	wg.Wait()

	p("ChannelwithWaitGroupBasic() completed")
}

func ChannelwithWaitGroupBasicV1() {
	wg := sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(2)
	go func() {
		ch <- 12
		wg.Done()
	}()

	go func() {
		p(<-ch)
		wg.Done()
	}()
	wg.Wait()

	p("ChannelwithWaitGroupBasicV1() completed")
}

func channelDeclaration() {
	ch := make(chan int)
	ch <- 110
}

func RecoverDeadlockNotWorksLikePanic() {
	ch := make(chan int)
	defer fmt.Println("defer exeecuted") //Normal defer statement too not called during deadlock.
	defer RecoverFunction()              //We can't recover deadlock error like panic error's recover().
	ch <- 10

}

func PanicRecoverGoRoutine() {

	//Both expected panic and recover function call should be same GoRoutine, Then only it recovers correctly
	go func() {
		defer RecoverFunction()
		panic("panic created")
	}()

	//After recovered from above panic in goRoutine, Below sample code executes normally
	p("PanicRecoverGoRoutine() continue executes")
	p("PanicRecoverGoRoutine() continue executes")

	a := 10
	for i := 0; i < 5; i++ {
		p("PanicRecoverGoRoutine() continue executes", a)
	}

}

func RecoverFunction() {
	if r := recover(); r != nil {
		p("recovered from panic")
	}
}

func ReadsSingleValueFromMultipleChannelValues() {
	ch := make(chan int)

	go func() {
		ch <- 10
	}()

	go func() {
		ch <- 11
	}()

	p(<-ch) //In above goRoutines we sends multiple channel values, But reading only one latest value from channel here. (IT not causing any deadlock error, Even for unbuffered channel, because after this line immediately this function completes with below print statement "p("ReadsSingleValueFromMultipleChannelValues() completed")" without further channel recieving values

	//Below ways we can read all the channel values
	// p(<-ch)
	// p(<-ch)

	// for i := 0; i < 2; i++ {
	// 	p(<-ch)
	// }

	p("ReadsSingleValueFromMultipleChannelValues() completed")
}

func ReadsSingleValueFromMultipleChannelValuesV1() {
	ch := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}

	}()

	go func() {
		for i := 10; i < 20; i++ {
			ch <- i
		}

	}()

	go func() {

		for i := 21; i < 31; i++ {
			ch <- i
		}
		time.Sleep(2 * time.Second) //forcing time.sleep so this goRoutine finishes last, So executing below "done" signal to return this parent function ReadsSingleValueFromMultipleChannelValuesV1()
		done <- true

	}()

	for {
		select {
		case value := <-ch:
			p("value receiving", value)
		case <-done:
			p("returning")
			return
		default:
			p("default statment")
			time.Sleep(1000 * time.Millisecond)
		}
	}

	p("ReadsSingleValueFromMultipleChannelValues() completed")
}

func InterruptORKillGoRoutine() {

	//refer the
	// wg := sync.WaitGroup{}

	// for i := 0; i < 5; i++ {
	// 	p(i)
	// 	wg.Add(1)
	// 	go func() {
	// 		if i == 3 {
	// 			defer wg.Done()
	// 		}
	// 	}()

	// }

	// go func() {
	// 	wg.Wait()
	// }()

	// p("main thread completes InterruptGoRoutine()")
}

func TimePackageWithConcurrency() {
	// MultipleGoRoutinesRunsOnDifferentFrequenies()

	TaskCompletedOrKilledOrTimeout()
}

func TaskCompletedOrKilledOrTimeout() {

	// ch := make(chan bool)
	ch := make(chan struct{})

	TimeOut := time.NewTimer(5 * time.Second) //setting the timeout after 5 seconds

	go func() {
		for i := 0; i < 10; i++ {
			// time.Sleep(1 * time.Second) //forcing the timeout
			p("task Running", i)

			//Killing or interrupt a GoRoutine, We should need a channel value/signal passing
			//Killing or interrupt a GoRoutine by sending dummy value in channel and returning this goRoutine.
			if i == 6 {
				p("task interrupted for some logic")
				ch <- struct{}{}
				return //This return is important to exit this goroutine immediately, Otherwise still this loops runs for remaining iteration
			}
		}
		// ch <- true //true is just a boolean value here, You can send "false" value as well for the same response
		p("task completed")
		ch <- struct{}{} //Here the real time example of empty struct, Here to send the dummy value for below select case just like above dummy boolean value "ch <- true"

	}()

	//Select case without infinite or normal range for loop- Will runs first received channel case and exits immediately
	select {
	case <-ch:
		p("Task Done successfully without timeouts")
		TimeOut.Stop() //task is done, So closing the timeout
	case timeOut := <-TimeOut.C:
		p("Timeout occured", timeOut)
	case timeout := <-time.After(2 * time.Second):
		p("timeout occured with time.After() function", timeout) //time.After() channel is more simpler way compares to time.NewTimer's channel
	}

}

func InfiniteForLoopWithSelectInGoRoutine() {

	ch := make(chan bool)
	ch1 := make(chan bool)
	// exit := make(chan bool)

	go func() {
		for { //If we don't want to run for forever, We can use "for loop with range" for exepected values only and then exits

			//Another way of returns this goRoutine with channel values, Once decided to close this goRoutine, We can enable this global variable.
			if globalFlagValue {
				p("globalFlagValue hitted")
				return
			}

			select {
			case <-ch:
				p("task 1 channel triggered")
			case <-ch1:
				p("task 2 channel triggered")

			//In the real use cases, We shouldn't return the infinite for loop with select case goRoutine in most of times. Instead we having the default case with some time.Sleep(), So goRoutine will be running forever with sleep.
			// case <-exit:
			// 	p("Exited the select case")
			// 	return

			default: //This default case will be running infinitely during the no channels value received in the above select cases.
				p("default case")
				time.Sleep(1 * time.Second) //This time.sleep will wait for one sleep, Even needed we can sleep in milliseconds. After time.sleep period this infinite for loop with select case will contine to receive channel values without any issues.
			}
		}
	}()

	//Another way of handling return in Goroutine with channel values
	// if globalFlagValue {
	// 	p("globalFlagValue is enabled, So above goRoutine is returned and not running now, So below channel sending values will cause deadlock panic due to the non-running above goRoutine")
	// 	return
	// }

	ch <- true
	time.Sleep(2 * time.Second)
	ch <- true //sending multiple same channel values due to the select case running inside infinite for loop.

	ch1 <- true

	// exit <- true
	//We already send signal in the exit channel and returned from above goRoutine, So below channel signal causes deadlock panic
	// ch <- true
	// time.Sleep(2 * time.Second)

	//Once this main function exits, Automatically above goRoutine also exits
	p("main thread InfiniteForLoopWithSelectInGoRoutine() completed")

}

func MultipleGoRoutinesRunsOnDifferentFrequenies() {
	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(2 * time.Second) //time.Sleep affects only the particular GoRoutine
			p("Time package with sleep", i)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 5; i++ {
			p("Time package without sleep", i)
		}
		wg.Done()
	}()

	for i := 0; i < 5; i++ {
		time.Sleep(5 * time.Second)
		p("Main thread running with sleep", i)
	}

	wg.Wait()
	//This wg.Wait() place below main thread for loop, So it makes two above goRoutines's for loops and main thread for loop, So all three loop are running at same time.
	//If wg.Wait() placed before main thread for loop, Then two above goRoutines's for loops only running same time and completes. Then main thread for loop runs.

	p("Time package main thread completed")
}

func MultipleGoRoutinesWithContextCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			select {
			// msg from other goroutine finish
			case <-ctx.Done():
				fmt.Println("goROutine1 closed -- ", ctx.Err().Error())
				return
				// end
			}
		}
	}()

	go GoRoutineWithContext(ctx, &wg)

	cancel() //This cancel with trigger the "ctx.Done()" channel value in all child goroutines with same context variable, Then that GoROutines gets retured.

	//If faced any goROutine related error, Then do the cancel() inside seperate goROutine
	// go func() {
	// 	defer wg.Done()
	// 	cancel()
	// }()

	wg.Wait() //Still have to follow the same waitGroup all methods.

	fmt.Println("Main thread completed")
}

func MultipleGoRoutinesWithContextCancelWithInbuiltTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel() //This cancel() will be called automatically, Once the given timeout reached. Thats the difference with context.WithCancel() function.
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			select {
			// msg from other goroutine finish
			case <-ctx.Done():
				fmt.Println("goROutine1 closed -- ", ctx.Err().Error())
				return
				// end
			}
		}
	}()

	go GoRoutineWithContext(ctx, &wg)

	time.Sleep(2 * time.Second) //Putting intentional time sleep delay, Otherwise cancel() will be called within 1 second timeout set

	wg.Wait() //Still have to follow the same waitGroup methods.
	fmt.Println("after", runtime.NumGoroutine())

	fmt.Println("Main thread completed")
}

func GoRoutineWithContext(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		// msg from other goroutine finish
		case <-ctx.Done():
			fmt.Println("goROutine2 closed", ctx.Err().Error())
			return
			// end
		}
	}

}

func MultipleGoRoutinesToSingleChannel() {
	ch := make(chan string)
	sum := 0
	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		defer close(ch)
		defer wg.Done()
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				// p("GO ROutine number", i)
				sum++
				ch <- fmt.Sprintf("Go Routine 1:%d ------ Value: %d", i, j)
			}
		}
	}()

	go func() {
		defer close(ch)
		defer wg.Done()
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				// p("GO ROutine number", i)
				sum++
				ch <- fmt.Sprintf("Go Routine 2:%d ------ Value: %d", i, j)
			}
		}
	}()

	// for i := 0; i < 5; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		for j := 0; j < 5; j++ {
	// 			// p("GO ROutine number", i)
	// 			sum++
	// 			ch <- fmt.Sprintf("Go Routine 2:%d ------ Value: %d", i, j)
	// 		}
	// 	}()
	// }

	// for i := 0; i < 50; i++ {
	// 	p("MultipleGoRoutinesToSingleChannel()", <-ch, sum)
	// }

	// for i := range ch {
	// 	p("MultipleGoRoutinesToSingleChannel()", i, sum)
	// }

	// go func() {
	// 	close(ch)
	// }()

	// for {
	// 	channelValue, ok := <-ch

	// 	if !ok {
	// 		close(ch)
	// 		break
	// 	}
	// 	p("MultipleGoRoutinesToSingleChannel()", channelValue, sum)

	// }
	flag := true
	for flag {
		select {
		case channelValue, ok := <-ch:
			if !ok {
				p("if case")
				flag = false
				break
			}
			p("MultipleGoRoutinesToSingleChannel()", channelValue, sum)
		default:
			p("default case")
			time.Sleep(100 * time.Millisecond)
		}
	}

	wg.Wait()
	// close(ch)

	p("MultipleGoRoutinesToSingleChannel() completed")
}

func SingleGoRoutineToMultipleChannels() {
	oddNum := make(chan int, 10)
	evenNum := make(chan int, 10)

	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 0 {
				evenNum <- i
			} else {
				oddNum <- i
			}
		}
	}()

	for i := 0; i < 10; i++ {
		select {
		case oddNum := <-oddNum:
			p("receiving odd numbers", oddNum)
		case evenNum := <-evenNum:
			p("receiving even numbers", evenNum)
		}
	}

	p("main thread SingleGoRoutineToMultipleChannels() completes")
}

func bufferedChannelBasic() {
	channel1 := make(chan int, 5)
	channel1 <- 20
	value, OK := <-channel1
	if OK {
		fmt.Println(value, OK)
	}
	close(channel1) //closing the channel
}

func SquareValueUsingChannelWaitgroup() {
	//***using channel to square each number
	wg := &sync.WaitGroup{}
	number := []int{2, 4, 6, 8, 10}
	ch := make(chan int)

	// wg.Add(len(number))   //If we want, Alternatively lt we can add the overall goroutine count wg.Add(len(number))instead of below wg.Add(1) for each iteration
	for _, j := range number {
		wg.Add(1)
		go square(wg, ch)
		ch <- j
	}
	wg.Wait()
}

func waitgroupWithoutChannel() {
	//***Making some code running as concurrent in two threads without any information passing using channels
	wg := &sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		fmt.Println("--------i", i)

		wg.Add(1) //For each outer for loop iteration, One goRoutine is created for inner for loop iteration
		go func() {
			for j := 0; j < 10; j++ {
				fmt.Println("=======j", j)
			}
			wg.Done()
		}()
		wg.Wait() //We don't need this wg.wait() inside loop, Both for loop are running independently. If there is no dependency logic between outer for loop and inner goRoutine for loop iterations then is fine to not using this wg.Wait() here. Otherwise should use wg.Wait() here.
	}
	wg.Wait()

}

func writeandReadDifferentGoRoutinesThroughChannel() {
	// var Chan10 = make(chan int)
	var Chan10 = make(chan int, 3)

	//In this scenario, We can use both the channel types,
	// unbuffered channel - After write occured for one iteration here, Will immediately push the data for read channel
	// buffered channel - It Waits and Stores all the write data the upto given length '5' once, Then its reads all the data once.
	// If we reduced the length to '3' EX :make(chan int, 3), Then stores all write value upto channel length '3' once and reads once. Again repeats the same write once and read once for the remaining data upto channel length '3'.

	var w10 sync.WaitGroup
	w10.Add(2)
	go func() {

		for i := 0; i < 5; i++ {
			fmt.Println("reading go")
			fmt.Println("reading go routine values", <-Chan10) //using int value channel, So at a time it can assign/get one integer value only, So that's why we printing channel values under the for loop.
		}
		w10.Done()
	}()

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("writing go routine values", i+10)
			Chan10 <- i + 10
		}
		w10.Done()
	}()

	//Not only the write and read values in channel EX: <- chan10 happens correcly, Also both for loops are running above on this channel doing exactly correct iterations. I think these loops waiting milliseconds internally to ensure both read and write channel is ready for operations, SO this is what channel synchronization working inbuilt.

	w10.Wait()
	fmt.Println("writeandReadDifferentGoRoutinesThroughChannel completed")
}

func structChannelHandle() {
	wg := &sync.WaitGroup{} //have to use waitgroup even for small scenarios, Ensure main thread waits until channel completes
	v := &sam1{}
	df := make(chan *sam1) //create channel for particular struct type
	v.a = 10
	wg.Add(1)
	go structChannel(df, wg)
	df <- v
	wg.Wait()
}

func structChannel(df chan *sam1, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("print struct value from channel", <-df)
}

type sam1 struct {
	a  int
	er *int
}

func ReadAllChannelValuesinMainThread() {

	wg1 := &sync.WaitGroup{}
	ch := make(chan string, 5)
	numberOfAPICalls := 5
	for i := 0; i < numberOfAPICalls; i++ {
		wg1.Add(1)
		go apiCall(ch, wg1)
	}

	wg1.Wait() //this all goRoutines wait() should be declared from below channel values reading. Otherwise below channel reading code gets executed before all GoRoutines complete its running.

	//If we know the number of expected channel values, we can set the limit in for loop, Otherwise need to use below infinite for loop
	for i := 0; i < numberOfAPICalls; i++ {
		select {
		case a := <-ch:
			fmt.Println("go routine received message from api", a)
		default:
			time.Sleep(2 * time.Second)
			fmt.Println("completed---")
			// break
		}
	}

	// for { //select under for loop, Its is always running even you adds break in default section, So we should use this only on inside seperate goRoutine in real cases, So these variables and channel 'a','ch' Should declare as "global variables"
	// 	select {
	// 	case a := <-ch:
	// 		fmt.Println("go routine received message from api", a)
	// 	default:
	// 		time.Sleep(2 * time.Second)
	// 		fmt.Println("completed---")
	// 		break
	// 	}
	// }

}

func GoRoutinesExecutionOrderWithMutexLockUnlock() {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	wg.Add(2)
	//Here we using mutex lock and unlock on both goRoutines, Once first goROutine holds the lock(), Second goroutine will wait until first Goroutine unlocks(), Once first Goroutine unlocks(), then second goRoutines takes lock and does its operation and unlocks it.
	go goRoutine10(&wg, &mu)
	go goRoutine20(&wg, &mu)

	for i := 0; i < 5; i++ {
		p("Main thread running", i)
	}

	wg.Wait()
	p("Time package main thread completed")
}

func goRoutine10(wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	for i := 0; i < 5; i++ {
		// time.Sleep(1 * time.Second) //time.Sleep affects only the particular GoRoutine
		p("First goRotuine", i)
	}
	mu.Unlock()
	wg.Done()
}

func goRoutine20(wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	for i := 0; i < 5; i++ {
		// time.Sleep(1 * time.Second)
		p("Second goRotuine")
	}
	mu.Unlock()
	wg.Done()
}

func ReadAllChannelValuesInanotherGoRoutine() {
	wg1 := &sync.WaitGroup{}
	ch := make(chan string)
	exitChannel := make(chan bool)

	wg1.Add(1)
	go readAllValuesFromChannel(ch, exitChannel, wg1) //passes the channel and waitgroup and this goRoutine runs always

	numberOfAPICalls := 5
	for i := 0; i < numberOfAPICalls; i++ {
		wg1.Add(1)
		go apiCall(ch, wg1)
	}

	wg1.Wait()
	close(ch) //channel should be closed after wg.wait(), So all channel's GoRoutines are completed.
	exitChannel <- true
	close(exitChannel)

}

func readAllValuesFromChannel(ch chan string, exitChannel chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for { //select under for loop, Its is always running even you adds break in default section, So we should use this only on inside seperate goRoutine in real cases
		select {
		case a := <-ch:
			fmt.Println("go routine received message from api---", a)
		case _, exitChannelStatus := <-exitChannel:
			if !exitChannelStatus {
				fmt.Println("exitChannelStatus")
				return
			}
		default:
			time.Sleep(2 * time.Second)
			fmt.Println("completed---")
			break //here "break" not works, this infinite for loop keeps on running

		}
	}
}

func apiCall(a chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("API called")
	a <- "hello message from API"
}

func goRoutinesWriteAndReadExample() {
	wg1 := &sync.WaitGroup{}
	wg1.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Println("write")
		}
	}(wg1)
	wg1.Wait()

	//write operation is completed, now starts the read operation
	wg1.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Println("read")
		}
	}(wg1)
	wg1.Wait()

	fmt.Println("main thread")
}

func bufferedChannel() {
	//buffered channel example
	fg := make(chan int, 5)
	defer close(fg)

	for i := 0; i < 3; i++ {
		fg <- i
		fmt.Println("ttttt")
	}

	fmt.Println(len(fg))

	for i := 0; i < 3; i++ { //Other code not concentrating buffered channel now
		fmt.Println("ttttt1111111")
	}

	// fmt.Println(<-fg) // reads from buffered channel
	// fmt.Println(<-fg)

	//always use this for loop and select for receiving multiple values from channel
	// for {
	// 	select {
	// 	case a := <-fg:
	// 		fmt.Println("values from channels-----", a)
	// 	default:
	// 		time.Sleep(2 * time.Second)
	// 	}
	// }
}

func square(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	channelValue := <-ch
	fmt.Println("square numbers", channelValue*channelValue)
}

func printSequenceUsingGoRoutines() {
	var wg = sync.WaitGroup{}

	ch := make(chan string)
	// ch1 := make(chan string)

	rowNumber := 7

	fmt.Println("Printing the sequence using two GoRoutines from main()--------")

	for i := 0; i < rowNumber; i++ {

		for j := 0; j <= i; j++ {
			wg.Add(2)

			if j%2 == 0 {
				go OneGoRoutine(&wg, ch)
				fmt.Print(<-ch)

			} else {
				go zeroGoRoutine(&wg, ch)
				fmt.Print(<-ch)
			}

		}

		fmt.Println()

	}

	go func() {
		wg.Wait()
		close(ch)
		// close(ch1)
	}()

	fmt.Println("completed------")
}

func zeroGoRoutine(wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	ch <- "0"
}

func OneGoRoutine(wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	ch <- "1"
}

func multipleSenderCloseBasic() {
	var wg = sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	go func() {
		defer wg.Done()
		for i := 11; i < 21; i++ {
			ch <- i
		}
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	for i := range ch {
		fmt.Println("receiving values from multiple channel sender 11----", i)
	}

	fmt.Println("Completed")
}

func multipleSenderCloseV1() {
	var wg = sync.WaitGroup{}
	ch := make(chan int)
	ch1 := make(chan int)

	wg.Add(4)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	go func() {
		defer wg.Done()
		for i := 11; i < 21; i++ {
			ch <- i
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
	}()

	go func() {
		defer wg.Done()
		for i := 11; i < 21; i++ {
			ch1 <- i
		}
	}()

	go func() {
		wg.Wait()
		close(ch)
		close(ch1)
	}()

	go func() {
		for i := range ch {
			fmt.Println("receiving values from multiple channel sender----", i)
		}
	}()

	for i := range ch1 {
		fmt.Println("receiving values from multiple channel sender 11----", i)
	}

	fmt.Println("Completed")
}

func GoRoutineLeakExample() {

	for i := 0; i < 4; i++ {
		GoRoutineLeak()
		fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
	}
	fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())

}

func GoRoutineLeak() {
	ch := make(chan int)

	go func() { ch <- 100 }()
	go func() { ch <- 100 }()
	go func() { ch <- 100 }()
	<-ch // Doing only one channel reading for one GoRoutine, So other two channel value reading from two goRoutines, Keep on waiting.
	//Normal simple .go, Once main() executes all GoRoutines going to terminate it. But in prod cases, Main() mostly running always. So this GoOrutine leak impact a lot.
}

// multiple go-routines order the execution by using channels.
// Each routine has seperate channel and By using these channels we ordering the go-routines' outpue flow

func goRoutineExecutionOrder() {
	// var wg sync.WaitGroup
	// defer fmt.Println("defer Main")
	// defer wg.Wait()

	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	// wg.Add(3)

	go func() {
		// defer wg.Done()
		ch1 <- "One"
	}()

	// go func() {
	// 	// defer wg.Done()
	// 	ch2 <- "Two"
	// }()

	secondTest(ch2)

	go func() {
		// defer wg.Done()
		ch3 <- "Three"
	}()

	fmt.Println(<-ch3)
	fmt.Println(<-ch1)
	fmt.Println(<-ch2)

	fmt.Println("main thread complete")
}

func secondTest(ch2 chan string) {
	ch2 <- "Two"
}

func checkRaceConditionWithAtomic() {
	wg := sync.WaitGroup{}
	atomic := func() {
		defer wg.Done()
		for i := 0; i < 2000; i++ {
			atomic.AddInt64(&counter, 1) //Increment one integer value
			// atomic.AddInt64(&counter, -1) //decrement one value
		}
	}

	wg.Add(3)
	go atomic()
	go atomic()
	go atomic()
	wg.Wait()

	fmt.Println("Final counter value After multiple Gorouitne ", counter)

}

// Assumes this is DB client struct with own waitgroup and mutexes on working with goRoutines
type MongoClientConnectionSample struct {
	DBValueSample int
	DBMap         map[int]int
	mutex         sync.Mutex
	wg            sync.WaitGroup
	rwmutex       sync.RWMutex
}

/*
Reading the varaible by multiple goRoutines/Concurrency safe with racing conditions:
If we are doing reading only among multiple goRoutines -- That is safe with racing conditions, Until that variable's value can be written/modified parallely.
*/

// sync.RWMutex - has normal lock as sync.mutex, As well as seperate exclusive read RLock() and read RUnLock(), This synchronize only with other GoRoutines read RLock() and read RUnLock() and prevent racing conditions on read operations alone, So doesn't interfere on normal mutex and RWmutex lock() and unlock().

func (m *MongoClientConnectionSample) LoadValues() {
	//We can use normal mutex lock() and unlock() as well.
	// defer m.mutex.Unlock()
	// m.mutex.Lock()

	defer m.rwmutex.Unlock()
	m.rwmutex.Lock()
	m.DBValueSample++
	m.DBMap[1] = m.DBValueSample
}

func (m *MongoClientConnectionSample) UnLoadValues() { //Should use pointer receiver for waitgroup and mutex works as pass by reference.

	defer m.rwmutex.Unlock()
	m.rwmutex.Lock()
	m.DBValueSample--
	m.DBMap[1] = m.DBValueSample
}

// This read the common value is not working even we used lock and unlock, So need better design for better real requirement to implement this readValue with multiple goROutines.
func (m *MongoClientConnectionSample) ReadValue() {
	defer m.wg.Done()
	defer m.rwmutex.RUnlock()
	m.rwmutex.RLock()
	fmt.Println("reading value 1", m.DBValueSample)
}

func CheckRaceConditionDBSample() {

	db := MongoClientConnectionSample{}
	db.DBMap = map[int]int{}

	closure := func() {
		defer db.wg.Done()
		for i := 0; i < 1000; i++ {
			db.LoadValues()
		}

	}

	Unloadclosure := func() {
		defer db.wg.Done()
		for i := 0; i < 1000; i++ {
			db.UnLoadValues()
		}
	}

	db.wg.Add(5)
	go closure()
	go Unloadclosure()
	go closure()
	go Unloadclosure()
	go closure()

	db.wg.Wait()

	fmt.Println("final DB value", db.DBValueSample, "--------", db.DBMap) //1000

}

func FanOut_FanIn_ConcurrencyPattern() {
	data := []int{1, 2, 3, 4, 5}
	input := make(chan int, len(data))
	for _, d := range data {
		input <- d
	}
	close(input)

	// Fan-out: Launch multiple worker goroutines for the task
	numWorkers := 3
	results := make(chan int, len(data))

	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for num := range input { //Reads from input buffered channel
				result := num * 2
				results <- result //writes into results buffered channel
			}
		}()
	}

	// Handling multiple sender goRoutines general
	go func() {
		wg.Wait()
		close(results)
	}()

	// Fan-In -- Aggregate results from workers and Process aggregated results
	for result := range results {
		fmt.Println("Fan-In aggregated result", result)
	}
}
