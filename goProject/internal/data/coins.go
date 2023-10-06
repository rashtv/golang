package data

import (
	"time"
)

type Coin struct {
	ID           int64        `json:"id"`
	CreatedAt    time.Time    `json:"-"`
	Title        string       `json:"title,omitempty"`
	Year         Year         `json:"year,omitempty"`
	Country      string       `json:"country,omitempty"`
	Status       string       `json:"status,omitempty"`
	Quantity     int64        `json:"quantity,omitempty"`
	Material     string       `json:"material,omitempty"`
	AuctionValue AuctionValue `json:"auction_value,omitempty"`
}
