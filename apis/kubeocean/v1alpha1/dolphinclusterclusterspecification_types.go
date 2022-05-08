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
	WaitToAssignDolphinClusterThreshold                    *int32 `json:"waitToAssigDolphinClusterThreshold,omitempty"`
	WaitToAssignKubesphereInstalledDolphinClusterThreshold *int32 `json:"waitToAssigKubesphereInstalledDolphinClusterThreshold,omitempty"`
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

// DolphinClusterSpecificationSpec defines the desired state of DolphinclusterClusterSpecification
type DolphinClusterSpecificationSpec struct {
	ControlPlaneSpecification ControlPlaneSpecification `json:"controlPlaneSpecification,omitempty"`
	Specifications            []Specification           `json:"specifications,omitempty"`
	ClusterMaxNodeCount       *int32                    `json:"clusterMaxNodeCount,omitempty"`
}

type DolphinClusterPool struct {
	Name        string   `json:"name,omitempty"`
	NotAssigned []string `json:"notAssigned,omitempty"`
}

type WorkerNodePool struct {
	Name         string   `json:"name,omitempty"`
	Converted    []string `json:"converted,omitempty"`
	NotConverted []string `json:"notConverted,omitempty"`
}

// DolphinClusterSpecificationStatus defines the observed state of DolphinClusterSpecification
type DolphinClusterSpecificationStatus struct {
	DolphinClusterPool                    []DolphinClusterPool `json:"dolphinClusterPool,omitempty"`
	KubesphereInstalledDolphinClusterPool []DolphinClusterPool `json:"kubesphereInstalledDolphinClusterPool,omitempty"`
	WorkerNodePool                        []WorkerNodePool     `json:"workerNodePool,omitempty"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:scope=Cluster,shortName=vcs
//+kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// DolphinClusterSpecification is the Schema for the DolphinClusterSpecifications API
type DolphinClusterSpecification struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DolphinClusterSpecificationSpec   `json:"spec,omitempty"`
	Status DolphinClusterSpecificationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//+kubebuilder:object:root=true

// DolphinClusterSpecificationList contains a list of DolphinClusterSpecification
type DolphinClusterSpecificationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DolphinClusterSpecification `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DolphinClusterSpecification{}, &DolphinClusterSpecificationList{})
}
