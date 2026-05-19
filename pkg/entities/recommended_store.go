package entities

import (
	"fmt"
	"slices"
	"strings"
	"time"
)

type RecommendedStore struct {
	ID                string    `json:"id" bson:"_id,omitempty"`
	SellerID          string    `json:"sellerId" bson:"sellerId"`
	ShopName          string    `json:"shopName" bson:"shopName"`
	ShopImageUrl      *string   `json:"shopImageUrl,omitempty" bson:"shopImageUrl,omitempty"`
	Distance          float64   `json:"distance" bson:"distance"`
	Rating            float64   `json:"rating" bson:"rating"`
	TotalReviews      int       `json:"totalReviews" bson:"totalReviews"`
	Location          Location  `json:"location" bson:"location"`
	Categories        []string  `json:"categories,omitempty" bson:"categories,omitempty"`
	IsOnline          bool      `json:"isOnline" bson:"isOnline"`
	AvailableProducts int       `json:"availableProducts" bson:"availableProducts"`
	RecommendedAt     time.Time `json:"recommendedAt" bson:"recommendedAt"`
	CachedAt          time.Time `json:"cachedAt" bson:"cachedAt"`
	IsCached          bool      `json:"isCached" bson:"isCached"`
}

func (r RecommendedStore) isFresh() bool {
	difference := time.Since(r.CachedAt)

	return difference < 24
}

func (r RecommendedStore) formattedDistance() string {
	if r.Distance < 1 {
		return fmt.Sprintf("%.0f m", r.Distance*1000)
	}
	return fmt.Sprintf("%.1f km", r.Distance)
}

func (r RecommendedStore) RatingString() string {
	return fmt.Sprintf("%.1f", r.Rating)
}

func (r RecommendedStore) HasCategory(category string) bool {
	lowerCategory := strings.ToLower(category)
	return slices.Contains(r.Categories, lowerCategory)
}

func (r RecommendedStore) StatusString() string {
	if !r.IsOnline {
		return "Offline"
	}
	if r.AvailableProducts == 0 {
		return "Produk Habis"
	}
	return "Online"
}
