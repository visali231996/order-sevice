package appis

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/visal/order-service/model"

	"github.com/IBM/sarama"
)

type Handler struct {
	biz IBizLogic
}

func NewHandler(db *sql.DB, producer sarama.SyncProducer) Handler {
	return Handler{biz: NewBizLogic(db, producer)}
}

func (h Handler) CreateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var order model.Order
		if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := h.biz.CreateOrderLogic(order); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusOK)
	}
}
