package api

import (
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/likimiad/wbtech0/api/docs"
	"github.com/likimiad/wbtech0/internal/database"
	httpSwagger "github.com/swaggo/http-swagger"
	"log/slog"
	"net/http"
	"time"
)

type Server struct {
	DB          *database.OrderService
	Router      *mux.Router
	Timeout     time.Duration
	IdleTimeout time.Duration
}

func getServer(db *database.OrderService) *Server {
	server := &Server{
		DB:     db,
		Router: mux.NewRouter(),
	}
	server.routes()
	return server
}

func NewServer(db *database.OrderService) *Server {
	defer func(start time.Time) {
		slog.Info("server initialized and routes configured", "duration", time.Since(start))
	}(time.Now())
	return getServer(db)
}

func (s *Server) Start(address string) error {
	slog.Info(fmt.Sprintf("The server is launched and available at http://%s", address))
	return http.ListenAndServe(address, s.Router)
}

func (s *Server) routes() {
	s.Router.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)

	s.Router.Handle("/get_orders", s.logger(s.handleGetOrders())).Methods("GET")
	s.Router.Handle("/get_order/{order_uid}", s.logger(s.handleGetOrder())).Methods("GET")
}

func (s *Server) logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func(start time.Time) {
			slog.Info("server request", "method", r.Method, "address", r.RemoteAddr, "url", r.URL.Path, "duration", time.Since(start))
		}(time.Now())
		next.ServeHTTP(w, r)
	})
}
