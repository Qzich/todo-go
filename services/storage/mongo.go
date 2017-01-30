package storage

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"errors"
	"github.com/qzich/todo/models"
	"time"
)

const DB_NAME string = "storage"
const COLLECTION_NAME string = "list"

type MongoStorage struct {
	Url string
}

func (this MongoStorage) Save(todoListModel *models.TodoListModel) (err error) {

	var command func(listCollection *mgo.Collection) error

	if todoListModel.Id == "" {
		todoListModel.Id = bson.NewObjectId()

		todoListModel.CreatedAt = time.Now()

		command = func(listCollection *mgo.Collection) error {
			return listCollection.Insert(todoListModel)
		}
	} else {
		todoListModel.UpdatedAt = time.Now()

		command = func(listCollection *mgo.Collection) error {
			return listCollection.UpdateId(todoListModel.Id, bson.M{"$set": &todoListModel})
		}
	}

	this.runCommand(command)

	return err
}

func (this MongoStorage) Get(objectId bson.ObjectId) (models.TodoListModel, error) {

	var err error
	todoListModel := &models.TodoListModel{}

	selectCommand := func(listCollection *mgo.Collection) error {
		return listCollection.FindId(objectId).One(todoListModel)
	}

	err = this.runCommand(selectCommand)

	return *todoListModel, err
}

func (this MongoStorage) GetList(limit int, sortBy string) ([]models.TodoListModel, error) {
	var err error

	var result []models.TodoListModel

	selectCommand := func(listCollection *mgo.Collection) error {
		return listCollection.Find(bson.M{}).Limit(limit).Sort(sortBy).All(&result)

	}

	err = this.runCommand(selectCommand)

	return result, err
}

func (this MongoStorage) DeleteList(removeTodoListModel models.RemoveTodoListModel) error {

	if removeTodoListModel.Id == "" {
		return errors.New("Id is required field for remove todo list.")
	}

	removeCommand := func(listCollection *mgo.Collection) error {
		return listCollection.Remove(removeTodoListModel)
	}

	return this.runCommand(removeCommand)
}

func (this MongoStorage) runCommand(command func(collection *mgo.Collection) error) (err error) {

	defer func() {
		panicMessage := recover()

		if panicMessage != nil {
			err = errors.New("It seems mongo server is down")
		}
	}()

	session, err := mgo.Dial(this.Url)

	if err != nil {
		panic(err)
	}

	listCollection := session.DB(DB_NAME).C(COLLECTION_NAME)

	defer session.Close()

	return command(listCollection)
}
