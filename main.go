package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/visal/order-service/appis"

	"github.com/IBM/sarama"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:Visali@23@tcp(127.0.0.1:3306)/lib_b?parseTime=True" //db file info
	db, err := sql.Open("mysql", dsn)                                //prepares the drivers and validate if dsn parameters are correct
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	producer, err := initKafkaProducer()
	if err != nil {
		log.Fatalf("Error creating Kafka producer: %v", err)
	}
	defer producer.Close()

	appis.RegisterRoutes(db, producer)

	//Start the HTTP server
	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func initKafkaProducer() (sarama.SyncProducer, error) {
	brokerList := []string{"localhost:9092"}

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		return nil, err
	}

	return producer, nil
}

//command to check which ports are listening windows: netstat -aon
