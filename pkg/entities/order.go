package entities

import (
	"time"
)

type OrderStatus string

const (
	OrderStatusPending    OrderStatus = "pending"
	OrderStatusConfirmed  OrderStatus = "confirmed"
	OrderStatusProcessing OrderStatus = "processing"
	OrderStatusShipped    OrderStatus = "shipped"
	OrderStatusDelivered  OrderStatus = "delivered"
	OrderStatusCancelled  OrderStatus = "cancelled"
	OrderStatusRefunded   OrderStatus = "refunded"
)

func ParseOrderStatus(value string) OrderStatus {
	switch value {
	case "confirmed":
		return OrderStatusConfirmed
	case "processing":
		return OrderStatusProcessing
	case "shipped":
		return OrderStatusShipped
	case "delivered":
		return OrderStatusDelivered
	case "cancelled":
		return OrderStatusCancelled
	case "refunded":
		return OrderStatusRefunded
	default:
		return OrderStatusPending
	}
}

func (s OrderStatus) String() string {
	return string(s)
}

type Order struct {
	ID            string      `json:"id" bson:"_id,omitempty"`
	UserID        string      `json:"userId" bson:"userId"`
	SellerID      *string     `json:"sellerId,omitempty" bson:"sellerId,omitempty"`
	MarketplaceID *string     `json:"marketplaceId,omitempty" bson:"marketplaceId,omitempty"`
	Items         []OrderItem `json:"items" bson:"items"`
	Status        OrderStatus `json:"status" bson:"status"`
	Payment       *Payment    `json:"payment,omitempty" bson:"payment,omitempty"`
	Shipment      *Shipment   `json:"shipment,omitempty" bson:"shipment,omitempty"`
	Subtotal      float64     `json:"subtotal" bson:"subtotal"`
	Tax           float64     `json:"tax" bson:"tax"`
	ShippingCost  float64     `json:"shippingCost" bson:"shippingCost"`
	Total         float64     `json:"total" bson:"total"`
	Notes         *string     `json:"notes,omitempty" bson:"notes,omitempty"`
	CreatedAt     time.Time   `json:"createdAt" bson:"createdAt"`
	UpdatedAt     time.Time   `json:"updatedAt" bson:"updatedAt"`
	LastSyncedAt  *time.Time  `json:"lastSyncedAt,omitempty" bson:"lastSyncedAt,omitempty"`
	IsSynced      bool        `json:"isSynced" bson:"isSynced"`
}

func (o Order) IsPaid() bool {
	return o.Payment != nil && o.Payment.IsCompleted()
}

func (o Order) IsShipped() bool {
	return o.Shipment != nil && o.Shipment.Status == "in_transit"
}

func (o Order) IsDelivered() bool {
	return o.Shipment != nil && o.Shipment.Status == "delivered"
}

func (o Order) ItemCount() int {
	total := 0
	for _, item := range o.Items {
		total += item.Quantity
	}
	return total
}

type OrderItem struct {
	ID        string    `json:"id" bson:"_id,omitempty"`
	OrderID   string    `json:"orderId" bson:"orderId"`
	ProductID string    `json:"productId" bson:"productId"`
	Product   Product   `json:"product" bson:"product"`
	Quantity  int       `json:"quantity" bson:"quantity"`
	UnitPrice float64   `json:"unitPrice" bson:"unitPrice"`
	Subtotal  float64   `json:"subtotal" bson:"subtotal"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}
