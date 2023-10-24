package main

import (
	"log"
	"github.com/streadway/amqp"

	"github.com/gin-gonic/gin"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func ping(c *gin.Context) {
    c.JSON(200, gin.H{
       "message": "pong",
    })

    conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
    	failOnError(err, "Failed to connect to RabbitMQ")
    	defer conn.Close()

    	ch, err := conn.Channel()
    	failOnError(err, "Failed to open a channel")
    	defer ch.Close()

    	q, err := ch.QueueDeclare(
    		"hello", // Queue name
    		false,   // Durable
    		false,   // Delete when unused
    		false,   // Exclusive
    		false,   // No-wait
    		nil,     // Arguments
    	)
    	failOnError(err, "Failed to declare a queue")

    	body := "Hello from Go!!!!!!!"
    	err = ch.Publish(
    		"",     // Exchange
    		q.Name, // Routing key
    		false,  // Mandatory
    		false,  // Immediate
    		amqp.Publishing{
    			ContentType: "text/plain",
    			Body:        []byte(body),
    		})
    	failOnError(err, "Failed to publish a message")

    	log.Printf(" [x] Sent %s", body)
}

func main() {

    r := gin.Default()
    r.GET("/ping", ping)

    r.Run()
}