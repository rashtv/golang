package data

import (
	"context"
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

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	return c.DB.QueryRowContext(ctx, query, args...).Scan(&coin.ID, &coin.CreatedAt, &coin.Version)
}

func (c CoinModel) Get(id int64) (*Coin, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
		SELECT id,
		       created_at,
		       title,
		       description,
		       country,
		       status,
		       quantity,
		       material,
		       auction_value,
		       version
		FROM coins
		WHERE id = $1`

	var coin Coin

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	err := c.DB.QueryRowContext(ctx, query, id).Scan(
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
	query := `
		UPDATE coins
		SET title = $1,
		    description = $2,
		    country = $3,
		    status = $4,
		    quantity = $5,
		    material = $6,
		    auction_value = $7,
		    version = version + 1
		WHERE id = $8 AND version = $6
		RETURNING version`

	args := []interface{}{
		coin.Title,
		coin.Description,
		coin.Country,
		coin.Status,
		coin.Quantity,
		coin.Material,
		coin.AuctionValue,
		coin.ID,
		coin.Version,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	err := c.DB.QueryRowContext(ctx, query, args...).Scan(&coin.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	return nil
}

func (c CoinModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
		DELETE FROM coins
		WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	result, err := c.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}

func (c CoinModel) GetAll(title string, country string, filters Filters) ([]*Coin, error) {
	query := `
		SELECT *
		FROM coins
		WHERE (to_tsvector('simple', title) @@ plainto_tsquery('simple', $1) OR $1 = '')
		AND (LOWER(country) = LOWER($2) OR $2 = '')
		ORDER BY id`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := c.DB.QueryContext(ctx, query, title, country)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var coins []*Coin

	for rows.Next() {
		var coin Coin

		err := rows.Scan(
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
			return nil, err
		}
		coins = append(coins, &coin)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return coins, nil
}
