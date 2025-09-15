package appis

import (
	"database/sql"
	"fmt"

	"github.com/visal/order-service/dataservice"
	"github.com/visal/order-service/model"
	"github.com/visal/order-service/queue"

	"github.com/IBM/sarama"
)

type IBizLogic interface {
	CreateOrderLogic(order model.Order) error
}

type BizLogic struct {
	DB       *sql.DB
	Producer sarama.SyncProducer
}

func NewBizLogic(db *sql.DB, producer sarama.SyncProducer) *BizLogic {
	return &BizLogic{DB: db, Producer: producer}
}

func (bl *BizLogic) CreateOrderLogic(order model.Order) error {
	// validation by making a get request
	//

	if err := dataservice.CreateOrder(bl.DB, order); err != nil {
		return err
	}

	// produce the message to kafka
	message := fmt.Sprintf("I love myself that is %s", order.Customer_Name)
	err := queue.ProduceKafkaMessage("Order_created_topic", message, bl.Producer)
	if err != nil {
		return fmt.Errorf("failed to produce kafka message: %v", err)
	}

	return nil
}
