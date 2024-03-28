package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "http://dummyjson.com/users"
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	byteResponse, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(byteResponse), response.StatusCode)
	var user Users
	err = json.Unmarshal(byteResponse, &user)
	if err != nil {
		panic(err)
	}
	// fmt.Println("users---", user)
	var outputUsers = Users{}
	receipesData := Receipe1()
	for _, i := range user.User {
		// fmt.Println(i.Address.City)
		// if i.Address.City == "Washington" {
		// 	fmt.Println("wahibngton", i)

		// }

		user := FindReceipeList(receipesData, i)
		if user != nil {
			outputUsers.User = append(outputUsers.User, *user)
		}
	}

	fmt.Println("users from washington---", outputUsers)

}

func FindReceipeList(receipesData Receipes, user User) *User {
	for _, i := range receipesData.Receipe {
		if user.Id == i.UserID {
			// fmt.Println("printing receipes from userID", i)
			user.Receipe = i
			return &user
		}
	}
	return nil
}

func Receipe1() Receipes {
	url := "http://dummyjson.com/recipes"
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	byteResponse, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(byteResponse), response.StatusCode)

	var receipe Receipes
	err = json.Unmarshal(byteResponse, &receipe)
	if err != nil {
		panic(err)
	}
	// fmt.Println("users---", receipe)
	return receipe

}

type Receipes struct {
	Receipe []Receipe `json:"recipes"`
}

type Receipe struct {
	UserID      int      `json:"userId"`
	Ingredients []string `json:"ingredients"`
}

type Users struct {
	User []User `json:"users"`
}

type User struct {
	Id      int     `json:"id"`
	Address Address `json:"address"`
	Receipe Receipe
}

type Address struct {
	City string `json:"city"`
}
