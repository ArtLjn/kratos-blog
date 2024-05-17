package mq

import (
	"github.com/streadway/amqp"
	"log"
)

type MQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func NewMQ(con *amqp.Connection, ch *amqp.Channel) *MQ {
	return &MQ{
		Conn:    con,
		Channel: ch,
	}
}

func ConnMq(url string) (*MQ, error) { // 修改返回值类型以包含错误处理
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		err = conn.Close()
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	return NewMQ(conn, channel), nil
}

// CreateExchange 创建交换机
func CreateExchange(mq *MQ, exchangeName string, exchangeType string) error {
	err := mq.Channel.ExchangeDeclare(
		exchangeName, // name
		exchangeType, // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		return err
	}

	log.Printf("Exchange %s declared", exchangeName)
	return nil
}

func CreateQueueAndBind(mq *MQ, queueName, routingKey string, exchangeName string) error {
	_, err := mq.Channel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return err
	}

	err = mq.Channel.QueueBind(
		queueName,    // queue name
		routingKey,   // routing key
		exchangeName, // exchange
		false,
		nil,
	)
	if err != nil {
		return err
	}

	log.Printf("Queue %s bound to exchange %s with routing key %s", queueName, exchangeName, routingKey)
	return nil
}

func PublicMsg(ch *amqp.Channel, exchange, queue string, body []byte) error {
	// 发送消息到RabbitMQ
	if err := ch.Publish(exchange, queue, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	}); err != nil {
		return err
	}
	return nil
}
