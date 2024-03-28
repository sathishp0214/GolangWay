package main

import (
	"fmt"
	"sort"
)

var P = fmt.Println

type firstStruct struct {
	intValue   int
	strValue   string
	slicevalue []string
	mapvalue   map[string]string
}

// Value receiver function use case- Can use for reading the struct data only,

// Pointer receiver Function  use case - Should use for writing the struct data, Can use for reading struct data as well.

func (o firstStruct) ValueReceiver() {
	//struct receiver(non-pointer) function
	o.intValue = 11 //int value call by value type default, So this value will not be reflect outside

	//both slice and map are default call by reference types, So these value will reflect
	o.slicevalue[0] = "z"
	o.mapvalue["d"] = "dd"

	//We know, slice append function will not work as call by reference check notes and GolangsinglefileRecollect.go
	o.slicevalue = append(o.slicevalue, "p")
}

func (o *firstStruct) PointerReceiver() {

	//Pointer receiver should use generally - If we modifying the struct's object data

	//struct receiver(pointer) function, So it works like call by reference defaulty for all data types even for default call by value data types.

	o.intValue = 20
	o.strValue = "hello"
}

func (o *firstStruct) pointerreceiver1() int {

	P("Inside pointer receiver111111", o, "=========", *o)
	o.intValue = 20
	o.strValue = "hello"

	return 0
	//Pointer receiver function can use return statement as well normally
}

func CallReferenceStructPointer(structValue *firstStruct) {
	//Struct passed as pointer, So below call by value data types works as call by reference
	structValue.intValue = 100
	structValue.strValue = "Nice"
}

type BasicStruct struct {
	a int
}

func (b BasicStruct) add() error {
	return nil
}

type BirdInterface interface {
	Fly() string
	MakeSound() string
}

type Pigeon struct{}

// If "pigeon" struct has function implementation of all interface methods EX:Fly(), MakeSound(), Then only this struct implements the interface, Otherwise We can't use/assign struct object with interface object.
func (p *Pigeon) Fly() string {
	return "Pigeon is flying."
}

func main() {
	// ReceiverPointerReceiverMethods()

	// BasicStructWithInterface()

	// basicCompositionStructInterface()

	d := Good{}
	// d := &Good{}  //Whether we using this pointer memory address or normal - d := Good{} , Still have to use the pointer receiver method to work as pass by reference in struct.
	d.SetName("sathish")
	fmt.Println(d)

}

type Good struct {
	Id   int
	Name string
}

func (g *Good) SetName(name string) {
	g.Name = name
}

// Setting default values to the struct fields
func (g *Good) DefaultValues() {
	if g.Name == "" {
		g.Name = "Sathish"
	}
}

// Another combination of creating struct variable with default values.
func NewGoodStruct(name string) *Good {
	g := &Good{
		Id:   100, //assigns default variable
		Name: name,
	}

	return g
}

func CommonStructForHandlingMultipleStruct() {
	//That comcast interview task
	c := C{}
	c.value = 1000   //nameless struct composition
	c.b.value1 = 101 //named struct composition
	c.checkFunction()
	fmt.Println("After compared A and B struct's result -- ", c.result)
}

type A struct {
	value int
}

type B struct {
	value1 int
}

type C struct {
	A        //nameless struct composition
	b      B //named struct composition
	result bool
}

func (c *C) checkFunction() {
	if c.value > c.b.value1 {
		c.result = true
		return
	}

	c.result = false
}

func basicCompositionStructInterface() {
	c := CompositionStruct{}

	//Structs have all struct variables and struct's methods
	// c.anotherStruct.a
	c.anotherStruct.add()

	//Interface has access to only methods

	c.anotherInterface = CheckStruct{} //Have to assign Interface implemented struct, Otherwise nil pointer error will occur
	c.anotherInterface.calm()
	c.anotherInterface.welcome()
}

type CompositionStruct struct {
	anotherStruct    BasicStruct
	anotherInterface Interface2
}

func EmbeddedInterfaceBasic() {
	var interfaceObject1 Interface1
	var interfaceObject Interface2 //Interface2 has embedded interface of Interface1, So Interface2 has access to two methods welcome(),calm()
	Check := CheckStruct{}
	interfaceObject = Check
	interfaceObject.welcome()
	interfaceObject.calm()

	interfaceObject1 = Check
	interfaceObject1.welcome()
}

func BasicStructWithInterface() {

	circleObject := circle{size: 100}
	getDetailsInterfaceFunction(circleObject)

	getDetailsInterfaceFunction(circle{size: 100})

	getDetailsInterfaceFunction(Rectangle{RectanleSize: 120})
	getDetailsInterfaceFunction(&Rectangle{RectanleSize: 120}) //Passing as Struct as pointer also behaves same

	//Another Combination of using interface
	var p Shape
	circleObject.size = 500
	p = circleObject
	p.getFunction()

	p = Rectangle{RectanleSize: 240}
	p.getFunction()

	p = CircleV1{}
	p.getFunction()

	p = CircleV1{circle{size: 1000}} //Passing Circle struct value inside CircleV1 embedded struct
	p.getFunction()

	//Another Combination of using interface with Slice of interface implemented structs
	for _, interfaceFunction := range []Shape{circle{}, Rectangle{}, CircleV1{}} {
		interfaceFunction.getFunction()
	}

}

func EmbeddedStructHandle() {
	CarObject := Car{WheelModel: "Test Wheel", Color: "Blue"}

	embeddedStruct := Honda{Engine: "Test Engine", Car: CarObject} //Passing Embedded struct's (or) normal struct object in composition

	CarObject.Color = "green" //This modifying colour Value in CarObject will not reflect in embeddedStruct object, even we Already passed "CarObject" in the above line. (So its just composition, Not an inheritance. Since inheritance is not supported in golang)

	embeddedStruct.Color = "greenV1" //This modifying colour Value works in embeddedStruct

	P(embeddedStruct)
	embeddedStruct.GetCar() //Test Wheel, greenV1
}

func ReceiverPointerReceiverMethods() {
	o := firstStruct{intValue: 10, slicevalue: []string{"a", "b"}, mapvalue: map[string]string{"c": "cc"}}
	o.ValueReceiver()
	P(o)

	o = firstStruct{intValue: 10, strValue: "sat"}
	o.PointerReceiver()
	P(o)

	CallReferenceStructPointer(&o)
	P(o)

	o1 := &firstStruct{intValue: 10, strValue: "sat111"}
	o1.pointerreceiver1() //Here we are using the pointer of firstStruct from above line, //We may or may not use the pointer of "o1" of firstStruct. But still  We need the pointer Receiver function "func (o *firstStruct) pointerreceiver1() {}" for works as Call by reference.

	P(o1)
}

type Car struct {
	WheelModel string
	Color      string
}

func (c Car) GetCar() {
	P(c.WheelModel, c.Color)
}

type Honda struct {
	Car    //Embedded struct -- It is not inheritance, Golang doesn't supports inheritance. Its just more easier way of representing "Composition".
	Engine string
}

type Shape interface {
	getFunction()
}

type circle struct {
	size int
}

func (c circle) getFunction() {
	P("From Circle getFunction", c)
}

type CircleV1 struct {
	circle
}

// Doing method overridding with circle's method.
func (c CircleV1) getFunction() {
	P("From CircleV1 Composition getFunction", c)
}

type Rectangle struct {
	RectanleSize int
}

func (c Rectangle) getFunction() {
	P("From Circle getFunction", c)
}

func getDetailsInterfaceFunction(s Shape) {
	P("Calling the respective struct function")
	s.getFunction()
}

type check struct {
	a int
	i checkInterface
}

type checkInterface interface {
	hello(int)
}

type Interface1 interface {
	welcome()
}

type Interface2 interface {
	calm()
	Interface1
}

type CheckStruct struct {
	v int
}

func (c CheckStruct) welcome() {
	fmt.Println("CheckStruct welcome function")
}

func (c CheckStruct) calm() {
	fmt.Println("CheckStruct calm function")
}

type sortInterface interface {
	getValue() int
}

type sortStruct struct {
	value int
}

func (p sortStruct) getValue() int {
	return p.value
}

type sortStruct1 struct {
	value int
}

func (p sortStruct1) getValue() int {
	return p.value
}

func SortStructsUsingInterface() {
	s := sortStruct{10}
	s1 := sortStruct1{15}

	interfaceSort := []sortInterface{s, s1} //slice of interface implemented structs

	fmt.Println("before sorted", interfaceSort)

	sort.SliceStable(interfaceSort, func(i, j int) bool {
		return interfaceSort[i].getValue() > interfaceSort[j].getValue() //descending order
	})

	fmt.Println("after sorted", interfaceSort)
}

type AA struct {
	name string
}

type BB struct {
	name string
}

// In nameless struct composition, Both structs having the same field "name", So we are using CCStructObject.name() - It will cause ambituity compilation error.
type CC struct {
	AA
	BB
}
