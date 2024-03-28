package main

import (
	"fmt"
	"sync"
)

//https://golangbyexample.com/singleton-design-pattern-go/

type Singleton struct {
	sampleValue int
}

// Creates Global variable for the Singleton struct, Creating normal (or) pointer variable is optional
var SingletonObject *Singleton

// Singleton - If already we have the SingleTon struct object, shares the same object otherwise Creates new object.

// USE CASES: Singleton used for DB/sessions connections,Configuration(Every time app starts) and logging etc.
func GetSingletonInstance() *Singleton {
	if SingletonObject == nil {
		fmt.Println("creating Singleton object 0")
		SingletonObject = &Singleton{}
		return SingletonObject
	}

	fmt.Println("Already created Singleton object 0")
	return SingletonObject
}

// ANOTHER METHOD-----------
// Below another method of creating and maintaining Singletion object with involvement of multiple GoRoutines, Sync.Once.DO() Ensures That func() calls only once for creating singleton object even with multiple GoRoutines
var syncOnce = sync.Once{}

func GetSingletonInstanceV1() *Singleton {
	if SingletonObject == nil {

		syncOnce.Do(func() {
			fmt.Println("creating Singleton object 1")
			SingletonObject = &Singleton{}
		})

	} else {
		fmt.Println("Already created Singleton object 1")

	}

	return SingletonObject

}

// ANOTHER METHOD-----------
// Another way of creating and maintaining Singletion object - If we put inside init(), It can be called only once.
// func init() {
// 	SingletonDesignPattern()
// }

func main() {

	SingletonDesignPattern()

}

func SingletonDesignPattern() {

	GetSingletonInstance()
	GetSingletonInstance() //Here we gets same object from above line function call

	for i := 0; i < 10; i++ {
		go GetSingletonInstanceV1()
	}

}
