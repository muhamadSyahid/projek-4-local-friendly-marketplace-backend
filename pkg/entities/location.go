package entities

type Location struct {
	Latitude   float64 `json:"latitude" bson:"latitude"`
	Longitude  float64 `json:"longitude" bson:"longitude"`
	Address    *string `json:"address,omitempty" bson:"address,omitempty"`
	City       *string `json:"city,omitempty" bson:"city,omitempty"`
	Province   *string `json:"province,omitempty" bson:"province,omitempty"`
	PostalCode *string `json:"postalCode,omitempty" bson:"postalCode,omitempty"`
	Country    *string `json:"country,omitempty" bson:"country,omitempty"`
}
