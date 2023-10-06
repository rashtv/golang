package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidYearFormat = errors.New("invalid year format")

type Year int32

func (y Year) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d year", y)

	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil
}

func (y *Year) UnmarshalJSON(jsonValue []byte) error {
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidYearFormat
	}

	parts := strings.Split(unquotedJSONValue, " ")

	if len(parts) != 2 || parts[1] != "year" {
		return ErrInvalidYearFormat
	}

	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidYearFormat
	}

	*y = Year(i)

	return nil
}
