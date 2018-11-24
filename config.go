package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Address  string         `toml:"address"`
	Database databaseConfig `toml:"database"`
	Redis    redisConfig    `toml:"redis"`
	Session  sessionConfig  `toml:"session"`
}

type databaseConfig struct {
	Host     string `toml:"host"`
	Name     string `toml:"name"`
	User     string `toml:"user"`
	Password string `toml:"password"`
}

type redisConfig struct {
	Address  string `toml:"address"`
	Password string `toml:"password"`
	DbNumber int    `toml:"db_number"`
}

type sessionConfig struct {
	SessionTimeout uint `toml:"session_timeout"`
}

func ReadConfig(configFile string) (*Config, error) {
	var config Config
	_, err := toml.DecodeFile(configFile, &config)
	if err != nil {
		logrus.Fatalf("bad config file: %v\n", configFile)
		return nil, err
	}

	return &config, nil
}

func (c Config) DataSourceName() string {
	if c.Database.Password != "" {
		return fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=5432 sslmode=disable",
			c.Database.User, c.Database.Name, c.Database.Password, c.Database.Host)
	} else {
		return fmt.Sprintf("user=%s dbname=%s host=%s port=5432 sslmode=disable",
			c.Database.User, c.Database.Name, c.Database.Host)
	}
}

func (c Config) RedisOptions() *redis.Options {
	return &redis.Options{
		Addr:     c.Redis.Address,
		Password: c.Redis.Password,
		DB:       c.Redis.DbNumber,
	}
}
