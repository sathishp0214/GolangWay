package trading

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	apiKey       = "c98d4774b51c458a8ce52626b231b7bc"
	tmpSecretKey = "2026.bcb223e8bfd643bcb2462285ec4431a38481760f58be45e3"
)

// func main() {
// 	InitiateFreshTokenGeneration()
// }

// run this function first, then initiate code generation on flattrade on browser
func InitiateFreshTokenGeneration() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// extract ?code=... from the URL
		code := r.URL.Query().Get("code")

		if code == "" {
			http.Error(w, "code parameter missing", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "Received code: %s", code)

		tmp := apiKey + code + tmpSecretKey

		data := []byte(tmp)

		// Calculate the hash: returns a [32]byte array
		hashArray := sha256.Sum256(data)

		shaToken := fmt.Sprintf("%x", hashArray)

		fmt.Println("completed", code, shaToken, "----------", tmp)

		generatingToken(code, shaToken)

		time.Sleep(1 * time.Second)

		return

	})

	fmt.Println("Listening on the port 8080")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))

}

func generatingToken(code string, shaToken string) {

	url1 := "https://authapi.flattrade.in/trade/apitoken"

	data := map[string]string{
		"api_key": apiKey, "request_code": code, "api_secret": shaToken,
	}

	payload, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", url1, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	fmt.Println("API Request for token generation:", resp)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("API Response for token generation:", string(body))
	// f.WriteString(fmt.Sprintf("API Response for token generation:%v\n", string(body)))
}
