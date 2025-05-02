KAFKA_HOME=~/Downloads/kafka_2.13-4.0.0


$KAFKA_HOME/bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --config retention.ms=10800000 --topic OrderReceived
$KAFKA_HOME/bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --config retention.ms=10800000 --topic OrderConfirmed
$KAFKA_HOME/bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --config retention.ms=10800000 --topic OrderPicked
$KAFKA_HOME/bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --config retention.ms=10800000 --topic Notification
$KAFKA_HOME/bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --config retention.ms=10800000 --topic DeadLetterQueue