package types

import (
	"encoding/json"
	"github.com/pkg/errors"
	"reflect"
)

type Operation interface {
	Type() OpType
}

type Operations []Operation

func (ops *Operations) UnmarshalJSON(b []byte) (err error) {
	// unmarshal array
	var o []json.RawMessage
	if err := json.Unmarshal(b, &o); err != nil {
		return err
	}

	// foreach operation
	for _, op := range o {
		var kv []json.RawMessage
		if err := json.Unmarshal(op, &kv); err != nil {
			return err
		}

		if len(kv) != 2 {
			return errors.New("invalid operation format: should be name, value")
		}

		var opType uint16
		if err := json.Unmarshal(kv[0], &opType); err != nil {
			return err
		}

		val, err := unmarshalOperation(OpType(opType), kv[1])
		if err != nil {
			return err
		}

		*ops = append(*ops, val)
	}

	return nil
}

func unmarshalOperation(opType OpType, obj json.RawMessage) (Operation, error) {
	op, ok := knownOperations[opType]
	if !ok {
		// operation is unknown wrap it as an unknown operation
		val := UnknownOperation{
			kind: opType,
			Data: obj,
		}
		return &val, nil
	} else {
		val := reflect.New(op).Interface()
		if err := json.Unmarshal(obj, val); err != nil {
			return nil, err
		}
		return val.(Operation), nil
	}
}

var knownOperations = map[OpType]reflect.Type{
	TransferOpType:         reflect.TypeOf(TransferOperation{}),
	LimitOrderCreateOpType: reflect.TypeOf(LimitOrderCreateOperation{}),
	LimitOrderCancelOpType: reflect.TypeOf(LimitOrderCancelOperation{}),
}

// UnknownOperation
type UnknownOperation struct {
	kind OpType
	Data json.RawMessage
}

func (op *UnknownOperation) Type() OpType { return op.kind }

// TransferOperation
type TransferOperation struct {
	From       ObjectID          `json:"from"`
	To         ObjectID          `json:"to"`
	Amount     AssetAmount       `json:"amount"`
	Fee        AssetAmount       `json:"fee"`
	Memo       Memo              `json:"memo"`
	Extensions []json.RawMessage `json:"extensions"`
}

type Memo struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Nonce   string `json:"nonce"`
	Message string `json:"message"`
}

func (op *TransferOperation) Type() OpType { return TransferOpType }

// LimitOrderCreateOperation
type LimitOrderCreateOperation struct {
	Fee          AssetAmount       `json:"fee"`
	Seller       ObjectID          `json:"seller"`
	AmountToSell AssetAmount       `json:"amount_to_sell"`
	MinToReceive AssetAmount       `json:"min_to_receive"`
	Expiration   Time              `json:"expiration"`
	FillOrKill   bool              `json:"fill_or_kill"`
	Extensions   []json.RawMessage `json:"extensions"`
}

func (op *LimitOrderCreateOperation) Type() OpType { return LimitOrderCreateOpType }

// LimitOrderCancelOpType
type LimitOrderCancelOperation struct {
	Fee              AssetAmount       `json:"fee"`
	FeePayingAccount ObjectID          `json:"fee_paying_account"`
	Order            ObjectID          `json:"order"`
	Extensions       []json.RawMessage `json:"extensions"`
}

func (op *LimitOrderCancelOperation) Type() OpType { return LimitOrderCancelOpType }