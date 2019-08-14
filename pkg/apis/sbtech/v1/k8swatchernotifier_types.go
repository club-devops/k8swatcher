package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// K8SWatcherNotifierSpec defines the desired state of K8SWatcherNotifier
// +k8s:openapi-gen=true
type DestSlack struct {
	SlackUrl string `json:"slack_url"`
	Channel  string `json:"channel"`
}
type K8SWatcherNotifierSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Slack DestSlack `json:"slack"`
}

// K8SWatcherNotifierStatus defines the observed state of K8SWatcherNotifier
// +k8s:openapi-gen=true
type K8SWatcherNotifierStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// K8SWatcherNotifier is the Schema for the k8swatchernotifiers API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type K8SWatcherNotifier struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   K8SWatcherNotifierSpec   `json:"spec,omitempty"`
	Status K8SWatcherNotifierStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// K8SWatcherNotifierList contains a list of K8SWatcherNotifier
type K8SWatcherNotifierList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []K8SWatcherNotifier `json:"items"`
}

func init() {
	SchemeBuilder.Register(&K8SWatcherNotifier{}, &K8SWatcherNotifierList{})
}
