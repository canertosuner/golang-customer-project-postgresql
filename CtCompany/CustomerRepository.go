package ctcompany

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)
func GetById(id int) Customer {

	db := initDB()
	defer db.Close()

	var result Customer

	sqlStatement := `SELECT id,firstname,lastname,email FROM "CtCompany"."Customer" where "id" = $1`

	err := db.QueryRow(sqlStatement, id).Scan(&result.Id,&result.FirstName,&result.LastName,&result.Email)

	checkErr(err)

	return result
}

func SelectAll() []Customer {

	db := initDB()
	defer db.Close()

	var results []Customer

	rows, err := db.Query(`SELECT id,firstname,lastname,email FROM "CtCompany"."Customer"`)

	checkErr(err)

	defer rows.Close()
	for rows.Next() {
		var cstmr Customer
		err := rows.Scan(&cstmr.Id, &cstmr.FirstName, &cstmr.LastName, &cstmr.Email)
		checkErr(err)

		c := append(results, cstmr)
		fmt.Printf("c :", c)
	}
	err = rows.Err()
	checkErr(err)

	return results
}

func Insert(customer Customer) int {

	db := initDB()
	defer db.Close()

	sqlStatement := `INSERT INTO "CtCompany"."Customer" ("firstname", "lastname", "email") 
					VALUES ($1, $2, $3) RETURNING id`
	lastInsertId := 0
	err := db.QueryRow(sqlStatement, customer.FirstName, customer.LastName, customer.Email).Scan(&lastInsertId)

	checkErr(err)
	fmt.Println("New record ID is:", lastInsertId)

	return lastInsertId
}

func Update(customer Customer) Customer {
	var newCustomer Customer
	db := initDB()
	defer db.Close()

	sqlStatement := `UPDATE "CtCompany"."Customer" SET "firstname" = $1, "lastname" = $2, "email" = $3 where "id" = $4 RETURNING id,"firstname","lastname","email";`

	err := db.QueryRow(sqlStatement, customer.FirstName, customer.LastName, customer.Email, customer.Id).Scan(&newCustomer.Id, &newCustomer.FirstName, &newCustomer.LastName, &newCustomer.Email)
	checkErr(err)

	fmt.Println("Row updated")

	return newCustomer
}

func initDB() *sql.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", Host, Port, User, Password, Dbname)

	db, err := sql.Open("postgres", psqlInfo)
	checkErr(err)

	err = db.Ping()
	checkErr(err)

	fmt.Println("Successfully connected!")

	return db
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type Customer struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
}
