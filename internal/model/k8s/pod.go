package k8s

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
	Spec   PodSpec   `json:"spec"`
	Status PodStatus `json:"status"`
}

type PodSpec struct {
	// Volumes                       []Volume                          `json:"volumes"`
	// Containers                    []Container                       `json:"containers"`
	// RestartPolicy                 corev1.RestartPolicy              `json:"restartPolicy"`
	TerminationGracePeriodSeconds *int64 `json:"terminationGracePeriodSeconds,omitempty"`
	ActiveDeadlineSeconds         *int64 `json:"activeDeadlineSeconds,omitempty"`
	// DNSPolicy                     corev1.DNSPolicy                  `json:"dnsPolicy,omitempty"`
	NodeSelector                 map[string]string `json:"nodeSelector,omitempty"`
	ServiceAccountName           string            `json:"serviceAccountName,omitempty"`
	AutomountServiceAccountToken *bool             `json:"automountServiceAccountToken,omitempty"`
	HostNetwork                  bool              `json:"hostNetwork,omitempty"`
	HostPID                      bool              `json:"hostPID,omitempty"`
	HostIPC                      bool              `json:"hostIPC,omitempty"`
	// SecurityContext               *corev1.PodSecurityContext        `json:"securityContext,omitempty"`
	SchedulerName      string `json:"schedulerName,omitempty"`
	PriorityClassName  string `json:"priorityClassName,omitempty"`
	Priority           *int32 `json:"priority,omitempty"`
	EnableServiceLinks *bool  `json:"enableServiceLinks,omitempty"`
	// PreemptionPolicy              *corev1.PreemptionPolicy          `json:"preemptionPolicy,omitempty"`
	// Overhead                      corev1.ResourceList               `json:"overhead,omitempty"`
	// TopologySpreadConstraints     []corev1.TopologySpreadConstraint `json:"topologySpreadConstraints,omitempty"`
}

type PodStatus struct {
	// Phase                      corev1.PodPhase          `json:"phase,omitempty"`
	// Conditions                 []corev1.PodCondition    `json:"conditions,omitempty"`
	Message string `json:"message,omitempty"`
	Reason  string `json:"reason,omitempty"`
	HostIP  string `json:"hostIP,omitempty"`
	PodIP   string `json:"podIP,omitempty"`
	// StartTime                  metav1.Time              `json:"startTime,omitempty"`
	// InitContainerStatuses      []corev1.ContainerStatus `json:"initContainerStatuses,omitempty"`
	// ContainerStatuses          []corev1.ContainerStatus `json:"containerStatuses,omitempty"`
	// QOSClass                   corev1.PodQOSClass       `json:"qosClass,omitempty"`
	// EphemeralContainerStatuses []corev1.ContainerStatus `json:"ephemeralContainerStatuses,omitempty"`
	// 필요에 따라 다른 필드 추가 가능
}
