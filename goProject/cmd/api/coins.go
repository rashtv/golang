package main

import (
	"errors"
	"fmt"
	"goProject/internal/data"
	"goProject/internal/validator"
	"net/http"
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

	coin, err := app.models.Coins.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"coin": coin}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateCoinHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	coin, err := app.models.Coins.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		Title        string            `json:"title"`
		Description  string            `json:"description"`
		Country      string            `json:"country"`
		Status       string            `json:"status"`
		Quantity     int64             `json:"quantity"`
		Material     string            `json:"material"`
		AuctionValue data.AuctionValue `json:"auction_value"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	coin.Title = input.Title
	coin.Description = input.Description
	coin.Country = input.Country
	coin.Status = input.Status
	coin.Quantity = input.Quantity
	coin.Material = input.Material
	coin.AuctionValue = input.AuctionValue

	v := validator.New()

	if data.ValidateCoin(v, coin); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Coins.Update(coin)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"coin": coin}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
