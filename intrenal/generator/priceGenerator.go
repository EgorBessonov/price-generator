//Package generator represents price generation in service
package generator

import (
	"encoding/json"
	"github.com/EgorBessonov/price-generator/intrenal/model"
	"math/rand"
	"time"
)

const (
	randSeed = 5
	AMAZON   = iota
	APPLE
	MICROSOFT
	NETFLIX
	PFIZER
)

// ShareList represents all  which could be generated
type ShareList struct {
	List []model.Share
}

//MarshalBinary method for ShareList
func (sl ShareList) MarshalBinary() ([]byte, error) {
	return json.Marshal(sl)
}

//UnmarshalBinary method for ShareList
func (sl ShareList) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &sl)
}

//NewShareList returns new ShareList instance
func NewShareList() *ShareList {
	return &ShareList{List: []model.Share{
		{
			Name:      AMAZON,
			Bid:       2874.16,
			Ask:       2878.31,
			UpdatedAt: time.Now().Format(time.RFC3339Nano),
		},
		{
			Name:      APPLE,
			Bid:       170.02,
			Ask:       171.71,
			UpdatedAt: time.Now().Format(time.RFC3339Nano),
		},
		{
			Name:      MICROSOFT,
			Bid:       307.90,
			Ask:       308.54,
			UpdatedAt: time.Now().Format(time.RFC3339Nano),
		},
		{
			Name:      NETFLIX,
			Bid:       382.95,
			Ask:       384.11,
			UpdatedAt: time.Now().Format(time.RFC3339Nano),
		},
		{
			Name:      PFIZER,
			Bid:       54.27,
			Ask:       55.16,
			UpdatedAt: time.Now().Format(time.RFC3339),
		},
	}}
}

//GeneratePrices updates share prices
func (sl *ShareList) GeneratePrices() {
	for i := range sl.List {
		rp := randPrice()
		if rp > 0 {
			sl.List[i].Bid += rp
			sl.List[i].Ask += 2 * rp
		} else {
			sl.List[i].Bid += rp
			sl.List[i].Ask += rp / 2
		}
		sl.List[i].UpdatedAt = time.Now().Format(time.RFC3339Nano)
	}
}

func randPrice() float32 {
	r := rand.New(rand.NewSource(randSeed))
	return r.Float32() - float32(1.0)
}
