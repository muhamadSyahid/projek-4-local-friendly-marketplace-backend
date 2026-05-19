package entities

import "time"

type User struct {
	ID              string     `json:"id" bson:"_id,omitempty"`
	Name            string     `json:"name" bson:"name"`
	Email           string     `json:"email" bson:"email"`
	Password        string     `json:"password,omitempty" bson:"password,omitempty"`
	Phone           *string    `json:"phone" bson:"phone,omitempty"`
	ProfileImageURL *string    `json:"profileImageUrl" bson:"profileImageUrl,omitempty"`
	Roles           []Roles    `json:"roles" bson:"roles"`
	MarketplaceID   *string    `json:"marketplaceId" bson:"marketplaceId,omitempty"`
	SellerID        *string    `json:"sellerId" bson:"sellerId,omitempty"`
	CreatedAt       time.Time  `json:"createdAt" bson:"createdAt"`
	UpdatedAt       time.Time  `json:"updatedAt" bson:"updatedAt"`
	LastSyncedAt    *time.Time `json:"lastSyncedAt" bson:"lastSyncedAt,omitempty"`
	IsSynced        bool       `json:"isSynced" bson:"isSynced"`
}

// HasRole checks if user has a specific role
func (u *User) HasRole(role Roles) bool {
	for _, r := range u.Roles {
		if r == role {
			return true
		}
	}
	return false
}

// IsSeller checks if user is a seller
func (u *User) IsSeller() bool {
	return u.HasRole(RoleSeller)
}

// IsAdmin checks if user is an admin
func (u *User) IsAdmin() bool {
	return u.HasRole(RoleAdmin)
}

// IsBuyer checks if user is a buyer (default role)
func (u *User) IsBuyer() bool {
	return u.HasRole(RoleBuyer)
}

// PrimaryRole returns the primary role for display purposes
func (u *User) PrimaryRole() Roles {
	if len(u.Roles) > 0 {
		return u.Roles[0]
	}
	return RoleBuyer
}
