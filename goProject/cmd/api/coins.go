package main

import (
	"fmt"
	"goProject/internal/data"
	"net/http"
	"time"
)

func (app *application) createCoinHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title        string            `json:"title"`
		Year         data.Year         `json:"year"`
		Country      string            `json:"country"`
		Status       string            `json:"status"`
		Quantity     int64             `json:"quantity"`
		Material     string            `json:"material"`
		AuctionValue data.AuctionValue `json:"auction_value"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
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
