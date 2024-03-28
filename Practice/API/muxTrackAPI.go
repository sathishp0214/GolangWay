package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {

	//apikey=95d956640088ece978d54b9b5f78da5b
	AccessToken := "95d956640088ece978d54b9b5f78da5b"
	// url1 := "http://api.musixmatch.com/ws/1.1/track.get/?commontrack_id=5920049"
	url1 := "http://api.musixmatch.com/ws/1.1/"

	// AccessToken := "a2a554c423c90786912e0fc8bdc99c1c" //last fm api key

	// url := "https://ws.audioscrobbler.com/2.0/?method=chart.gettopartists&api_key=a2a554c423c90786912e0fc8bdc99c1c&format=json" //last fm url
	// /2.0/?method=chart.gettopartists&api_key=YOUR_API_KEY&format=json

	// url1 := "https://ws.audioscrobbler.com/2.0/?method=chart.gettopartists&format=json"
	// url1 := "https://ws.audioscrobbler.com/2.0/?"

	// Create a Bearer string by appending string access token
	// var bearer = "Bearer " + AccessToken

	baseURL, _ := url.Parse(url1)
	parm := url.Values{}
	parm.Add("apikey", AccessToken)
	parm.Add("method", "track.get")
	parm.Add("commontrack_id", "5920049")
	parm.Add("format", "json")
	parm.Add("limit", "2")
	// parm.Add("name", "Taylor Swift")
	baseURL.RawQuery = parm.Encode()

	fmt.Println("============================", baseURL.String())

	// Create a new request using http
	// req, err := http.NewRequest("GET", url1, strings.NewReader(parm.Encode()))
	response, err := http.Get(baseURL.String())
	fmt.Println("error----------", err)

	// add authorization header to the req
	// req.Header.Add("Authorization", "apikey="+AccessToken)
	// req.Header.Add("Authorization", "Basic "+AccessToken)
	// req.Header.Add("Accept", "application/json")
	// req.Header.Add("x-api-key", AccessToken)
	// req.Header.Add("apikey", AccessToken)
	// req.Header.Set("apikey", AccessToken)

	// req.Header.Add("api_key", AccessToken)

	// Send req using http Client
	// client := &http.Client{}
	// response, err := http.DefaultClient.Do(req)
	fmt.Println("error----------", err)
	fmt.Println("response--------", response)

	// response, _ := http.Get(url)
	byteResponse, err := io.ReadAll(response.Body)
	fmt.Println("error----------", err)
	fmt.Println(string(byteResponse), "=========", response.Status)
}
