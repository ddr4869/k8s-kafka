package internal

import (
	"log"
	"net/http"
	"os"
)

func RouteSetUp(s *Server) {
	r := s.router
	api := r.Group("/api")

	api.GET("/pods", s.GetAllPods)
	api.GET("/pods/:namespace", s.GetNamespacePodListValid, s.GetNamespacePods)
	api.GET("/pod/:namespace", s.GetPodValid, s.GetPod)
	api.GET("/pod/:namespace/log", s.GetPodValid, s.GetPodLog)
	api.POST("/pods", s.CreatePodValid, s.CreateNamespacePods)

	api.GET("/namespaces", s.GetAllNamespace)
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
