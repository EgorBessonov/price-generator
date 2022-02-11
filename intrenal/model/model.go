// Package model represents all structures which are used in price generator
package model

import "encoding/json"

//Share type represent share object
type Share struct {
	Name      int
	Bid       float32
	Ask       float32
	UpdatedAt string
}

//MarshalBinary realisation for share model
func (sh Share) MarshalBinary() ([]byte, error) {
	return json.Marshal(sh)
}

//UnmarshalBinary realisation for share model
func (sh Share) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &sh)
}
