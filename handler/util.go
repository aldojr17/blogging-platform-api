package handler

import (
	"blogging-platform-api/utils/pagination"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PaginationSuccessResponse(c *gin.Context, data *pagination.Page, message string) {
	result := map[string]interface{}{
		"is_success": true,
		"data": map[string]interface{}{
			"data":         data.Data,
			"current_page": data.CurrentPage,
			"total":        data.Total,
			"total_page":   data.TotalPage,
			"limit":        data.Limit,
		},
		"message": message,
	}
	c.JSON(http.StatusOK, result)
}

func sendResponse(c *gin.Context, isSuccess bool, data interface{}, message string, code int) {
	response := map[string]interface{}{
		"is_success": isSuccess,
		"message":    message,
		"data":       data,
	}

	c.JSON(code, response)
}

func ResponseOK(c *gin.Context, data interface{}, message string) {
	sendResponse(c, true, data, message, http.StatusOK)
}

func ResponseCreated(c *gin.Context, data interface{}, message string) {
	sendResponse(c, true, data, message, http.StatusCreated)
}

func ResponseBadRequest(c *gin.Context, err error) {
	sendResponse(c, false, nil, err.Error(), http.StatusBadRequest)
}

func ResponseUnauthorized(c *gin.Context, err error) {
	sendResponse(c, false, nil, err.Error(), http.StatusUnauthorized)
}

func ResponseNotFound(c *gin.Context, err error) {
	sendResponse(c, false, nil, err.Error(), http.StatusNotFound)
}

func ResponseUnsupportedMediaType(c *gin.Context, err error) {
	sendResponse(c, false, nil, err.Error(), http.StatusUnsupportedMediaType)
}

func ResponseInternalServerError(c *gin.Context, err error) {
	sendResponse(c, false, nil, err.Error(), http.StatusInternalServerError)
}

func queryParamOrNull(req *http.Request, key string) interface{} {
	if value := req.FormValue(key); value != "" {
		return value
	}
	return nil
}

func queryLikeParamOrNull(req *http.Request, key string) interface{} {
	likeParam := queryParamOrNull(req, key)
	if likeParam != nil {
		return "%" + likeParam.(string) + "%"
	} else {
		return "%%"
	}
}
