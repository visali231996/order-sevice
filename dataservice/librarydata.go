package dataservice

import (
	"database/sql"

	"github.com/visal/order-service/model"
)

func CreateOrder(db *sql.DB, order model.Order) error {
	query := "INSERT INTO tables(id, customer_name, paymentmethod,placedon,deliveredon,item,address,amount) VALUES (?,?,?,?,?,?,?,?)"
	_, err := db.Exec(query, order.Id, order.Customer_Name, order.PaymentMethod, order.PlacedOnDate, order.DeliveredOn, order.Item, order.Address, order.Amount)
	if err != nil {
		return err
	}
	return nil
}
