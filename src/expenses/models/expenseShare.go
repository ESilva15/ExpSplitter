package models

import (
	"encoding/json"
	dec "github.com/shopspring/decimal"
)

type Share struct {
	ExpShareID int32       `json:"ExpShareID"`
	User       User        `json:"User"`
	Share      dec.Decimal `json:"Share"`
	Calculated dec.Decimal `json:"Calculated"`
}
type Shares []Share

// ShareFromJSON takes []byte and returns an *Share
func ShareFromJSON(data []byte) (*Share, error) {
	var share Share

	err := json.Unmarshal(data, &share)

	return &share, err
}

func (sh Shares) Equal(other Shares) bool {
	if len(sh) != len(other) {
		return false
	}

	for k := range sh {
		if sh[k].User != other[k].User ||
			!sh[k].Calculated.Equal(other[k].Calculated) ||
			!sh[k].Share.Equal(other[k].Share) {
			return false
		}
	}

	return true
}
