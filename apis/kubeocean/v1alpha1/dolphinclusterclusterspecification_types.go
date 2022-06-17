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
	Name          *string `json:"name,omitempty"`
	CpuCurrent    *int32  `json:"cpuCurrent,omitempty"`
	MemoryCurrent *int32  `json:"memoryCurrent,omitempty"`
	DiskSize      *int32  `json:"diskSize,omitempty"`
	ImageID       *string `json:"imageID,omitempty"`
	SshKey        *string `json:"sshKey,omitempty"`
}

type ControlPlaneSpecification struct {
	Specification `json:",inline"`
	Replicas      *int32 `json:"replicas,omitempty"`
}

type WorkerNodeSpecification struct {
	Specification                                          `json:",inline"`
	InitReplica                                            *int32 `json:"initReplicas,omitempty"`
	WaitToAssignDolphinClusterThreshold                    *int32 `json:"waitToAssigDolphinClusterThreshold,omitempty"`
	WaitToAssignKubesphereInstalledDolphinClusterThreshold *int32 `json:"waitToAssigKubesphereInstalledDolphinClusterThreshold,omitempty"`
	WaitToConvertWorkerNodeThreshold                       *int32 `json:"waitToAssigWorkerNodeThreshold,omitempty"`
}

type Cluster struct {
	ControlPlaneSpecification ControlPlaneSpecification `json:"controlPlaneSpecification,omitempty"`
	WorkerNodeSpecification   []WorkerNodeSpecification `json:"workerNodeSpecification,omitempty"`
	MaxWorkerNodeCount        *int32                    `json:"maxWorkerNodeCount,omitempty"`
	Version                   *string                   `json:"version,omitempty"`
	Zone                      *string                   `json:"zone,omitempty"`
}

type ClusterResource struct {
	Cluster         Cluster `json:"cluster,omitempty"`
	Namespace       *string `json:"namespace,omitempty"`
	MaxClusterCount *int32  `json:"maxClusterCount,omitempty"`
	ExternalDomain  *string `json:"externalDomain,omitempty"`
}

// DolphinClusterSpecificationSpec defines the desired state of DolphinclusterClusterSpecification
type DolphinClusterSpecificationSpec struct {
	ClusterResource ClusterResource `json:"clusterResource"`
}

type WorkerNode struct {
	NodeName    string `json:"nodeName,omitempty"`
	ClusterName string `json:"clusterName,omitempty"`
}

type WorkerNodePool struct {
	Name         string       `json:"name,omitempty"`
	Converted    []WorkerNode `json:"converted,omitempty"`
	NotConverted []WorkerNode `json:"notConverted,omitempty"`
}

type DolphinClusterPool struct {
	Name        string   `json:"name,omitempty"`
	NotAssigned []string `json:"notAssigned,omitempty"`
}

// DolphinClusterSpecificationStatus defines the observed state of DolphinClusterSpecification
type DolphinClusterSpecificationStatus struct {
	WorkerNodePool                        []WorkerNodePool     `json:"workerNodePool,omitempty"`
	DolphinClusterPool                    []DolphinClusterPool `json:"dolphinClusterPool,omitempty"`
	KubesphereInstalledDolphinClusterPool []DolphinClusterPool `json:"kubesphereInstalledDolphinClusterPool,omitempty"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:scope=Cluster,shortName=dcs
//+kubebuilder:printcolumn:name="ClusterMaxWorkerNodeCount",type="number",JSONPath=".spec.clusterResource.cluster.maxWorkerNodeCount"
//+kubebuilder:printcolumn:name="MaxClusterCount",type="number",JSONPath=".spec.clusterResource.maxClusterCount"
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
