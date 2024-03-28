package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //this is postgres driver works within "sqlx" package
)

var (

	//db client connection
	db, _ = sqlx.Connect("postgres", "user=postgres password=234403 dbname=dvdrental sslmode=disable")
)

func main() {

	// db, err := sqlx.Connect("postgres", "user=postgres password=234403 dbname=dvdrental sslmode=disable")
	// fmt.Println(db, err)

	SelectSingleRow()

}

func InsertData() {

	//using fmt.Sprintf -- We can pass the arguments
	query := `
	insert
	into
	address (address_id,
	address,
	address2,
	district,
	city_id,
	postal_code,
	phone,
	last_update,
	dummy)
values (10000,
'address value',
'address value 2',
'chennai',
'500',
'353553',
'354264657788',
'2024-03-02 11:25:11.116',
3000)`

	result := db.MustExec(query)
	fmt.Println("insert result", result, result.RowsAffected)

	Query := "insert into apitable(id,name,age) values (%s,'%s',%s)"
	Query = fmt.Sprintf(Query, user.ID, user.Name, user.Age)

	result := Postgresdb.MustExec(Query)
	fmt.Println(result)
}

func UpdateQuery() {
	query := `update address set dummy = 1000 where dummy is null;`
	result := db.MustExec(query)
	fmt.Println("update result", result, result.RowsAffected)

}

func SelectSingleRow() {
	// //queries only one row
	// var err error
	// var address Address
	// inputDistrict := "Texas"

	// //passing argument into query
	// query := "select * from address where district='%s'"
	// query = fmt.Sprintf(query, inputDistrict)
	// row := db.QueryRowx(query)
	// err = row.StructScan(&address)
	// fmt.Println("Single row Data ------", address, err)

	getQuery := "select * from apitable where id = %s"
	getQuery = fmt.Sprintf(getQuery, "1")
	fmt.Println(getQuery)
	row := db.QueryRowx(getQuery)
	fmt.Println(row)
	var user User
	err := row.StructScan(&user)
	fmt.Println(err, user)

}

func SelectMultipleRows() {
	rows, err := db.Queryx("select * from address limit 5")

	//gets column info
	// fmt.Println(rows.Columns())
	// fmt.Println(rows.ColumnTypes())

	for rows.Next() {
		var address Address

		//another method is parsing sql data into golang struct
		// err = rows.Scan(&address.AddressId, &address.Address, &address.Address2, &address.District, &address.cityID, &address.PostalCode, &address.Phone, &address.Phone, &address.Dummy)
		// fmt.Println(err)
		// fmt.Println("result", address, reflect.TypeOf(address.cityID))

		err = rows.StructScan(&address)
		fmt.Println(err)
		fmt.Println("result", address)
	}

}

type Address struct {
	AddressId   int    `db:"address_id"`
	Address     string `db:"address"`
	Address2    string `db:"address2"`
	District    string `db:"district"`
	CityID      int    `db:"city_id"`
	PostalCode  string `db:"postal_code"`
	Phone       string `db:"phone"`
	LastUpdated string `db:"last_update"`
	Dummy       string `db:"dummy"`
}

type User1 struct {
	ID        int
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
}

type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

/*
SQL relationship in golang structs:

With orm usage-- We have Both structs and its fields and table and its columns are same.
With raw query usage -- We may have different struct fields and different table columns, Its about in .go file, how we are using raw sql queries like join queries or multiple select queries from different tables and how we are processing and stores the data into golang structs

type Publisher struct {
	Id int
	Name string
	Books []Book
}

type Book struct {
	Name string
	publisherID int
}

In this scenario,
i) USing two select queries on two tables publisher and books -- We can select query in publisher table particular publisher id and loads data into publisher struct. Then again query with Books table and appends books rows data with Publisher structs Books slice.

ii)Using single select join query -- Loads the data into above structs relationship
*/
