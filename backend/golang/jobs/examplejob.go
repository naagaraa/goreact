package main

import (
    "fmt"
    "log"
    "os"
    "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
    if err != nil {
        log.Fatalf("%s: %s", msg, err)
    }
}

func main() {
   // Read RabbitMQ settings from environment variables
   rabbitMQHost := os.Getenv("RABBITMQ_HOST")
   rabbitMQPort := os.Getenv("RABBITMQ_PORT")
   rabbitMQUser := os.Getenv("RABBITMQ_USER")
   rabbitMQPassword := os.Getenv("RABBITMQ_PASSWORD")
   rabbitMQVHost := os.Getenv("RABBITMQ_VHOST")
   rabbitMQQueue := os.Getenv("RABBITMQ_QUEUE")

   if rabbitMQHost == "" || rabbitMQPort == "" || rabbitMQUser == "" || rabbitMQPassword == "" || rabbitMQVHost == "" || rabbitMQQueue == "" {
       log.Fatal("Missing required environment variables")
   }

   // Construct the connection URL dynamically from environment variables
   rabbitMQURL := fmt.Sprintf("amqp://%s:%s@%s:%s%s", rabbitMQUser, rabbitMQPassword, rabbitMQHost, rabbitMQPort, rabbitMQVHost)

   // Connect to RabbitMQ
   conn, err := amqp091.Dial(rabbitMQURL)
   failOnError(err, "Failed to connect to RabbitMQ")
   defer conn.Close()

   ch, err := conn.Channel()
   failOnError(err, "Failed to open a channel")
   defer ch.Close()

   // Declare the queue using the environment variable value
   q, err := ch.QueueDeclare(
       rabbitMQQueue, // Queue name from environment variable
       true,          // Durable
       false,         // Delete when unused
       false,         // Exclusive
       false,         // No-wait
       nil,           // Arguments
   )
   failOnError(err, "Failed to declare a queue")

    msgs, err := ch.Consume(
        q.Name, // Queue name
        "",     // Consumer tag
        true,   // Auto-ack
        false,  // Exclusive
        false,  // No-local
        false,  // No-wait
        nil,    // Arguments
    )
    failOnError(err, "Failed to register a consumer")

    forever := make(chan bool)

    go func() {
        for d := range msgs {
            var payload JobPayload
            if err := json.Unmarshal(d.Body, &payload); err != nil {
                log.Println("Error decoding JSON:", err)
                continue
            }

            // Process the job
            fmt.Printf("Processing job with data: %s\n", payload.Data)

            // Simulate heavy processing
            // Replace this with actual logic
            processJob(payload)

            fmt.Println("Job completed!")
        }
    }()

    <-forever
}

func processJob(payload JobPayload) {
    // Simulate a task that takes time
    fmt.Printf("Simulating heavy task with data: %s\n", payload.Data)
    // Add actual processing logic here
}