package dto

import (
	v1 "k8s.io/api/core/v1"
)

type NamespaceList struct {
	Name []string `json:"name"`
}

func V1NamespaceToJson(v1Ns *v1.NamespaceList) (*NamespaceList, error) {
	var nsList NamespaceList
	for _, item := range v1Ns.Items {
		nsList.Name = append(nsList.Name, item.Name)
	}
	return &nsList, nil
}
