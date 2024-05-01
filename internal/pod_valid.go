package internal

import (
	"net/http"

	"github.com/ddr4869/k8s-kafka/internal/dto"
	"github.com/gin-gonic/gin"
)

func (s *Server) GetNamespacePodValid(c *gin.Context) {
	var req dto.GetNamespacePodRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "Failed to bind request")
		return
	}
	c.Set("req", req)
	c.Next()
}
