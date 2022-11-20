package handler

import (
	"disease-api/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *handler) CreateDisease(c echo.Context) error {
	var req models.Disease
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	ctx := c.Request().Context()
	code, err := h.service.DiseaseService.CreateDisease(ctx, &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorMessage{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code": code,
	})
}

func (h *handler) UpdateDisease(c echo.Context) error {
	var req models.Disease
	diseaseCode := c.Param("code")

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	req.Code = diseaseCode

	ctx := c.Request().Context()
	err = h.service.DiseaseService.UpdateDisease(ctx, &req)
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

func (h *handler) DeleteDisease(c echo.Context) error {
	diseaseCode := c.Param("code")

	ctx := c.Request().Context()
	err := h.service.DiseaseService.DeleteDisease(ctx, diseaseCode)
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

func (h *handler) GetDisease(c echo.Context) error {
	diseaseCode := c.Param("code")

	ctx := c.Request().Context()
	resp, err := h.service.DiseaseService.GetDisease(ctx, diseaseCode)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorMessage{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *handler) GetAllDiseases(c echo.Context) error {
	ctx := c.Request().Context()
	resp, err := h.service.DiseaseService.GetAllDiseases(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorMessage{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *handler) CreateDiseaseType(c echo.Context) error {
	var req models.DiseaseType
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	ctx := c.Request().Context()
	id, err := h.service.DiseaseService.CreateDiseaseType(ctx, &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorMessage{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *handler) UpdateDiseaseType(c echo.Context) error {
	var req models.DiseaseType
	diseaseTypeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	err = c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	req.ID = diseaseTypeID

	ctx := c.Request().Context()
	err = h.service.DiseaseService.UpdateDiseaseType(ctx, &req)
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

func (h *handler) DeleteDiseaseType(c echo.Context) error {
	diseaseTypeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	ctx := c.Request().Context()
	err = h.service.DiseaseService.DeleteDiseaseType(ctx, diseaseTypeID)
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

func (h *handler) GetDiseaseType(c echo.Context) error {
	diseaseTypeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	ctx := c.Request().Context()
	resp, err := h.service.DiseaseService.GetDiseaseType(ctx, diseaseTypeID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorMessage{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *handler) GetAllDiseaseTypes(c echo.Context) error {
	ctx := c.Request().Context()
	resp, err := h.service.DiseaseService.GetAllDiseaseTypes(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorMessage{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resp)
}
