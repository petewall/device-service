package lib

//go:generate counterfeiter -generate

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-redis/redis/v8"
)

//counterfeiter:generate . DBInterface
type DBInterface interface {
	GetDevices() ([]*Device, error)
	GetDevice(mac string) (*Device, error)
}

type DBConfig struct {
	Host string
	Port int
}

type DB struct {
	client *redis.Client
	ctx    context.Context
}

func Connect(config *DBConfig) *DB {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: "",
		DB:       0,
	})

	db := &DB{
		client: client,
		ctx:    context.Background(),
	}

	return db
}

func macToKey(mac string) string {
	return "devices:" + strings.ReplaceAll(mac, ":", "")
}

func (db *DB) GetDevices() ([]*Device, error) {
	var devices []*Device
	iter := db.client.Scan(db.ctx, 0, "devices:*", 0).Iterator()
	for iter.Next(db.ctx) {
		device, err := db.getDevice(iter.Val())
		if err != nil {
			return nil, err
		}
		devices = append(devices, device)
	}
	if iter.Err() != nil {
		return nil, iter.Err()
	}

	return devices, nil
}

func (db *DB) GetDevice(mac string) (*Device, error) {
	return db.getDevice(macToKey(mac))
}

func (db *DB) getDevice(key string) (*Device, error) {
	res := db.client.HGetAll(db.ctx, key)
	if res.Err() != nil {
		return nil, res.Err()
	}

	device := &Device{}
	err := res.Scan(device)
	if err != nil {
		return nil, err
	}

	return device, nil
}
