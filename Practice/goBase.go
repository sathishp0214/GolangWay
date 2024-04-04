package main

import (
	"encoding/json"
	"errors"
	"fmt"
	. "fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

/*
Python vs Golang:

Golang doesn't support exceptions,inheritance,oop constructor like python.
Golang anonymous is more powerful than python anonymous functions. python anonymous functions supports only single line statements and return it.
Python doesn't supports pointers, defer statements. Python supports exception "finally" and "with" statement instead of "defer".
Python has better package support and community support than golang.
Golang code is more faster than python EX: compare normal same "hello world" program to same algorithm programs.
Golang has inbuilt and more efficent concurrency like GoRoutine with channels support. Python multithreading is not light weight and not effcient as Golang and have to use "multithreading" package.
Python thread is preferred for I/O bound tasks only, Where python multiprocessing is preferred for CPU-bound tasks, Where GoROutine is fine for both I/O bound and CPU bound tasks, So Routines don't need multiprocessing seperately.
Golang is preferred in development of high performance required application, microservices, cloud computing, system programming like direct OS hardware applications like drivers.
Python is definitely preferred in Data analytics/Data scientist/Machine learning areas and Web application development beacuse of more mature with more inbuilt features frameworks like django and flask.


Golang Top features:
Inbuilt concurrency support
Supports first class functions.
Powerful anonymous function support.
Defer
Panic and recover.
*/

var p = fmt.Println //using "fmt.Println" as global variable "p" for convenience usage of print statements as p()

/*Golang doesn't supports constant/(literal means constant) variable in array,slice,maps. */

// enum -- Generally useful for setting auto-increment constant values, Golang uses "iota" like below for implementing enum

const (
	North int = iota //int defaultly starts on 0, north=0, south=1 and auto-increments
	South
	East = 10 //breaks this auto-increment iota and assigns value as 10, So all the next below values set as 10
	West
	SouthEast
)

//enum iota - We can't use for string data type

// // Global variables declaration for all data types
var a1 = 10

// var b1 int = 100
// var pe = false
// var w1 = []int{2, 3, 4, 5}
// var ep = map[int]int{1: 11, 2: 22}
// var Globalstruct SampleStruct
// var GlobalstructPointer *SampleStruct
// var GlobalInterface interface{}
// var globalIntPointer *int
// var globalBoolPointer *bool
// var globalSlicePointer *[]int
// var channel1 chan string

var GlobalAnonmyousStruct = struct {
	a int
	b string
}{10, "Sat"}

//NEED TO DO:
// mutable and immutable X
// call by value and call by reference X
// call by value and call by reference with struct receivers X
// channel,GoRoutine combinations/exercises practice
// first class functions practice
// struct and interface exercises practice
// Solid principles
// Design patterns
// unit testing practice
// http client request GET and POST
// Mongo connection and running a query
// SQL connection and running a query
// bufio package
// flag package
// runtime package
// context package
// recall all packages
// Recall programs
// new programs
// Recall docker,kafka,jenkins,aws
// data structures algorithms

func main() {

	//data types
	// int - int8,int16,int32,int64
	// float - float32, float64,
	// string
	// bool
	// rune - int32
	// byte - int8
	// array
	// slice
	// map
	// struct
	// interface - its also called as "any" dataType
	// chan

	//short variable declaration/dynamic variable declaration/type inference
	// f := 10

	//If we used more number range than allowed capacity, Will get compilation error
	// var i int
	// i = 345265373576868896096  //This big number not allowed in int, If we reduced 1 or 2 digits we can still use with int data type
	// fmt.Println(i + i + i + i + i*1000)

	//a1 - this Already global variable, But still we can create new local scope variable with same variable name
	// a1 := 20
	// fmt.Println(a1)

	// checkCallByValueReference()

	// panicAndRecover()
	// p("Recovered from panic and main function continues to run")

	// FirstClassFunctions()

	// emptyInterface()

	// p(North, South, East, West)

	// NestedFunction()

	// SliceMakeFunction()

	// panicAndRecover()

	// InbuiltGolangFunctionsLatest()

	varDeclaration()

}

func NestedFunction() {

	//We can't have normal nested function definition/declaration
	// func hello() {
	// 	func hello1() {
	// 		p("hello")
	// 	}
	// }

	//Nested function is possible in inline functions
	func() {
		for i := 0; i < 5; i++ {
			func() {
				p("inside inner function")
			}()
		}
		p("inside outer function")
	}()
}

func VariableShadowingInGolang() {
	x := 10
	//VariableShadowingInGolang - If same Variables are using on both outer scope and inner scope, Inner scope variables shadows/hides the outer scope same variable.
	func() {
		x := 11
		fmt.Println(x) //11
	}()

	fmt.Println(x) //10

	//also variables declared in side for loops, if conditions etc
	if true {
		x := 25
		fmt.Println(x) //25
	}

	fmt.Println(x) //10
}

func errorDataType() {
	//Creating "error" type messages
	error1 := errors.New("Returning error message")
	error2 := fmt.Errorf("Returning error message one")
	if error1 != nil && error2 != nil {
		p(error1, error2)
	}
}

func mutable_Immutable_DataTypes() {

	//Don't try to print and compare the memory address like below, It will confuse you, Have simple understanding in below simple words.
	c := 10
	d := 10
	fmt.Printf("%p %p", &c, &d)

	fmt.Println()

	a := []int{1, 2}
	b := []int{1, 2}
	fmt.Printf("%p %p", a, b)

	//mutable data types  - These data types values can be modified directly without need to allocte any new memory for the modifying operation.

	//In Simple Words - We can directly modify the values EX: array[0] = 2, map[2] = 22 etc

	// array and array of any datatypes
	// slice ''
	// map	  ''
	// chan    ''

	//immutable data types

	//In simple Words - We can't directly modify the values EX: string[1]="a"

	//For modifying this We need to allocate new memory Example -- Either we need to change this string into any other data type like slice and then modify this (or) We can use the strings package's replace() funtion, Again this function use the new memory inside for modifying.

	// int - int8,int16,int32,int64
	// float - float32, float64,
	// string
	// bool
	// rune - int32
	// byte - int8
	// struct

	//another Value (or) New value assigning is possible for all data types

	//For me -- Pointer is not a data type, It just Golang's feature for holding the memory address for any data type. So in this mutuable or immutable, Pointer just hold the memory address, So we can't modify the memory addres directly. So we have to dereference into value, then have to modify, EX: *pointer = 10, Then that particular data type EX: (int) nature is automatically came inside.

}

func InbuiltGolangFunctionsLatest() {

	// make()  //USed to create slice, map, channel data types

	//new() - Used to create new pointer for all data types
	Intpointer := new(int)
	SlicePointer := new([]int)
	structPointer := new(SampleStruct)

	p("new() pointers", Intpointer, SlicePointer, structPointer)
	p("new() creates non-nil pointers", Intpointer == nil) //false, new() creates non-nil pointers

	f := []int{1, 2, 3}
	g := map[int]int{1: 11, 2: 22}

	//max and min currently only useful for passed values like below - Not much useful in passing another data types like slice etc
	p(min("c", "a", "b"))
	//max

	//Clear() useful in empty the map, In slice used to turn all values into zeroes
	clear(f)
	p("clear values in slice to zeroes", f, len(f)) //[0 0 0] converted into zeroes only, Still length is 3 only

	clear(g)
	p("emptied Map", g, len(g))

	g1 := map[int]int{1: 11, 2: 22}
	delete(g1, 2)
	p("deleted the key in map", g1)

	//copy()
	sourceSlice := []int{1, 2, 3, 4, 5}
	//copy() does the deep copy in slices.
	//copy Destination slice should be in same length of source slice, So we can create using below make() with length or normal slice creation EX: destinationSlice := []int{40, 41, 42, 43, 44}
	destinationSlice := make([]int, 5)
	copy(destinationSlice, sourceSlice)
	p(destinationSlice, sourceSlice)
	destinationSlice[0] = 111 //This does deep copy, Its vary from other slice normal copying sliceCOpying()
	p(destinationSlice, sourceSlice)

	//If the length of source and destination slice is different in copy()
	//If source slice has 5 values, Destination slice has 3 values or viceversa. copy() copies only the destination slice length. (not useful in all situations)
	df := []int{1, 2, 3, 4, 5}
	df1 := []int{10, 11, 12}
	copy(df1, df)
	p(df1, df) //df1 ==> [1 2 3]   //destination slice length is 3, only copied first three values

	initialSlice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	itemsNeedToBeAppended := []int{9, 10, 11, 12, 13}
	//using copy() to merge two slices, This advised for more faster and memory efficient rather than simple append()
	p(appendToSliceMoreEfficentMethod(initialSlice, itemsNeedToBeAppended))

}

func appendToSliceMoreEfficentMethod(initial []int, items []int) []int {
	result := make([]int, len(initial)+len(items)) //creates the length of two input slices to get result slice
	p(len(initial)+len(items), result)
	copy(result, initial)              //First copying the initial slice into result
	copy(result[len(initial):], items) //Then copying items slice into result
	return result
}

func slicing() {
	Slice := []int{2, 3, 4, 5, 6, 7, 8}
	p(Slice[2:4])
}
func sliceShallowDeepCopyDefault() {
	var emptySlice []int
	baseSlice := []int{2, 3, 4, 5, 6, 7, 8}
	Second := baseSlice[:]
	third := baseSlice
	emptySlice = baseSlice

	// Second[2] = 133 //This changes 2nd index in all slices, "why?? Shallow copy - We just passes the baseSlice reference to other slices in different above ways.
	// p("before", baseSlice, Second, third, emptySlice)
	// //before [2 3 133 5 6 7 8] [2 3 133 5 6 7 8] [2 3 133 5 6 7 8] [2 3 133 5 6 7 8]

	//At this momemt, All above slices has the same memory address because every slice got values passed from "baseSlice" in different ways. Once the other individual slices has some modifications like appending in below lines, Then that particular slice memory address change from the baseSlice memory address, After that any modifications on either or both baseSlice and otherSlice(s) values Will not be reflected(Like shallow Copy).

	//After we doing the append in below slices, That creates new memory (Due to re-slicing), So Hereafter whatever changes will not reflect with "baseSlice", So DeepCopy works like this in slices.
	Second = append(Second, 100)
	third = append(third, 200)
	emptySlice = append(emptySlice, 300)

	//Below modifications will reflect deep copy the values in any other slices, Because all other slices had append modifications on above

	baseSlice[1] = 1000
	Second[0] = 33

	p(baseSlice, Second, third, emptySlice)
	//[2 1000 4 5 6 7 8] [33 3 4 5 6 7 8 100] [2 3 4 5 6 7 8 200] [2 3 4 5 6 7 8 300] -- This value got, Assume "Second[2] = 133" above code is commented

	// InbuiltGolangFunctionsLatest() //refer this function for slice deep copy

}

// Interface without any method signatures called as empty interface.
func emptyInterface() {

	//using interface as map key is not supported
	// var outputMap map[interface]interface{}

	//validating interface data type value example
	var i interface{}
	i = 10
	value, ok := i.(string) //actual int value pre-validates as string value
	if !ok {
		fmt.Println("error in interface data type handling", value)
	}

	//Type assertion - Type assertion access data as particular data type in the interface data.
	var T interface{}
	T = 10
	T = "stringValue"
	T = sample{}
	T = []int{1, 2, 3, 4}
	T = &sample{}
	T = map[int]int{1: 11, 2: 22}
	switch T.(type) { //Using switch case, Type assertion access data type from interface{}
	case int:
		p("int type", T.(int))
	case string:
		p("string type", T.(string))
	case []int:
		p("int slice type", T.([]int))
	case map[int]int:
		p("map type", T.(map[int]int))
	case sample:
		p("sample struct type", T.(sample))
	case *sample:
		p("sample struct pointer type", T.(*sample))
	}

	// //Interface map of different types
	interfaceMap := map[int]interface{}{}

	//Another ways of delaring map interface
	// var interfaceMap map[int]interface{}
	// interfaceMap = map[int]interface{}{}

	// interfaceMap := make(map[int]interface{})

	interfaceMap[0] = 000
	interfaceMap[1] = "sathish"
	interfaceMap[2] = []int{1, 2, 3}
	interfaceMap[3] = sample{}
	interfaceMap[4] = true

	for i, j := range interfaceMap {
		p(i, j, reflect.TypeOf(j))
	}

	//interface slice
	interfaceSlice := []interface{}{}
	interfaceSlice = append(interfaceSlice, "ss")
	interfaceSlice = append(interfaceSlice, 00)

}

func ReflectionInGolang() {
	//Reflection generally inspect the variables and variables data type in runtime.
	//Reflection in golang, If we are using empty interface, Then we have to inspect the variables and variables data type in runtime using reflect packages reflect.Typeof(), reflect.ValueOf() etc or simple type assertion.

	var x interface{}
	x = 12
	fmt.Println(reflect.TypeOf(x), reflect.ValueOf(x)) //int, 12

	x = s{10, "sat"}
	fmt.Println(reflect.TypeOf(x), reflect.ValueOf(x))               //main.s {10 sat}
	d := x.(s)                                                       //type assertion
	fmt.Println(d.a, d.b, reflect.TypeOf(d.a), reflect.ValueOf(d.a)) //10 sat int 10

}

type s struct {
	a int
	b string
}

func defaultMemoryBytesSizeDataTypes() {
	//gets the default bytes size Ex: 8,16,24 bytes of different data types, It may not accurate size, But understanding of memory size, So we prefer the less memory data types for supported functions

	stringValue := "s hello well sat hello well sat hello well sat hello well sat hello well sat hello well sat hello well sat hello well sat hello well sat hello well sat hello well sat hello well sat hello well sat hello well sat hello well sat hello well sat hello well sat hello well sat hello well sat hello well sat hello well"
	stringSlice := []string{"11", "22", "33", "44"}
	boolValue := true
	intValue := 100
	intSlice := []int{1, 2, 3, 4}

	p("bytes sizes--------------")
	fmt.Println(reflect.Type.Size(reflect.TypeOf(stringValue)), unsafe.Sizeof(stringValue)) //string 16 bytes, Even its very small or big string
	fmt.Println(reflect.Type.Size(reflect.TypeOf(stringSlice)), unsafe.Sizeof(stringSlice)) //slice 24 bytes

	p("bool size", unsafe.Sizeof(boolValue)) //bool size 1 byte
	p(unsafe.Sizeof(intValue))               //int value 8 byte
	p(unsafe.Sizeof(intSlice))               //int slice 24 bytes same as string slice

	runeSize := rune('s')
	ByteSize := byte('f')
	p(reflect.Type.Size(reflect.TypeOf(runeSize)), unsafe.Sizeof(runeSize)) // rune 4 bytes
	p(reflect.Type.Size(reflect.TypeOf(ByteSize)), unsafe.Sizeof(ByteSize)) // byte data type size- 1 byte

	map1 := map[int]int{1: 11, 2: 22}
	map2 := map[string]string{"a": "aaa", "b": "bbb"}
	p("map size", unsafe.Sizeof(map1), unsafe.Sizeof(map2)) //map size is 8 bytes

	p("Normal struct size", unsafe.Sizeof(SampleStruct{})) //normal struct size is depends upon the summation of all struct's variable datatypes size.

	p("empty struct size", unsafe.Sizeof(emptyStruct{})) //Size of empty struct is 0 bytes, So that's why its is used in map's dummy key's value, Using channel's dummy signal value.

	p("global variable size", unsafe.Sizeof(a1)) //global variable a1 int has as same 8 bytes like local int variable

	fmt.Println(unsafe.Sizeof(&stringSlice), unsafe.Sizeof(&SampleStruct{}), unsafe.Sizeof(&boolValue), unsafe.Sizeof(&intValue)) //pointer 8 bytes common for different data types, Even normal bool size is 1 byte, But using bool as pointer size is 8 bytes.
}

type emptyStruct struct{}

func pointerDeclaration() {

	//For me -- Pointer is not a data type, It just Golang's feature for holding the memory address for any data type. So in this golang's mutuable or immutable data types, Pointer just hold the memory address

	t := 100
	structPointer := &SampleStruct{}
	structPointer.er = &t
	//dereferencing struct pointer variable and changes the value
	*structPointer.er = 1000
	fmt.Println(structPointer, *structPointer)

	t2 := 200
	tp := SampleStruct{}
	tp.er = &t2
	*tp.er = 300
	fmt.Println(tp.er, *tp.er)

	h := []*SampleStruct{} //slice of struct pointers
	h = append(h, &SampleStruct{a: 10, er: &t2}, &SampleStruct{a: 20, er: &t}, structPointer)

	//CheckCallByValueDataTypes()   -- check this function for more dereference examples

}

//variadicFunction(1, 20, 30, 40, 50)
//variadicFunction(1) //this 1 for "a" argument, no argument passed for params argument.

func variadicFunction(a int, params ...int) { //params is an optional argument and also take any number of int arguments or no arguments at all, Here params argument is a []int data type

	for _, v := range params {
		p(v) //20, 30, 40, 50
	}

}

func recoverFunction() {
	r := recover() //This alone recovers from panic() function
	if r != nil {  //This used for logging, Whether recover is happened or not
		p("recovered successfully")
	}

}

func panicAndRecover() {
	//different below ways for defer recover function

	// defer recoverFunction()

	func() {
		//defer with anonymous function for recover
		defer func() {
			r := recover() //This alone recovers from panic() function, Then function which called panic occured function, Will continue to run.
			if r != nil {
				fmt.Println("Panic recoverd") //This used for logging, Whether recover is happened or not
			}
		}()

		panic("panic created")
	}()

	fmt.Println("recovered and running further")

}

func ByteAndRune() {
	//rune vs byte
	//byte - uint8 -- 8 bites -- it works only upto ascii values (0-255)
	//rune -- int32 -- 32 bits -- It works ASCII and more broader unicode characters upto around 65000 types of characters

	//converts character to ascii/unicode and ascii/unicode to character using both (rune() and byte())
	d := rune('A')
	fmt.Println(rune('A'), d)     // 65
	fmt.Println(string(rune(65))) // A

	v := byte('A')
	fmt.Println(byte('A'), v)
	fmt.Println(string(byte(65)))

	//fp := 'P' //short varibale declaration, It considers a rune type automatically

	//runeslice, Similarly can do byteSlice as well.
	runeSlice := []rune{'1', 'a', 2} //this int 2 takes direct ascii value
	runeSlice = append(runeSlice, 'd')
	Println("runeSlice", runeSlice) //prints ascii/unicode values -- [49 97 2 100]

	fg11 := "sathish is hello"

	bytesArray := []byte(fg11)
	RuneArray := []rune(fg11)
	bytesArray[0] = 'Z'
	RuneArray[0] = 'Z'
	fmt.Println(bytesArray, RuneArray, string(bytesArray), string(RuneArray))

}

func varDeclaration() {

	//declaring variables with "var" keyword with below data type returns nil, So needs to intialize and use it for avoiding nil pointer errors.
	var a int
	var b string
	var c bool
	var e map[int]int
	var e1 []int
	var p interface{}
	var p1 chan int
	if e == nil && e1 == nil && p == nil && p1 == nil {
		fmt.Println("declare variables with var keywords returns nil--", e, e1, p, p1, a, b, c)
	}

	//now intialized the values, Now it will not return nil
	e = map[int]int{1: 1, 2: 2}
	e1 = []int{1, 2, 3}
	p = 100
	p1 = make(chan int)

	fmt.Println("initialzed values with var variables--", e, e1, p, p1)
}

func typeCasting() {
	//Type_Casting - Converts the variable from one dat type to another data type.

	var f float64 = 6.44
	fmt.Println(int(f)) //converts float into integer

	//converts int to string
	NumberInStringFormat := fmt.Sprintf("%v", 10)
	fmt.Println(NumberInStringFormat, reflect.TypeOf(NumberInStringFormat))

	n2 := 154
	n3 := strconv.Itoa(n2) //Converts integer into string (This is recommended way) can not use "string(int_number)"
	fmt.Println(n3, reflect.TypeOf(n3))

	//strconv.Atoi - returns in 'int' datatype (This is more convenient for integer operations)
	var str1 string = "101958738909596"
	newInt, _ := strconv.Atoi(str1) //converts string into integer
	fmt.Println(newInt)

	//converts string to string slice, byte slice and rune slice

	//Strings are immutable, So we can't assign values directly Ex: fg11[2] = "Z", So should use string slice, bytes slice, array slice for achieving that.

	//using byte slice is preferred here, Because its uint8 data type and its uses less memory

	//This is Rune and Bytes most used real use case.

	fg11 := "sathish is hello"

	bytesArray := []byte(fg11)
	RuneArray := []rune(fg11)
	bytesArray[0] = 'Z'
	RuneArray[0] = 'Z'
	fmt.Println(bytesArray, RuneArray, string(bytesArray), string(RuneArray))

	stringslice := strings.Split(fg11, "") //converts string into string slice
	stringslice[2] = "ttttt"
	fmt.Println(stringslice, strings.Join(stringslice, "")) //converts string slice into string

}

func deferFunction() {
	//defer not works in deadlock errors
	//defer works in "panic" also. (if "defer statements" comes only before the "panic statement")
	// defer fmt.Println("PPPPPPPP")   //This prints
	// panic("Stop")
	// defer fmt.Println("HHHHH")   //this defer will not work

	//-----------------------------
	// i := 0
	// defer fmt.Println("defer: ", i)  //prints i=0, This defer statement prints last, But stores the value in the line program execution flow.
	// i++
	// fmt.Println("normal: ", i)

	//defer behaviour
	fmt.Println("before")
	defer fmt.Println("defer hitted")
	if true { //False condition defer statement will not execute.
		defer fmt.Println("defer if condition hitted") //even its inside if condition stil executes after this statement "fmt.Println("after")"
		fmt.Println("if condition")
	}
	fmt.Println("after")

	for i := 0; i < 3; i++ {
		defer fmt.Println("defer hitted", i) //This defer statement executes everytime for each loop iteration
		fmt.Println()
	}

	// defer fmt.Println("defer statement executed")
	// os.Exit(100) //Due to this os.Exit() program exits this line, Above defer statement even too not executed
	// fmt.Println("normal statement")
}

func CallByValueDataTypes() {

	intValue := 10
	floatValue := 5.5
	strValue := "sat"
	boolValue := true
	arrayValue := [3]int{1, 2, 3}
	structValue := SampleStruct{a: 100, er: &intValue}

	//Since this all call by value data types, So passing all these data types as reference to works as pass by reference

	p("before-------", intValue, floatValue, strValue, boolValue, arrayValue, structValue)

	CheckCallByValueDataTypes(&intValue, &floatValue, &strValue, &boolValue, &arrayValue, &structValue)

	p("after-------", intValue, floatValue, strValue, boolValue, arrayValue, structValue, *structValue.er)

}

// Memory creation in function arguments - Always creates new different memory for all function arguments even for pointer arguments, SomeHow these different pointer memory address is linked with actual pointer memory address for acheiving pass by reference effects.

func CheckCallByValueDataTypes(intValue *int, floatValue *float64, strValue *string, boolValue *bool, arrayValue *[3]int, structValue *SampleStruct) {
	*intValue = *intValue * *intValue
	*floatValue = 10.5
	*strValue = *strValue + "hello"

	if *boolValue {
		*boolValue = false
	}

	//This way to dereference in both array and slices
	for i := 0; i < len(*arrayValue); i++ {
		(*arrayValue)[i] = i + 10
	}

	//This way to dereference in both array and slices
	(*arrayValue)[0] = 2

	structValue.a = 200
	//dereference struct's pointer value
	*structValue.er = 250

}

func CallByReferenceDataTypes() {
	sliceValue := []int{10, 20, 30}
	mapValue := map[int]int{1: 11, 2: 22}

	// channelValue := make(chan int)
	byteSlice := []byte{'a', 'b', 'c'}

	p("before reference datatypes------", sliceValue, mapValue, byteSlice)

	CheckCallByReference(sliceValue, mapValue, byteSlice)

	p("After---------", sliceValue, mapValue, byteSlice)

}

func CheckCallByReference(sliceValue []int, mapValue map[int]int, byteSlice []byte) {

	sliceValue[0] = 11 //Slice value assigns works as pass by reference

	sliceValue = append(sliceValue, 40) //Slice append() not works as pass by reference,Works same for all data types slice. Refer google docs notes

	mapValue[3] = 33
	byteSlice[2] = 'd'

}

func StructCallByValueReferenceReceiver() {
	o := firstStruct{intValue: 10, slicevalue: []string{"a", "b"}, mapvalue: map[string]string{"c": "cc"}}
	o.receiver()
	p(o)

	o = firstStruct{intValue: 10, strValue: "sat"}
	o.PointerReceiver()
	p(o)

	CallReferenceStructPointer(&o)
	p(o)
}

type firstStruct struct {
	intValue   int
	strValue   string
	slicevalue []string
	mapvalue   map[string]string
}

func (o firstStruct) receiver() {
	//struct receiver(non-pointer) function
	o.intValue = 11 //int value call by value type default, So this value will not be reflect outside

	//both slice and map are default call by reference types, So these value will reflect
	o.slicevalue[0] = "z"
	o.mapvalue["d"] = "dd"

	//We know, slice append function will not work as call by reference check notes and GolangsinglefileRecollect.go
	o.slicevalue = append(o.slicevalue, "p")
}

func (o *firstStruct) PointerReceiver() { //Always advised to use pointerReceiver as much as possible
	//struct receiver(pointer) function, So it works like call by reference defaulty for all data types even for default call by value data types.
	o.intValue = 20
	o.strValue = "hello"
}

func CallReferenceStructPointer(structValue *firstStruct) {
	//Struct passed as pointer, So below call by value data types works as call by reference
	structValue.intValue = 100
	structValue.strValue = "Nice"
}

type SampleStruct struct {
	a  int
	er *int
	v  string
}

func anonymousInlineFunction() {
	//inline function, Don't need to pass the local variables, they have the scope automatically
	var data int
	data = 3
	func() {
		fmt.Println("have the scope for the local variable", data)
		data++
		ScopeOnlyToThisInlineFunction := 10 //this variable don't have access/scope to outside this inline function func()
		ScopeOnlyToThisInlineFunction++
	}()
	fmt.Println("latest value", data)

	//closure in golang -- anonymous function can be assigned to variable, Then call that anonymous function using that closure variable

	//closure helps to get the returned value

	//closure - We can pass anonymous function as argument(callback function) into another function

	closureVariable := func() int {
		p("anonymouns function print statement")
		return 1000
	} //Shouldn't use () here, It will be used with closure variable

	returnedValue := closureVariable() //calls the above anonymous function in one or multiple times
	returnedValue = closureVariable()
	returnedValue = closureVariable()
	fmt.Println(returnedValue)

}

func SliceDatatype() {
	er1 := []int{1, 2, 3}
	er1 = append(er1, 20, 40, 60) //we can append multiple values also

	//clearing values in slice
	t1 := []int{1, 2, 3, 4, 5}
	t1 = []int{}           // t = nil --> for returning nil
	fmt.Println(t1 == nil) // [], false

	SliceMakeFunction()

	sliceShallowDeepCopyDefault()
}

func SliceMakeFunction() {

	//working of make() with slice

	// sliceCap1 := make([]int, 5)  // Capacity value is optional, Not given here, So capacity value is automatically equal to given length.
	//Created slice with length of 5, This create the zero values of length like array type[0,0,0,0,0]
	//So we can directly assign values sliceCap1[1] = 10
	//Here Appending values will place it as 6th value EX: [0,0,0,0,0,4]

	//Slice size/length and capacity:
	//Current capacity is doubled, once the slice size is exceeded the capacity.
	//Capacity can be equal or greater than size/length of slice

	sliceCap2 := make([]int, 0) //This works same as normal slice declaration EX: sliceCap2 := []int{}
	sliceCap2 = append(sliceCap2, 18, 19, 20)
	p(sliceCap2, len(sliceCap2), cap(sliceCap2))

	sliceCap := make([]int, 5, 10)
	sliceCap[0] = 11
	p(sliceCap, len(sliceCap), cap(sliceCap))
	sliceCap = append(sliceCap, 1, 2, 3, 4, 5, 6) //Once the given length of slice is exceeded,The overall capacity will be doubled like Internally The given capacity of the new underlying array will be created.
	p(sliceCap, len(sliceCap), cap(sliceCap))

	//normal empty slice declaration -- If capacity not mentioned, capacity is same as size.
	slice := []int{}
	for i := 0; i < 20; i++ {
		fmt.Println(len(slice), cap(slice))
		slice = append(slice, i)
	}
	/* Size and capacity increasing for each value appending flow
	0 0
	1 1
	2 2
	3 4
	4 4
	5 8
	6 8
	7 8
	8 8
	9 16
	10 16
	11 16
	12 16
	13 16
	14 16
	15 16
	16 16
	17 32
	18 32
	19 32
	last 20 32
	*/
	fmt.Println("last size and capacity", len(slice), cap(slice))

	//normal non-empty slice declaration
	gh1 := []int{1, 2, 3, 4, 5}
	p(len(gh1), cap(gh1)) //5,5

	gh1 = append(gh1, 1, 2)
	p(len(gh1), cap(gh1)) //7,10  -- Current capacity is doubled, once the slice size is exceeded the capacity.

	gh1 = append(gh1, 1, 2)
	p(len(gh1), cap(gh1)) //9,10  -- Here still capacity is 10, Because we appeneded just two more values
}

func arrayDataTypeFunction() {
	//can compare only arrays not slice
	er := [3]int{1, 2, 3}
	er2 := [3]int{1, 2, 3}
	if er == er2 {
		fmt.Println("arrays are equal")
	}

	//2-D dimensional array
	array_2 := [5][5]int{} //fixed sized 2-D array, Slice array ex: [][]int{}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			array_2[i][j] = j
		}
	}
	fmt.Println(array_2) //[[0 1 2 3 4] [0 1 2 3 4] [0 1 2 3 4] [0 1 2 3 4] [0 1 2 3 4]]
}

func anonmyousStructFunction() {

	//anonmyous Struct means struct with no name, Similar to anonmyous function func(){}, We can create and use this within a function EX: decode and encode jsons, testing package test inputs.
	anonmyousStruct := struct {
		a int
		b string
	}{10, "Sat"}

	anonmyousStruct1 := []struct {
		a int
		b string
	}{{10, "Sat"}, {11, "Sathoo"}}

	fmt.Println("anonmyousStruct------", anonmyousStruct.a, anonmyousStruct.b, anonmyousStruct1)
}

func channelFunction() {

	fg := make(chan map[int]int)
	fg1 := make(chan []int)
	fg2 := make(chan interface{})
	fmt.Println("different datatype channels", fg, fg1, fg2)

	//cap()   -- Returns the capacity of unbuffered and buffered channel
	p := make(chan int, 10) //returns 10
	// p := make(chan int)     //unbuffered channel -- returns 0
	fmt.Println(cap(p))
}

func forLoopExamples() {

	for i := 0; i < 5; i++ {
		p("Normal for loop", i)
	}

	for i := 0; i < 5; {
		p("control the increasing or decreasing iteration loop", i)
		i = i + 2
		i = i - 1
	}

	// for {
	// 	p("infinite while loop")
	// }

	j := 0
	for j < 5 {
		p("while loop", j)
		j = j + 1
	}

	d := "hello"
	for index, value := range d {
		p("Range for loop - using in string, all slices, all map upto the value's length", index, value)
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		p("Printing only odd numbers using continue", i)
	}

	for i := 0; i < 10; i++ {
		if i == 9 {
			fmt.Println("breaking the loop using break")
			break
		}
	}

}

func printfExamples() {

	//different printf representation

	// %s - For string
	// %d - For int

	// %T - Data type of variable

	// %p - memory address of variable, Can use "pointer" or "&" prefix with variable
	var r *int
	f := 100
	slice := []int{1, 2, 3}
	fmt.Printf("%p %p %p %p", &f, r, &slice, &slice[1])

	fmt.Printf("%s %s %T %v %p", "saa", "34", "dfg", 134, &SampleStruct{})

	sa := sample{10, "SSS"}
	fmt.Printf("%v", sa)  //{10 SSS}  //%v -- Useful for print for all data types EX: slices
	fmt.Printf("%+v", sa) //{a:10 b:SSS}
	fmt.Printf("%#v", sa) //main.sample{a:10, b:"SSS"}

	fmt.Printf("Now you have %f problems.\n", math.Sqrt(7)) //2.645751 -- float value
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7)) //2.6457513110645907 -- decimal value bigger precision than float value
}

func switchCaseExaamples() {

	n2 := 12
	switch n2 { //using switch expression, //switch case with declared integer type "n2", So all switch cases only support integer value

	case 11: //so here we cannot use logical/conditional expressions like ==, >=, &&, || etc. So can use only integer type case due to integer type switch expression .
		p("not equal")
	case 12:
		p("its equal")
	}

	//ASCII values categoring using switch case without any switch expression - (Ex: switch {})
	n1 := 12
	switch { // not defining any expression in switch statement(Ex: switch {}), So we can use true or false data type with logical/conditional expressions like ==, >=, &&, || etc
	case n1 == 32:
		p("space value")
	case n1 >= 48 && n1 <= 57:
		p("numbers 0 - 9")
	case n1 >= 65 && n1 <= 90:
		p("Capital alphabets")
	case n1 >= 97 && n1 <= 122:
		p("Non-Capital alphabets")
	case true:
		p("true case")
	default:
		p("other conditions")
	}

	//Switch case - Type assertion access data type from interface{}
	var T interface{}
	T = 10
	T = "stringValue"
	T = sample{}
	T = []int{1, 2, 3, 4}
	T = map[int]int{1: 11, 2: 22}
	switch T.(type) { //Type assertion access data type from interface{}
	case int:
		p("int type")
	case string:
		p("string type")
	case []int:
		p("int slice type")
	case map[int]int:
		p("map type")
	case sample:
		p("sample struct type")
	}

	//fallthrough in switch case
	t1 := 0
	switch t1 {
	case 0:
		p("zero case")
		fallthrough //If code entered this case, Fallthrough statements runs next below case too, Ex: it runs case 1 too.
	case 1:
		p("one case")
	default:
		p("default case")
	}

	switch 1 {
	case 1:
		break //Can have break in switch cases
		fmt.Println("i printed")
	}

}

func labelsInLoop() {

	num := 10
LabelControl0:
	fmt.Println(num)
	num--
	if num > 0 {
		goto LabelControl0 //this works like loop and will print 10 times
	}

LabelControl:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				p("breaked the outer for loop using Labels, Useful in While can't/don't want to return from function")
				break LabelControl
			}
		}

	}

LabelControl1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				p("continue label control in both inner and outer for loop")
				continue LabelControl1 //This "continue" works on both inner and outer for loop, In normal "continue" that works only on inner loop.
			}
			p(i, j)
		}

	}

}

//Different GenericFunction calls with predined argument types

// GenericFunction([]int{5, 6, 7, 8})
// GenericFunction([]string{"a", "b", "c", "d"})
// GenericFunction([]sample{sample{a: 1, b: "11"}, sample{a: 2, b: "22"}})

func GenericFunction[genericType int | string | sample](s []genericType) {
	for index, value := range s {
		p("GenericFunction", index, value)
	}

	// 	We can make a function generic to different defined data types.
	// Empty Interface type simply takes all data types and have to type assert everytime EX: i.(int), But generics have more control on the input parameter values and less chances of unexpected data types with empty Interface and don't have to do type asserting.

	//Can't use switch statements as well
	//So use the logic that is common for all specified arguments data types, EX: s = append(s,10) - This is not allowed, This supports only int type slice.

}

func FirstClassFunctions() {

	//i)Functions assigned to the variable on below methods.

	pr := fmt.Println //fmt.Println is one most popular use case for this
	pr("Functions assigned to the variable, Using that variables we can call the functions")

	//ii) Passing function as a argument to another function - This type of function also called as “higher order function or callback function” - Used to reuse the common function for handling multiple functions like below subsitue()

	sub1 := Substitue(sampleAdd, 10, 10)
	sub2 := Substitue(sampleMultiply, 10, 10)
	pr(sub1, sub2) //20,100

	//Sending multiple functions as arguments
	functionsExecuteOneByOne(sampleSubstract, sampleSubstract1)

	//iii) Function returns from another function - Can be useful in situations, We handles some pre-validation logics and after that returning from another function

	pr(PreValidationFunction(40, 40, 40))

}

func PreValidationFunction(a, b, c int) int {
	if a < 10 || b < 10 { //This some pre-validation logic before returning from below function
		return -1
	}
	return sampleAdd(a, b)
}

func functionsExecuteOneByOne(functions ...func()) { //handles as variadic functions, We can handle normal function argument as well

	for _, function := range functions {
		function()
	}
}

func sampleSubstract() {
	p("From sampleSubstract")
}

func sampleSubstract1() {
	p("From sampleSubstract1")
}

func Substitue(functionArg func(a int, b int) int, c int, d int) int { //func(a int, b int) int --> function signature like arguments and return type of sampleAdd(), sampleMultiply()

	return functionArg(c, d) //This actually calls the sampleAdd(c,d), sampleMultiply(c,d)
}

func sampleAdd(a int, b int) int {
	return a + b
}

func sampleMultiply(a int, b int) int {
	return a * b
}

func CustomInDataType() {
	d := stringCustomType("sathish")
	p("adding the custom string length", d.customLogic())

	f := sampleString{"a", "b", "c"}
	p("adding the custom string slice length", f.customLogic())
}

// Can create type with inbuilt data type like string, But Currently I don't see any unique need for this.
type stringCustomType string

func (s stringCustomType) customLogic() int {
	return len(s) + 2 //EX: s + "aaa" -- >"stringCustomType" even Don't supports "+", Because "+" normally supports string data type.

}

// inbuilt string struct type for custom functions
type sampleString []string

func (s sampleString) customLogic() int {
	return len(s) + 2
}

type sample struct {
	a int
	b string
}

type Student1 struct {
	Name  string
	Score int
}

func ConvertStructIntoMapViceVersa() {

	//We can use struct itself as key in map
	// map1 := map[Student]interface{}{}
	// map1[Student{"sat", 10}] = "sathish"
	// map1[Student{"sat1", 10}] = "sathish1"
	// fmt.Println(map1)

	s1 := Student1{Name: "sat1", Score: 100}
	var map1 map[string]interface{}
	byteData, err := json.Marshal(s1) //Above "Student1" struct fields should be exported, Otherwise Marshal and unmarshal can't access these unexported fields and so marshal and unmarshal not work.
	fmt.Println(err)
	err = json.Unmarshal(byteData, &map1)
	fmt.Println(err, map1)

	//converts map into slice of struct -- We can use interface{} data type in both map and struct for dynamic data types handling
	m := map[string]int{"1": 11, "2": 22}
	structSlice := []Student1{}
	for key, value := range m {
		tmp := Student1{key, value}
		structSlice = append(structSlice, tmp)

	}
	fmt.Println(structSlice)

}
