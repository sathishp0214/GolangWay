package main

import "fmt"

type OldInterface interface {
	OldMethod()
}

type OldStruct struct {
	OldValue int
}

func (o OldStruct) OldMethod() {
	fmt.Println("From OldMethod")
}

type NewAdapterInterface interface {
	NewAdapterMethod()
}

//Adapter design pattern - Using different interface's method. Here we reusing the OldInterface's OldMehod() method's in another NewAdapterInterface's method.

//Real life use case: Convert 2 pin socket into 3 pin socket. So 3 pin socket(adapter struct) which plugs top of the 2 pin socket

//USE CASE: Reusing another Interface's method in another struct or another struct which implements another interface's method
type NewAdapterStruct struct {
	Value int
	Old   OldInterface //Adding oldInterface -- Here Assumes this OldInterface method logic is different from  NewAdapterInterface's NewAdapterMethod(), Otherwise we can directly use OldInterface method.
}

func (newAdapter NewAdapterStruct) NewAdapterMethod() {
	fmt.Println("From NewMethod")
	newAdapter.Old.OldMethod() //Calling the OldInterface's OldMehod(),
}

func main() {
	oldStruct := OldStruct{}
	NewAdapter := NewAdapterStruct{Old: oldStruct} //Passing oldStruct which implements OldInterface's OldMehod()
	NewAdapter.NewAdapterMethod()
}
