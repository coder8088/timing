package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/tedux/timing/pkg/app"
	"github.com/tedux/timing/pkg/model"
)

func AddActionGroupHandler(app app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var actionGroup model.ActionGroup
		if err := c.ShouldBindJSON(&actionGroup); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}
		id, err := app.AddActionGroup(&actionGroup)
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

func GetActionGroupHandler(app app.App) gin.HandlerFunc {
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
		action, err := app.GetActionGroup(id)
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

func ListActionGroupsHandler(app app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		groups, err := app.ListActionGroups()
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Code:  http.StatusInternalServerError,
				Error: err.Error(),
			})
			return
		}
		if groups == nil {
			groups = make([]*model.ActionGroup, 0)
		}
		c.JSON(http.StatusOK, BaseResponse{
			Code: 0,
			Data: groups,
		})
	}
}

func DeleteActionGroupHandler(app app.App) gin.HandlerFunc {
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
		err = app.DeleteActionGroup(id)
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
