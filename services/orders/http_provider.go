package orders

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Service interface {
	GetOrderByID(id string) (*Order, error)
	DeleteOrderByID(id string) error
}

type HandlersHTTP struct {
	Find   http.Handler
	Delete http.Handler
}

type httpProvider struct {
	svc Service
}

func (p *httpProvider) find(rw http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	order, err := p.svc.GetOrderByID(id)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		_, _ = rw.Write([]byte(err.Error()))
	} else {
		rw.WriteHeader(http.StatusOK)
		rw.Header().Set("Content-Type", "application/json; charset=utf-8")
		_ = json.NewEncoder(rw).Encode(order)
	}
}

func (p *httpProvider) delete(rw http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	err := p.svc.DeleteOrderByID(id)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		_, _ = rw.Write([]byte(err.Error()))
	} else {
		rw.WriteHeader(http.StatusOK)
	}
}

func NewHTTPProvider(svc Service) HandlersHTTP {
	provider := &httpProvider{svc}
	return HandlersHTTP{
		Find:   http.HandlerFunc(provider.find),
		Delete: http.HandlerFunc(provider.delete),
	}
}
