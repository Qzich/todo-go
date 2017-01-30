package handlers

import (
	"github.com/qiangxue/fasthttp-routing"
	"gopkg.in/mgo.v2/bson"
	"github.com/qzich/todo/models"
)

func AddTodoListItemHandler(context *routing.Context) error {
	var responseModel interface{}
	var requestModel models.AddTodoListItemRequest
	var errorModel models.ErrorResponse

	unmarshalJsonError := bson.UnmarshalJSON(context.PostBody(), &requestModel)

	if unmarshalJsonError != nil {
		errorModel.Errors = unmarshalJsonError.Error()
	} else {
		todoListIdHex := context.Param("listId")

		if bson.IsObjectIdHex(todoListIdHex) {
			domainModel, _ := listStorage.Get(bson.ObjectIdHex(todoListIdHex))

			todoItemIdHex := context.Param("itemId")

			var domainItemModel *models.TodoListItemModel

			if bson.IsObjectIdHex(todoItemIdHex) {

				for _, todoListItemModel := range domainModel.Items {
					if todoListItemModel.Id.Hex() == todoItemIdHex {
						domainItemModel = todoListItemModel
					}
				}

				if domainItemModel == nil {
					errorModel.Errors = "Todo item not found"
				} else {
					domainItemModel.Done = requestModel.Done
					domainItemModel.Title = requestModel.Title
				}

			} else {

				domainItemModel = &models.TodoListItemModel{
					Id:    bson.NewObjectId(),
					Title: requestModel.Title,
					Done:  requestModel.Done,
				}

				domainModel.Items = append(domainModel.Items, domainItemModel)
			}

			saveError := listStorage.Save(&domainModel)

			if saveError != nil {
				errorModel.Errors = saveError.Error()
			}

			responseModel = models.AddTodoListResponse{
				Id: domainItemModel.Id.Hex(),
			}
		}
	}

	if errorModel.Errors != nil {
		responseModel = errorModel
	}

	return sendResponse(responseModel, context)
}

func RemoveTodoListItemHandler(context *routing.Context) error {
	var responseModel interface{}

	var errorModel models.ErrorResponse

	todoListIdHex := context.Param("listId")

	if bson.IsObjectIdHex(todoListIdHex) {
		domainModel, _ := listStorage.Get(bson.ObjectIdHex(todoListIdHex))

		todoItemIdHex := context.Param("itemId")

		if bson.IsObjectIdHex(todoItemIdHex) {

			var domainItemModel *models.TodoListItemModel

			for todoListItemModelIndex, todoListItemModel := range domainModel.Items {
				if todoListItemModel.Id.Hex() == todoItemIdHex {
					domainItemModel = todoListItemModel

					if domainItemModel == nil {
						errorModel.Errors = "Todo item not found"
					} else {
						domainModelItems := domainModel.Items
						domainModel.Items = append(domainModelItems[:todoListItemModelIndex], domainModelItems[todoListItemModelIndex + 1:]...)
					}

				}
			}

		}

		saveError := listStorage.Save(&domainModel)

		if saveError != nil {
			errorModel.Errors = saveError.Error()
		} else {
			responseModel = models.SuccessResponse{Status:"ok"}
		}
	}

	if errorModel.Errors != nil {
		responseModel = errorModel
	}

	return sendResponse(responseModel, context)
}

func GetTodoListItemHandler(context *routing.Context) error {
	var responseModel interface{}

	var errorModel models.ErrorResponse

	todoListIdHex := context.Param("listId")

	if bson.IsObjectIdHex(todoListIdHex) {
		domainModel, getError := listStorage.Get(bson.ObjectIdHex(todoListIdHex))

		if getError != nil {
			errorModel.Errors = getError.Error()
		} else {

			todoItemIdHex := context.Param("itemId")

			if bson.IsObjectIdHex(todoItemIdHex) {

				var domainItemModel *models.TodoListItemModel

				for _, todoListItemModel := range domainModel.Items {
					if todoListItemModel.Id.Hex() == todoItemIdHex {
						domainItemModel = todoListItemModel
					}
				}

				responseModel = models.GetTodoListResponseItem{
					Id:    domainItemModel.Id.Hex(),
					Title: domainItemModel.Title,
					Done:  domainItemModel.Done,
				}

			}
		}
	}

	if errorModel.Errors != nil {
		responseModel = errorModel
	}

	return sendResponse(responseModel, context)
}

func GetTodoListItemsHandler(context *routing.Context) error {
	var responseModel interface{}

	var errorModel models.ErrorResponse

	todoListIdHex := context.Param("listId")

	if bson.IsObjectIdHex(todoListIdHex) {
		domainModel, getError := listStorage.Get(bson.ObjectIdHex(todoListIdHex))

		if getError != nil {
			errorModel.Errors = getError.Error()
		} else {

			var result []models.GetTodoListResponseItem

			for _, todoListItemModel := range domainModel.Items {
				result = append(result, models.GetTodoListResponseItem{
					Id:    todoListItemModel.Id.Hex(),
					Title: todoListItemModel.Title,
					Done:  todoListItemModel.Done,
				})
			}

			responseModel = result

		}
	}

	if errorModel.Errors != nil {
		responseModel = errorModel
	}

	return sendResponse(responseModel, context)
}
