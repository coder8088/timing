package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/tedux/timing/pkg/app"
	"github.com/tedux/timing/pkg/model"
)

func AddActionHandler(app app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var action model.Action
		if err := c.ShouldBindJSON(&action); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}
		id, err := app.AddAction(&action)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Code:  http.StatusInternalServerError,
				Error: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, BaseResponse{
			Code: 0,
			Data: gin.H{"id": id},
		})
	}
}

func GetActionHandler(app app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}
		action, err := app.GetAction(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Code:  http.StatusInternalServerError,
				Error: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, BaseResponse{
			Code: 0,
			Data: action,
		})
	}
}

func ListActionsHandler(app app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		actions, err := app.ListActions()
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Code:  http.StatusInternalServerError,
				Error: err.Error(),
			})
			return
		}
		if actions == nil {
			actions = make([]*model.Action, 0)
		}
		c.JSON(http.StatusOK, BaseResponse{
			Code: 0,
			Data: actions,
		})
	}
}

func ListActionsByGroupIdHandler(app app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		groupIdStr := c.Query("group")
		groupId, err := strconv.ParseInt(groupIdStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}
		actions, err := app.ListActionsByGroupId(groupId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Code:  http.StatusInternalServerError,
				Error: err.Error(),
			})
			return
		}
		if actions == nil {
			actions = make([]*model.Action, 0)
		}
		c.JSON(http.StatusOK, BaseResponse{
			Code: 0,
			Data: actions,
		})
	}
}

func DeleteActionHandler(app app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}
		err = app.DeleteAction(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Code:  http.StatusInternalServerError,
				Error: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, BaseResponse{
			Code: 0,
			Data: gin.H{"id": id},
		})
	}
}

func UpdateActionGroupIdHandler(app app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}
		groupIdStr := c.Query("group")
		groupId, err := strconv.ParseInt(groupIdStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}
		err = app.UpdateGroupId(id, groupId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Code:  http.StatusInternalServerError,
				Error: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, BaseResponse{
			Code: 0,
			Data: gin.H{"id": id},
		})
	}
}
