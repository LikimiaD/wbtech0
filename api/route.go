package api

import (
	"errors"
	"github.com/gorilla/mux"
	"github.com/likimiad/wbtech0/internal/database"
	"net/http"
)

var (
	ErrLoadOrders = errors.New("error load orders from database")
	ErrLoadOrder  = errors.New("error load order from database")
)

// @Summary Get list of all orders
// @Description Retrieves a list of all orders from the database
// @Tags orders
// @Accept  json
// @Produce  json
// @Success 200 {array} database.Order "A list of orders"
// @Failure 400 "Bad request if specific error on getting orders"
// @Failure 500 "Internal server error if orders cannot be loaded"
// @Router /get_orders [get]
func (s *Server) handleGetOrders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		orders, err := s.DB.GetAllOrders()
		if err != nil {
			if errors.Is(err, database.ErrGetOrders) {
				s.respondWithError(w, http.StatusBadRequest, err)
			} else {
				s.respondWithError(w, http.StatusInternalServerError, ErrLoadOrders)
			}
			return
		}

		s.respondAny(w, http.StatusOK, orders)
	}
}

// @Summary Get a single order by UID
// @Description Retrieves details of an order by order UID
// @Tags orders
// @Accept  json
// @Produce  json
// @Param   order_uid   path      string     true  "Order UID"
// @Success 200 {object} database.Order "Detailed order information"
// @Failure 400 "Bad request if no order found or other request error"
// @Failure 500 "Internal server error if the order cannot be loaded"
// @Router /get_order/{order_uid} [get]
func (s *Server) handleGetOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		orderUID := vars["order_uid"]

		order, err := s.DB.GetOrder(orderUID)
		if err != nil {
			if errors.Is(err, database.ErrGetOrder) {
				s.respondWithError(w, http.StatusBadRequest, err)
			} else {
				s.respondWithError(w, http.StatusInternalServerError, ErrLoadOrder)
			}
			return
		}
		s.respondAny(w, http.StatusOK, order)
	}
}
