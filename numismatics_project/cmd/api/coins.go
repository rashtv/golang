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
		Title        *string            `json:"title"`
		Description  *string            `json:"description"`
		Country      *string            `json:"country"`
		Status       *string            `json:"status"`
		Quantity     *int64             `json:"quantity"`
		Material     *string            `json:"material"`
		AuctionValue *data.AuctionValue `json:"auction_value"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	if input.Title != nil {
		coin.Title = *input.Title
	}

	if input.Description != nil {
		coin.Description = *input.Description
	}

	if input.Country != nil {
		coin.Country = *input.Country
	}

	if input.Status != nil {
		coin.Status = *input.Status
	}

	if input.Quantity != nil {
		coin.Quantity = *input.Quantity
	}

	if input.Material != nil {
		coin.Material = *input.Material
	}

	if input.AuctionValue != nil {
		coin.AuctionValue = *input.AuctionValue
	}

	v := validator.New()

	if data.ValidateCoin(v, coin); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Coins.Update(coin)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
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

func (app *application) deleteCoinHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.Coins.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusNoContent, envelope{"message": "coin successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listCoinsHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string `json:"title"`
		Country string `json:"country"`
		data.Filters
	}

	v := validator.New()

	qs := r.URL.Query()

	input.Title = app.readString(qs, "title", "")
	input.Country = app.readString(qs, "country", "")

	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)
	input.Filters.Sort = app.readString(qs, "sort", "id")
	input.Filters.SortSafelist = []string{
		"id", "title", "country", "quantity", "auction_value",
		"-id", "-title", "-country", "-quantity", "-auction_value"}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	coins, err := app.models.Coins.GetAll(input.Title, input.Country, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"coins": coins}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
