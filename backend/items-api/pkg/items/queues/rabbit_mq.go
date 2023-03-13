package queues

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/rabbitmq/amqp091-go"
	"items-api/internal/apierrors"
	"items-api/internal/logger"
	"items-api/pkg/items"
	"time"
)

const (
	encodingJSON = "application/json"
	encodingUTF8 = "UTF-8"
)

type itemsRabbitMQ struct {
	connection *amqp091.Connection
	channel    *amqp091.Channel
	queue      *amqp091.Queue
	logger     *logger.Logger
	appName    string
}

// NewItemsRabbitMQ instances a new items' queues against RabbitMQ
func NewItemsRabbitMQ(ctx context.Context, host string, port int, user string, password string, queueName string, appName string, logger *logger.Logger) (itemsRabbitMQ, error) {
	connection, err := amqp091.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", user, password, host, port))
	if err != nil {
		logger.Errorf(ctx, "Error dialing RabbitMQ conenction: %s", err.Error())
		return itemsRabbitMQ{}, err
	}
	channel, err := connection.Channel()
	if err != nil {
		logger.Errorf(ctx, "Error getting RabbitMQ connection channel: %s", err.Error())
		return itemsRabbitMQ{}, err
	}
	queue, err := channel.QueueDeclare(queueName, false, false, false, false, nil)
	if err != nil {
		logger.Errorf(ctx, "Error declaring RabbitMQ channel queue: %s", err.Error())
		return itemsRabbitMQ{}, err
	}
	return itemsRabbitMQ{
		connection: connection,
		channel:    channel,
		queue:      &queue,
		logger:     logger,
		appName:    appName,
	}, nil
}

// PublishItemNotification notifies an item new
func (publisher itemsRabbitMQ) PublishItemNotification(ctx context.Context, action items.Action, priority items.Priority, id int64) apierrors.APIError {
	message := map[string]interface{}{
		"action": action.String(),
		"id":     id,
	}

	bytes, err := json.Marshal(message)
	if err != nil {
		return apierrors.NewInternalServerError(fmt.Sprintf("error generating message for RabbitMQ: %s", err.Error()))
	}

	if err := publisher.channel.PublishWithContext(ctx, "", publisher.queue.Name, false, false, amqp091.Publishing{
		ContentType:     encodingJSON,
		ContentEncoding: encodingUTF8,
		DeliveryMode:    amqp091.Transient,
		Priority:        priority.Value(),
		MessageId:       uuid.New().String(),
		Timestamp:       time.Now().UTC(),
		AppId:           publisher.appName,
		Body:            bytes,
	}); err != nil {
		return apierrors.NewInternalServerError(fmt.Sprintf("error publishing item with RabbitMQ: %s", err.Error()))
	}
	return nil
}
