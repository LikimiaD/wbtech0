package data_generator

import (
	"fmt"
	"github.com/likimiad/wbtech0/internal/database"
	"math/rand"
	"time"
)

func GenerateRandomOrder() *database.Order {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	trackNumber := fmt.Sprintf("WBILMTESTTRACK%d", rand.Intn(10000))
	order := &database.Order{
		OrderUID:    fmt.Sprintf("b563feb7b2b84b6%d", rand.Intn(100)),
		TrackNumber: trackNumber,
		Entry:       "WBIL",
		Delivery: database.Delivery{
			Name:    "Test Testov",
			Phone:   "+9720000000",
			Zip:     rand.Int63(),
			City:    "Kiryat Mozkin",
			Address: "Ploshad Mira 15",
			Region:  "Kraiot",
			Email:   "test@gmail.com"},
		Payment: database.Payment{
			Transaction:  fmt.Sprintf("b563feb7b2b84b6test%d", rand.Int63()),
			Currency:     "USD",
			Provider:     "wbpay",
			Amount:       rand.Int63(),
			PaymentDT:    time.Now(),
			Bank:         "alpha",
			DeliveryCost: rand.Int63n(2000),
			GoodsTotal:   rand.Int63n(500),
			CustomFee:    rand.Int63n(10)},
		Items: []database.Item{{
			ChrtID:      uint64(rand.Int63n(10000)),
			TrackNumber: trackNumber,
			Price:       rand.Int63(),
			Rid:         "ab4219087a764ae0btest",
			Name:        "Mascaras",
			Sale:        rand.Int63n(100),
			Size:        "0",
			TotalPrice:  rand.Int63n(10000),
			NmID:        rand.Int63(),
			Brand:       "Vivienne Sabo",
			Status:      202}},
		Locale:            "en",
		InternalSignature: "",
		CustomerID:        uint64(rand.Int63n(10000)),
		DeliveryService:   "meest",
		ShardKey:          "9",
		SmID:              rand.Intn(100),
		DateCreated:       time.Now(),
		OofShard:          "1",
	}
	return order
}
