package services

import (
	"github.com/qzich/todo/services/storage/mongo"
	"github.com/qzich/todo/services/storage"
)

var ListStorage = &Storage{
	New: mongo.New,
}

type Storage struct {
	New func(url string) *storage.MongoStorage
}
