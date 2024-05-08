package database

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	ID                uint64    `gorm:"column:id;primaryKey;autoIncrement"         json:"id"`
	OrderUID          string    `gorm:"column:order_uid;uniqueIndex"               json:"order_uid"`
	TrackNumber       string    `gorm:"column:track_number;notNull"                json:"track_number"`
	Entry             string    `gorm:"column:entry;notNull"                       json:"entry"`
	Delivery          Delivery  `gorm:"foreignKey:OrderUID;references:OrderUID"    json:"delivery"`
	Payment           Payment   `gorm:"foreignKey:OrderUID;references:OrderUID"    json:"payment"`
	Items             []Item    `gorm:"foreignKey:OrderUID;references:OrderUID"    json:"items"`
	Locale            string    `gorm:"column:locale;notNull"                      json:"locale"`
	InternalSignature string    `gorm:"column:internal_signature"                  json:"internal_signature"`
	CustomerID        uint64    `gorm:"column:customer_id;notNull"                 json:"customer_id"`
	DeliveryService   string    `gorm:"column:delivery_service;notNull"            json:"delivery_service"`
	ShardKey          string    `gorm:"column:shardkey;notNull"                    json:"shard_key"`
	SmID              int       `gorm:"column:sm_id;notNull;check:sm_id >= 0"      json:"sm_id"`
	DateCreated       time.Time `gorm:"column:date_created;notNull"                json:"date_created"`
	OofShard          string    `gorm:"column:oof_shard;notNull"                   json:"oof_shard"`
}

type Delivery struct {
	gorm.Model
	OrderUID string `gorm:"primaryKey;column:order_uid;notNull" json:"order_uid"`
	Name     string `gorm:"column:name;notNull"                 json:"name"`
	Phone    string `gorm:"column:phone;notNull"                json:"phone"`
	Zip      int64  `gorm:"column:zip;notNull"                  json:"zip"`
	City     string `gorm:"column:city;notNull"                 json:"city"`
	Address  string `gorm:"column:address;notNull"              json:"address"`
	Region   string `gorm:"column:region;notNull"               json:"region"`
	Email    string `gorm:"column:email;notNull"                json:"email"`
}

type Payment struct {
	gorm.Model
	OrderUID     string    `gorm:"primaryKey;column:order_uid;notNull"                   json:"order_uid"`
	Transaction  string    `gorm:"column:transaction;notNull"                            json:"transaction"`
	RequestID    uint64    `gorm:"column:request_id"                                     json:"request_id"`
	Currency     string    `gorm:"column:currency;notNull"                               json:"currency"`
	Provider     string    `gorm:"column:provider;notNull"                               json:"provider"`
	Amount       int64     `gorm:"column:amount;notNull;check:amount >= 0"               json:"amount"`
	PaymentDT    time.Time `gorm:"column:payment_dt;notNull"                             json:"payment_dt"`
	Bank         string    `gorm:"column:bank;notNull"                                   json:"bank"`
	DeliveryCost int64     `gorm:"column:delivery_cost;notNull;check:delivery_cost >= 0" json:"delivery_cost"`
	GoodsTotal   int64     `gorm:"column:goods_total;notNull;check:goods_total >= 0"     json:"goods_total"`
	CustomFee    int64     `gorm:"column:custom_fee;notNull;check:custom_fee >= 0"       json:"custom_fee"`
}

type Item struct {
	gorm.Model
	ItemID      uint   `gorm:"primaryKey;column:item_id"                         json:"item_id"`
	OrderUID    string `gorm:"column:order_uid;notNull"                          json:"order_uid"`
	ChrtID      uint64 `gorm:"column:chrt_id;notNull;check:chrt_id >= 0"         json:"chrt_id"`
	TrackNumber string `gorm:"column:track_number;notNull"                       json:"track_number"`
	Price       int64  `gorm:"column:price;notNull;check:price >= 0"             json:"price"`
	Rid         string `gorm:"column:rid;notNull"                                json:"rid"`
	Name        string `gorm:"column:name;notNull"                               json:"name"`
	Sale        int64  `gorm:"column:sale;check:Sale >= 0"                       json:"sale"`
	Size        string `gorm:"column:size;notNull"                               json:"size"`
	TotalPrice  int64  `gorm:"column:total_price;notNull;check:total_price >= 0" json:"total_price"`
	NmID        int64  `gorm:"column:nm_id;notNull;check:nm_id >= 0"             json:"nm_id"`
	Brand       string `gorm:"column:brand;notNull"                              json:"brand"`
	Status      int64  `gorm:"column:status;notNull;check:status >= 0"           json:"status"`
}
