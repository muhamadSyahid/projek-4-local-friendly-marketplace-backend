package entities

import "time"

type Marketplace struct {
	ID           string     `json:"id" bson:"_id,omitempty"`
	Name         string     `json:"name" bson:"name"`
	Description  *string    `json:"description,omitempty" bson:"description,omitempty"`
	LogoUrl      *string    `json:"logoUrl,omitempty" bson:"logoUrl,omitempty"`
	WebsiteUrl   *string    `json:"websiteUrl,omitempty" bson:"websiteUrl,omitempty"`
	ContactEmail *string    `json:"contactEmail,omitempty" bson:"contactEmail,omitempty"`
	ContactPhone *string    `json:"contactPhone,omitempty" bson:"contactPhone,omitempty"`
	Address      *string    `json:"address,omitempty" bson:"address,omitempty"`
	IsActive     bool       `json:"isActive" bson:"isActive"`
	CreatedAt    time.Time  `json:"createdAt" bson:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt" bson:"updatedAt"`
	LastSyncedAt *time.Time `json:"lastSyncedAt,omitempty" bson:"lastSyncedAt,omitempty"`
	IsSynced     bool       `json:"isSynced" bson:"isSynced"`
}
