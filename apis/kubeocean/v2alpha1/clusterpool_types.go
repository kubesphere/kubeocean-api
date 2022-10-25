/*
Copyright 2020 The KubeSphere Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v2alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type Specification struct {
	ClusterSpec  `json:",inline"`
	WaitToAssign int32    `json:"waitToAssign,omitempty"`
	Tags         []string `json:"tags,omitempty"`
}

// ClusterPoolSpec defines the desired state of ClusterPool
type ClusterPoolSpec struct {
	Specifications []Specification `json:"specifications,omitempty"`
	MaxClusters    int32           `json:"maxClusters,omitempty"`
}

type ClusterName string
type SpecificationStatus map[ClusterPhase][]ClusterName

// ClusterPoolStatus defines the observed state of ClusterPool
type ClusterPoolStatus struct {
	SpecificationStatuses map[SpecificationID]SpecificationStatus `json:"specificationStatuses,omitempty"`
	CountByPhase          map[ClusterPhase]int32                  `json:"countByPhase,omitempty"`
	TotalCount            int32                                   `json:"totalCount,omitempty"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:scope=Cluster,shortName=cp
//+kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// ClusterPool is the Schema for the clusterpools API
type ClusterPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterPoolSpec   `json:"spec,omitempty"`
	Status ClusterPoolStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterPoolList contains a list of ClusterPool
type ClusterPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterPool `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterPool{}, &ClusterPoolList{})
}
