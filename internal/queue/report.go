package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/orensimple/otus_go_project/internal/domain/models"
	"github.com/orensimple/otus_go_project/internal/logger"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

// implements domain.interfaces.ReportQueue
type ReportQueue struct {
	ch *amqp.Channel
}

func NewReportQueue() (*ReportQueue, error) {

	dsn := fmt.Sprintf("amqp://%s:%s@%s:%s/", viper.GetString("amqp.user"), viper.GetString("amqp.passwd"), viper.GetString("amqp.ip"), viper.GetString("amqp.port"))

	connAMQP, err := amqp.Dial(dsn)
	if err != nil {
		logger.ContextLogger.Errorf("Failed to connect to RabbitMQ, retry after 30 second", err.Error())
		timer1 := time.NewTimer(30 * time.Second)
		<-timer1.C
		connAMQP, err = amqp.Dial(dsn)
		if err != nil {
			logger.ContextLogger.Errorf("Failed to retry connect to RabbitMQ", err.Error())
		}
	}

	ch, err := connAMQP.Channel()
	if err != nil {
		logger.ContextLogger.Errorf("Failed to open a channel", err.Error())
	}

	err = ch.ExchangeDeclare(
		"rotationBanners", // name
		"direct",          // type
		true,              // durable
		false,             // auto-deleted
		false,             // internal
		false,             // no-wait
		nil,               // arguments
	)
	if err != nil {
		logger.ContextLogger.Infof("Failed to declare an exchange", err.Error())
	}

	_, err = ch.QueueDeclare(
		"currentStat", // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		logger.ContextLogger.Infof("Problem QueueDeclare", err.Error())
	}

	err = ch.QueueBind(
		"currentStat",     // name
		"showAndClick",    // key
		"rotationBanners", // exchange
		false,             // no-wait
		nil,               // arguments
	)
	if err != nil {
		logger.ContextLogger.Infof("Problem bind queue", err.Error())
	}

	return &ReportQueue{ch: ch}, nil
}

func (rq *ReportQueue) PublicReport(ctx context.Context, report models.ReportQueueFormat) error {

	body, err := json.Marshal(report)
	if err != nil {
		logger.ContextLogger.Errorf("Error encoding JSON", err.Error())
	}

	err = rq.ch.Publish(
		"rotationBanners", // exchange
		"showAndClick",    // routing key
		false,             // mandatory
		false,             // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})

	if err != nil {
		logger.ContextLogger.Errorf("Failed to publish a message", err.Error())
	}
	logger.ContextLogger.Infof("Publish message type: ", report.Type)

	return nil

}
