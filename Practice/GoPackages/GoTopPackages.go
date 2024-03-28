package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"os/signal"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"sync"
	"syscall"
	"time"

	"math/rand"

	"log"

	"github.com/google/go-cmp/cmp"
	"github.com/sirupsen/logrus"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func main() {

	//check this below function for http package
	// httpCheck()

	// sortPackage()

	// et := map[int]int{1: 111, 0: 000, 2: 222}
	// d := 0
	// d1 := 1
	// fmt.Println(et[d[0]] > et[d1[0]])

	simpleHtttpRequests()

	runtimePackage()

}

func OSPackage() {

	//import os --its a big package we can do most of OS and terminal operations

	// create/open/delete/move/copy file/folder
	//read and write file
	//change/set file/folder permissions
	//kill pid process
	//set and get environment variables
	//run OS terminal commands
	//gets into file/folder path

	//read and write file
	filedata := []byte("good hello this is text data added into the file new version data Go has excellent built-in support for file operations. Using the os package, you can easily open, read from, write to and close the file.In this example, we focus on writing data to a file. We show you how you can write text and binary data in different ways - entire data at once, line by line, as an array of bytes in a specific place, or in a buffered manner")
	err := os.WriteFile(filename, filedata, 0666) //clears existing file data and adds the data
	fmt.Println(err)

	readfile, err := os.ReadFile(filename)
	fmt.Println("read file data from OS package----------------", err, string(readfile))

	//By using this function, We can set different flags like create file, readonly,writeonly,readAndWrite,Append
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666) //os.O_CREATE - if file not there it creates new one
	file.WriteString("adding new data =======================================")
	readfile, err = os.ReadFile(filename)
	fmt.Println("appended new data into file from OS package----------------", err, string(readfile))

	//example for reading file
	file, err = os.OpenFile(filename, os.O_RDONLY, 0666)
	fmt.Println("Againing reading the data----------------", err, string(readfile))

	//Arguments can be passed while running go file
	//go run "d:\Golang_practice_personal\Interview\interview1.go" 20 10   -- passing arguments like this
	fmt.Println(os.Args)
	fmt.Println(len(os.Args))
	fmt.Println(os.Args[1])

	// fmt.Println(os.Hostname())
	// fmt.Println(os.Environ()) //environment variables

	// os.Getpid()  //current process's ID
	// os.Getppid() //current process's parent ID
	// os.Chmod() //changing particular file/folder mod
	// os.Chown() //changing particular file/folder ownership

	// os.Truncate()  //reduces the file in particular size or completely empty
	// os.Stat() //returns file is found or not and file's info

	// defer fmt.Println("defer statement executed")
	// os.Exit(100) //Due to the os.Exit(), Above defer statement will not execute
	// fmt.Println("normal statement not executes")

	// changeModeError := os.Chmod("D:/Documents/Golang_practice/practice/goPackages", 0777) //changing file mode
	// // changeModeError1 := file1.Chmod(777)  //changing current file mode in another method
	// fmt.Println(changeModeError)

	// execute basic os commands
	// cmd := exec.Command("hostname") //  "os/exec" package - Run raw os commands
	// out, err := cmd.Output()
	// fmt.Println(err, string(out))

	//creating os new/child process
	// startingProcess := exec.Command("sleep", "10") //running
	// startingProcess.Start()                        //starting the OS process
	// startingProcess.Wait()                         //wait the process to complete
	// startingProcess.Run()                   //combination of start() and Wait()
	// startingProcess.Cancel()
}

func cmpInbuiltPackage() {
	//import cmp
	fmt.Println(cmp.Equal([]int{2, 3, 4}, []int{2, 3, 4, 5})) //similar to reflect.DeepEqual, checks both values are same
	fmt.Println(cmp.Diff([]int{2, 3, 4}, []int{2, 3, 4, 5}))  //If any difference between it prints difference like github difference page like below
	// 	[]int{
	//         2,
	//         3,
	//         4,
	// +       5,
	//   }
}

func sortPackage() {
	//import sort
	// sort.Ints(arr)     //sorts integer slice
	// sort.Strings(arr1) //sorts string slice

	// sa := []int{1, 2, 3, 4, 5, 6}
	// fmt.Println(sort.IntsAreSorted(sa)) //If integer slice are already sorted in ascending
	// fmt.Println(sort.StringsAreSorted(arr1)) //If string slice are already sorted in ascending

	// //Descending order
	// sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	// fmt.Println("Descending order map_keys in int array", arr) //[9 8 7 3 1]

	// sort.Sort(sort.Reverse(sort.StringSlice(arr1)))
	// fmt.Println("Descending order map_keys in string array", arr1) //[hhhh aaa a PPPPP KKKK]

	Student := []struct {
		s_name  string
		s_marks int
		s_id    int
	}{
		{"Ranjan", 280, 1098},
		{"Ajay", 280, 300},
		{"Anita", 330, 109},
		{"Kavita", 100, 107},
		{"Ashima", 444, 258},
		{"Rohit", 450, 188},
		{"Vijay", 289, 118},
		{"Dhanush", 239, 329},
		{"Priya", 400, 123},
		{"Nikita", 312, 111},
	}

	sort.Slice(Student, func(i, j int) bool {
		return Student[i].s_marks > Student[j].s_marks //descending order for s_marks
	})

	sort.Slice(Student, func(i, j int) bool {
		return Student[i].s_marks < Student[j].s_marks //ascending order for s_marks
	})

	//sort.SliceStable() - this is same as sort.Slice(), we should use this this is stable function compare to sort.Slice()
	sort.SliceStable(Student, func(i, j int) bool {
		return Student[i].s_marks < Student[j].s_marks //ascending order for s_marks
	})

	//can use normal slice also for sorting and reverse sorting(this is more easy to remember).
	k := []int{3, 4, 2, 4, 6, 9}
	sort.SliceStable(k, func(i, j int) bool {
		return k[i] < k[j]
	})

	//we can't sort map by this function
	// et := map[int]int{1: 111, 0: 000, 2: 222}
	// sort.SliceStable(et, func(i, j int) bool {
	// 	return et[i] < et[j]
	// })

	fmt.Println("after sorting", Student, k)

	//finds a number in slice with index using sort.Search(), We can use in structs etc
	a := []int{1, 3, 10, 15, 21, 28, 36, 45, 55}
	xy := 6

	//uses binary search for searching
	i := sort.Search(len(a), func(i int) bool { return a[i] >= xy })
	if i < len(a) && a[i] == xy {
		fmt.Printf("found %d at index %d in %v\n", xy, i, a)
	} else {
		fmt.Printf("%d not found in %v\n", xy, a)
	}
}

func slicePackage() {
	f1 := []int{1, 2, 3, 4, 10, 101}

	//slices import package
	fmt.Println(slices.IsSorted(f1))
	fmt.Println(slices.Contains(f1, 20))
	fmt.Println(slices.Index(f1, 10)) //returns value index
	//slices.Delete(slice,delete_starting_index, delete_ending_index)
	deleted_Values_slice := slices.Delete(f1, 3, len(f1)) //deletes from 3rd index to end of index, It uses slicing like f1[3:len(f1)]
	fmt.Println(deleted_Values_slice, f1)                 //slices packages returns new slice, So not modifies on the source slice.

	fmt.Println("replacing", slices.Replace(f1, 3, len(f1), 2000, 2001, 2002)) //like above delete, it replaces the values(2000, 2001, 2002), in index f1[3:index]

	fmt.Println(slices.Insert(f1, 0, 2000)) //[2000 1 2 3 4 10 101] inserts 0th index

	// slices.Reverse(f1) //not returns new slice, reverses on source slice itself
	// fmt.Println(f1)

	//copy the slice
	copied_slice := slices.Clone(f1)
	fmt.Printf("%p %p", f1, copied_slice) //both seperate memory address
	fmt.Println(f1, copied_slice)

	fmt.Println(slices.Equal(f1, copied_slice)) //both slices are equal or not

	fmt.Println(slices.Min(f1))
	fmt.Println(slices.Max(f1))

	slices.Sort(f1) //ascending order
	fmt.Println(f1)
	// slices.Reverse(f1) //for decending order, after sorted we can use reverse()

	fmt.Println(slices.BinarySearch(f1, 101)) //5 true --> //binarySearch returns index of the value. For binarySearch generally input slice/array should be sorted.We can use this for faster operation.

	//slices packages Func inbuilt functions like SortFunc,ContainsFunc etc, This useful for doing custom operations/logics in those inbuilt func functions

	fmt.Println("slices contains function------------", slices.ContainsFunc(f1, func(a int) bool {
		fmt.Println("Inside containsFunc function every slice value", a)
		//func(a int) bool  -- every slice values is passed inside this function to a argument.

		// if a%2 == 0 {    //return false, if any slice value is even number
		// 	return false    //custom conditions
		// }

		return true //simply passing constant condition
	}))

	slices.SortFunc(f1, func(a int, b int) int {
		return 0
	})

	fmt.Println("First negative value's index", slices.IndexFunc([]int{0, 42, -10, 8}, func(n int) bool {
		fmt.Println("print values", n)
		return n < 0 //returns if any slice value is lesser than 0
	}))
}

func mapPackage() {
	//maps inbuilt package
	ap := map[int]string{1: "11", 2: "22"}
	fmt.Println("map keys", maps.Keys(ap))
	fmt.Println("map values", maps.Values(ap))

	ap1 := map[int]string{3: "33", 4: "44", 5: "55"}
	maps.Copy(ap1, ap)
	fmt.Println("after copying there source map into destination map", ap1) //map[1:11 2:22 3:33 4:44 5:55]  so its works like concating two maps together

	// maps.Clear()  //empties map
	// maps.Clone()   //shallow copy of a map
	// maps.Equal()  //checks two maps are equal

	// delete(ap,1)  //deletes particular map key and its value

}

func reflectPackage() {
	//reflect package
	var ep s1
	ep = "1000"

	w := sample{}
	w.a = 20
	w.b = "sat"
	fq := "dfdfdf"

	fmt.Println(reflect.TypeOf(fq)) //gets the data type

	fmt.Println(reflect.TypeOf(fq).Kind(), "compares the data type---", reflect.TypeOf(fq).Kind() == reflect.String) //get and compares the data type

	fmt.Println(reflect.TypeOf(w), "----", reflect.TypeOf(w).Kind()) //main.sample ---- struct
	//i)typeOf() - gets struct name main.sample, ii)kind() gets type -- struct

	//string --gets interface actual data type
	fmt.Println(reflect.TypeOf(ep).Kind())

	// fmt.Println(reflect.DeepEqual(value,value)) //compares two values are same like two strings, two slice,maps, two struct objects

	fmt.Println(reflect.DeepEqual("saa", "saa"))
	fmt.Println(reflect.DeepEqual([]int{2, 3, 4}, []int{2, 3, 4}))
	//
	// reflect.DeepEqual(map1, map2)

	//compare the struct's pointer variable and struct's non pointer variable also
	// c := &SampleStruct{}
	// d := &SampleStruct{}
	// println(reflect.DeepEqual(c, d))

	//Elem() used to gets the values from the struct pointer
	fmt.Println(reflect.ValueOf(&w).Elem())

	//Gets a function’s arguments and return type informations
	// t := reflect.TypeOf(appendToSliceMoreEfficentMethod)
	// t := reflect.TypeOf(defaultMemoryBytesSizeDataTypes) //func()
	t := reflect.TypeOf(strings.Contains) //func(string, string) bool
	fmt.Println(t)

	//Get only struct's all Exported functions(functions start with capital letter) name.
	t := reflect.TypeOf(&struct{})
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m.Name)
	}
}

func timePackage() {
	//time package
	fmt.Println(time.Now())        ///IST time default
	fmt.Println(time.Now().UTC())  //utc time
	fmt.Println(time.Now().Unix()) //1698000880 -- timestamp

	//time formatting --"02-01-2006 15:04:05" this is constant time, should used the same date and time for expected time format
	fmt.Println(time.Now().Format("02-01-2006 15:04:05")) //23-10-2023 00:52:27
	fmt.Println(time.Now().Format("01-02-2006 15:04:05")) //10-23-2023 00:52:27 (month/day/year format)\

	fmt.Println(time.Now().Format("02-01 15:04")) //23-10 00:52
	fmt.Println(time.Now().Format("02-01-2006"))  //23-10-2023
	fmt.Println(time.Now().Format("02-01"))       //23-10

	//converts string into time.Time object
	//"23-10-2013 13:10" string time
	timeObject, _ := time.Parse("02-01-2006 15:04", "23-10-2013 13:10")
	fmt.Println(timeObject, reflect.TypeOf(timeObject))

	//creates date time.Time value from int values
	//2023-10-23 00:00:00 +0000 UTC
	fmt.Println(time.Date(2023, 10, 23, 0, 0, 0, 0, time.UTC))

	fmt.Println(time.Now().Year())    //2023
	fmt.Println(time.Now().Month())   //october
	fmt.Println(time.Now().Weekday()) //monday -- which day
	fmt.Println(time.Now().Day())     //23 -- which day of month
	fmt.Println(time.Now().Hour())    //which hour of day
	fmt.Println(time.Now().Minute())
	fmt.Println(time.Now().Second())

	fmt.Println(time.Now().Clock()) // 0 55 39  hr/min/sec useful for integer operations

	fmt.Println(time.Now().Add(time.Hour))   //adds 1 hour from current time
	fmt.Println(time.Now().Add(time.Minute)) //adds 1 minute from current time
	fmt.Println(time.Now().AddDate(0, 0, 1)) //adds one day from current time
	fmt.Println(time.Now().AddDate(0, 1, 1)) //adds one month and one day from current time

	fmt.Println(time.Now().Add(-time.Hour))        //subtracts 1 hour from current time
	fmt.Println(time.Now().Add(-time.Minute * 30)) //substracts 30 minutes from current time
	fmt.Println(time.Now().AddDate(0, -1, -1))     //substracts one month and one day from current time

	timeObject1 := time.Now().AddDate(0, 0, -1) //subtracts 1 day time
	fmt.Println(time.Now().Sub(timeObject1))    //24h0m0s  -- difference between two times
	fmt.Println(time.Since(timeObject1))        //another easy way of getting the difference time

	//time ticking it runs like infinite for/while loop with time sleep
	// for tick := range time.Tick(2 * time.Second) {
	// 	fmt.Println("time ticking", tick, time.Now())
	// }

	// fmt.Println(time.Now().After()
	// time.Now().Before()
	//time.NewTimer()
}

func mathPackage() {
	//import math package
	math.Abs(-0.25)                                    //similar to python abs() removes negative sign
	fmt.Println(math.Max(10, 20))                      //20
	fmt.Println(math.Min(10, 20))                      //10
	fmt.Println(math.Pow(5, 3))                        //125
	fmt.Println(math.Sqrt(25))                         //5
	fmt.Println(math.Cbrt(125))                        //5  cube root
	fmt.Println(math.Round(121.609450476548765076056)) //122 ---> 121.60 > 121.5 so rounding off to 122
	fmt.Println(math.Ceil(4.01))                       // 5 -- here input greater than 4, So always returns 5
}

// math/rand packages
func RandomNumbersStrings() {
	//rand.Seed(time.Now().UnixNano())  //rand.Seed() -- We can update the seed value often EX: pass the current timestamp, So less chances for the same random number generation.

	// fmt.Println("Any Random int number", rand.Int())

	fmt.Println(rand.Float64()) //0.10421671051352199

	fmt.Println(rand.Intn(100)) //54 -- random number between 0 to 100

	//random any four digit number
	fmt.Println("random four digit number", rand.Intn(9)*1000+rand.Intn(999))

	fmt.Println(rand.Perm(20)) //prints 0 to 20 numbers in random/shuffle order -- [8 0 1 16 4 9 18 11 14 12 19 2 17 15 5 6 13 3 10 7]

	fg := []int{21, 13, 4, 7, 8, 10, 5}
	rand.Shuffle(len(fg), func(i, j int) {
		fg[i], fg[j] = fg[j], fg[i]
	})
	fmt.Println("shuffle slice", fg) //shuffling the input slice -- [13 4 21 5 7 10 8]

	fmt.Println("any random capital letter", string(65+rand.Intn(25))) //Capital letter ASCII starts from 65 then adding any random number from 0 to 25

	fmt.Println("any random non-capital letter", string(97+rand.Intn(25)))

	fmt.Println("random 4 digit capital letters", string(65+rand.Intn(25))+string(65+rand.Intn(25))+string(65+rand.Intn(25))+string(65+rand.Intn(25))) //can use for loop as well.

	// de := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// rand.Shuffle(len(de), func(i, j int) {
	// 	de[i], de[j] = de[j], de[i]
	// })
	// fmt.Println("shuffle slice", de[:4]) //IDEA : after shuffled you can expected random string length by slicing,

}

// bytes package -- Almost similar to strings package's all inbuilt functions
func bytesPackage() {

	fmt.Println(bytes.Count([]byte{'a', 'd', 'd', 'f'}, []byte{'d'}))
	fg := []byte{'a', 'd', 'd', 'f'}
	fg[0] = 'p'
	fmt.Println(fg, string(fg))

}

func runtimePackage() {
	// runtime.GC()          //manually forcing a garbage collection now

	// runtime.GOROOT()      //go root location
	// runtime.Version()     //go version

	// runtime.Goexit()  //Should place inside a goRoutine, It terminates that current go Routine only.

	d := runtime.MemStats{}          //returns the heap memory in bytes
	fmt.Println(d.Sys, d.TotalAlloc) //d.sys - Total system memory bytes available, d.TotalAlloc - total allocated heap memory bytes

	pr(runtime.NumCPU()) //number of system total threads, My personal laptop 1215U processor has 6 cores, 4 physical cores and 8 threads. So each physical core has two threads.

	// runtime.GOMAXPROCS(4) //Go runtime defaulty uses all available system threads - "runtime.NumCPU()". In this function GOMAXPROCS(), we set limit the thread usage number for go program.

	runtime.NumGoroutine() //Prints the number of current goRoutines running - main() function also goRoutine

	//These below lock and unlock goRoutine like python threads, Aware of these may useful - May Need to practice this
	// runtime.LockOSThread()
	// runtime.UnlockOSThread()

	//runtime/trace package ---> runtime packages also has the option to lots of trace options like, goRoutine traces, goRoutines blocking, heap memory allocations etc.

	runtimeDebugPAckage()

}

func runtimeDebugPAckage() {
	// debug.PrintStack() //prints the stack trace in standard output

	// debug.SetGCPercent() //sets the threshold limit, Which triggers the Garbage collector to work once given threshold reaches
	// debug.SetMaxStack()  //setting maximum memory a goRoutine can take.

	// debug.SetTraceback("all")  //prints detailed info about runtime errors and goROutine crashes etc
}

func contextPackage() {
	//contexts are used to pass request-scoped values, deadlines, and cancellation signals between apis/functions and Go-routines.

	//powerful package for managing concurrent operations

	//Main use cases:
	//cancellation signal of two or apis/functions and multiple/child goRoutines
	//Auto timeout with cancellation signal in two or  apis/functions and multiple/child goroutines
	//pass values in  apis/functions and goroutines

	//Context cancelling (<-ctx.Done()) vs channel dummy signal value(<-close) for closing the GoROutine:
	//Context is more easy for killing/closing Multiple GoRoutines.
	//killing/closing single GoRoutine - Both are almost similar

	//Below are the basic ways of creating a context package's object
	// context1:= context.Background()
	// context2 := context.TODO() //This function declaration normally used as temporary creation, Expected to be updated with better context's package functions in future.

	//cancellation signal of two or multiple goRoutines
	ctx, cancelSignal := context.WithCancel(context.Background())
	pr(ctx.Err(), cancelSignal)

	go func() {
		pr("sends cancel signal")
		cancelSignal()
	}()

	time.Sleep(1 * time.Second)
	select {
	case <-ctx.Done(): //This is triggered, When context's cancel signal function is called.
		pr("task done - ", ctx.Err())
	default:
		pr("default case")
	}

	// http.NewRequestWithContext(ctx, method, url, body)  //http request with context

	//Refer this functions with multiple GOROuitnes implementation
	// MultipleGoRoutinesWithContextCancel()
	// MultipleGoRoutinesWithContextCancelWithInbuiltTimeout()

	//Pass the key-value pair data using context
	ctx = context.TODO()
	ctx = context.WithValue(ctx, "name", "Sat")
	PassValueWithContext(ctx)

}

func PassValueWithContext(ctx context.Context) {
	fmt.Println(ctx.Value("name"))
}

func contextWithTimeout() {
	//timeout in goroutines two or multiple goRoutines - Similar to above withCancel(), Here we adding the timeout
	ctx1, cancelSignal1 := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelSignal1()

	go func() {
		time.Sleep(2 * time.Second)
		pr("inside goRoutine 1")
	}()

	pr(ctx1)

}

func LiveHeapMemoryGarbageCollectionOperation() {
	d := &runtime.MemStats{}
	runtime.ReadMemStats(d)
	pr("before large data-----------", d.HeapInuse) //gets current heap memory

	s := make([]int, 1000000000) //creating local big dynamic heap memory
	s[0] = 1
	runtime.ReadMemStats(d)
	pr("after large data -----------", d.HeapInuse)

	runtime.GC() //Manually running garbage collection.

	runtime.ReadMemStats(d)
	pr("heap memory cleared-----------", d.HeapInuse)

	/* WORKING OF INBUILT GARBAGE COLLECTION */
	//Assume, If not used large slice "s" again in this function, So our manual garbage collection clears that slice heap memory.
	pr(s[0]) //Here again uses that slice here, so manual garbage collection not clears that slice heap memory, Because we using this slice again, So again have to call "runtime.GC()" Once this slice not used anymore in this function for clearing the heap memory.

	runtime.ReadMemStats(d)
	pr("Final memory -----------", d.HeapInuse)
}

func bufioPackage() {
	//bufio : Buffered Input/Output --
	//bufio I/O -- It accumulates data upto the buffer size and then doing the flush(for performs acutal I/O operation) once. So it reduces the number of system calls and works efficiently overall.

	//Example bufio buffer write operation-- It accumulate(collects without directly write data on file), So it reduces the small data write system calls into files and works efficiently overall.

	//default buffer size is 4096K, Bufio can hold the data upto this capacity

	var lines = []string{
		"Go",
		"is",
		"the",
		"best",
		"programming",
		"language",
		"in",
		"the",
		"world",
	}

	//bufio write
	f, err := os.Create("bufiofile.txt")
	bufioWriter := bufio.NewWriter(f) //have to use the created file from "os" package, This also clears existing file data

	i, err := bufioWriter.Write(filedata)
	fmt.Println(i, err)
	for _, v := range lines {
		//can do each write as different data type
		bufioWriter.Write([]byte(v))  //write data as byte data type
		bufioWriter.WriteString("\n") //Adds new line //write data as string data type
	}
	err1 := bufioWriter.Flush() //After doing flush(), Then only it writes data on the file until buffer default capacity size 4096K.
	fmt.Println(err1)
	f.Close() //should close it, If we are going to read it immeditely

	//reads bufio
	readfile1, _ := os.Open("bufiofile.txt")
	bufioScanner := bufio.NewScanner(readfile1)

	for bufioScanner.Scan() {
		fmt.Println("read file data----------", bufioScanner.Text()) //reading file data line by line
	}
}

func LogDefaultBasicPackage() {
	file, err := os.OpenFile("logfile.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(file) //All the log messages will be appended in this file

	//If we not declared above setoutput file, Then it will print the messages in standard output as "fmt.Println"

	// log.Panic("panic error") //prints panic error with this message
	// log.Fatal("fatal error") //prints this line and program exits from this line itself
	log.Println("Log message 1")
	log.Println("Log message 2")
}

func LogrusPackage() {
	//Logrus package is one of top popular log package and well maintained with package updates.
	//We can update from default "log" package into logrus package by import as --> "log "github.com/sirupsen/logrus"""

	//Can generate with various log message category loke warning,info,error
	logrus.Warning("warning message")
	logrus.Error("error message")

	//Can generate key value json log data, USed in Rave
	logrus.WithFields(
		logrus.Fields{"key": "value for log",
			"key1": "value1 for log"}).Info("log information with json key:value pairs")

	logrus.WithFields(
		logrus.Fields{"key": "value for log",
			"key1": "value1 for log"}).Error("error information with json key:value pairs")
}

func flagPackage() {
	pr(flag.Args()) //will get the command line arguments - go run check.go hello sat -- "hello,sat" args values

	//runs flag command line arguments or default below flag values
	intValue := flag.Int("n", 5, "number of iterations") //go run check.go -n=3, If we didn't passed this -n=3 flag argument value, Then default flag value "n" as 5 is taken.
	pr(*intValue, reflect.TypeOf(intValue))              //We should redeference for value

	strValue := flag.String("flagarg", "sat", "Just giving name") //go run check.go -flagarg=sathish
	pr(*strValue)

	flag.Parse() //This will activate all above flag values declarations.

}

type s1 interface{}

type sample struct {
	a int
	b string
}

func httpCheck() {
	request, err := http.NewRequestWithContext(context.Background(), "GET", "https://httpbin.org/", nil)
	fmt.Println(request, err)
	if err != nil {
		fmt.Println(request.Body)
		return
	}
	client := &http.Client{}
	response, error := client.Do(request)
	fmt.Println(response, error)
	if error != nil {
		return
	}
	defer response.Body.Close()

	//prints the response
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))

	//stores into map of struct (or) slice of struct mostly
	var responseResult map[string]string
	decode := json.NewDecoder(response.Body)
	decode.Decode(responseResult) //storing values

	fmt.Println(responseResult)

}

func signalPackage() {
	//This signal works for unix/linux environments only. (Not sure about windows environment)

	signalChan := make(chan os.Signal, 1) //should be buffered channel with length of 1
	signal.Notify(signalChan)             //Passing the signal channel into signal

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		for {
			fmt.Println("signal listening for control+c")
			SignalReceived := <-signalChan        //Once signal is triggered, Signal value will be received in channel
			if SignalReceived == syscall.SIGINT { //Compares the signal for action
				fmt.Println("control+C signal received")
				wg.Done()
				os.Exit(1) //terminates this program
				return
			}

		}
	}()

	wg.Wait()
	fmt.Println("main thread completed")

}

func simpleHtttpRequests() {
	response, _ := http.Get("https://google.com")
	byteResponse, _ := io.ReadAll(response.Body)
	fmt.Println(string(byteResponse), "=========", response.Status)

	// PostData := map[string]string{"ID":"1","Data":"TestData"}
	// testPostBodyInJson, _ := json.Marshal(PostData)
	// res, err := http.Post("https://google.com","application/json",PostData)
	// byteResponse1, _ := io.ReadAll(res.Body)
	// fmt.Println(string(byteResponse1),"=========",res.Status)
}
