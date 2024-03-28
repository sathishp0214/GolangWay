package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {

}

// Generally HTTP POST IS similar to HTTP PUT/PATCH requests
func HTTPPOSTRequest() {

	ServerUrl := "http://localhost:10000/post/"

	POSTRequestData := []byte(`{"name":"sat","age":18,"country":"India"}`) //This is a way to pass the json data as multi line string into byte slice

	request, err := http.NewRequest(http.MethodPost, ServerUrl, bytes.NewBuffer(POSTRequestData))

	//adding key-value pair in http request headers
	request.Header.Add("Content-Type", "application/json")

	//similarly can add key-value pair in url paramters also
	// baseURL, _ := url.Parse(ServerUrl)
	// parm := url.Values{}
	// parm.Add("api_key", "AccessToken")
	// parm.Add("format", "json")
	// baseURL.RawQuery = parm.Encode()

	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	byteResponse, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("http response", string(byteResponse))
}

// This is almost similar to get request
func HTTPDELETERequest() {
	ServerUrl := "http://localhost:10000/Delete/10" //deleting 10th resource data

	request, err := http.NewRequest(http.MethodDelete, ServerUrl, nil)

	//adding key-value pair in http request headers
	request.Header.Add("Content-Type", "application/json")

	//similarly can add key-value pair in url paramters also
	// baseURL, _ := url.Parse(ServerUrl)
	// parm := url.Values{}
	// parm.Add("api_key", "AccessToken")
	// parm.Add("format", "json")
	// baseURL.RawQuery = parm.Encode()

	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	byteResponse, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("http response", string(byteResponse))
}
