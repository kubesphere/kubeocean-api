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

type SpecificationPluginType string

type InstanceTypeType string

type Specification struct {
	SpecificationID       SpecificationPluginType `json:"specificationId,omitempty"`
	InstanceType          InstanceTypeType        `json:"instanceType,omitempty"`
	SshKey                string                  `json:"sshKey,omitempty"`
	CpuCurrent            int32                   `json:"cpuCurrent,omitempty"`
	MemoryCurrent         int32                   `json:"memoryCurrent,omitempty"`
	DiskSize              int32                   `json:"diskSize,omitempty"`
	WaitToAssignThreshold int32                   `json:"waitToAssignThreshold,omitempty"`
}

type OsImageType string

// KindClusterPoolSpec defines the desired state of KindClusterPool
type KindClusterPoolSpec struct {
	Specifications                   map[OsImageType]Specification `json:"specificaiton"`
	TargetClusterMaxKindClusterCount int32                         `json:"targetClusterMaxKindClusterCount"`
}

type SpecificationStatus struct {
	Assigned    []string `json:"assigned"`
	NotAssigned []string `json:"notAssigned"`
}

// KindClusterPoolStatus defines the observed state of KindClusterPool
type KindClusterPoolStatus struct {
	SpecificaitonsStatus map[InstanceTypeType]map[SpecificationPluginType]SpecificationStatus `json:"specificaitonsStatus"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:scope=Cluster,shortName=kcp
//+kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// KindClusterPool is the Schema for the kindclusterpools API
type KindClusterPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KindClusterPoolSpec   `json:"spec,omitempty"`
	Status KindClusterPoolStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KindClusterPoolList contains a list of KindClusterPool
type KindClusterPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KindClusterPool `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KindClusterPool{}, &KindClusterPoolList{})
}
