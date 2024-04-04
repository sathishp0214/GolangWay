package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

/*
Unit testing:

go test file .go name format-- sample_test.go

//function Name format should start with Test
func TestFunction(t *testing.T) {}

//run all go test files in directory
go test -v

go test ./...   //run all test function in entire codebase directory

go test -cover  //Written unit test functions for overall code coverage percentage


Where Don't need for unit testing:
Very Simple logic function
simple create struct/class objects, get,set methods
3rd party package functions
Real Database CRUD -- These are two components, client and server - This suits more on integration testing.
http client and server -- ""
Testing multiple components/classes - Should not do integration testing with unit testing.


very difficult for unit testing:
Legacy codebase without proper comments

How many test cases should write:
How many return statement occurances are in a function.
In each return multiple values are returning, Have combination test case for each multiple value returning.


Different test case types:

Positive test case:
Passing valid/Non-error inputs to whether a function works correctly or not according to the positive input.

Negative test case:
Intentionally passing invalid/error inputs to check the whether a function handles invalid/error inputs correctly or not.

Boundary test case:
If a function works for values in the range of (1-100), So we are passing values just beyond this range EX: passing 0,101,102 etc

Base test case:
In recursion function, Testing the base condition.

Edge test case:
Non-ordinary or extreme or very big value test inputs. EX: passing 100000000000000000 input to integer32 value'

Corner test case: (bigger version of edge case)
If two or more edge case inputs.






Unit testing vs Integration testing:
Unit testing - Testing single function/component/class/struct. Since we are doing testing on single unit,  So depends upon the function arguments and return values, We are mocking the input and output test data.


integration testing - Combing/integrating two different function/classes/struct/component.

Example scenario of seperating both unit and integration testing:
User accessing REST API GET url and Then reads and displays the data from DB. So here we have two components REST API REquest/response and DB accessing and value reading.  So here we are doing Integration testing because we are using API compoenent and DB component. If in case of unit testing, We should do the seperate unit test for both API compoenent and DB component.

In integration testing, We can reuse by calling the seperate unit test functions inside integration tests.

For both unit and integration testing - Can use the same "testing" package

Its subjective, Depends upon the individul and team practice choice- How simple and how extend we can use either of unit or integration testing.

Example:
DB simple all CRUD functions - Can do it in a single unit test
DB complicated single function - Can do it in a seperate single unit test
Struct with complex method - Can do it in a seperate single unit test.

Two different dependent functions - Can do it in integration testing

Code coverage:
go test -cover  //Can measure the code coverage percentage EX: coverage: 91.9% of statements

go test -coverprofile=coverage.out
go tool cover -html=coverage.out  //This command will highlight the actual our code lines, where our test cases are covered and not-covered. Then we can write test case for not-covered areas.

*/

/*
Unit testing GoRoutine without channel inside - We can test like normal function without using "go" functioncall().

Unit testing in channels - Channels working between goRoutines, So unit testing we does on single function, My opinion channels is fit for unit testing. So we can do integration testing.
*/

/*
Errorf() vs Fail()
Fail() - Just fails the case, Cannot display the failure message, But Errorf() Supports display the failure message

*/

func TestSimple(t *testing.T) {
	input := 10
	output := 20

	//If no fail test case functions called, Then its considered as Success
	if input*2 == output {

		t.Logf("Log message %v", 1) //prints this message in terminal

		t.Errorf("error message %v", 1) //This also Failing the test with message in terminal

		t.Skipf("skipping all the code below %v", 1)

	}
}

func TestFatal(t *testing.T) {
	var a error
	if a == nil {
		// t.Fatalf("fatal message check")
		t.Errorf("Error message check")
	}

	//Fatalf vs Errorf - Fatalf will stop the code execution at that line itsef, Even code hits Errorf, Still code execution not stops. Anyways Both fails the test case.
	//should use fatalf for pre-validation before getting into expensive/long code operations

	t.Logf("Checks whether this line prints")
}

func TestCallAnotherTestFunction(t *testing.T) {
	TestSimple(t) //We are calling another test function inside another test function
}

func TestFunctionWithSubTest(t *testing.T) {
	/*
		t.Run() - We can run the subtest using t.Run

		//can run particular subtest alone
		go test -run=TestFunctionWithSubTest/Subtest_function_name_1 -v
	*/

	t.Run("Subtest_function_name_1", TestSimple) //Calling another test function "TestSimple" in another function and gets the return with Success or fail case from that test function "TestSimple".

	t.Run("Subtest_function_name_2", func(t *testing.T) {
		a := 10
		if a < 10 {
			t.Errorf("failed case in value")
		}
	})
}

// Example actual function with multiple returns
func getNumbers(n int) bool {
	if n <= 0 {
		return false
	}

	if n%2 == 0 {
		return true
	}

	return false
}

// This test functions, passes different inputs according to the different return values
// If none of error condition is called, Then one or all test cases are passed.
// Here Am used three test cases in one test function. (or)Only one test case in one test function.
func TestGetNumbers(t *testing.T) {
	n := getNumbers(0)
	if n {
		t.Errorf("case 1 - Should return true")
	} else {
		t.Logf("case 1 - test passed")
	}

	n1 := getNumbers(2)
	if !n1 {
		t.Errorf("case 2 - Should return true")
	}

	n3 := getNumbers(3)
	assert.Falsef(t, n3, "case 3 - Should return false")

	n4 := getNumbers(-1)
	if n4 {
		t.Errorf("case 4 - Should return true")
	}
}

type TableDrivenTestGetNumbers struct {
	InputNumber int
	Output      bool
}

// table driven test (or) Data driven test -- Passing multiple input test cases into test function using struct.
// table driven test type reduces the code duplicate. Instead of writing muliple test function for multiple test inputs, We can use single table driven test function with passing multiple test inputs.
func TestTableDrivenGetNumbers(t *testing.T) {
	inputs := []TableDrivenTestGetNumbers{{0, false}, {2, true}, {3, false}}

	for _, testCase := range inputs {
		funcOutput := getNumbers(testCase.InputNumber)
		if funcOutput != testCase.Output {
			//This type reduces the code duplicate, TestGetNumbers() - Above This function uses multiple conditions and its error statements for each getNumbers() call.
			t.Errorf("Test case failed -- %v -- %v", testCase.InputNumber, testCase.Output)
		}
	}
}

type DrivenTestCaseExample struct {
	Description            string
	number                 int
	flag                   bool
	ExpectedOutputFunction func() (interface{}, error)
	insertFunction         func() error
}

func TestDrivenFunction(t *testing.T) {
	testCases := []DrivenTestCaseExample{
		//test case 1
		{
			Description: "test description",
			number:      12,
			flag:        false,
			//inside these functions, We can put the needed actual code to test
			ExpectedOutputFunction: func() (interface{}, error) {
				e := "value1"
				return e[0], nil
			},
			insertFunction: func() error {
				return errors.New("test error")
			},
		},
		//test case 2
		{
			Description: "test description",
			number:      12,
			flag:        false,
			ExpectedOutputFunction: func() (interface{}, error) {
				e := "value"
				return e, nil
			},
			insertFunction: func() error {
				return nil
			},
		},
	}

	for _, testCase := range testCases {
		if testCase.insertFunction() != nil {
			t.Logf("insert successed")
		}

		if testCase.number > 10 {
			t.Errorf("number should not greter than 10")
		}
	}

}

// func TestGoRoutineSample(t *testing.T) {
// 	wg := &sync.WaitGroup{}
// 	for i := 0; i < 10000; i++ {
// 		wg.Add(1)
// 		go func() {
// 			for i := 0; i < 10000000; i++ {
// 				_ = 10
// 			}
// 			wg.Done()
// 		}()
// 	}

// 	wg.Wait()
// 	t.Logf("GoRoutine completed")
// }

// func TestGoRoutineSample1(t *testing.T) {
// 	for i := 0; i < 10000; i++ {
// 		go func() {
// 			for i := 0; i < 10000000; i++ {
// 				_ = 10
// 			}
// 		}()
// 	}
// 	t.Logf("GoRoutine completed")
// }

func TestAssertPackage(t *testing.T) {

	//Assert package -- Used to check different conditions, Which Has lot of inbuilt functions to support it. So we can directly use these inbuilt functions and get the pass/fail test case directly.

	//Without using this assert package, We have to write some code manually and put the conditions manually to pass or fail test case.

	//most popular functions
	//equal or not equal, empty or not empty, True or false, len of array/slice, Nil/Not nil, compare data type, HTTP related functions and lot more

	assert.Equal(t, 10, 10, "Should be equal") //gets pass test case directly

	// assert.Equal(t, 10, 11, "Should be equal") //intentional fail test case,

	assert.NotEqual(t, 10, 11, "Should be not equal")

	var pointer *int
	assert.Nil(t, pointer, "Nil pointer")

	// err := functionError()
	// if assert.Error(t, err, "error happend or not") {
	// 	assert.Fail(t, "Function returned error")  //intentionally fails the test case
	// }

	condition := assert.Panics(t, func() { //Checks function returns panic or not
		panic("panic created")
	})

	assert.True(t, condition, "panic is true") //compares the "true" boolean case.

	//some most needed function
	// assert.IsType()
	// assert.True()
	// assert.False()
	// assert.Len()
	// assert.Contains()
	// assert.ElementsMatch()  //compare array,slice etc
	// assert.Empty()

	//assert some http related
	// assert.HTTPBodyContains()
	// assert.HTTPError()
	// assert.HTTPStatusCode()
	// assert.HTTPSuccess()

}

func functionError() error {
	return errors.New("Error is true")
}

// Testing the interface function(EX:Interface as function argument/return etc) with mock struct:
// Declaring the mock sample struct with interface declared method and passing into interface method to check it.
type SampleInterface interface {
	Save(int) int //Assume these function methods are DB methods, API methods, Data write methods, So we can't do it directly on actual struct methods, So we create a mockStruct and its methods for unit testing
}

type MockStruct struct{}

func (m *MockStruct) Save(a int) int {
	//Doing some mock here
	return 0
}

func TestInterfaceFunction(t *testing.T) {

	mockStruct := &MockStruct{}
	err := InterfaceFunction(mockStruct, 10)
	if err != nil {
		t.Error("Test casse failed")
	}

}

func InterfaceFunction(i SampleInterface, value int) error {
	i.Save(value)
	return nil
}

//Can use method overriding with mock struct methods for testing other struct methods.

/*
Mocking a database in unit testing:

Other approaches:

We can directly use our testing db server for CRUD unit/integration testing operations. But its not considered best and faster way, If we are going to run our test functions so frequently like every git commit, frequent jenkins build test cases.



go-sqlmock package - Used for mocking sql database. It returns "*sql.DB" sql package's dummy object. We can pass this dummy "*sql.DB" object with test data into our actual CRUD functions. So this will not read/write in actual db. Check below example.

mongomock package - For mocking mongo database connection.
*/

func TestSQLMockDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "data"}).AddRow(1, "mocked data")
	mock.ExpectQuery("SELECT (.+) FROM table").WillReturnRows(rows)                                         //runs mock/dummy select query
	mock.ExpectExec("INSERT INTO product_viewers").WithArgs(2, 3).WillReturnResult(sqlmock.NewResult(1, 1)) //runs dummy select query
	mock.ExpectExec("UPDATE products").WillReturnResult(sqlmock.NewResult(1, 1))                            //dummy update queries

	// Invoke your actual database function and check for errors.
	// err = ActualDBInsertFunction(db, 2, 1)
	// if err != nil {
	// 	t.Errorf("db Insert failed")
	// }
}

// func ActualDBInsertFunction(db *sql.db, productID int, UserID int) error {
// 	//doing our actual insert logics
// 	return nil
// }

// func TestMongoMockDB(t *testing.T) {
// 	db := mongomock.NewDB() //mongo mock package --simialr to above sql mock package, We can pass this object to actual mongo CRUD functions to test it.
// 	collection := db.Collection("users")
// 	err := collection.InsertOne(User{
// 		ID:    1,
// 		Name:  "test",
// 		Email: "example@example.org",
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

type User struct {
	ID    int    `bson:"_id" json:"id"`
	Name  string `bson:"username"`
	Email string `bson:"email"`
}

// Uni testing scenario
type Client struct{}

func NewClient() Client {
	return Client{}
}

func (c Client) GetData() (string, error) {
	return "data", nil
}

func Controller() error {
	externalClient := NewClient()
	fromExternalAPI, err := externalClient.GetData()
	if err != nil {
		return err
	}
	// do some things based on data from external API
	if fromExternalAPI != "data" {
		return errors.New("unexpected data")
	}
	return nil
}

func Controller1(externalClient Client) error {
	// externalClient := NewClient()
	fromExternalAPI, err := externalClient.GetData()
	if err != nil {
		return err
	}
	// do some things based on data from external API
	if fromExternalAPI != "data" {
		return errors.New("unexpected data")
	}
	return nil
}

type IexternalClient interface {
	GetData() (string, error)
}

type MockSClientStruct struct {
}

func (c MockSClientStruct) GetData() (string, error) {
	return "MockSClientStruct data", nil
}

func Controller2(externalClient IexternalClient) error {
	// externalClient := NewClient()
	fromExternalAPI, err := externalClient.GetData()
	if err != nil {
		return err
	}
	// do some things based on data from external API
	if fromExternalAPI != "data" {
		return errors.New("unexpected data")
	}
	return nil
}

func TestController(t *testing.T) {
	//We have different Controller() functions for unit testing

	err := Controller() //In this function code, We don't have the access for the "Client" struct object, So we can't check Success or failure from "Client" struct method.
	if err != nil {
		t.Errorf("error failed")
	}

	//In this case, We can pass "Client" struct object can test the success scenario, But still can't test it for failure scenarios.
	err = Controller1(Client{})
	if err != nil {
		t.Errorf("error failed")
	}

	err = Controller2(Client{})
	if err != nil {
		t.Errorf("error failed")
	}

	//In this case, Now we have interface as function argument with Mock struct, So we load failure case value to the MockStruct and test for Failure case scenario.
	err = Controller2(MockSClientStruct{})
	if err != nil {
		t.Errorf("error failed --- %v", err)
	}

}
