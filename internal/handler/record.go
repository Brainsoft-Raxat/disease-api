package handler

import (
	"disease-api/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *handler) CreateRecord(c echo.Context) error {
	var req models.Record
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	ctx := c.Request().Context()
	email, err := h.service.RecordService.CreateRecord(ctx, &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorMessage{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"email": email,
	})
}

func (h *handler) UpdateRecord(c echo.Context) error {
	var req models.Record
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	ctx := c.Request().Context()
	err = h.service.RecordService.UpdateRecord(ctx, &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorMessage{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.StatusResponse{
		Code:    200,
		Message: "updated",
	})
}

func (h *handler) DeleteRecord(c echo.Context) error {
	var req models.Record
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	ctx := c.Request().Context()
	err = h.service.RecordService.DeleteRecord(ctx, req.Email, req.CName, req.DiseaseCode)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorMessage{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.StatusResponse{
		Code:    200,
		Message: "deleted",
	})
}

func (h *handler) GetRecordsFilter(c echo.Context) error {
	var req models.Record
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	ctx := c.Request().Context()
	resp, err := h.service.RecordService.GetRecordsFilter(ctx, req.Email, req.CName, req.DiseaseCode)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorMessage{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resp)
}
