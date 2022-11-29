package main

import (
	"fmt"
	"github.com/joho/godotenv"
	natslib "github.com/nats-io/nats.go"
	"log"
	"time"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("[ENV LOAD ERROR]: %s\n", err.Error())
	}

	nc, _ := natslib.Connect(natslib.DefaultURL)

	log.Println("Connected to " + natslib.DefaultURL)

	data := `{"pub_id": 4, "sub_id": 3, "amount": 10000}`

	if _, err := nc.QueueSubscribe("tr_response", "transactions", handler); err != nil {
		log.Fatalf("[SUB ERR]: %v", err)
	}

	if err := nc.Publish("transfer", []byte(data)); err != nil {
		log.Fatalf("Pub error %s", err.Error())
	}

	time.Sleep(15 * time.Second)
}

func handler(msg *natslib.Msg) {
	fmt.Println(string(msg.Data))
}
