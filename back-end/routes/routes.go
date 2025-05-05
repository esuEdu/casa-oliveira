package routes

import (
	"net/http"

	"github.com/esuEdu/casa-oliveira/internal/handlers"
	"github.com/esuEdu/casa-oliveira/internal/repositories"
	"github.com/esuEdu/casa-oliveira/internal/service"
	"gorm.io/gorm"
)

func SetupRoutes(r *http.ServeMux, db *gorm.DB) {
	setupAuthRoutes(r)
	setupProductRoutes(r, db)
}

func setupAuthRoutes(r *http.ServeMux) {
	service := service.NewAuthService()
	handler := handlers.NewAuthHandler(service)

	r.Handle("/api/signup", withMiddleware(
		http.HandlerFunc(handler.SingUp),
	))

	r.Handle("/api/signin", withMiddleware(
		http.HandlerFunc(handler.SignIn),
	))
}

func setupProductRoutes(r *http.ServeMux, db *gorm.DB) {

	repo := repositories.NewProductRepo(db)
	service := service.NewProductService(repo)
	handler := handlers.NewProductHandler(service)

	// /api/product/
	r.Handle("/api/product", withMiddleware(
		http.HandlerFunc(handler.HandleProduct),
	))

	// /api/product/:id
	r.Handle("/api/product/", withMiddleware(
		http.HandlerFunc(handler.HandleProductByID),
	))
}

func withMiddleware(h http.Handler, m ...func(http.Handler) http.Handler) http.Handler {
	for _, middleware := range m {
		h = middleware(h)
	}
	return h
}
