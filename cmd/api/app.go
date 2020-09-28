package main

import (
	"database/sql"
	"net"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/risentveber/example-di-go/services/clients"
	"github.com/risentveber/example-di-go/services/orders"
)

type App struct {
	listener net.Listener
	server   *http.Server
}

func (a *App) Start() error {
	return a.server.Serve(a.listener)
}

func NewDB() (*sql.DB, error) {
	return sql.Open("postgres", os.Getenv("DB_PATH"))
}

func NewListener() (net.Listener, error) {
	return net.Listen("tcp", ":8080")
}

func NewRouter(clientsHandlers clients.HandlersHTTP, ordersHandlers orders.HandlersHTTP) http.Handler {
	r := mux.NewRouter().StrictSlash(true)
	r.Handle("/v1/clients/{id}/", clientsHandlers.Find).Methods("GET")
	r.Handle("/v1/orders/{id}/", ordersHandlers.Delete).Methods("DELETE")
	r.Handle("/v1/orders/{id}/", ordersHandlers.Find).Methods("GET")
	return r
}

func NewApp() (*App, error) {
	app := &App{}
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
	app.listener, err = NewListener()
	app.server = &http.Server{Handler: router}
	return app, err
}
