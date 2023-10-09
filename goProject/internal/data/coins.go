package data

import (
	"database/sql"
	"goProject/internal/validator"
	"time"
)

type Coin struct {
	ID           int64        `json:"id"`
	CreatedAt    time.Time    `json:"-"`
	Title        string       `json:"title,omitempty"`
	Description  string       `json:"description,omitempty"`
	Country      string       `json:"country,omitempty"`
	Status       string       `json:"status,omitempty"`
	Quantity     int64        `json:"quantity,omitempty"`
	Material     string       `json:"material,omitempty"`
	AuctionValue AuctionValue `json:"auction_value,omitempty"`
}

func ValidateCoin(v *validator.Validator, coin *Coin) {
	v.Check(coin.Title != "", "title", "must be provided")
	v.Check(len(coin.Title) <= 256, "title", "must not be more than 256 bytes long")

	v.Check(len(coin.Description) <= 1024, "title", "must not be more than 1024 bytes long")

	v.Check(len(coin.Status) <= 256, "status", "must not be more than 256 bytes long")
	v.Check(coin.Quantity > 0, "quantity", "cannot be equal to zero")
	v.Check(len(coin.Material) <= 256, "material", "must not be more than 256 bytes long")
}

type CoinModel struct {
	DB *sql.DB
}

func (c CoinModel) Insert(coin *Coin) error {
	return nil
}

func (c CoinModel) Get(id int64) (*Coin, error) {
	return nil, nil
}

func (c CoinModel) Update(coin *Coin) error {
	return nil
}

func (c CoinModel) Delete(id int64) error {
	return nil
}
