package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/stretchr/testify/mock"
)

type Smartphone struct {
	name string
}

type Geek struct {
	smartphone *Smartphone
}

func replaceByNG(s **Smartphone) {
	*s = &Smartphone{"Galaxy Nexus"}
}

func replaceByIPhone(s *Smartphone) {
	s = &Smartphone{"IPhone 4S"}
	// s.name = "DDDDDDDdd"
}

// func main() {
// 	// geek := Geek{&Smartphone{"Nexus S"}}
// 	// println(geek.smartphone.name)

// 	// replaceByIPhone(geek.smartphone)
// 	// println(geek.smartphone.name)

// 	// replaceByNG(&geek.smartphone)
// 	// println(geek.smartphone.name)

// 	// d := []int{2, 3, 4, 5}
// 	// e := SliceFn{}
// 	// e.Len(d)

// 	sl := make([]int, 20, 100)
// 	fmt.Println(sl, len(sl))

// 	runtime.Gosched()

// }

func task1() {

	for i := 0; i < 10; i++ {
		runtime.Gosched()
		fmt.Println("Task 1:", i)
	}
}

func task2() {
	for i := 0; i < 10; i++ {
		fmt.Println("Task 2:", i)
	}

}

func task3() {
	for i := 0; i < 10; i++ {
		fmt.Println("Task 3:", i)
	}
}

func GoScheduleRuntime() {
	go task1()
	go task2()
	go task3()

	time.Sleep(time.Second * 3) // Wait for both tasks to finish (approximately)
}

func longRunningTask(name string) {
	for i := 0; i < 5; i++ {
		fmt.Println(name, ":", i)
		time.Sleep(1 * time.Second) // Simulate some work
		runtime.Gosched()           // Yield control after each iteration
	}
}

type MockUserService struct {
	mock.Mock
}

func main() {
	// go longRunningTask("Routine A")
	// go longRunningTask("Routine B")

	// time.Sleep(10 * time.Second) // Wait for both routines to finish

	d := MockUserService{}
	d.Called()

}

type SliceFn[T any] struct {
	s    []T
	less func(T, T) bool
}

func (s SliceFn[T]) Len() int {
	return len(s.s)
}
func (s SliceFn[T]) Swap(i, j int) {
	s.s[i], s.s[j] = s.s[j], s.s[i]
}
func (s SliceFn[T]) Less(i, j int) bool {
	return s.less(s.s[i], s.s[j])
}
