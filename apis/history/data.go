package history

import "github.com/scorum/openledger-go/types"

type Bucket struct {
	ID          string    `json:"id"`
	Key         BucketKey `json:"key"`
	HighBase    uint32    `json:"high_base"`
	HighQuote   uint32    `json:"high_quote"`
	LowBase     uint32    `json:"low_base"`
	LowQuote    uint32    `json:"low_quote"`
	OpenBase    uint32    `json:"open_base"`
	OpenQuote   uint32    `json:"open_quote"`
	CloseBase   uint32    `json:"close_base"`
	CloseQuote  uint32    `json:"close_quote"`
	BaseVolume  uint32    `json:"base_volume"`
	QuoteVolume uint32    `json:"quote_volume"`
}

type BucketKey struct {
	Base    string     `json:"base"`
	Quote   string     `json:"quote"`
	Seconds uint32     `json:"seconds"`
	Open    types.Time `json:"open"`
}

type OrderHistory struct {
	ID        string     `json:"id"`
	Key       OrderKey   `json:"key"`
	Time      types.Time `json:"time"`
	Operation Operation  `json:"op"`
}

type OrderKey struct {
	Base     types.ObjectID `json:"base"`
	Quote    types.ObjectID `json:"quote"`
	Sequence int32          `json:"sequence"`
}

type Operation struct {
	Fee       types.AssetAmount `json:"fee"`
	Pays      types.AssetAmount `json:"pays"`
	Receives  types.AssetAmount `json:"receives"`
	FillPrice types.Price       `json:"fill_price"`
	OrderID   types.ObjectID    `json:"order_id"`
	AccountID types.ObjectID    `json:"account_id"`
	IsMaker   bool              `json:"is_maker"`
}
