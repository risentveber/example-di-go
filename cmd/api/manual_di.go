package main

import (
	"github.com/risentveber/example-di-go/services/clients"
	"github.com/risentveber/example-di-go/services/orders"
)

func BuildManually() (*App, error) {
	db, err := NewDB()
	if err != nil {
		return nil, err
	}
	ordersRepo := orders.NewRepository(db)
	ordersSvc := orders.NewService(ordersRepo)
	ordersHTTP := orders.NewHTTPProvider(ordersSvc)
	clientsRepo := clients.NewRepository(db)
	clientsSvc := clients.NewService(clientsRepo)
	clientsHTTP := clients.NewHTTPProvider(clientsSvc)
	router := NewRouter(clientsHTTP, ordersHTTP)
	listener, err := NewListener()
	if err != nil {
		return nil, err
	}

	return NewApp(listener, router), err
}
