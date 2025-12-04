package queue

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
	"github.com/zeromicro/go-zero/core/logx"
)

// RabbitMQConfig holds RabbitMQ configuration
type RabbitMQConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	VHost    string
}

// RabbitMQClient wraps RabbitMQ connection and channel
type RabbitMQClient struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

// NewRabbitMQClient creates a new RabbitMQ client
func NewRabbitMQClient(config RabbitMQConfig) (*RabbitMQClient, error) {
	dsn := fmt.Sprintf("amqp://%s:%s@%s:%d/%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.VHost,
	)

	conn, err := amqp.Dial(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to rabbitmq: %w", err)
	}

	channel, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to open channel: %w", err)
	}

	logx.Infof("Successfully connected to RabbitMQ at %s:%d", config.Host, config.Port)

	return &RabbitMQClient{
		conn:    conn,
		channel: channel,
	}, nil
}

// DeclareQueue declares a queue
func (r *RabbitMQClient) DeclareQueue(name string, durable, autoDelete, exclusive, noWait bool) error {
	_, err := r.channel.QueueDeclare(
		name,       // name
		durable,    // durable
		autoDelete, // auto-delete
		exclusive,  // exclusive
		noWait,     // no-wait
		nil,        // arguments
	)
	return err
}

// Publish publishes a message to a queue
func (r *RabbitMQClient) Publish(queue string, message interface{}) error {
	body, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	return r.channel.Publish(
		"",    // exchange
		queue, // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}

// Consume consumes messages from a queue
func (r *RabbitMQClient) Consume(queue string, consumer string, autoAck, exclusive, noLocal, noWait bool) (<-chan amqp.Delivery, error) {
	return r.channel.Consume(
		queue,     // queue
		consumer,  // consumer
		autoAck,   // auto-ack
		exclusive, // exclusive
		noLocal,   // no-local
		noWait,    // no-wait
		nil,       // args
	)
}

// DeclareExchange declares an exchange
func (r *RabbitMQClient) DeclareExchange(name, kind string, durable, autoDelete, internal, noWait bool) error {
	return r.channel.ExchangeDeclare(
		name,       // name
		kind,       // kind
		durable,    // durable
		autoDelete, // auto-delete
		internal,   // internal
		noWait,     // no-wait
		nil,        // arguments
	)
}

// PublishToExchange publishes a message to an exchange
func (r *RabbitMQClient) PublishToExchange(exchange, routingKey string, message interface{}) error {
	body, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	return r.channel.Publish(
		exchange,   // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}

// Close closes the RabbitMQ connection
func (r *RabbitMQClient) Close() error {
	if r.channel != nil {
		r.channel.Close()
	}
	if r.conn != nil {
		return r.conn.Close()
	}
	return nil
}

// GetChannel returns the underlying channel
func (r *RabbitMQClient) GetChannel() *amqp.Channel {
	return r.channel
}

