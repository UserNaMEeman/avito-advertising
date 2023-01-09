package handler

import (
	"advertising/internal/service"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/newproduct", h.CreateProduct)
	router.Post("/getproducts", h.ReturnProducts)
	return router
}
