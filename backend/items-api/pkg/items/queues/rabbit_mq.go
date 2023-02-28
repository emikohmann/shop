package queues

import (
	"context"
	"fmt"
	"github.com/emikohmann/shop/backend/items-api/internal/apierrors"
	"github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"time"
)

type itemsRabbitMQ struct {
	connection *amqp091.Connection
	channel    *amqp091.Channel
	queue      *amqp091.Queue
	logger     *logrus.Logger
}

// NewItemsRabbitMQ instances a new items' queues against RabbitMQ
func NewItemsRabbitMQ(host string, port int, user string, password string, queueName string, logger *logrus.Logger) (itemsRabbitMQ, error) {
	connection, err := amqp091.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", user, password, host, port))
	if err != nil {
		logger.Errorf("Error dialing RabbigMQ conenction: %s", err.Error())
		return itemsRabbitMQ{}, err
	}
	channel, err := connection.Channel()
	if err != nil {
		logger.Errorf("Error getting RabbitMQ connection channel: %s", err.Error())
		return itemsRabbitMQ{}, err
	}
	queue, err := channel.QueueDeclare(queueName, false, false, false, false, nil)
	if err != nil {
		logger.Errorf("Error declaring RabbitMQ channel queue: %s", err.Error())
		return itemsRabbitMQ{}, err
	}
	return itemsRabbitMQ{
		connection: connection,
		channel:    channel,
		queue:      &queue,
		logger:     logger,
	}, nil
}

// SendItem notifies an item new
func (publisher itemsRabbitMQ) SendItem(ctx context.Context, id int64) apierrors.APIError {
	if err := publisher.channel.PublishWithContext(ctx, "", publisher.queue.Name, false, false, amqp091.Publishing{
		Headers:         nil,
		ContentType:     "",
		ContentEncoding: "",
		DeliveryMode:    0,
		Priority:        0,
		CorrelationId:   "",
		ReplyTo:         "",
		Expiration:      "",
		MessageId:       "",
		Timestamp:       time.Now().UTC(),
		Type:            "",
		UserId:          "",
		AppId:           "",
		Body:            nil,
	}); err != nil {
		return apierrors.NewInternalServerError(fmt.Sprintf("error publishing item with RabbitMQ: %s", err.Error()))
	}
	publisher.logger.Infof("Successfully new published for item: %d", id)
	return nil
}
