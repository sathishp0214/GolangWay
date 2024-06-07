package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
)

type MyService interface {
	GetUser(id int) (User, error)
}

type User struct {
	ID   int
	Name string
}

/*
Need for mocking:
Assume function is reading/writing data from DB, doing http requests, Sending email, sms, notifications, writing the data somewhere. We normally runs our unit test functions for n number of times, So our actual function should not write the data and Even for reading purposes the overhead/load is high and we can eliminate with the use of mocking functions

*/

// Actual GetUSer function, Assume this function is reading data from DB, HTTP GET request etc - So in the unit test, We are mocking this function.
func (u User) GetUser(id int) (User, error) {
	return User{10, "sathish"}, nil
}

type MockService struct {
	mock.Mock //nameless composition of mock
}

func (m *MockService) GetUser(id int) (User, error) {
	args := m.Called(id) //This marks, this mock "GetUser" function is called,  If no arguments passed in this function then EX: args := m.Called()

	return args.Get(0).(User), args.Error(1)
	//Return(User{ID: 1, Name: "Alice"}, nil)   -- Actually this is the mock two argument values we passed here.
	//args.Get(0) -- 0 means 1st function argument
	//args.Error(1)  -- 1 means 2nd function argument, "Error" is the error data type
}

// Not tested this, Test once
func TestGetUser(t *testing.T) {
	ErrUserNotFound := errors.New("error not found")
	mockService := new(MockService)

	// On("GetUser", 1) -- "On" function sets the mock function's configuration like functionname and arguments like "1" is that "id" function argument
	// Return(User{ID: 1, Name: "Alice"}, nil) -- configuring what mock values should return from above mock GetUser function

	mockService.On("GetUser", 1).Return(User{ID: 1, Name: "Alice"}, nil) //Here
	mockService.On("GetUser", 2).Return(User{}, ErrUserNotFound)

	user1, err1 := mockService.GetUser(1) //calling above actual mock GetUser function with function arguments
	//user1, err1 - catches the mocked values passed into by above -- Return(User{ID: 1, Name: "Alice"}, nil)
	user2, err2 := mockService.GetUser(2)

	// here doing the expected output's validation for unit test errors
	if err1 != nil {
		t.Errorf("Unexpected error for user 1: %v", err1)
	}
	if user1.ID != 1 || user1.Name != "Alice" {
		t.Errorf("Unexpected user data for user 1")
	}

	if err2 != nil && err2 != ErrUserNotFound {
		t.Errorf("Unexpected error for user 2: %v", err2)
	}
	if user2.ID != 0 || user2.Name != "" {
		t.Errorf("Unexpected user data for user 2 (should be empty)")
	}

	mockService.AssertExpectations(t) //verifying whether our mock methods are called with these values "mockService.On("GetUser", 1).Return(User{ID: 1, Name: "Alice"}, nil)" or not.

}
