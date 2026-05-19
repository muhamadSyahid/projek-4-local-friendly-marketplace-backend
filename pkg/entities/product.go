package entities

import "time"

type Product struct {
	ID           string     `json:"id" bson:"_id,omitempty"`
	SellerID     string     `json:"sellerId" bson:"sellerId"`
	Name         string     `json:"name" bson:"name"`
	Description  string     `json:"description" bson:"description"`
	Price        float64    `json:"price" bson:"price"`
	Quantity     int        `json:"quantity" bson:"quantity"`
	Category     string     `json:"category" bson:"category"`
	Images       []string   `json:"images,omitempty" bson:"images,omitempty"`
	SKU          *string    `json:"sku,omitempty" bson:"sku,omitempty"`
	Weight       *float64   `json:"weight,omitempty" bson:"weight,omitempty"`
	Unit         *string    `json:"unit,omitempty" bson:"unit,omitempty"`
	IsAvailable  bool       `json:"isAvailable" bson:"isAvailable"`
	CreatedAt    time.Time  `json:"createdAt" bson:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt" bson:"updatedAt"`
	LastSyncedAt *time.Time `json:"lastSyncedAt,omitempty" bson:"lastSyncedAt,omitempty"`
	IsSynced     bool       `json:"isSynced" bson:"isSynced"`
	IsLocalOnly  bool       `json:"isLocalOnly" bson:"isLocalOnly"`
}

func (p Product) NeedsSync() bool {
	return !p.IsSynced || p.IsLocalOnly
}

func (p Product) InStock() bool {
	return p.IsAvailable && p.Quantity > 0
}

func (p Product) StockStatus() string {
	if !p.IsAvailable {
		return "Tidak Tersedia"
	}
	if p.Quantity <= 0 {
		return "Stok Habis"
	}
	if p.Quantity < 5 {
		return "Stok Terbatas"
	}
	return "Stok Tersedia"
}
