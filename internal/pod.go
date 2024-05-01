package internal

import (
	"context"
	"net/http"

	"github.com/ddr4869/k8s-kafka/internal/dto"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (s *Server) GetDefaultsPods(c *gin.Context) {
	pods, _ := s.K8sClient.CoreV1().Pods("default").List(context.TODO(), v1.ListOptions{})
	podList, err := dto.V1PodListToJson(pods)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, podList)
}

func (s *Server) GetNamespacePods(c *gin.Context) {
	req := c.MustGet("req").(dto.GetNamespacePodRequest)
	pods, _ := s.K8sClient.CoreV1().Pods(req.Namespace).List(context.TODO(), v1.ListOptions{})
	podList, err := dto.V1PodListToJson(pods)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, podList)
}
