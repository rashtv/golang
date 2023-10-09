package main

import (
	"fmt"
	"goProject/internal/data"
	"goProject/internal/validator"
	"math"
	"net/http"
	"strconv"
	"time"
)

func (app *application) createCoinHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title        string            `json:"title"`
		Description  string            `json:"description"`
		Country      string            `json:"country"`
		Status       string            `json:"status"`
		Quantity     int64             `json:"quantity"`
		Material     string            `json:"material"`
		AuctionValue data.AuctionValue `json:"auction_value"`
		Version      int32             `json:"version"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	coin := &data.Coin{
		Title:        input.Title,
		Description:  input.Description,
		Country:      input.Country,
		Status:       input.Status,
		Quantity:     input.Quantity,
		Material:     input.Material,
		AuctionValue: input.AuctionValue,
		Version:      input.Version,
	}

	v := validator.New()

	if data.ValidateCoin(v, coin); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	err = app.models.Coins.Insert(coin)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("v1/coins/%d", coin.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"coin": coin}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
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
		Title:        "Coin " + strconv.FormatInt(id, 10),
		Description:  "Coin's description",
		Country:      "Coin's country",
		Status:       "Coin's status of usability",
		Quantity:     1,
		Material:     "Coin's material",
		AuctionValue: math.MaxInt,
		Version:      1,
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"coin": coin}, nil)
	if err != nil {
		app.logger.Println(err)
		app.serverErrorResponse(w, r, err)
	}
}
