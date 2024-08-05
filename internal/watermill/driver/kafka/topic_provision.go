package kafka

import (
	"context"
	"log/slog"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type AutoProvisionTopic struct {
	Topic         string
	NumPartitions int32
}

// provisionTopics creates the topics if they don't exist. This relies on the confluent kafka lib, as the sarama doesn't seem to
// properly support interacting with the confluent cloud.
func provisionTopics(ctx context.Context, logger *slog.Logger, config kafka.ConfigMap, topics []AutoProvisionTopic) error {
	// This is not supported on admin client, so we need to remove it
	delete(config, "go.logs.channel.enable")

	adminClient, err := kafka.NewAdminClient(&config)
	if err != nil {
		return err
	}

	defer adminClient.Close()

	for _, topic := range topics {
		result, err := adminClient.CreateTopics(ctx, []kafka.TopicSpecification{
			{
				Topic:         topic.Topic,
				NumPartitions: int(topic.NumPartitions),
			},
		})
		if err != nil {
			return err
		}

		for _, r := range result {
			code := r.Error.Code()

			if code == kafka.ErrTopicAlreadyExists {
				logger.Debug("topic already exists", slog.String("topic", topic.Topic))
			} else if code != kafka.ErrNoError {
				return r.Error
			}
		}
	}

	return nil
}
