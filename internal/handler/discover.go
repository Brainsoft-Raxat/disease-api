package handler

import (
	"disease-api/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *handler) CreateDiscover(c echo.Context) error {
	var req models.Discover
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	ctx := c.Request().Context()
	code, err := h.service.CreateDiscover(ctx, &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorMessage{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"disease_code": code,
	})
}

func (h *handler) UpdateDiscover(c echo.Context) error {
	var req models.Discover
	diseaseCode := c.Param("code")

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	req.DiseaseCode = diseaseCode

	ctx := c.Request().Context()
	err = h.service.UpdateDiscover(ctx, &req)
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

func (h *handler) DeleteDiscover(c echo.Context) error {
	country := c.QueryParam("cname")
	diseaseCode := c.QueryParam("code")

	ctx := c.Request().Context()
	err := h.service.DeleteDiscover(ctx, country, diseaseCode)
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

func (h *handler) GetDiscover(c echo.Context) error {
	country := c.QueryParam("cname")
	diseaseCode := c.QueryParam("code")

	ctx := c.Request().Context()
	resp, err := h.service.GetDiscover(ctx, country, diseaseCode)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorMessage{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *handler) GetAllDiscovers(c echo.Context) error {
	ctx := c.Request().Context()
	resp, err := h.service.GetAllDiscovers(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorMessage{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resp)
}
