package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Coins       CoinModel
	Permissions PermissionModel
	Tokens      TokenModel
	Users       UserModel
}

func NewModel(db *sql.DB) Models {
	return Models{
		Coins:       CoinModel{DB: db},
		Permissions: PermissionModel{DB: db},
		Tokens:      TokenModel{DB: db},
		Users:       UserModel{DB: db},
	}
}
