package order

import "github.com/2u35-Wooje/gorder-v2/common/genproto/orderpb"

type Order struct {
	ID          string
	CustomerID  string
	Status      string
	PaymentLink string
	Items       []*orderpb.Item
}
