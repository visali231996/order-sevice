package model

type Order struct {
	Id            int    `json:"id"`
	Customer_Name string `json:"customer_name"`
	PaymentMethod string `json:"paymentmethod"`
	PlacedOnDate  string `json:"placedondate"`
	DeliveredOn   string `json:"deliveredon"`
	Item          string `json:"item"`
	Address       string `json:"address"`
	Amount        int    `json:"amount"`
}
