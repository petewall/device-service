package deviceservice

//go:generate counterfeiter -generate

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

//counterfeiter:generate . DBInterface
type DBInterface interface {
	GetDevices() (error, []*Device)
}

type DB struct {
	client *redis.Client
}

func Connect(config *Config) *DB {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.DBHost, config.DBPort),
		Password: "",
		DB:       0,
	})

	db := &DB{
		client: client,
	}

	return db
}

func (db *DB) GetDevices() (error, []*Device) {
	// db.client.HKeys()
	// db.GetDevice()
	return nil, []*Device{}
}

func (db *DB) GetDevice(mac string) (error, *Device) {
	return nil, nil
}

// func (db *DB) UpdateDevice()
