package dto

type GetNamespacePodRequest struct {
	Namespace string `uri:"namespace" binding:"required"`
}

type GetPodNameRequest struct {
	Name string `form:"name" binding:"required"`
}

type CreateNamespacePodRequest struct {
	FileName string `json:"file_name" binding:"required"`
}
