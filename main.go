package main

import "customer/ctcompany"
import "fmt"

func main() {

	//insert
	newCustomer := ctcompany.Customer{
		FirstName: "Caner",
		LastName:  "Tosuner",
		Email:     "canertosuner@gmail.com",
	}
	ctcompany.Insert(newCustomer)


	//update
	updateCustomer := ctcompany.Customer{
		FirstName: "Berker",
		LastName:  "Sönmez",
		Email:     "berkersonmez@gmail.com",
		Id:        3,
	}
	updatedCustomer := ctcompany.Update(updateCustomer)
	fmt.Println(updatedCustomer)


	//selectAll
	customerList := ctcompany.SelectAll()
	fmt.Println(customerList)


	//getById
	cust := ctcompany.GetById(3)
	fmt.Println(cust)
}
