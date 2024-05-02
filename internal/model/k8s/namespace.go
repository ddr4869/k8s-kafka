package k8s

type NamespaceList struct {
	ListMeta struct {
		TotalItems int    `json:"totalItems"`
		SelfLink   string `json:"selfLink"`
	} `json:"metadata"`
	Items []Namespace `json:"items"`
	// Add other fields as needed
}

type Namespace struct {
	InnterMeta struct {
		Name string `json:"name"`
		UID  string `json:"uid"`
	} `json:"metadata"`
}
