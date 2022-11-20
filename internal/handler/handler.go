package handler

import (
	"disease-api/internal/service"
	"github.com/labstack/echo/v4"
)

type handler struct {
	service *service.Service
}

type Handler interface {
	Register(e *echo.Echo)
}

func New(services *service.Service) Handler {
	return &handler{service: services}
}

func (h *handler) Register(e *echo.Echo) {
	e.Use()
	dt := e.Group("/disease_type")
	{
		dt.POST("/create", h.CreateDiseaseType)
		dt.PUT("/:id", h.UpdateDiseaseType)
		dt.DELETE("/:id", h.DeleteDiseaseType)
		dt.GET("/:id", h.GetDiseaseType)
		dt.GET("/all", h.GetAllDiseaseTypes)
	}

	d := e.Group("/disease")
	{
		d.POST("/create", h.CreateDisease)
		d.PUT("/:code", h.UpdateDisease)
		d.DELETE("/:code", h.DeleteDisease)
		d.GET("/:code", h.GetDisease)
		d.GET("/all", h.GetAllDiseases)
	}

	dsc := e.Group("/discover")
	{
		dsc.POST("/create", h.CreateDiscover)
		dsc.PUT("/:code", h.UpdateDiscover)
		dsc.DELETE("", h.DeleteDiscover)
		dsc.GET("", h.GetDiscover)
		dsc.GET("/all", h.GetAllDiscovers)
	}

	dc := e.Group("/doctor")
	{
		dc.POST("/create", h.CreateDoctor)
		dc.PUT("/:email", h.UpdateDoctor)
		dc.DELETE("/:email", h.DeleteDoctor)
		dc.GET("/:email", h.GetDoctor)
		dc.GET("/all", h.GetAllDoctors)
	}

	c := e.Group("/country")
	{
		c.POST("/create", h.CreateCountry)
		c.PUT("/:cname", h.UpdateCountry)
		c.DELETE("/:cname", h.DeleteCountry)
		c.GET("/:cname", h.GetCountry)
		c.GET("/all", h.GetAllCountries)
	}

	ps := e.Group("/public_servant")
	{
		ps.POST("/create", h.CreatePublicServant)
		ps.PUT("/:email", h.UpdatePublicServant)
		ps.DELETE("/:email", h.DeletePublicServant)
		ps.GET("/:email", h.GetPublicServant)
		ps.GET("/all", h.GetAllPublicServants)
	}
	r := e.Group("/record")
	{
		r.POST("/create", h.CreateRecord)
		r.PUT("/update", h.UpdateRecord)
		r.DELETE("/delete", h.DeleteRecord)
		r.GET("/filter", h.GetRecordsFilter)
	}
}
