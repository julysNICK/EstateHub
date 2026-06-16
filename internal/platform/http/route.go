package http

//OLHE O USER PACKAGE PARA VER O QUE FOI EDITADO POIS ESTÁ 100%
import (
	"log/slog"
	nethttp "net/http"
)

type LeadHandler interface {
	List(nethttp.ResponseWriter, *nethttp.Request)
	GetByID(nethttp.ResponseWriter, *nethttp.Request)
	Create(nethttp.ResponseWriter, *nethttp.Request)
	Update(nethttp.ResponseWriter, *nethttp.Request)
	Delete(nethttp.ResponseWriter, *nethttp.Request)
}

type UserHandler interface {
	Create(nethttp.ResponseWriter, *nethttp.Request)
	List(nethttp.ResponseWriter, *nethttp.Request)
	GetByID(nethttp.ResponseWriter, *nethttp.Request)
	Update(nethttp.ResponseWriter, *nethttp.Request)
	Delete(nethttp.ResponseWriter, *nethttp.Request)
}
type PropertyHandler interface {
	List(nethttp.ResponseWriter, *nethttp.Request)
	GetByID(nethttp.ResponseWriter, *nethttp.Request)
	Create(nethttp.ResponseWriter, *nethttp.Request)
	Update(nethttp.ResponseWriter, *nethttp.Request)
	Delete(nethttp.ResponseWriter, *nethttp.Request)
}
type RouterConfig struct {
	Logger          *slog.Logger
	HealthHandler   *HealthHandler
	LeadHandler     LeadHandler
	UserHandler     UserHandler
	PropertyHandler PropertyHandler
}

func NewRouter(cfg RouterConfig) nethttp.Handler {
	mux := nethttp.NewServeMux()

	healthHandler := cfg.HealthHandler
	leadHandler := cfg.LeadHandler
	userHandler := cfg.UserHandler
	propertyHandler := cfg.PropertyHandler

	mux.HandleFunc("/healtz", healthHandler.Healthz)
	mux.HandleFunc("/readyz", healthHandler.Readyz)
	mux.HandleFunc("Get /leads", leadHandler.List)
	mux.HandleFunc("Get /leads/{id}", leadHandler.GetByID)
	mux.HandleFunc("Post /leads", leadHandler.Create)
	mux.HandleFunc("Put /leads/{id}", leadHandler.Update)
	mux.HandleFunc("Delete /leads/{id}", leadHandler.Delete)

	mux.HandleFunc("Post /users", userHandler.Create)
	mux.HandleFunc("Get /users", userHandler.List)
	mux.HandleFunc("Get /users/{id}", userHandler.GetByID)
	mux.HandleFunc("Post /users", userHandler.Create)
	/* mux.HandleFunc("Get /agencies", userHandler.GetAgencies)
	mux.HandleFunc("Get /brokers", userHandler.GetBrokers) */
	mux.HandleFunc("Put /users/{id}", userHandler.Update)
	mux.HandleFunc("Delete /users/{id}", userHandler.Delete)

	mux.HandleFunc("Get /properties", propertyHandler.List)
	mux.HandleFunc("Get /properties/{id}", propertyHandler.GetByID)
	mux.HandleFunc("Post /properties", propertyHandler.Create)
	mux.HandleFunc("Put /properties/{id}", propertyHandler.Update)
	mux.HandleFunc("Delete /properties/{id}", propertyHandler.Delete)

	return RequestLogger(cfg.Logger, mux)
}
