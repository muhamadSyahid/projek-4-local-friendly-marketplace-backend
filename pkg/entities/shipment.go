package entities

import "time"

type ShippingMethod string

const (
	ShippingMethodStandard  ShippingMethod = "standard"
	ShippingMethodExpress   ShippingMethod = "express"
	ShippingMethodOvernight ShippingMethod = "overnight"
)

// Value returns the string value of the shipping method
func (s ShippingMethod) Value() string {
	switch s {
	case ShippingMethodStandard:
		return "standard"
	case ShippingMethodExpress:
		return "express"
	case ShippingMethodOvernight:
		return "overnight"
	default:
		return "standard"
	}
}

type ShippingStatus string

const (
	ShippingStatusPending    ShippingStatus = "pending"
	ShippingStatusProcessing ShippingStatus = "processing"
	ShippingStatusShipped    ShippingStatus = "shipped"
	ShippingStatusInTransit  ShippingStatus = "in_transit"
	ShippingStatusDelivered  ShippingStatus = "delivered"
	ShippingStatusFailed     ShippingStatus = "failed"
)

// Value returns the string value of the shipping status
func (s ShippingStatus) Value() string {
	switch s {
	case ShippingStatusPending:
		return "pending"
	case ShippingStatusProcessing:
		return "processing"
	case ShippingStatusShipped:
		return "shipped"
	case ShippingStatusInTransit:
		return "in_transit"
	case ShippingStatusDelivered:
		return "delivered"
	case ShippingStatusFailed:
		return "failed"
	default:
		return "pending"
	}
}

type Shipment struct {
	ID                    string         `json:"id" bson:"_id,omitempty"`
	OrderID               string         `json:"orderId" bson:"orderId"`
	Method                ShippingMethod `json:"method" bson:"method"`
	Status                ShippingStatus `json:"status" bson:"status"`
	TrackingNumber        *string        `json:"trackingNumber" bson:"trackingNumber,omitempty"`
	Carrier               *string        `json:"carrier" bson:"carrier,omitempty"`
	ShippingAddress       *string        `json:"shippingAddress" bson:"shippingAddress,omitempty"`
	EstimatedDeliveryDate time.Time      `json:"estimatedDeliveryDate" bson:"estimatedDeliveryDate"`
	ActualDeliveryDate    *time.Time     `json:"actualDeliveryDate" bson:"actualDeliveryDate,omitempty"`
	Notes                 *string        `json:"notes" bson:"notes,omitempty"`
	CreatedAt             time.Time      `json:"createdAt" bson:"createdAt"`
	UpdatedAt             time.Time      `json:"updatedAt" bson:"updatedAt"`
}

// IsDelivered checks if shipment is delivered
func (s *Shipment) IsDelivered() bool {
	return s.Status == ShippingStatusDelivered
}

// IsInTransit checks if shipment is in transit
func (s *Shipment) IsInTransit() bool {
	return s.Status == ShippingStatusInTransit
}
