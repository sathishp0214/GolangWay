package main

import "fmt"

// func main() {

// }

//Open-closed Principle - Open for extension and closed for modification.
//General Interface with struct example - Structs implements interface function declaration methods, Having common interface logic, Where all structs type methods can be handled in the same way.
type Payment interface {
	Process()
}

type CreditCardPayment struct {
	Account string
	Amount  float64
}

func (ccp CreditCardPayment) Process() {
	fmt.Printf("Processing credit card payment of $%.2f from %s\n", ccp.Amount, ccp.Account)
	// Perform credit card payment specific logic
}

type BankTransferPayment struct {
	Account string
	Amount  float64
}

func (btp BankTransferPayment) Process() {
	fmt.Printf("Processing bank transfer payment of $%.2f from %s\n", btp.Amount, btp.Account)
	// Perform bank transfer specific logic
}

func ProcessPayment(payment Payment) {
	payment.Process()
}

func OpenClosedPrinciple() {

	creditCardPayment := CreditCardPayment{Account: "123456789", Amount: 100.0}
	ProcessPayment(creditCardPayment)

	bankTransferPayment := BankTransferPayment{Account: "987654321", Amount: 200.0}
	ProcessPayment(bankTransferPayment)
}

func main() {

}

//https://medium.com/@quicktechlearn/golang-solid-principles-dependency-inversion-principle-dip-b76b897a4aa4
func DependencyInversionPrinciple() {

	d := Department{}
	d.AddEmployee(Worker{1, "sathish"})
	d.AddEmployee(Supervisor{2, "Vel"})
	fmt.Println(d.GetEmployee(1))
	fmt.Println(d.GetEmployee(2))
}

type Worker struct {
	ID   int
	Name string
}

func (w Worker) GetID() int {
	return w.ID
}

func (w Worker) GetName() string {
	return w.Name
}

type Supervisor struct {
	ID   int
	Name string
}

func (s Supervisor) GetID() int {
	return s.ID
}

func (s Supervisor) GetName() string {
	return s.Name
}

type Employee interface {
	GetID() int
	GetName() string
}

//Dependency Inversion Principle - Higher level structs/modules should not depend on Lower level structs/modules.
//Higher Level structs - Worker, Supervisor
//Lower Level structs - Department
//Here both higher and lower level structs not depends on each other directly, Instead depends on "Employee" interface

type Department struct {
	employee []Employee
}

//Same composition example instead of Interface
// type Department struct {
// 	Workers []Worker
// 	Supervisors []Supervisor
// }

func (d *Department) AddEmployee(e Employee) { //If we used composition in Department struct instead of interface, We should handle both Worker and Supervisor structs objects seperately by more code use. We should modify Seperately for any changes in both Worker and Supervisor structs.
	d.employee = append(d.employee, e)
}

func (d Department) GetEmployee(id int) Employee {
	for _, i := range d.employee {
		if i.GetID() == id {
			return i
		}
	}

	return nil
}

type SInterface interface {
	Ask()
}

type S1 struct {
}

func (s S1) Ask() {

}

type S11 struct {
}

func (s S11) Ask() {

}

type S2 struct {
	s1 S1
	sp SInterface
}

func (s S2) Check(s1 SInterface) {
	s1.Ask()
}

func InterfaceSegregationPrinciple() {
	//Interface Segregation Principle - Interface should not be bulk, Since Structs should define all interface methods for implementing an interface. So if an interface with more number of methods, Structs whether its needed or not needed, Structs are forced to define all multiple interface methods to implement an interface. That's why we should have light weight interface to avoid forcing the structs to define unneeded methods definition.
	var db DBSampleLightInterface
	var dbQ DBQueryExecuteLightInterface
	S := StructInterfaceSegregation{} //It define two interface all methods, So this struct implements both interfaces

	//Below assigns struct object with both implmented interface
	db = S
	dbQ = S

	db.Connect()
	db.Disconnect()
	dbQ.Query()
	dbQ.Execute()

	S1 := StructInterfaceSegregationV1{} //This struct implmemnts only lighter weight "DBQueryExecute" interface. If the interface is bulk, Then This struct forced to idefine other methods
	db = S1
	db.Connect()
	db.Disconnect()
}

type BulkInterface interface {
	Connect()
	Disconnect()
	Query()
	Execute()
}

type DBSampleLightInterface interface {
	Connect()
	Disconnect()
}

type DBQueryExecuteLightInterface interface {
	Query()
	Execute()
}

type StructInterfaceSegregation struct {
}

func (s StructInterfaceSegregation) Connect() {

}

func (s StructInterfaceSegregation) Disconnect() {

}

func (s StructInterfaceSegregation) Execute() {

}

func (s StructInterfaceSegregation) Query() {

}

type StructInterfaceSegregationV1 struct {
}

func (s StructInterfaceSegregationV1) Connect() {

}

func (s StructInterfaceSegregationV1) Disconnect() {

}

//Liskov Substituion Principle - Objects of Superclass/Parent Struct should be replacable with objects of subClass/Child struct without modification in the code.
func LiskovSubstituionPrinciple() {
	Animal := AnimalStruct{}
	bird := Bird{}
	Sound(Animal)
	Sound(bird)
}

type Wild interface {
	MakeSound()
}

type AnimalStruct struct{}

func (b AnimalStruct) MakeSound() {
	fmt.Println("Animal sound")
}

type Bird struct {
	AnimalStruct //compositon
}

func (b Bird) MakeSound() {
	fmt.Println("Bird sound")
}

func Sound(w Wild) {
	//Both parent and child struct implements the "wild" interface, So depends upon the Parent and child struct's objects, We calling respective struct's MakeSound() -- So this is what this principle is -- Objects of Superclass/Parent Struct should be replacable with objects of subClass/Child struct without modification in the code.
	w.MakeSound()
	// Animal sound
	// Bird sound
}

// Interface as argument in Normal function or struct method
// func (s struct)InterfaceFunction(i interface) {
// }

//Interface as function return type

// //Using interface as with named or nameless composition
// type Struct struct {
// 	interfaceObject interface
// 	interface
// }
