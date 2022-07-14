# Golang RabbitMQ Consumer

## Install

``` go
go get github.com/fothoncode/golang_rabbitmq
```

## Example code

``` go
import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	mqConsumer "github.com/fothoncode/golang_rabbitmq"
)

func ConsumeCallback(d amqp.Delivery) {
	fmt.Println("msgs recieved!")

    /* Handler code here */
}

func Main() {
    RabbitMQString := "xxx"
    QueueKey := "queueName"

	consumeConfig := mqConsumer.Config{
		RabbitMQString: RabbitMQString,
		QueueKey:       QueueKey,
	}
	mqConsumer.Consume(consumeConfig, ConsumeCallback)
}

```