package internal

import (
	"log"
	"net/http"
	"os"
)

func RouteSetUp(s *Server) {
	r := s.router
	api := r.Group("/api")

	api.GET("/pods", s.GetDefaultsPods)
	api.GET("/pods/:name", s.GetNamespacePodValid, s.GetNamespacePods)
}

func (s *Server) Start() error {

	srv := &http.Server{
		Addr:    ":" + os.Getenv("SERVER_PORT"),
		Handler: s.router,
	}

	log.Printf("Listening and serving HTTP on %s\n", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Printf("listen: %s\n", err)
		return err
	}

	return nil
}
