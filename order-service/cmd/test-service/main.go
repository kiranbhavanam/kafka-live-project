package main

import (
    "context"
    "log"
    "time"
    "order-service/publisher"
)

func main() {
    pub := publisher.NewPublisher([]string{"localhost:9092"})
    defer pub.Close()

    order := publisher.OrderReceived{
        OrderId:    "order123",
        CustomerID: "cust456",
        Amount:     99.99,
        Timestamp:  time.Now().Format(time.RFC3339),
    }

    err := pub.PublishOrder(context.Background(), order)
    if err != nil {
        log.Fatalf("Failed to publish: %v", err)
    }

    log.Println("Successfully published order event")
}