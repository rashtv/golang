package data

import (
	"database/sql"
	"errors"
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
	Version      int32        `json:"version"`
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
	query := `
		INSERT INTO coins (title, description, country, status, quantity, material, auction_value)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at, version`

	args := []interface{}{
		coin.Title,
		coin.Description,
		coin.Country,
		coin.Status,
		coin.Quantity,
		coin.Material,
		coin.AuctionValue,
	}

	return c.DB.QueryRow(query, args...).Scan(&coin.ID, &coin.CreatedAt, &coin.Version)
}

func (c CoinModel) Get(id int64) (*Coin, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `SELECT * FROM coins WHERE id = $1`

	var coin Coin

	err := c.DB.QueryRow(query, id).Scan(
		&coin.ID,
		&coin.CreatedAt,
		&coin.Title,
		&coin.Description,
		&coin.Country,
		&coin.Status,
		&coin.Quantity,
		&coin.Material,
		&coin.AuctionValue,
		&coin.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &coin, nil
}

func (c CoinModel) Update(coin *Coin) error {
	return nil
}

func (c CoinModel) Delete(id int64) error {
	return nil
}
