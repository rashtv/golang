package main

import (
	"fmt"
	"goProject/internal/data"
	"net/http"
	"time"
)

func (app *application) createCoinHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new coin")
}

func (app *application) showCoinHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	coin := data.Coin{
		ID:           id,
		CreatedAt:    time.Now(),
		Title:        "100 KZT",
		Country:      "Kazakhstan",
		Status:       "Still in Production",
		Quantity:     100000,
		Material:     "Base metal alloys",
		AuctionValue: 1000,
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"coin": coin}, nil)
	if err != nil {
		app.logger.Println(err)
		app.serverErrorResponse(w, r, err)
	}
}
