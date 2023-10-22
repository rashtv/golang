package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Coins CoinModel
}

func NewModel(db *sql.DB) Models {
	return Models{
		Coins: CoinModel{DB: db},
	}
}
