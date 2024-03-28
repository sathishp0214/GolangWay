package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//sample token authentication through header through middleware
//logging middleware
//crud api handler server methods
//postgres db crud

var Postgresdb, _ = sqlx.Connect("postgres", "user=postgres password=234403 dbname=dvdrental sslmode=disable")

func main() {
	router := chi.NewRouter()

	//middlewares should be defined before api router
	router.Use(LogMiddlewareFunction)
	// router.Use(ValidationMiddleware)

	router.Get("/api/user/{id}", GetFunction)
	router.Get("/api/user/", GetAllFunction)
	router.Post("/api/user/", PostFunction)
	router.Delete("/api/user/{id}", DeleteFunction)
	router.Put("/api/user/{id}", PutFunction)

	http.ListenAndServe(":9000", router)
}

//middlewares defaults prints all each request's handlerfunction's print,error,log messages and exclusively we are printing some info inside below middlewares function.

func LogMiddlewareFunction(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("middleware printing", request.URL.Path)
		fmt.Println("middleware printing", request.UserAgent())
		fmt.Println("middleware printing", request.Body)
		fmt.Println("middleware printing", request.Body)
		handler.ServeHTTP(writer, request) //This sending the input's request to next middleware or request's handling function
	})
}

func ValidationMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		//doing some validation, Similarly we can validate "authentication" from request headers, username/password from request's credentials etc
		if request.Method == "GET" {
			writer.WriteHeader(400)
			writer.Write([]byte("not allowed method"))
			return //Returning the response here itself, So the request will not reach to the actual handler function.
		}

		fmt.Println("not a get request, SO proceeding further")
		handler.ServeHTTP(writer, request)
	})
}

func GetFunction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("recieved get request")
	idParamater := chi.URLParam(r, "id")
	w.Header().Set("content-type", "application/json")

	getQuery := "select * from apitable where id = %s"
	getQuery = fmt.Sprintf(getQuery, idParamater)
	fmt.Println(getQuery)
	var user User
	row := Postgresdb.QueryRowx(getQuery)
	err := row.StructScan(&user)
	fmt.Println(err)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("User Id not found"))
		return
	}
	//writing data into response
	w.WriteHeader(200)
	_ = json.NewEncoder(w).Encode(user)
	fmt.Println("sending get response")
}

func GetAllFunction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	getQuery := "select * from apitable"
	var users []User
	rows, _ := Postgresdb.Queryx(getQuery)

	for rows.Next() {
		var user User
		_ = rows.StructScan(&user)
		users = append(users, user)
	}

	//writing data into response
	w.WriteHeader(200)
	_ = json.NewEncoder(w).Encode(users)

}

func PostFunction(w http.ResponseWriter, r *http.Request) {

	fmt.Println("inside post request", r.Body)
	w.Header().Set("content-type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	Query := "insert into apitable(id,name,age) values (%d,'%s',%d)"

	Query = fmt.Sprintf(Query, user.ID, user.Name, user.Age)
	fmt.Println(Query)
	result, err := Postgresdb.Exec(Query)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(400)
		return
	}
	fmt.Println(result)
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected > 0 {
		w.Write([]byte("Created the given resource data"))
		w.WriteHeader(200)
	} else {
		w.Write([]byte("Not Created the given resource data"))
		w.WriteHeader(400)
	}

}

func DeleteFunction(w http.ResponseWriter, r *http.Request) {
	idParamater := chi.URLParam(r, "id")
	w.Header().Set("content-type", "application/json")

	getQuery := "delete from apitable where id = %s"
	getQuery = fmt.Sprintf(getQuery, idParamater)
	fmt.Println(getQuery)
	result, err := Postgresdb.Exec(getQuery)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(400)
		return
	}

	fmt.Println(result)
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected > 0 {
		w.Write([]byte("deleted the given resource data"))
		w.WriteHeader(200)
	} else {
		w.Write([]byte("Resource not found for deletion"))
		w.WriteHeader(204)
	}

}

// still not completed, Need to think on query upsert handling with table column constraints
func PutFunction(w http.ResponseWriter, r *http.Request) {

	fmt.Println("inside put request", r.Body)
	w.Header().Set("content-type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	Query := "insert into apitable(id,name,age) values (4,'sathish',34) on conflict (id) do update set id = 5,name='sathish1',age=30"

	Query = fmt.Sprintf(Query, user.ID, user.Name, user.Age)
	fmt.Println(Query)
	result, err := Postgresdb.Exec(Query)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(400)
		return
	}
	fmt.Println(result)
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected > 0 {
		w.Write([]byte("Created the given resource data"))
		w.WriteHeader(200)
	} else {
		w.Write([]byte("Not Created the given resource data"))
		w.WriteHeader(400)
	}

}

func PatchFunction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside patch request", r.Body)
	w.Header().Set("content-type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	query := "update apitable $set name='Given name in request body' where id = Given id in request url"
	fmt.Println(query)
}

type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}
