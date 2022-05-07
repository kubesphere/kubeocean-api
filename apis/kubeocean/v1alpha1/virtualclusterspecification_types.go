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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Specification struct {
	Name                                                   string `json:"name,omitempty"`
	WaitToAssignVirtualClusterThreshold                    *int32 `json:"waitToAssigVirtualClusterThreshold,omitempty"`
	WaitToAssignKubesphereInstalledVirtualClusterThreshold *int32 `json:"waitToAssigKubesphereInstalledVirtualClusterThreshold,omitempty"`
	WaitToConvertWorkerNodeThreshold                       *int32 `json:"waitToAssigWorkerNodeThreshold,omitempty"`
	CpuCurrent                                             *int32 `json:"cpuCurrent,omitempty"`
	MemoryCurrent                                          *int32 `json:"memoryCurrent,omitempty"`
	DiskSize                                               *int32 `json:"diskSize,omitempty"`
}

type ControlPlaneSpecification struct {
	Name          string `json:"name,omitempty"`
	CpuCurrent    *int32 `json:"cpuCurrent,omitempty"`
	MemoryCurrent *int32 `json:"memoryCurrent,omitempty"`
	DiskSize      *int32 `json:"diskSize,omitempty"`
	Replicas      *int32 `json:"replicas,omitempty"`
}

// VirtualClusterSpecificationSpec defines the desired state of VirtualClusterSpecification
type VirtualClusterSpecificationSpec struct {
	ControlPlaneSpecification ControlPlaneSpecification `json:"controlPlaneSpecification,omitempty"`
	Specifications            []Specification           `json:"specifications,omitempty"`
	ClusterMaxNodeCount       *int32                    `json:"clusterMaxNodeCount,omitempty"`
}

type VirtualClusterPool struct {
	Name        string   `json:"name,omitempty"`
	NotAssigned []string `json:"notAssigned,omitempty"`
}

type WorkerNodePool struct {
	Name         string   `json:"name,omitempty"`
	Converted    []string `json:"converted,omitempty"`
	NotConverted []string `json:"notConverted,omitempty"`
}

// VirtualClusterSpecificationStatus defines the observed state of VirtualClusterSpecification
type VirtualClusterSpecificationStatus struct {
	VirtualClusterPool                    []VirtualClusterPool `json:"virtualClusterPool,omitempty"`
	KubesphereInstalledVirtualClusterPool []VirtualClusterPool `json:"kubesphereInstalledVirtualClusterPool,omitempty"`
	WorkerNodePool                        []WorkerNodePool     `json:"workerNodePool,omitempty"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:scope=Cluster,shortName=vcs
//+kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// VirtualClusterSpecification is the Schema for the virtualclusterspecifications API
type VirtualClusterSpecification struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualClusterSpecificationSpec   `json:"spec,omitempty"`
	Status VirtualClusterSpecificationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//+kubebuilder:object:root=true

// VirtualClusterSpecificationList contains a list of VirtualClusterSpecification
type VirtualClusterSpecificationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualClusterSpecification `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VirtualClusterSpecification{}, &VirtualClusterSpecificationList{})
}
