package main

import (
	"context"
	"fmt"
	"github.com/EgorBessonov/price-generator/intrenal/config"
	"github.com/EgorBessonov/price-generator/intrenal/generator"
	"github.com/EgorBessonov/price-generator/intrenal/producer"
	"github.com/EgorBessonov/price-generator/intrenal/service"
	"github.com/caarlos0/env"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("can't parse configs")
	}
	shList := generator.NewShareList()
	redisClient, err := newRedisClient(&cfg)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("can't create redis client instance")
	}
	rCli := producer.NewRedis(redisClient, cfg.RedisStreamName)
	defer func() {
		if err := redisClient.Close(); err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err,
			}).Error("error while closing redis client")
		}
	}()
	s := service.NewService(rCli, shList)
	log.Println("price generator started...")
	exitContext, cancelFunc := context.WithCancel(context.Background())
	exitChan := make(chan os.Signal, 1)
	signal.Notify(exitChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM)
	s.StartPriceGenerator(exitContext)
	<-exitChan
	cancelFunc()
}

func newRedisClient(cfg *config.Config) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisURL,
		Password: "",
		DB:       0,
	})
	if _, err := redisClient.Ping().Result(); err != nil {
		return nil, fmt.Errorf("producer: can't create new instance - %e", err)
	}
	return redisClient, nil
}
