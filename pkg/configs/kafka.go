package configs

type KafkaConfig struct {
	BrokersHost       string
	ConsumerGroupID   string
	NotificationTopic string
}

type ConsumerConfig struct {
	Topic           string
	ConsumerGroupID string
}

func initializeKafkaConfig() {
	if KafkaCfg == nil {
		KafkaCfg = &KafkaConfig{
			BrokersHost:       getEnv("KAFKA_BROKER_HOST", ""),
			ConsumerGroupID:   ApplicationCfg.AppName,
			NotificationTopic: getEnv("KAFKA_NOTIFICATION_TOPIC", ""),
		}
	}
}

func NotificationConsumerConfig() *ConsumerConfig {
	return &ConsumerConfig{
		Topic:           getEnv("KAFKA_NOTIFICATION_TOPIC", ""),
		ConsumerGroupID: getEnv("CONSUMER_GROUP_ID", ""),
	}
}
