package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/tedux/timing/pkg/app"
)

type ErrorResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type BaseResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func RegisterHandler(r gin.IRouter, app app.App) {
	v1 := r.Group("/v1")
	//groups
	v1.GET("/groups", ListActionGroupsHandler(app))
	v1.GET("/groups/:id", GetActionGroupHandler(app))
	v1.POST("/groups", AddActionGroupHandler(app))
	v1.DELETE("/groups/:id", DeleteActionGroupHandler(app))
	//actions
	v1.GET("/actions", ListActionsByGroupIdHandler(app))
	v1.GET("/actions/:id", GetActionHandler(app))
	v1.POST("/actions", AddActionHandler(app))
	v1.PATCH("/actions/:id", UpdateActionGroupIdHandler(app))
	v1.DELETE("/actions/:id", DeleteActionHandler(app))
	//timings
	v1.GET("/timings", ListTimingByActionAndDtHandler(app))
	v1.POST("/timings", StartTimingHandler(app))
	v1.PUT("/timings/:id", StopTimingHandler(app))
}
