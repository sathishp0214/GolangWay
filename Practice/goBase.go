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


interpreted vs compiled programming language execution process and performance:

Compiled - Compiling whole Code converted/compiled directly into machine binary code(0’s and 1’s) single time, so that the OS processor can understand and execute. So compiled language removes the interpreter middle man and there is no need to transalate code into OS processor for undertanding, So it increases the performance. EX: two native persons speaks directly in there native language, So communication went efficiently.

Interpreted languages are also called scripting languages, because they use an interpreter to translate code line by line at runtime. (Like movie scripts happens scene by scene)
Interpreted - In run time interpreting, virtual machine changes the actual code line by line into Some other format (ex: opcode(or) 0’s and 1’s  in python ), then the OS processor understands and executes code.
(For each line, again and again It executes and converts that line code into binary code 0’s and 1’s, So it executes slower than compiled languages)
Even in a for loop with 100 iterations, the interpreted freshly translates the code for each iteration again and again.

python's .pyc is bytecode(Which is more compact and optimized representation of the Python code) that's not equalent to the compiled language's compiled binary file.

Choosing a interpreted vs dynamic programming language for application/task development:

i)current and future requirements interms of developer's current knowledge and developer's confidence/comfort/speed of doing in a programming language or programming language's frameworks, packages support EX: pandas, numpy etc, inbuilt features like decorators, inheritance, exceptions etc is very important. (if dynamic language is preferred here, then choose dynamic language otherwise choose compiled language)


https://stackoverflow.com/questions/38491212/difference-between-compiled-and-interpreted-languages

Hash Table data structure:
A Hash table is defined as a data structure used to insert, look up/read, and remove key-value pairs quickly.

Both python's dictonary and Golang's Map both uses the hash table data structure internally.

Hash table uses O(1) for both data inserting and searching/reading. Array uses O(log n) for data searching/reading, So hash table is more faster.


Golang supports Self-Contained Binaries -

A self-contained binary is a single executable(.exe) file that includes:
Your code
The Go runtime
All dependencies (standard and third-party packages) EX: packages like net, Gin, mux

Once we builded the golang compiled binary file using "go build" with expected executing OS architecture like linux or windows, Then We don't need any dependenies source code files/libraries and go.mod/go.sum etc. we can execute the golang compiled binary file directly on any machines like any other .exe executable file.


Python vs Golang:

Golang doesn't support exceptions,decorator,inheritance,oop constructor like python.
Golang anonymous is more powerful than python anonymous functions. python anonymous functions supports only single line statements and automatically return it.
Python doesn't supports pointers, defer statements. Python supports exception's "finally" or "with" statement instead of "defer".
currently Python has more number of package support and community support than golang.
Golang production compiled code is more faster than python production code EX: compare normal same "hello world" program to same higher algorithm programs.
Golang has inbuilt and more efficent concurrency like GoRoutine with channels support. Python multithreading is not light weight and not effcient as Golang and have to use "multithreading" package.
Python thread is preferred for I/O bound tasks only, Where python multiprocessing is preferred for CPU-bound tasks, Where GoROutine is fine for both I/O bound and CPU bound tasks, So Routines don't need multiprocessing seperately.
Python is definitely preferred in Data analytics/Data scientist/Machine learning areas and Web application development beacuse of more mature with more inbuilt features and frameworks like django and flask.

golang preferred development:
Golang is generally preferred in development of high performance required real time applications, microservices, high concurrency application, Due to golangs's single binary file which is useful in developing containerization applications which uses docker and kubernetes, devops tools development like docker, kubernetes, cloud computing, system programming like direct OS hardware applications like drivers.

golang currently non-preferred development so far:
Mobile app development
GUI application
Machine learning


Golang Top features:
Inbuilt concurrency support
Supports first class functions.
Powerful anonymous function support.
Defer
Panic and recover.


Golang doesn't support exceptions reasons by Golang creators - golang wants to handle each error instead of multiple errors with single exception scope. And due to the design contrainst of implementing exceptions in golang.

go doesn't support ternary '?:' operator
*/

/*

/*
--------Go environment:


Package - (directory/folder of go file(s)).
Module -- If a package has a go.mod file, Then it is considered as module.
multiple packages can also have a common go.mod file.


//create go.mod file under directory
go mod init directory_name/folder_name

go.mod keeps the list of import packages(our own go packages and third party go packages) in the particular package(s).
go.mod - For installing/updating/uninstalling packages will be automatically updated in go.mod. (Like python pip requirements.txt)


Running a go file -- go run file.go:
package main  --- Package should be the "main"

func main() {   --- Function should be the "main"

}

both package and function should be "main" for executing. It is considered an entrypoint and this main() is called first, from here we can call other functions.

init()
All declared init(), Called only once before main(), When we running particular package. We can use for load initial values and initializing client connections etc.

Golang package structure:

Simple structure:
repoWorkspace/main.go
	  		  app.go (optional)
	  		  go.mod
	  		  go.sum



intermediate/large project:
repoWorkspace/
			cmd/myapp/main.go
					  app.go (optional)
			internal/data/data.go
					/businessLogic/logic.go
					/db/mongo.db
					/model/model.go
			go.mod
			go.sum

cmd package - Go convention to put main.go for the application's entrypoint.
internal package - Go convention to avoid outside projects/workspace to import files under internal package.



//Download other 3rd party package/repo EX: pip install
go get Outside_package_name
go get -u github.com/spf13/cobra   //upgrade to latest version
go get github.com/spf13/cobra@v1.5.1 //particular package version
go get github.com/spf13/cobra@4rf356 //package particular git commit

go mod download  – this download all listed packages in go.mod file and stores in GOPATH/pkg/mod

go.mod file:

require (
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/cobra v1.2.1
	github.com/spf13/pflag v1.0.5 // indirect
)

// indirect  --- These "indirect" -- this mousetrap package we are not downloaded. This package is used by our other installed package "cobra". So we use this package indirectly.


go mod tidy: (syncs the package information in go.mod and go.sum file automatically and download and installs packages)
If we added/remove an import package in a go file. Then run "go mod tidy", it will add/remove package informations in go.mod and go.sum

go.sum: Stores the package checksum informations. (Checksum means generally - like hash id - creates id for current software package, if its modified or corrupted from checksum id we can compare and find it)

go.mod and go.sum file works together. "go mod tidy" -- syncs the both go.mod and go.sum files.Mostly we are works directly with go.sum file

For a single package from go.mod -- In go.sum, keeping that package's single or multiple versions hash checksum.

Example:
github.com/eapache/go-resiliency v1.1.0/go.mod h1:kFI+JgMyC7bLPUVY133qvEBtVayf5mFgVsvEsIPBvNs=
github.com/eapache/go-resiliency v1.3.0 h1:RRL0nge+cWGlxXbUzJ7yMcq6w2XBEr19dCN6HECGaT0=
github.com/eapache/go-resiliency v1.3.0/go.mod h1:5yPzW0MIvSe0JDsv0v+DvcjEv2FyD6iZYSs1ZI+iQho=


GOROOT  -- Path for Go language's OS installation and its inbuilts files/packages. In linux $GOROOT mostly -- /usr/local/go

GOPATH -- Its for Go project path, Where we have project's source code, project's 3rd party packages source code, cache etc.
Default GOPATH linux -- \home\go (or) \home\user_account\go

GOPATH's folders:

GOPATH/src -- contains project source/repo code

GOPATH/pkg/mod -- In mod directory, stores the project's go.mod packages files.

go clean - Used to remove already builded binary file and internal temporary files that can be deleted

go clean -modcache   – this used to remove all go.mod packages in GOPATH/pkg/mod

go get vs go get -u vs go get -d:

go get -- download and install go package

go get -d -- only download the package

go get -u -- download and install the update of that go package

Removing a go package:
go mod tidy  //if you remove a package in all imports in go files, while running this command, deletes that package automatically.

go get package@none

go get vs go run vs go build vs go install:

go get -- download and install package and maintain dependency with go.mod file. Stores packages in GOPATH/pkg

go run -- single step of compiling and executing go file. It will not store any binary file. (its for local development env purpose)

go build -- creates executable compiled binary file in the same directory. If any change in code, again I have to do the go build. Overall avoids repeated compiling time. (its for production environment purpose)
To Run a go builded file in terminal -- ./compiled_file

go install -- It installs the go package in the system's level. we can access from cli like tool/application ex- "python" code.py


Go build/install useful for production environments

-----Golang package helper:

go doc strings   //This terminal command list all functions in strings function with function arguments and return types

go doc strings contains  //Gives the brief instruction for “contains” function signature in strings package



Gets a function’s arguments and return type informations
// t := reflect.TypeOf(appendToSliceMoreEfficentMethod)
// t := reflect.TypeOf(defaultMemoryBytesSizeDataTypes) //func()
t := reflect.TypeOf(strings.Contains) //func(string, string) bool
fmt.Println(t)

Go inbuilt core tools:

Go vet:

go vet package/singleFile.go  //Gives warnings in code, Still code will be compiled and executed.
Ex: Unreachable code - Code after the return statements, break statements

go vet main.go

game_version := 3
EX: fmt.Printf("Super Mario %s\n",game_version)
./main.go:6:2: Printf format %s has arg 3 of wrong type int    //go vet response wrong data type value passed in printf

Goimports  - Code formatting, Auto package import/unimport on file

golint - Gives opinion about the code conventions like variable naming, comments added on top of function definitions etc


Go terminal help commands:
go help   //list go help commands

go help run  //detailed info of each command
go help build
go help doc
go help vet

go doc -help   // -help will give detail info of other set of commands
goimports -help  //Manually do go imports on pacakge/file


-----go vendor:

can have vendor folder in same application’s repository directory
go mod vendor - This create/update the latest go.mod packages in the vendor folder.

Vendor folder has all required current go.mod packages.
Vendor folder contains go.mod packages backup which is copying from the GOPATH/go/mod/.
To have to manually update the go vendor everytime we add or update or delete the packages in go.mod.

Uses:

Maintain particular extact package versions for stable build or to avoid conflict issues with other packages or environment.
can do faster build during the CI pipeline stages because of don't need to download package everytime.
Reproducable builds with the extact package versions because we maintains the same package and its versions through go vendor.
Cases we can reduce the docker image size by reusing the vendor folder of packages and faster builds.

However still go.mod way of automatically managing packages withoug go vendor is advisable in the production environment.


how to set application should use go vendor packages instead of GOPATH/pkg/mod packages from go.mod file?

By Using Vendor Mode:

To instruct Go to use the vendor directory instead of the go.mod GOPATH/pkg/mod directory:

i)Set the GOFLAGS environment variable:
export GOFLAGS=-mod=vendor

ii)Alternatively, you can pass the -mod=vendor flag directly when running your Go commands:

go build -mod=vendor
go test -mod=vendor
go run -mod=vendor main.go
/*

----Need/uses of pointers:
In golang, Pointer has the same fixed memory size of 8 bytes for even bigger data types like structs with multiple fields.
The size of a pointer variable is 8-bytes for 64-bit machines and 4-bytes for 32-bit machines.

Can do "Pass as reference" in function - So modifications happening inside the function reflects outside also. (If we do "Pass as Value" in function, Will create a new memory/copy for that variable and then process.)
So overall its memory efficient.

Empty pointer has nil value, So we can use pointers for validating by nil value.

using pointers for large structs.

----Golang doesn't support function overloading, method overloading.

-----Abstraction:

-- means not a physical one, just a representation of a physical one EX: All thoughts we are getting are abstraction of real/physical things.

Abstraction in golang can be achieved by interface with abstract method declarations.

As above mentioned the abstraction name meaning --
type Sample interface {
	function()
} -- This is just represention of real/physical function(), the implementation of function() will be in other structs methods.

---------
Variable declaration  – Creating a variable — Var i int
Variable Assigning   – Assigns value to variable —  i:=10, var t = 20
Variable initializers   EX: i:=10, var i int, var i int = 100   – initial value of a variable.


-------golang formatting in fmt.printf():

fmt.Printf("%s %s %T %v %p", "saa", "34", "dfg", 134, &struct)

%s - string value
%d - int value
%v - prints any data type value dynamically -- This is more useful
%+v - prints structs with key:value format (useful in pointer structs also), remaining all same with "%v"
%T - prints data type -- int, string etc
%p - prints memory address, while passing value like this -- &slice, &struct etc


 ---dot (.) import:
 import . "fmt"     // We can directly use packages function without "fmt." -- Println("hello")

 -------“Blank identifier” underscore “_” also known as “Blank identifier” - To avoid declared and not used compilation error.
Used for unused variables mostly in for range loops and err return value from function
Used for unused import in a go file

-----Debugging in GO

GDB debugging in go:

//build all files in a package folder
go build -gcflags "-N -l" -o .
run in terminal "gdb compiled_file"

Then we can set breakpoints, run the program and prints variables/pointers and go routines.

delve (dlv) debugging in go: (This debugger only we uses in VSCode debugging for golang)

https://golang.cafe/blog/golang-debugging-with-delve.html
https://vtimothy.com/posts/debugging-goroutines/

dlv debug main.go   //This way we using this debugger with terminal

Then in the terminal, we can set breakpoints, prints variables/go routines and change variable values.
Delve is better and has more features than gdb debugging.



-----Golang Naming conventions and rules:
Local variable names can be start with lowercase - user, userName

Global and constant variables, function name, struct name and its fields and its methods depend upon exported (or) unexported, Can start with lowercase or uppercase.

func printEmployeeDetails(employeeID int, employeeName string) {}  - Function arguments can be start with lowercase

Error variables can be started as ErrValue, ErrMessage

Bool variables can be started as IsTrue, IsContains

Avoid repetitive Words like
widget.NewWidget -> widget.New
widget.NewWidgetWithName -> widget.NewWithName
db.LoadFromDatabase -> db.Load

All the go directory/package names and .go file names should be starting with a small letter.  (this will not affect the Exported/Unexported)
EX: github.com/go-chi/chi  // Here both package and filename are started in small letters.
99% of golang packages/filenames follow this only

Test file name format – filename_test.go

--------Normal variable vs short variable declaration usage:
Normal variable declaration - If don’t know the initial value of the variable
EX: Use “var a int” instead of  “a := 0”
Short Variable declaration - If know the initial value of the variable. EX: Height := 10

----Function types in go: (Don’t know any use case as of now)
Different functions have the same number of arguments and argument types and same number of return values and types.

-------Incomparable data types in GO: - Slice,Map
We can’t compare slices,maps. EX: slice1 == slice2
We can’t use these types as map keys.


-------Memory allocation in golang:

Both stack and heap memory are controlled by the GO Runtime in golang.

Stack memory(LIFO) in GO and broad terms – Used to store all function calls and complete the function calls in LIFO order and functions' memory like arguments,returns and function's local variables. Once a particular function call is completed, The same is removed in stack memory.

EX: func main() {
	log()
}

func log() {
	another_function()
}

In this above code stack LIFO order, First completing function is another_function() and then log() function completes,

Default maximum stack size in Linux os is 8MB. But in golang every function call including goRoutines takes a few kbs size in stack memory.

If the stack size memory exceeds, Generally “Stackoverflow” errors return.




------Heap memory in GO and broad terms: Used for handling dynamic memory and global variables/long term memory having scope beyond stack memory’s function call.

Heap memory using often data  - go Slice, Maps and Slice/map inside struct uses heap memory, new(), make() function variables uses heap memory in golang, function returning values, channel passing values etc
Still Go Runtime will decide accordingly and store other data in heap as well.



---------Golang Garbage Collection:

Go's inbuilt garbage collector (GC) uses the Tricolor Mark and Sweep algorithm

golang inbuilt Garbage collection is used in heap memory to free the unneeded memory(even python also has inbuilt garbage collection).

Stack memory is generally faster than heap memory.

Manually force/trigger the inbuilt garbage collection in a program:
runtime.GC()   //runtime package has this code
We can manually trigger, Once you are clear/delete the large data types like above mentioned “Heap memory using often data types”.

------When Garbage collection triggers automatically:
Golang auto garbage collection does memory cleaning periodically like every 2 minutes or Once heap memory reaches the auto predefined threshold limit, Assume if a particular heap memory variable is used and not required anymore inside a function. Then auto garbage collection cleans those heap memory variables.

We can turn off auto garbage collection entriely and also can put custom threshold values in environment variables EX: GOGC=off, GOGC=80, Defaulty GOGC=100


Memory Leak Generally - Code assigns the memory but not releases back that memory, So the memory keep on increasing and affects the performance and can crash the application. Even garbage collector can't releases it, because more chances the memory still referenced for use.

Memory Leak scenarios:
Global variable slice,Map - Keep on adding the values, But not removes it. (Add the expiration logic and periodically check it for deletion)
Functions thats keep on running infinitely and slices,maps inside it keep on storing like above scenerio - Should take care more for memory leak possibility.
GoRoutines leak - check this function GoRoutineLeakExample()

bottleneck means in generally - particular piece of code slows down the entire app because High CPU usage / slow response ex: huge computations, slow db connections and executions,


Profiling in GO and broad term:
Used to analyze the performance/efficiency to find out memory leaks/bottlenecks(particular code takes more time/memory) of a Go program/function execution.

CPU profiling - CPU usage and time to execute a program/function
Memory profiling - Memory usage to execute a program/function
Concurrency profiling - performance while multiple goRoutines are used.

Inbuilt go pprof packages - runtime/pprof , net/http/pprof and other packages available

Go inbuilt profiling for “testing” package testing functions:
go test -cpuprofile cpu.prof -memprofile mem.prof -bench .

https://hackernoon.com/go-the-complete-guide-to-profiling-your-code-h51r3waz

--------- package, import package, Init() and global variables and main() execution order:

EXample single main.go file execution order ---> package main ---> import packages -->constant variables --> global variables --> init() ---> main()


https://david-yappeter.medium.com/init-in-go-programming-31e2c2bc2371


------Unsafe pointer:
Generally pointer are more restricted and safer standards in golang like
i)can’t do mathematical/conditional/logical operations directly on pointers like &a + 2, &a>2 etc
ii)Can’t convert one pointer data type into another datatype

Using unsafe package, We can do operations without the above restriction and adds more flexibility on golang pointer on unsafe manner. Overall it's risky and not recommended.



------GOPRIVATE environmental variable - We can set our project module in GOPRIVATE, Then that module considers a private module not available for public users to download it from github, gitlab etc.


----CGO in golang: CGO helps to cross-compilation - Helps to use C language libraries from golang. The CGO_enabled option enables the CGO in our go environment, Defaultly its disabled CGO_enabled=0.

chi vs mux vs gin golang framework:
chi and mux are almost similar, Mux is not maintained and almost deprecated. Both chi and mux are good for routing activities for api's CRUD operations.
Gin has more features and is more popular than chi and Gin claimed themselves faster than Chi/mux frameworks.

Struct field tags:
name string `json:”name_of_user”`  //Example struct field name

From struct’s variable names, Used to set key(or)field names for json, xml,yaml, sql column fields,bson(mongodb document key name) data

Even we can use to validate more conditions EX: field is empty or not, valid email etc

*/

/*

Golang official documentation facts:
efficient compilation, efficient execution, ease of programming in single language

Go combines the ease of programming of an interpreted, dynamically typed language with the efficiency and safety of a statically typed, compiled language.

golang develped with the reference of multiple langunages like C,Pascal etc.

golang used for developing in docker, kubernetes, terraform, Google cloud, consul etc.

From golang code, We can call C language writtern packages/libaries using "cgo" and C++ language writtern packages/libaries using "SWIG"

Golang may lack some features, intentionally due to the golang's uncompromised compilation speed and affects other available features.

Golang normal map is not safe in race conditions, We Can use sync package's map for the same or lock/unlock mechanisms.

*/

/*
every golang own (or) 3rd party EX:sarama import package should contain the go.mod file. That go.mod file only has the dependency modules and module version information, That is mandatory for installing and maintaing/upgrade/downgrade the module and module's dependency modules from go.mod file.

circular dependency (or) cycle imports:

If two packages are imported on each other causes circular dependency compilation error. EX: Package A imports Package B and Package B imports Package A.
Its is a design issue. We should aware of this while designing and coding the package.

We can eliminate this issue by design the depency in the seperate package like below Examples:
A package depends/imports C
B package depends/imports C

A package depends/imports on B
B depends on C
D depends on A and B

golang var a int declaration - Defaulty it will be int64 incase of 64 bit machine or int32 incase of 32 bit machine

Go runtime compiler - determines the usage of stack and heap memory.

Golang goroutines and channels are refered from Communicating Sequential Processes(CSP).(Communicating Sequential Processes (CSP) is a formal language for describing the interactions in concurrent systems, which was proposed by Tony Hoare in 1978. It models concurrent processes as sequential programs that communicate through synchronized message passing via channels.)

Go supports both concurrency and parallelism.
If we set GOMAXPROCS=1, then it eliminates the parallelism.

Sometimes adding more CPUs can slow a program down. In practical terms, programs that spend more time synchronizing or communicating(I/O tasks) than doing useful computation may experience performance degradation when using multiple OS threads.

Python uses lock/unlock based techniques for read/write the common data with the multiple threads involved.
Python also has thread-safe data structure called queue.

Go compiler(GC) written in golang.

Golang is developed to match the performance of C language and Golang creators believes golang is very competitive interms of performance with other compiled programming languages. However they believes garbage collector is not fast enough and compiler is good and can be better.

Golang doesn't supports pointer arithmetic as per offical documentation. pointer arithmetic means doing arithmetic operations like increment/decrement/add/substract on pointers directly.
However In golang we can dereference the pointers and then only we can do the arithmethic operations on integer pointers.

Golang generics use case:

reverse
sorting
Find smallest/largest element in different slice types
Find average/standard deviation of different slice types
Compute union/intersection of different map types
Find shortest path in node/edge graph
Apply transformation function to slice/map, returning new slice/map

Read from a channel with a timeout
Combine two channels into a single channel
Call a list of functions in parallel, returning a slice of results
Call a list of functions, using a Context, return the result of the first function to finish, canceling and cleaning up extra goroutines



channels vs mutex lock/unlock use case among multiple goRoutines:

Channel should be used for sending data EX: distribute data, result data, signal values between the goRoutines.
Combination of value passing/receing from channels. EX: Also whenever we received the channel value, We can trigger another goRoutine call also.

lock/Unlock can be used in increment/decrement,boolean flag, dB read/write, slice/maps read/write.

lock/Unlock has better performance than channels overall.

Depends upon the simplicity of use case implementation, We can prefer either one of these.




https://go.dev/doc/faq
https://go.dev/doc/effective_go


security vulnerbilities sample examples in golang.
https://github.com/aws-samples/amazon-codeguru-golang-detectors/



Gin framework - https://gin-gonic.com/docs/examples/

system programming means in golang

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

	// InbuiltGolangFunctionsLatest()

	// GenericsInStructExample()

	StructCallByValueReferenceReceiver()

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

	//Pointer is also a data type, It holding the memory address for any data type. we can't modify the memory addres directly. So we have to dereference into value, then have to modify, EX: *pointer = 10, Then that particular data type EX: (int) nature is automatically came inside.

	//immutable data types

	//In simple Words - We can't directly modify the values with same memory EX: string[1]="a"

	//For modifying this We need to allocate new memory Example -- Either we need to change this string into any other data type like slice and then modify this (or) We can use the strings package's replace() funtion, Again this function use the new memory inside for modifying.

	// int - int8,int16,int32,int64
	// float - float32, float64,
	// string
	// bool
	// rune - int32
	// byte - int8
	// struct

	//another Value (or) New value assigning is possible for all data types

}

func PointerAndDoublePointers() {

	value := 10
	var singlePointer *int
	fmt.Println(singlePointer == nil) //This is true, Because we not assigned any other variable's memory address.In golang,Every pointer defaulty has its own memory address we can get that address by "&singlePointer" like all variables.
	singlePointer = &value
	fmt.Println("address of value variable", &value)
	fmt.Println("address of SinglePointer variable", &singlePointer)
	fmt.Println("value of SinglePointer variable", singlePointer) //This address of Value variable and SinglePointer value is same

	//doulePointer(this is rare) - Holds the address of the another pointer. That another pointer holds the address of the another variable.
	//doulePointer use case - If we have to update the pointer's memory address in function and works pointer as call by value. Refer this https://stackoverflow.com/questions/8768344/what-are-pointers-to-pointers-good-for
	var doublePointer **int
	doublePointer = &singlePointer

	fmt.Println("address of DoublePointer variable", &doublePointer)
	fmt.Println("value of DoublePointer variable", doublePointer) //similarly address of singlePointer and value of doublePointer is same.
	fmt.Println("deference value from doublePointer", **doublePointer)
}

func InbuiltGolangFunctionsLatest() {

	// make()  //USed to create slice, map, channel data types

	//new() - Used to create not-nil pointer for all data types
	Intpointer := new(int)
	SlicePointer := new([]int)
	structPointer := new(SampleStruct)
	// Intpointer1 := new(*int)  // If we try to pass pointer EX:*int in new(), It will create double pointer EX:**int

	p("new() created pointers", Intpointer, SlicePointer, structPointer)
	p("new() creates non-nil pointers", SlicePointer == nil) //false, new() creates non-nil pointers

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
	//copy() does the deep copy of slices.
	//copy Destination slice should be in same length of source slice, So we can create using below make() with length or normal slice creation EX: destinationSlice := []int{40, 41, 42, 43, 44}
	destinationSlice := make([]int, 5)
	copy(destinationSlice, sourceSlice)
	p(destinationSlice, sourceSlice)
	destinationSlice[0] = 111 //This does deep copy, Its vary from other slice normal copying sliceCOpying()
	p(destinationSlice, sourceSlice)

	//If the length of source and destination slice is different in copy()
	//If source slice has 5 values, Destination slice has 3 values or viceversa. copy() copies only the destination slice length.
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
	//go doesn't support reverse/backward slicing like python
	// Go doesn’t support slice stepping s[0:len(s):2]
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

	//Below modifications will reflect deep copy the values in any other slices, Because all other slices had append modifications on above, due to re-slicing memory address is updated.

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

	//Type assertion generally -  Type assertion is a programming technique where you tell the compiler to treat a variable as a specific type
	// in golang acheive Type assertion access using interface.
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
	//generally 1 byte = 8 bits
	//byte - uint8 -- 8 bites (1 byte in size) -- it works only upto ascii values (0-255)
	//rune -- int32 -- 32 bits --(4 bytes in size) It works ASCII and more broader unicode characters upto around 65000 values of characters

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

	//Below data types returns not nil
	// var a int  //default value 0
	// var b string //default value empty string
	// var c bool //default value false
	// var d1 SampleStruct //struct field's default values {0}

	//Below same data types as pointers returns nil
	var a *int
	var b *string
	var c *bool
	var d *SampleStruct

	if a == nil && b == nil && c == nil && d == nil {
		fmt.Println("check nil with pointer", a, b, c, d)
	}

	//declaring variables with "var" keyword with below data type returns nil, So needs to intialize and theh should use it for avoiding nil pointer errors.

	var e map[int]int
	var e1 []int
	var p interface{}
	var p1 chan int
	// var p3 chan interface{}  //allowed - can use interface as data in channel
	// var p2 chan *int  -- different type pointer declaration also allowed in channel, will handle &int data through channels

	if e == nil && e1 == nil && p == nil && p1 == nil { //condition passes true
		fmt.Println("declare variables with var keywords returns nil--", e, e1, p, p1)
	}

	//now intialized the values, Now it will not return nil
	e = map[int]int{1: 1, 2: 2}
	e1 = []int{}
	p = 100
	p1 = make(chan int)

	fmt.Println("initialzed values with var variables--", e, e1, p, p1)
}

func typeCasting() {
	//Type_Casting - Converts the variable from one data type to another data type.

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

	//This is Rune and Bytes type most used real use case.

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
	//defer works in "panic". (if "defer statements" comes only before the "panic statement")
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

	//multiple defer flow:

	// defer fmt.Println("three")
	// defer fmt.Println("two")
	// defer fmt.Println("one")

	// fmt.Println("main function body")

	// Output:
	// main function body
	// one
	// two
	// three

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

// Memory creation in function arguments - Always creates new different memory for all function arguments even for pointer arguments. SomeHow in design, these different pointer memory address is linked with actual pointer memory address for acheiving pass by reference effects.

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

	sliceValue = append(sliceValue, 40) //Slice append() not works as pass by reference because of re-slicing, Works same for all data types slice.

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

	//We know, slice append function will not work as same as call by reference, so this below value will not reflect
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
	//This functions also called as function literals.
	//inline function, Don't need to pass the local variables, they have the outer function scope automatically
	var data int
	data = 3
	func() {
		fmt.Println("have the scope for the local variable", data)
		data++
		ScopeOnlyToThisInlineFunction := 10 //this variable don't have access/scope to outside from this inline function func()
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

	//closureVariable - This variable currently holds the reference of the above lambda function. EX: p = fmt.Println

	returnedValue := closureVariable() //calls the above anonymous function in one or multiple times
	returnedValue = closureVariable()
	returnedValue = closureVariable()
	fmt.Println(returnedValue)

	//another pattern from the function closures: Lambda function as return type
	//https://www.calhoun.io/5-useful-ways-to-use-closures-in-go/
	//https://code101.medium.com/understanding-closures-in-go-encapsulating-state-and-behaviour-558ac3617671

	FunctionClosurePatternExample()

}

func FunctionClosurePatternExample() {
	counter1 := createCounter() //these counter1, counter2 has the references of the returned inline function from the createrCounter()
	counter2 := createCounter()

	//both these counters maintains the seperate state. We can use like python's generator. Still have some issue in understanding this pattern completely
	fmt.Println(counter1()) // Output: 1
	fmt.Println(counter1()) // Output: 2
	fmt.Println(counter2()) // Output: 1
	fmt.Println(counter2()) // Output: 2
}

func createCounter() func() int {
	count := 0
	increment := func() int {
		count++
		return count
	}
	return increment
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
	//Current capacity is doubled, once the current slice length is exceeded the capacity.
	//Capacity can be equal or greater than size/length of slice

	sl := make([]int, 20, 100)
	fmt.Println(sl, len(sl)) //[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] 20

	sliceCap2 := make([]int, 0) //This works same as normal slice declaration EX: sliceCap2 := []int{}
	sliceCap2 = append(sliceCap2, 18, 19, 20)
	p(sliceCap2, len(sliceCap2), cap(sliceCap2))

	sliceCap := make([]int, 5, 10)
	sliceCap[0] = 11
	p(sliceCap, len(sliceCap), cap(sliceCap))
	sliceCap = append(sliceCap, 1, 2, 3, 4, 5, 6) //Once the given length of slice is exceeded,The overall capacity will be doubled because Internally The given capacity of the new underlying array will be created.(-----this is called re-slicing)
	p(sliceCap, len(sliceCap), cap(sliceCap))

	//another example
	f := make([]int, 2, 5)
	fmt.Println(f)
	f = append(f, 10)
	f = append(f, 20)
	f = append(f, 30)
	fmt.Println(f, len(f), cap(f)) //[0 0 10 20 30] 5 5
	f = append(f, 40)
	fmt.Println(f, len(f), cap(f)) //0 0 10 20 30 40] 6 10

	//normal empty slice declaration -- If capacity not mentioned, capacity is same as slice length, Once the given length of slice is exceeded,The overall capacity will be doubled.
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

	//normal non-empty slice declaration with values
	gh1 := []int{1, 2, 3, 4, 5}
	p(len(gh1), cap(gh1)) //5,5

	gh1 = append(gh1, 1, 2)
	p(len(gh1), cap(gh1)) //7,10  -- Current capacity is doubled, once the slice size is exceeded the capacity.

	gh1 = append(gh1, 1, 2)
	p(len(gh1), cap(gh1)) //9,10  -- Here still capacity is 10, Because we appeneded just two more values

	/*
		Using the proper capacity in slice can give better performance - If we have the idea of the length of slice definetely or approximetely, You can use capacity accordingly.
		EX: If you are dealing with slice with large expected length, You can set capacity at intialization itself high.
		https://stackoverflow.com/questions/45423667/what-is-the-point-in-setting-a-slices-capacity
	*/
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
	// 	p("infinite while loop alike")
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

// Generics stack struct
type Stack[T any] struct {
	values []T
}

// NewStack creates a new empty stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{values: []T{}}
}

// Push adds an element to the top of the stack
func (s *Stack[T]) Push(value T) {
	s.values = append(s.values, value)
}

// Pop removes and returns the top element from the stack
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.values) == 0 {
		var zeroValue T
		fmt.Println("zerValue here", zeroValue)
		return zeroValue, false
	}
	value := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return value, true
}

func GenericsInStructExample() {
	// Stack of integers

	intStack := NewStack[int]()
	intStack.Pop()
	intStack.Push(10)
	intStack.Push(20)
	value, ok := intStack.Pop()
	fmt.Println("Popped from int stack:", intStack, "------", value, ok) // Output: Popped from int stack: 20 true

	// Stack of strings
	stringStack := NewStack[string]()
	stringStack.Push("Hello")
	stringStack.Push("World")
	value1, ok := stringStack.Pop()
	fmt.Println("Popped from string stack:", stringStack, "------", value1, ok) // Output: Popped from string stack: World true
}
