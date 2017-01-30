package handlers

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/asaskevich/govalidator"
	"github.com/qzich/todo/models"
	"github.com/qiangxue/fasthttp-routing"
)

func AddTodoListHandler(context *routing.Context) error {

	var requestModel models.AddTodoListRequest
	var responseModel interface{}
	var domainModel models.TodoListModel
	var errorModel models.ErrorResponse

	todoIdHex := context.Param("id")

	unmarshalJsonError := bson.UnmarshalJSON(context.PostBody(), &requestModel)

	if unmarshalJsonError != nil {
		errorModel.Errors = unmarshalJsonError.Error()
	} else {

		isAddTodoListRequestValid, validationError := govalidator.ValidateStruct(requestModel)

		if validationError != nil {
			errorModel.Errors = govalidator.ErrorsByField(validationError)
		}

		if isAddTodoListRequestValid {

			domainModel.Name, domainModel.Color = requestModel.Name, requestModel.Color

			if bson.IsObjectIdHex(todoIdHex) {
				domainModel.Id = bson.ObjectIdHex(todoIdHex)
			}

			saveError := listStorage.Save(&domainModel)

			if saveError != nil {
				errorModel.Errors = saveError.Error()
			}

			responseModel = models.AddTodoListResponse{
				Id: domainModel.Id.Hex(),
			}
		}
	}

	if errorModel.Errors != nil {
		responseModel = errorModel
	}

	return sendResponse(responseModel, context)
}

func GetTodoListHandler(context *routing.Context) error {
	var responseModel interface{}

	var getTodoListResponse models.GetTodoListResponse

	var errorModel models.ErrorResponse

	todoIdHex := context.Param("id")

	if todoIdHex == "" {
		errorModel.Errors = "Todo list id has not be empty"
	} else {

		todoObjectId := bson.ObjectIdHex(todoIdHex)

		todoListModel, getError := listStorage.Get(todoObjectId)

		if getError != nil {
			errorModel.Errors = getError.Error()
		} else {
			getTodoListResponse = models.GetTodoListResponse{
				Name:      todoListModel.Name,
				Color:     todoListModel.Color,
				CreatedAt: todoListModel.CreatedAt.Unix(),
				UpdatedAt: todoListModel.UpdatedAt.Unix(),
			}

			for _, todoListItemModel := range todoListModel.Items {
				getTodoListResponse.Items = append(getTodoListResponse.Items, models.GetTodoListResponseItem{
					Id:    todoListItemModel.Id.Hex(),
					Title: todoListItemModel.Title,
					Done:  todoListItemModel.Done,
				})
			}
		}
	}

	if errorModel.Errors != nil {
		responseModel = errorModel
	} else {
		responseModel = getTodoListResponse
	}

	return sendResponse(responseModel, context)
}

func GetTodoListsHandler(context *routing.Context) error {
	var responseModel interface{}

	var re []models.GetTodoListResponse

	result, _ := listStorage.GetList(3, "-created_at")

	for _, todoListModel := range result {
		re = append(re, models.GetTodoListResponse{
			Id:        todoListModel.Id.Hex(),
			Name:      todoListModel.Name,
			Color:     todoListModel.Color,
			CreatedAt: todoListModel.CreatedAt.Unix(),
			UpdatedAt: todoListModel.UpdatedAt.Unix(),
		})

	}

	responseModel = re

	return sendResponse(responseModel, context)
}

func RemoveTodoListHandler(context *routing.Context) error {
	var responseModel interface{}

	var errorModel models.ErrorResponse

	todoIdHex := context.Param("id")

	if todoIdHex == "" {
		errorModel.Errors = "Todo list id has not be empty"
	} else {
		todoObjectId := bson.ObjectIdHex(todoIdHex)

		removeTodoListModel := models.RemoveTodoListModel{Id:todoObjectId}

		deleteError := listStorage.DeleteList(removeTodoListModel)

		if deleteError != nil {
			errorModel.Errors = deleteError.Error()
		} else {
			responseModel = models.SuccessResponse{Status:"ok"}
		}
	}

	if errorModel.Errors != nil {
		responseModel = errorModel
	}

	return sendResponse(responseModel, context)
}
