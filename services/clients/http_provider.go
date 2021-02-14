package clients

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Service interface {
	GetClientByID(id string) (*Client, error)
}

type httpProvider struct {
	svc Service
}

type HandlersHTTP struct {
	Find http.Handler
}

func (p *httpProvider) find(rw http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	client, err := p.svc.GetClientByID(id)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		_, _ = rw.Write([]byte(err.Error()))
	} else {
		rw.WriteHeader(http.StatusOK)
		rw.Header().Set("Content-Type", "application/json; charset=utf-8")
		_ = json.NewEncoder(rw).Encode(client)
	}
}

func NewHTTPProvider(svc Service) HandlersHTTP {
	provider := &httpProvider{svc}
	return HandlersHTTP{
		Find: http.HandlerFunc(provider.find),
	}
}
