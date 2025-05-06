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
	setupProductRoutes(r, db)
	setupUserRoutes(r, db)
}

func setupUserRoutes(r *http.ServeMux, db *gorm.DB) {
	repo := repositories.NewUserRepo(db)
	service := service.NewUserService(repo)
	handler := handlers.NewUserHandler(service)

	// /api/signup
	r.Handle("/api/signup", withMiddleware(
		http.HandlerFunc(handler.SignUp),
	))

	// /api/signin
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
		middleware.AuthMiddleware,
	))

	// /api/product/:id
	r.Handle("/api/product/", withMiddleware(
		http.HandlerFunc(handler.HandleProductByID),
		middleware.AuthMiddleware,
	))
}

func withMiddleware(h http.Handler, m ...func(http.Handler) http.Handler) http.Handler {
	for _, middleware := range m {
		h = middleware(h)
	}
	return h
}
