package producer

import (
	"fmt"
	"github.com/EgorBessonov/price-generator/intrenal/generator"
	"github.com/go-redis/redis"
	"strconv"
)

//Producer type represent redis stream producer
type Producer struct {
	RedisClient *redis.Client
	StreamName  string
}

//NewRedis method return new producer instance
func NewRedis(redisCli *redis.Client, streamName string) *Producer {
	return &Producer{RedisClient: redisCli, StreamName: streamName}
}

//SendPricesToStream send prices to redis stream
func (rCli *Producer) SendPricesToStream(prices *generator.ShareList) error {
	shareMap := make(map[string]interface{})
	for i, sh := range prices.List {
		shareMap[strconv.Itoa(i)] = sh
	}
	result := rCli.RedisClient.XAdd(&redis.XAddArgs{
		Stream: rCli.StreamName,
		Values: shareMap,
	})
	if _, err := result.Result(); err != nil {
		return fmt.Errorf("producer: can't send message to stream - %e", err)
	}
	return nil
}
