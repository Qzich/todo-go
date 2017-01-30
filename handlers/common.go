package handlers

import (
	"github.com/qzich/todo/helpers"
	"fmt"
	"github.com/qzich/todo/models"
	"github.com/qiangxue/fasthttp-routing"
	"encoding/json"
	"github.com/qzich/todo/services"
)

var listStorage = services.ListStorage.New("192.168.33.10")

const JSON_CONTENT_TYPE = "application/json"

func sendResponse(responseModel interface{}, context *routing.Context) error {

	context.SetContentType(JSON_CONTENT_TYPE)

	responseBytes := helpers.Must(json.Marshal(responseModel)).([]byte)

	switch responseModel.(type) {
	default:
		context.SetStatusCode(200)
		fmt.Fprint(context, string(responseBytes))
	case models.ErrorResponse:
		context.SetStatusCode(400)
		fmt.Fprint(context, string(responseBytes))
	}

	return nil
}
