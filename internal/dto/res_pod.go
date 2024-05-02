package dto

import (
	"encoding/json"
	"fmt"

	"github.com/ddr4869/k8s-kafka/internal/model/k8s"
	v1 "k8s.io/api/core/v1"
)

func V1PodToJson(v1Pod *v1.Pod) (*k8s.Pod, error) {
	jsonBytes, err := json.Marshal(v1Pod)
	if err != nil {
		fmt.Println("JSON marshaling error:", err)
		return nil, err
	}
	var pod k8s.Pod
	err = json.Unmarshal(jsonBytes, &pod)
	if err != nil {
		fmt.Println("JSON Unmarshaling error:", err)
		return nil, err
	}
	return &pod, nil
}

func V1PodListToJson(v1Pods *v1.PodList) (*k8s.PodList, error) {
	jsonBytes, err := json.Marshal(v1Pods)
	if err != nil {
		fmt.Println("JSON marshaling error:", err)
		return nil, err
	}
	var podList k8s.PodList
	err = json.Unmarshal(jsonBytes, &podList)
	if err != nil {
		fmt.Println("JSON Unmarshaling error:", err)
		return nil, err
	}
	return &podList, nil
}
