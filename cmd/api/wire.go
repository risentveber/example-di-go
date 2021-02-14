//+build wireinject

package main

import (
	"github.com/google/wire"

	"github.com/risentveber/example-di-go/services/clients"
	"github.com/risentveber/example-di-go/services/orders"
)

func BuildAppCompileTime() (*App, error) {
	wire.Build(NewDB,
		orders.NewRepository,
		orders.NewService,
		orders.NewHTTPProvider,
		clients.NewRepository,
		clients.NewService,
		clients.NewHTTPProvider,
		wire.Bind(new(orders.Repository), new(*orders.SQLRepository)),
		wire.Bind(new(clients.Repository), new(*clients.SQLRepository)),
		wire.Bind(new(clients.Service), new(*clients.PlainService)),
		wire.Bind(new(orders.Service), new(*orders.PlainService)),
		NewRouter,
		NewListener,
		NewApp)
	return &App{}, nil
}
