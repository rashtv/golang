package data

import (
	"fmt"
	"strconv"
)

type AuctionValue int64

func (av AuctionValue) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d USD", av)

	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil
}
