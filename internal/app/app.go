package app

import (
	"encoding/json"
	"fmt"
	natslib "github.com/nats-io/nats.go"
	"log"
	"transaction-service/internal/delivery/dto"
	"transaction-service/internal/service"
	"transaction-service/pkg/nats"
)

type App struct {
	*nats.Streaming
	*service.Service
	subs []*natslib.Subscription
}

func NewApp(streaming nats.Streaming, services *service.Service) *App {
	return &App{&streaming, services, nil}
}

func (a *App) Start() error {
	sb, err := a.NC.QueueSubscribe("transfer", "transactions", a.handler)
	if err != nil {
		return err
	}

	a.subs = append(a.subs, sb)

	return nil
}

func (a *App) Stop() error {
	for _, sub := range a.subs {
		if err := sub.Unsubscribe(); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) handler(msg *natslib.Msg) {
	var transferDto dto.CreateTransferDto
	if err := json.Unmarshal(msg.Data, &transferDto); err != nil {
		log.Printf("[Marshall]: %v\n", err)
		a.publish(`{ "success": false, "message": "invalid json" }`)
		return
	}

	if err := a.Service.Transfer(transferDto.PubId, transferDto.SubId, transferDto.Amount); err != nil {
		resp := fmt.Sprintf(`{ "success": false, "message": %v }`, err)
		a.publish(resp)
		return
	}

	a.publish(`{ "success": true }`)
}

func (a *App) publish(msg string) {
	if err := a.NC.Publish("tr_response", []byte(msg)); err != nil {
		fmt.Printf("[PUBLISH ERR: %v", err)
		return
	}
}

func (a *App) publishJS(msg string) {
	if _, err := a.JS.Publish("test_res", []byte(msg)); err != nil {
		fmt.Printf("[PUBLISH ERR: %v", err)
		return
	}
}
