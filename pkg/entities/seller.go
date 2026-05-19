package entities

import "time"

type Seller struct {
	ID              string     `json:"id" bson:"_id,omitempty"`
	UserID          string     `json:"userId" bson:"userId"`
	ShopName        string     `json:"shopName" bson:"shopName"`
	ShopDescription *string    `json:"shopDescription" bson:"shopDescription,omitempty"`
	ShopImageURL    *string    `json:"shopImageUrl" bson:"shopImageUrl,omitempty"`
	ShopAddress     *string    `json:"shopAddress" bson:"shopAddress,omitempty"`
	ShopPhone       *string    `json:"shopPhone" bson:"shopPhone,omitempty"`
	Location        *Location  `json:"location" bson:"location,omitempty"`
	Categories      []string   `json:"categories" bson:"categories"`
	Rating          float64    `json:"rating" bson:"rating"`
	TotalReviews    int        `json:"totalReviews" bson:"totalReviews"`
	TotalProducts   int        `json:"totalProducts" bson:"totalProducts"`
	IsVerified      bool       `json:"isVerified" bson:"isVerified"`
	IsActive        bool       `json:"isActive" bson:"isActive"`
	IsOnline        bool       `json:"isOnline" bson:"isOnline"`
	CreatedAt       time.Time  `json:"createdAt" bson:"createdAt"`
	UpdatedAt       time.Time  `json:"updatedAt" bson:"updatedAt"`
	LastSyncedAt    *time.Time `json:"lastSyncedAt" bson:"lastSyncedAt,omitempty"`
	IsSynced        bool       `json:"isSynced" bson:"isSynced"`
}
