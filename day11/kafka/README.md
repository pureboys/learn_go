### 启动docker 并且创建topic 

```
docker-composer up -d

cd /opt/bitnami/kafka/bin/

./kafka-topics.sh --zookeeper zookeeper:2181 --create --topic nginx_topic --partitions 3 --replication-factor 1
```