package data

import (
	"fmt"
	"strconv"
	"strings"
)

type AuctionValue int64

func (av AuctionValue) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d USD", av)

	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil
}

func (av *AuctionValue) UnmarshalJSON(jsonValue []byte) error {
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidYearFormat
	}

	parts := strings.Split(unquotedJSONValue, " ")

	if len(parts) != 2 || parts[1] != "USD" {
		return ErrInvalidYearFormat
	}

	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidYearFormat
	}

	*av = AuctionValue(i)

	return nil
}
