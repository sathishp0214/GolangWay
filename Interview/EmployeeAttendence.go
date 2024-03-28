package main

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var EmpAttendenceID = 0

var PostgresClient, _ = sqlx.Connect("postgres", "user=postgres password=234403 dbname=dvdrental sslmode=disable")

func main() {

	officeAttendence, employeeAttendence := EmployeeRegistry(100)
	employeeAttendence.CheckInEmployee()
	employeeAttendence.CheckOutEmployee()
	employeeAttendence.InsertEmployeeAttendence()
	officeAttendence.Date = time.Now()
	officeAttendence.EmpID = employeeAttendence.EmpID
	officeAttendence.InsertOfficeAttendence()

	officeAttendence, employeeAttendence = EmployeeRegistry(102)
	employeeAttendence.CheckInEmployee()
	employeeAttendence.CheckOutEmployee()
	employeeAttendence.InsertEmployeeAttendence()
	officeAttendence.Date = time.Now()
	officeAttendence.EmpID = employeeAttendence.EmpID
	officeAttendence.InsertOfficeAttendence()

	officeAttendence, employeeAttendence = EmployeeRegistry(102)
	employeeAttendence.CheckInEmployee()
	employeeAttendence.CheckOutEmployee()
	employeeAttendence.InsertEmployeeAttendence()
	officeAttendence.Date = time.Now()
	officeAttendence.EmpID = employeeAttendence.EmpID
	officeAttendence.InsertOfficeAttendence()

	//Generating attendence report for all employees
	officeAttendence.AttendenceReport()

}

func EmployeeRegistry(ID int) (OfficeAttendence, EmployeeAttendence) {

	officeAttendence := OfficeAttendence{Date: time.Now(), EmpID: ID}
	EmployeeAttendence := EmployeeAttendence{Id: EmpAttendenceID + 1, EmpID: ID}

	return officeAttendence, EmployeeAttendence

}

func (i *OfficeAttendence) InsertOfficeAttendence() {
	Query := "insert into OfficeAttendence values(%s,%s)"
	Query = fmt.Sprintf(Query, i.Date, i.EmpID)
	PostgresClient.MustExec(Query)
}

func (i *EmployeeAttendence) InsertEmployeeAttendence() {
	Query := "insert into EmployeeAttendence values(%s,%s,%s,%s )"
	Query = fmt.Sprintf(Query, i.Id, i.EmpID, i.CheckIn, i.CheckOut)
	PostgresClient.MustExec(Query)
}

func (i *OfficeAttendence) AttendenceReport() {
	Query := `select
	sum(checkOut-checkIn),
	OfficeAttendence.empId
from
	OfficeAttendence
left join EmployeeAttendence on OfficeAttendence.empId = EmployeeAttendence.Emp_id 
group by OfficeAttendence.empId`

	rows, error := PostgresClient.Queryx(Query)
	fmt.Println(error)
	var reports []Report
	for rows.Next() {
		var report Report
		rows.StructScan(&report)
		reports = append(reports, report)
	}

	fmt.Println("Overall office Attendence report -------", reports)

}

type Report struct {
	// TotalHours int `db:"TotalHours"`
	EmpID int `db:"empid"`
}

type OfficeAttendence struct {
	Date time.Time `db:"officeDate"`
	// TotalNumberOFHours int
	EmpID int `db:"empId"`
}

type EmployeeAttendence struct {
	Id    int `db:"id"`
	EmpID int `db:"Emp_id"`
	// Name     string
	CheckIn  time.Time `db:"checkIn`
	CheckOut time.Time `db:"checkOut`
}

func (attendance *EmployeeAttendence) CheckInEmployee() {
	attendance.CheckIn = time.Now()

}

func (attendance *EmployeeAttendence) CheckOutEmployee() {
	attendance.CheckOut = time.Now().Add(4 * time.Hour)
}
