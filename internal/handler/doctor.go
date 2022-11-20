package handler

import (
	"disease-api/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *handler) CreateDoctor(c echo.Context) error {
	var req models.Doctor
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	ctx := c.Request().Context()
	email, err := h.service.DoctorService.CreateDoctor(ctx, &req)
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

func (h *handler) UpdateDoctor(c echo.Context) error {
	var req models.Doctor
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
	err = h.service.DoctorService.UpdateDoctor(ctx, &req)
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

func (h *handler) DeleteDoctor(c echo.Context) error {
	email := c.Param("email")

	ctx := c.Request().Context()
	err := h.service.DoctorService.DeleteDoctor(ctx, email)
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

func (h *handler) GetDoctor(c echo.Context) error {
	email := c.Param("email")

	ctx := c.Request().Context()
	resp, err := h.service.DoctorService.GetDoctor(ctx, email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorMessage{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *handler) GetAllDoctors(c echo.Context) error {
	ctx := c.Request().Context()
	resp, err := h.service.DoctorService.GetAllDoctors(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorMessage{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resp)
}
