package server

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (s *Server) Serve() error {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.HandleFunc("/", s.healthCheck)
	r.Options("/*", handleOptions)
	r.Post("/auth", s.auth)

	// https://miaorg.atlassian.net/wiki/spaces/BOPP/pages/1420918992/Transaction+Monitoring
	r.Route("/v1", func(r chi.Router) {
		r.Use(middleware.Logger)
		r.Get("/quotes", s.getQuotes)
		r.Post("/categories", s.getCategories)
		r.Get("/categories", s.postCategories)
		// r.Get("/mlr/over-amount-paid-pi/{interval}/{amount}/{times}", s.ReportOverAmount)
	})

	log.Println("Listening on :" + s.Port)
	return http.ListenAndServe(":"+s.Port, r)

}
