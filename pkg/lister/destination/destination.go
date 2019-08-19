package destination

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Destination struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec DestinationSpec `json:"spec"`
}

type DestinationSpec struct {
	IpRange string `json:"ipRange"`
	Port    int32  `json:"port"`
}
