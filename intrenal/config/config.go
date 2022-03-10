//Package config represents config info of price generator
package config

// Config type store all env info of price generator
type Config struct {
	RedisURL        string `env:"REDISURL" envDefault:"localhost:6379"`
	RedisStreamName string `env:"RSTREAMNAME" envDefault:"PriceGenerator"`
	GrpcPort        string `env:"GRPCPORT" envDefault:"localhost:8083"`
}
