package handlers

import (
	"fmt"
	"net/http"
)

type Order struct{}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) { 
	fmt.Println("Create an order")
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) { 
	fmt.Println("List all orders")
}

func (o *Order) GetByID(w http.ResponseWriter, r *http.Request) { 
	fmt.Println("Get an order by id")
}

func (o *Order) UpdateByID(w http.ResponseWriter, r *http.Request) { 
	fmt.Println("Update an order by id")
}

func (o *Order) DeleteByID(w http.ResponseWriter, r *http.Request) { 
	fmt.Println("Delete an order by id")
}