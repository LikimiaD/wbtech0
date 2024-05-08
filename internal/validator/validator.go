package validator

import "github.com/likimiad/wbtech0/internal/database"

func checkOrderParams(o *database.Order) bool {
	if o.OrderUID == "" || o.TrackNumber == "" || o.Entry == "" ||
		o.Locale == "" || o.CustomerID == 0 || o.DeliveryService == "" ||
		o.ShardKey == "" || o.SmID < 0 || o.DateCreated.IsZero() || o.OofShard == "" {
		return false
	}
	return true
}

func checkDelivery(d *database.Delivery) bool {
	return !(d.Name == "" || d.Phone == "" || d.Zip == 0 || d.City == "" ||
		d.Address == "" || d.Region == "" || d.Email == "")
}

func checkPayment(p *database.Payment) bool {
	return !(p.Transaction == "" || p.Currency == "" || p.Provider == "" || p.Amount < 0 ||
		p.PaymentDT.IsZero() || p.Bank == "" || p.DeliveryCost < 0 || p.GoodsTotal < 0 || p.CustomFee < 0)
}

func checkItems(items []database.Item) bool {
	for _, item := range items {
		if item.TrackNumber == "" || item.Price < 0 ||
			item.Rid == "" || item.Name == "" || item.Size == "" || item.TotalPrice < 0 ||
			item.NmID < 0 || item.Brand == "" || item.Status < 0 {
			return false
		}
	}
	return true
}

func IsValidOrder(o *database.Order) bool {
	if !checkOrderParams(o) {
		return false
	}
	if !checkDelivery(&o.Delivery) {
		return false
	}
	if !checkPayment(&o.Payment) {
		return false
	}
	if !checkItems(o.Items) {
		return false
	}
	return true
}
