package database

import (
	"errors"
	"gorm.io/gorm"
	"sync"
)

var (
	ErrExists    = errors.New("order already exists")
	ErrGetOrder  = errors.New("no order found with UID")
	ErrGetOrders = errors.New("no orders found")
)

type OrderService struct {
	db    *Database
	Cache map[string]*Order
	mutex sync.RWMutex
}

func NewOrderService(db *Database) *OrderService {
	return &OrderService{
		db:    db,
		Cache: make(map[string]*Order),
	}
}

func (s *OrderService) CreateOrder(order *Order) error {
	if err := s.db.Create(order).Error; err != nil {
		if errors.Is(err, gorm.ErrRegistered) {
			return ErrExists
		}
		return err
	}

	s.cacheOrder(order)
	return nil
}

func (s *OrderService) GetOrder(orderUID string) (*Order, error) {
	s.mutex.RLock()
	if order, exists := s.Cache[orderUID]; exists {
		s.mutex.RUnlock()
		return order, nil
	}
	s.mutex.RUnlock()

	var order Order
	if err := s.db.Preload("Delivery").Preload("Payment").Preload("Items").First(&order, "order_uid = ?", orderUID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrGetOrder
		}
		return nil, err
	}

	s.cacheOrder(&order)
	return &order, nil
}

func (s *OrderService) GetAllOrders() ([]Order, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var orders []Order
	if err := s.db.Preload("Delivery").Preload("Payment").Preload("Items").Find(&orders).Error; err != nil {
		return nil, err
	}

	if len(orders) == 0 {
		return nil, ErrGetOrders
	}

	return orders, nil
}

func (s *OrderService) cacheOrder(order *Order) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.Cache[order.OrderUID] = order
}

func (s *OrderService) RestoreCache() error {
	var orders []Order
	if err := s.db.Preload("Delivery").Preload("Payment").Preload("Items").Find(&orders).Error; err != nil {
		return err
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()
	for _, order := range orders {
		s.Cache[order.OrderUID] = &order
	}
	return nil
}
