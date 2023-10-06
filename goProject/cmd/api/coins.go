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

	coin := &data.Coin{
		Title:        input.Title,
		Description:  input.Description,
		Year:         input.Year,
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
		Title:        "Coin " + strconv.FormatInt(id, 10),
		Description:  "Coin's description",
		Year:         2023,
		Country:      "Coin's country",
		Status:       "Coin's status of usability",
		Quantity:     1,
		Material:     "Coin's material",
		AuctionValue: math.MaxInt,
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"coin": coin}, nil)
	if err != nil {
		app.logger.Println(err)
		app.serverErrorResponse(w, r, err)
	}
}
