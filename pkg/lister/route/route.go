package source

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Route struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec RouteSpec `json:"spec"`
}

type labelSelector struct {
	MatchLabels map[string]string `json:"matchLabels"`
}

type RouteSpec struct {
	Destination string        `json:"dest"`
	Selector    labelSelector `json:"selector"`
}
