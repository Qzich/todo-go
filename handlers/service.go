package handlers

import (
	"github.com/qiangxue/fasthttp-routing"
	"github.com/qzich/todo/models"
)

func ServiceHeartbeatHandler(context *routing.Context) error {
	return sendResponse(models.SuccessResponse{Status: "ok"}, context)
}
