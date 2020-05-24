package routers

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	converter "github.com/kiransabne/converter/converter/controllers"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/converter", converter.ConvertFileRoutes())
	})

	return router
}
