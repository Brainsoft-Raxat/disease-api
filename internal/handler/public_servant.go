package handler

import (
	"disease-api/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *handler) CreatePublicServant(c echo.Context) error {
	var req models.PublicServant
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	ctx := c.Request().Context()
	email, err := h.service.PublicServantService.CreatePublicServant(ctx, &req)
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

func (h *handler) UpdatePublicServant(c echo.Context) error {
	var req models.PublicServant
	email := c.Param("email")

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	req.Email = email

	ctx := c.Request().Context()
	err = h.service.PublicServantService.UpdatePublicServant(ctx, &req)
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

func (h *handler) DeletePublicServant(c echo.Context) error {
	email := c.Param("email")

	ctx := c.Request().Context()
	err := h.service.PublicServantService.DeletePublicServant(ctx, email)
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

func (h *handler) GetPublicServant(c echo.Context) error {
	email := c.Param("email")

	ctx := c.Request().Context()
	resp, err := h.service.PublicServantService.GetPublicServant(ctx, email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorMessage{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *handler) GetAllPublicServants(c echo.Context) error {
	ctx := c.Request().Context()
	resp, err := h.service.PublicServantService.GetAllPublicServants(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorMessage{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resp)
}
