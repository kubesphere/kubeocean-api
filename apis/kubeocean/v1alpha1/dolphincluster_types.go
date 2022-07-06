/*
Copyright 2022.

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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DolphinClusterSpec defines the desired state of DolphinCluster
type DolphinClusterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	VirtualCluster VirtualCluster `json:"virtualCluster,omitempty"`
	ExternalDomain string         `json:"externalDomain,omitempty"`
}

type VirtualClusterNode struct {
	Name          string `json:"name,omitempty"`
	CpuCurrent    *int32 `json:"cpuCurrent,omitempty"`
	MemoryCurrent *int32 `json:"memoryCurrent,omitempty"`
	DiskSize      *int32 `json:"diskSize,omitempty"`
}

type VirtualClusterInitNodeSpecification struct {
	VirtualClusterNode `json:",inline"`
	NodeName           string `json:"nodeName,omitempty"`
}

type VirtualClusterNodeSpecification struct {
	VirtualClusterNode `json:",inline"`
	Replicas           *int32 `json:"replicas,omitempty"`
}

type Kubesphere struct {
	Enable  bool   `json:"enable,omitempty"`
	Version string `json:"version,omitempty"`
}

type VirtualCluster struct {
	K3s            ImageRepoVersion                    `json:"k3s,omitempty"`
	Vcluster       ImageRepoVersion                    `json:"vcluster,omitempty"`
	Plugins        ImageRepoVersion                    `json:"plugins,omitempty"`
	InitNode       VirtualClusterInitNodeSpecification `json:"initNode,omitempty"`
	Specifications []VirtualClusterNodeSpecification   `json:"specifications,omitempty"`
	Kubesphere     Kubesphere                          `json:"kubesphere,omitempty"`
}

type ImageRepoVersion struct {
	ImageRepo string   `json:"imageRepo,omitempty"`
	Version   string   `json:"version,omitempty"`
	Args      []string `json:"args,omitempty"`
}

type VirtualClusterNodeSpecificationStatus struct {
	VirtualClusterNode `json:",inline"`
	Nodes              []string `json:"nodes,omitempty"`
}

type DolphinClusterKubespherePhase string

const (
	DolphinClusterKubesphereInstalling    DolphinClusterKubespherePhase = "installing"
	DolphinClusterKubesphereInstalled     DolphinClusterKubespherePhase = "installed"
	DolphinClusterKubesphereInstallFailed DolphinClusterKubespherePhase = "failed"
)

type KubesphereStatus struct {
	Enable            bool                          `json:"enable,omitempty"`
	Phase             DolphinClusterKubespherePhase `json:"phase,omitempty"`
	KubesphereVersion string                        `json:"version,omitempty"`
	KsConsoleAddress  string                        `json:"ksConsoleAddress,omitempty"`
}

type DolphinClusterPhase string

const (
	DolphinClusterCreating DolphinClusterPhase = "createing"
	DolphinClusterRunning  DolphinClusterPhase = "running"
	DolphinClusterFailed   DolphinClusterPhase = "failed"
	DolphinClusterDeleting DolphinClusterPhase = "deleting"
)

// DolphinClusterStatus defines the observed state of DolphinCluster
type DolphinClusterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Phase                DolphinClusterPhase                     `json:"phase,omitempty"`
	InClusterKubeconfig  string                                  `json:"inClusterKubeconfig,omitempty"`
	OutClusterKubeconfig string                                  `json:"outClusterKubeconfig,omitempty"`
	InitNode             VirtualClusterInitNodeSpecification     `json:"initNode,omitempty"`
	Specifications       []VirtualClusterNodeSpecificationStatus `json:"specifications,omitempty"`
	Kubesphere           KubesphereStatus                        `json:"kubesphere,omitempty"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:scope=Cluster,shortName=dc
//+kubebuilder:printcolumn:name="InitNode",type="string",JSONPath=".status.initNode.nodeName",description="The virtual cluster init node name"
//+kubebuilder:printcolumn:name="KubesphereEnabled",type="boolean",JSONPath=".status.kubesphere.enable",description="Whether install kubesphere in the virtual cluster enable"
//+kubebuilder:printcolumn:name="KubespherePhase",type="string",JSONPath=".status.kubesphere.phase",description="Kubesphere phase in the virtual cluster"
//+kubebuilder:printcolumn:name="Phase",type="string",JSONPath=".status.phase",description="The virtual cluster running phase"
//+kubebuilder:printcolumn:name="ExternalDomain",type="string",JSONPath=".spec.externalDomain",description="The external domain for virtual cluster"
//+kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// DolphinCluster is the Schema for the dolphinclusters API
type DolphinCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DolphinClusterSpec   `json:"spec,omitempty"`
	Status DolphinClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//+kubebuilder:object:root=true

// DolphinClusterList contains a list of DolphinCluster
type DolphinClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DolphinCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DolphinCluster{}, &DolphinClusterList{})
}
