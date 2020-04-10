package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/tedux/timing/pkg/app"
	"github.com/tedux/timing/pkg/model"
)

func StartTimingHandler(app app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		actionName := c.Query("action")
		if len(actionName) == 0 {
			c.JSON(http.StatusBadRequest, ErrorResponse{
				Code:  http.StatusBadRequest,
				Error: "action name is required",
			})
		}
		id, err := app.StartTiming(actionName)
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

func StopTimingHandler(app app.App) gin.HandlerFunc {
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
		err = app.StopTiming(id)
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

func ListTimingByActionAndDtHandler(app app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		actionName := c.Query("action")
		dt := c.Query("dt")
		timings, err := app.SearchTimingsByActionAndDt(actionName, dt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Code:  http.StatusInternalServerError,
				Error: err.Error(),
			})
			return
		}
		if timings == nil {
			timings = make([]*model.Timing, 0)
		}
		c.JSON(http.StatusOK, BaseResponse{
			Code: 0,
			Data: timings,
		})
	}
}
