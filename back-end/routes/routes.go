package routes

import (
	"net/http"

	"github.com/esuEdu/casa-oliveira/internal/handlers"
	"github.com/esuEdu/casa-oliveira/internal/middleware"
	"github.com/esuEdu/casa-oliveira/internal/repositories"
	"github.com/esuEdu/casa-oliveira/internal/service"
	"gorm.io/gorm"
)

func SetupRoutes(r *http.ServeMux, db *gorm.DB) {
	SetupProductRoutes(r, db)
}

func SetupProductRoutes(r *http.ServeMux, db *gorm.DB) {

	repo := repositories.NewProductRepo(db)
	service := service.NewProductService(repo)
	handler := handlers.NewProductHandler(service)

	// /api/product/
	r.Handle("/api/product", withMiddleware(
		http.HandlerFunc(handler.HandleProduct),
		middleware.EnsureValidToken(),
	))

	// /api/product/:id
	r.Handle("/api/product/", withMiddleware(
		http.HandlerFunc(handler.HandleProductByID),
		middleware.EnsureValidToken(),
	))
}

func withMiddleware(h http.Handler, m ...func(http.Handler) http.Handler) http.Handler {
	for _, middleware := range m {
		h = middleware(h)
	}
	return h
}
