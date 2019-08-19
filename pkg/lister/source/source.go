package source

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Source struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec SourceSpec `json:"spec"`
}

type SourceSpec struct {
	IpRange string `json:"ipRange"`
	Port    int32  `json:"port"`
}
