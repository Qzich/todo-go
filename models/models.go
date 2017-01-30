package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type TodoListRequestColor uint8

type AddTodoListRequest struct {
	Name  string `json:"name" valid:"required,alphanum,length(5|100)"`
	Color TodoListRequestColor `json:"color" valid:"TodoValidColors~Valid colors are 1|2|3"`
}

type UpdateTodoListRequest struct {
	Name  string `json:"name" valid:"alphanum,length(5|100),optional"`
	Color TodoListRequestColor `json:"color" valid:"TodoValidColors~Valid colors are 1|2|3"`
}

type ErrorResponse struct {
	Errors interface{} `json:"errors"`
}

type SuccessResponse struct {
	Status interface{} `json:"status"`
}

type AddTodoListResponse struct {
	Id string `json:"id"`
}

type RemoveTodoListModel struct {
	Id bson.ObjectId `bson:"_id"`
}

type GetTodoListResponseItem struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Done  bool `json:"done"`
}

type GetTodoListResponse struct {
	Id        string `json:"id,omitempty"`
	Name      string `json:"name"`
	Color     TodoListRequestColor `json:"color"`
	Items     []GetTodoListResponseItem `json:"items,omitempty"`
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

type TodoListModel struct {
	Id        bson.ObjectId `bson:"_id"`
	Name      string `bson:"name,omitempty"`
	Color     TodoListRequestColor `bson:"color,omitempty"`
	Items     []*TodoListItemModel `bson:"items"`
	CreatedAt time.Time `bson:"created_at,omitempty"`
	UpdatedAt time.Time `bson:"updated_at,omitempty"`
}

type AddTodoListItemRequest struct {
	Title string `json:"title"`
	Done  bool `json:"done"`
}

type TodoListItemModel struct {
	Id    bson.ObjectId `bson:"_id"`
	Title string `bson:"title,omitempty"`
	Done  bool `bson:"done,omitempty"`
}
