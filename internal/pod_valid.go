package internal

import (
	"net/http"

	"github.com/ddr4869/k8s-kafka/internal/dto"
	"github.com/gin-gonic/gin"
)

func (s *Server) GetNamespacePodListValid(c *gin.Context) {
	var reqUri dto.GetNamespacePodRequest
	if err := c.ShouldBindUri(&reqUri); err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "Failed to bind request")
		return
	}
	c.Set("reqUri", reqUri)
	c.Next()
}

func (s *Server) GetPodValid(c *gin.Context) {
	var reqUri dto.GetNamespacePodRequest
	if err := c.ShouldBindUri(&reqUri); err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "Failed to bind request")
		return
	}
	var req dto.GetPodNameRequest
	if err := c.ShouldBind(&req); err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "Failed to bind request")
		return
	}

	c.Set("reqUri", reqUri)
	c.Set("req", req)
	c.Next()
}

func (s *Server) CreatePodValid(c *gin.Context) {
	var req dto.CreateNamespacePodRequest
	if err := c.ShouldBind(&req); err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "Failed to bind request")
		return
	}
	c.Set("req", req)
	c.Next()
}
