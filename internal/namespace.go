package internal

import (
	"context"
	"net/http"

	"github.com/ddr4869/k8s-kafka/internal/dto"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (s *Server) GetAllNamespace(c *gin.Context) {
	ns, _ := s.K8sClient.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	nsList, err := dto.V1NamespaceToJson(ns)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, nsList)
}
