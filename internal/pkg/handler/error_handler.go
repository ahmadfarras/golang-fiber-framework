package handler

import (
	"ahmadfarras/fiberframework/internal/pkg/domain/model/response"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

func HandleError(err error) response.Response {

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return response.BuildErrorResponse(http.StatusNotFound, err.Error())
	}

	return response.BuildErrorResponse(http.StatusInternalServerError, err.Error())
}
