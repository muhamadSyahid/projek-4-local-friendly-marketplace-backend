package entities

import "time"

type StoreCategory struct {
	ID           string    `json:"id" bson:"_id,omitempty"`
	Name         string    `json:"name" bson:"name"`
	Icon         *string   `json:"icon" bson:"icon,omitempty"`
	Description  *string   `json:"description" bson:"description,omitempty"`
	ProductCount int       `json:"productCount" bson:"productCount"`
	CreatedAt    time.Time `json:"createdAt" bson:"createdAt"`
}
