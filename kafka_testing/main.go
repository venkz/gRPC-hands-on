package main

import (
	"gRPC-hands-on/kafka_testing/strmessages"
	broker "gRPC-hands-on/utils"
)

func main() {
	strmessages.Producer(broker.KafkaWriter)
	strmessages.Consumer(broker.KafkaReader)
}