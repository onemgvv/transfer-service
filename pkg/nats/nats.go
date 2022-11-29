package nats

import (
	nts "github.com/nats-io/nats.go"
	"log"
)

type Streaming struct {
	NC *nts.Conn
	JS nts.JetStreamContext
}

const (
	streamName     = "TRANSACTIONS"
	streamSubjects = "TRANSACTIONS.*"
)

func NewStreaming() Streaming {
	nc, _ := nts.Connect(nts.DefaultURL)
	js, _ := nc.JetStream()

	return Streaming{NC: nc, JS: js}
}

// createStream creates a stream by using JetStreamContext
func createStream(js nts.JetStreamContext) error {
	// Check if the ORDERS stream already exists; if not, create it.
	stream, err := js.StreamInfo(streamName)
	if err != nil {
		log.Println(err)
	}
	if stream == nil {
		log.Printf("creating stream %q and subjects %q", streamName, streamSubjects)
		_, err = js.AddStream(&nts.StreamConfig{
			Name:     streamName,
			Subjects: []string{streamSubjects},
		})
		if err != nil {
			return err
		}
	}
	return nil
}
