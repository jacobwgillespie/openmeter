package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type NotificationServiceConfiguration struct {
	Consumer ConsumerConfiguration
}

func (c NotificationServiceConfiguration) Validate() error {
	if err := c.Consumer.Validate(); err != nil {
		return fmt.Errorf("consumer: %w", err)
	}
	return nil
}

func ConfigureNotificationService(v *viper.Viper) {
	ConfigureConsumer(v, "notificationService.consumer")
	v.SetDefault("notificationService.consumer.dlq.topic", "om_sys.notification_service_dlq")
	v.SetDefault("notificationService.consumer.consumerGroupName", "om_notification_service")
}
