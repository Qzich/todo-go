package validators

import (
	"github.com/qzich/todo/models"
)

func ValidateTodoListColor(field interface{}, context interface{}) bool {
	isValid := true

	todoListRequestColor := field.(models.TodoListRequestColor)

	var validColors = map[models.TodoListRequestColor]string{
		1: "green",
		2: "red",
		3: "blue",
	}

	if _, ok := validColors[todoListRequestColor]; !ok {
		isValid = false
	}

	return isValid
}
