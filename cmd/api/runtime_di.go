package main

import (
	"go.uber.org/dig"

	_ "github.com/lib/pq"
	"github.com/risentveber/example-di-go/services/clients"
	"github.com/risentveber/example-di-go/services/orders"
)

func BindOrdersRepository(sqlRepo *orders.SQLRepository) orders.Repository {
	return sqlRepo
}
func BindClientsRepository(sqlRepo *clients.SQLRepository) clients.Repository {
	return sqlRepo
}
func BindClientsService(service *clients.PlainService) clients.Service {
	return service
}
func BindOrdersService(service *orders.PlainService) orders.Service {
	return service
}

func BuildInRuntime() (*App, error) {
	c := dig.New()
	servicesConstructors := []interface{}{
		NewDB,
		orders.NewRepository,
		orders.NewService,
		orders.NewHTTPProvider,
		clients.NewRepository,
		clients.NewService,
		clients.NewHTTPProvider,
		BindOrdersRepository,
		BindClientsRepository,
		BindClientsService,
		BindOrdersService,
		NewRouter,
		NewListener,
		NewApp,
	}

	for _, service := range servicesConstructors {
		if err := c.Provide(service); err != nil {
			return nil, err
		}
	}

	var result *App
	err := c.Invoke(func(a *App) {
		result = a
	})
	return result, err
}
