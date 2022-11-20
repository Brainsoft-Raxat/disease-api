package handler

import (
	"disease-api/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *handler) CreateCountry(c echo.Context) error {
	var req models.Country
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	ctx := c.Request().Context()
	cname, err := h.service.CountryService.CreateCountry(ctx, &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorMessage{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"cname": cname,
	})
}

func (h *handler) UpdateCountry(c echo.Context) error {
	var req models.Country
	cname := c.Param("cname")

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	req.Cname = cname

	ctx := c.Request().Context()
	err = h.service.CountryService.UpdateCountry(ctx, &req)
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

func (h *handler) DeleteCountry(c echo.Context) error {
	cname := c.Param("cname")

	ctx := c.Request().Context()
	err := h.service.CountryService.DeleteCountry(ctx, cname)
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

func (h *handler) GetCountry(c echo.Context) error {
	cname := c.Param("cname")

	ctx := c.Request().Context()
	resp, err := h.service.CountryService.GetCountry(ctx, cname)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorMessage{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *handler) GetAllCountries(c echo.Context) error {
	ctx := c.Request().Context()
	resp, err := h.service.CountryService.GetAllCountries(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorMessage{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resp)
}
