package mongo

import "github.com/qzich/todo/services/storage"

func New(url string) *storage.MongoStorage {

	return &storage.MongoStorage{
		Url: url,
	}
}
