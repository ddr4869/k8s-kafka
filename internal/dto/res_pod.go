package dto

import (
	"encoding/json"
	"fmt"

	v1 "k8s.io/api/core/v1"
)

type PodList struct {
	ListMeta struct {
		TotalItems int    `json:"totalItems"`
		SelfLink   string `json:"selfLink"`
	} `json:"metadata"`
	Items []Pod `json:"items"`
}

type Pod struct {
	ObjectMeta struct {
		Name      string            `json:"name"`
		Namespace string            `json:"namespace"`
		Labels    map[string]string `json:"labels"`
	} `json:"metadata"`
	// Other fields omitted for brevity
}

func V1PodListToJson(pods *v1.PodList) (*PodList, error) {
	jsonBytes, err := json.Marshal(pods)
	if err != nil {
		fmt.Println("JSON marshaling error:", err)
		return nil, err
	}
	var podList PodList
	err = json.Unmarshal(jsonBytes, &podList)
	if err != nil {
		fmt.Println("JSON Unmarshaling error:", err)
		return nil, err
	}
	return &podList, nil
}
