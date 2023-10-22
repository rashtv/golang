package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidAuctionValueFormat = errors.New("invalid auction value format")

type AuctionValue int64

func (av AuctionValue) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d USD", av)

	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil
}

func (av *AuctionValue) UnmarshalJSON(jsonValue []byte) error {
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidAuctionValueFormat
	}

	parts := strings.Split(unquotedJSONValue, " ")

	if len(parts) != 2 || parts[1] != "USD" {
		return ErrInvalidAuctionValueFormat
	}

	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidAuctionValueFormat
	}

	*av = AuctionValue(i)

	return nil
}
