package publisher

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
)

type OrderReceived struct{
	OrderId string 	`json:"orderId"`
	CustomerID string 	`json:"customerId"`
	Amount float64	`json :"amount`
	Timestamp string `json:"timestamp`

}
type Publisher struct {
	writer *kafka.Writer
}

func NewPublisher(brokers []string) *Publisher{
	w:=&kafka.Writer{
		Addr:kafka.TCP(brokers...),
		Topic:"OrderReceived",
		Balancer:&kafka.LeastBytes{},
	}
	return &Publisher{writer:w}
}

func (p *Publisher) PublishOrder(ctx context.Context, order OrderReceived) error {
    value, err := json.Marshal(order)
    if err != nil {
        return err
    }

    return p.writer.WriteMessages(ctx, kafka.Message{
        Value: value,
    })
}

func (p *Publisher) Close() error {
    return p.writer.Close()
}