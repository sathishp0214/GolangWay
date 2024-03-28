package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/http/cookiejar"

	"github.com/go-chi/chi"
)

/*
Every http handler function is http server function, Each server handler function have http request and http response
*/

func main() {
	router := chi.NewRouter()

	//different http requests.
	// router.Get("/get",handlerfunction)
	router.Post("/post", CreatePost)
	// router.Put()

	//form value parsing and handling
	router.HandleFunc("/formHandling", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Println("prints the Username from input request's form --", r.FormValue("Username"))
		fmt.Println("prints the email from input request's form --", r.FormValue("Email"))
		fmt.Println("prints the Password from input request's form --", r.FormValue("Password"))

	})

	router.Handle("/", http.FileServer(http.Dir("./static/index.html"))) //Assume "static" is the folder which holds html pages. Serving static html web page in main home url "/"

	router.HandleFunc("/redirect", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "http://redirectUrl.com/", 301) //redirect code
	})

	router.HandleFunc("/dynamicData", DynamicDataHandler)

	//can use middlewares function
	router.Use(LogMiddlewareFunction)
	// router.Use(SecurityMiddlewareFunction)

	//need to check

	//chi subroute
	// chi login and logout

	router.HandleFunc("/setcookie", setCookieHandler) //creates cookie on http server handler function returns cookie on http response
	router.HandleFunc("/getcookie", getCookieHandler) //gets and reads the cookie from http request

	//listening above all url routing on this port.
	http.ListenAndServe(":8080", router)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	// var db *sql.DB //using sql db here - Use can use any sql db like mysql and postgres, Have to incluse its respective driver with sql package

	// var post PostStruct //Post struct for DB schema
	// json.NewDecoder(r.Body).Decode(&post)

	// query, err := db.Prepare("Insert posts SET title=?, content=?")  //Using raw queries
	// if err != nil {
	// 	panic(err)
	// }

	// _, er := query.Exec(post.Title, post.Content)
	// if er != nil {
	// 	panic(er)
	// }
	// defer query.Close()

	// respondwithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"}) //handling response data in below

}
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// func LogMiddlewareFunction(handler http.Handler) http.Handler {
// 	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
// 		log.Println(request.URL.Path)
// 		handler.ServeHTTP(writer, request)
// 	})
// }

var tpl = template.Must(template.ParseFiles("index.html"))

func DynamicDataHandler(w http.ResponseWriter, r *http.Request) {
	// 	data := NotificationDataStructSample{
	// 		Name:    "John Doe",
	// 		Message: "Your account has been updated.",
	// 	}
	// 	tpl.Execute(w, data)

	// 	//Passing this above dynamic data into respone for index.html web page.

	// 	// Dear {{.Name}},
	// 	// You have a new notification: {{.Message}}
	// 	// Sincerely,
	// 	// The Team

	// 	//Assume the above content in html page, Our dynamic data will be displayed.
}

func httpRequestTopInfoHandling() {

	r, _ := http.NewRequest("GET", "http://localhost:10000", nil)

	//We have to handle these below request operations in client side and server side seperately like below.

	//--------------------Request client side handling on below--------------------------
	r.Header.Add("Adds new header keyName", "sets keys value")
	r.Header.Set("header keyName", "set new header key's value")

	//adds cookie to the request
	// r.AddCookie()

	//Can read form data from Which is sent in the http post/put/patch request body, We can read like below ParseForm()
	// r.PostForm.Add("Username")

	//-----------------Http request's server side handling on below----------------
	//We have server side request handling based on client side request inputs

	//gets request body data
	// r.GetBody()

	// r.UserAgent()
	// r.Host

	//can writes the request like header,body etc
	// r.Write()

	//once form parsed, Gets form value from key
	r.ParseForm()
	r.FormValue("Username") //Gets username value from input request's from

	//Can read upload/sent File, Which is sent in the http post/put/patch request body with "multipart/form-data" content-Type similar to application/json
	// file := r.MultipartForm.File

	//reads username and password from http request's basic authentication type.
	// username,password, bool := r.BasicAuth()

	//read particular header value
	r.Header.Get("api key") //gets api key passed in request headers, So we can verify authentication.
	r.Header.Get("header keyName")

	//we gets the complete url, We can read the url parameters with regex, Or we have beter packages as well for url parsing
	// IncomingrequestUrl := r.URL.Path

	//gets the request cookies
	// r.Cookie()
	// r.Cookies()

}

func httpResponseTopHandling() {
	response := http.Response{}

	//response header objects
	// response.Header.Get()
	// response.Header.Set()

	response.StatusCode = 200

	//from response object we have request object info
	// response.Request

}

func DefaultHTTPHandleFunc() {
	http.HandleFunc("/hello", httpHandlerFunction) //This works for all http requests
	http.ListenAndServe(":9000", nil)
}

func httpHandlerFunction(http.ResponseWriter, *http.Request) {
	fmt.Println("http packages default handler function")
}

func setCookieHandler(w http.ResponseWriter, r *http.Request) {
	// Initialize a new cookie containing the string "Hello world!" and some
	// non-default attributes.
	cookie := http.Cookie{
		Name:     "exampleCookie",
		Value:    "Hello world!",
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	//another example cookie
	// cookie := http.Cookie{}
	// cookie.Name = "Test cookie"
	// cookie.Value = "Test cookie value"
	// cookie.Expires = time.Now().Add(time.Duration(time.Now().Year()))
	// cookie.Secure = true

	// Use the http.SetCookie() function to send the cookie to the client.
	// Behind the scenes this adds a `Set-Cookie` header to the response
	// containing the necessary cookie data.
	http.SetCookie(w, &cookie)

	// Write a HTTP response as normal.
	w.Write([]byte("cookie set!"))
}

func getCookieHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the cookie from the request using its name (which in our case is
	// "exampleCookie"). If no matching cookie is found, this will return a
	// http.ErrNoCookie error. We check for this, and return a 400 Bad Request
	// response to the client.
	cookie, err := r.Cookie("exampleCookie")
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			http.Error(w, "cookie not found", http.StatusBadRequest)
		default:
			log.Println(err)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return
	}

	// Echo out the cookie value in the response body.
	w.Write([]byte(cookie.Value))
}

// https://www.golangprograms.com/how-do-you-set-cookies-in-an-http-request-with-an-http-client-in-go.html
func CookieHndlingInCLientSide() {
	cookie := &http.Cookie{
		Name:  "session_id",
		Value: "12345",
	}

	client := &http.Client{
		Jar:       &cookiejar.Jar{},
		Transport: &http.Transport{},
	}

	req, err := http.NewRequest("GET", "https://www.example.com", nil)
	if err != nil {
		// handle error
	}

	client.Jar.SetCookies(req.URL, []*http.Cookie{cookie}) //adds the cookei in the cookie jar

	resp, err := client.Do(req) //client while sending request with above added cookie
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
}

/*
Sessions in http.cookie: -- Cookie can be used to holds the session information

type Cookie struct {
	Name  string
	Value string

	Path       string    // optional
	Domain     string    // optional

	//Can set the cookie expiry time, So the cookie's expiration expires the sessions too
	Expires    time.Time // optional
	RawExpires string    // for reading cookies only

	// MaxAge=0 means no 'Max-Age' attribute specified.
	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
	// MaxAge>0 means Max-Age attribute present and given in seconds
	MaxAge   int
	Secure   bool
	HttpOnly bool
	SameSite SameSite
	Raw      string
	Unparsed []string // Raw text of unparsed attribute-value pairs
}

//Example
expiration := time.Now().Add(365 * 24 * time.Hour)
//assume "astaxie" is session validation value
cookie := http.Cookie{Name: "username", Value: "astaxie", Expires: expiration}
http.SetCookie(w, &cookie)
*/

/*
ORM in golang:
Famous go orm package - "gorm" -supports orm with sql databases
It is a big package comes with lot of functions and own styles.

Can covert golang structs into sql tables. From structs field itself can define columns with primary keys, constraints, validators etc

Can automigrate -- convert structs into sql tables.

Have CRUD inbuilt functions

supports transactions, pre and post functions after CRUD functions.

*/
