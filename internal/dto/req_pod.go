package dto

type GetNamespacePodRequest struct {
	Namespace string `form:"namespace" binding:"required"`
}
