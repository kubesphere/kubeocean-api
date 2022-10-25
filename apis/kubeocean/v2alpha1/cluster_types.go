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
	"github.com/kubesphere/kubeocean-api/v2/constant/qingcloud"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type SpecificationID string

type OsImageID string

type InstanceSpec struct {
	InstanceType qingcloud.InstanceType `json:"instanceType,omitempty"`
	Cpu          int32                  `json:"cpu,omitempty"`    // in cores
	Memory       int32                  `json:"memory,omitempty"` // in MB
}

type NodeSpec struct {
	DiskSize          int32  `json:"diskSize,omitempty"`
	KubernetesVersion string `json:"kubernetesVersion,omitempty"`
	KubesphereVersion string `json:"kubesphereVersion,omitempty"`
	KubespherePlugins string `json:"kubespherePlugins,omitempty"` // minimal/full/etc.
	Masters           int32  `json:"masters,omitempty"`
	Workers           int32  `json:"workers,omitempty"`
}

type OsImageSpec struct {
	OsImageID OsImageID `json:"osImageID,omitempty"`
	NodeSpec  `json:",inline"`
}

type ProvisionMethod string

const (
	// ProvisionMethodDynamic means dynamically create the cluster
	ProvisionMethodDynamic ProvisionMethod = "dynamic"

	// ProvisionMethodStatic means boot up cluster from os image, this is the default
	ProvisionMethodStatic ProvisionMethod = "static"
)

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	SpecificationID SpecificationID `json:"specificationID,omitempty"`
	SshKeyID        string          `json:"sshKeyID,omitempty"`
	OsImage         OsImageSpec     `json:"osImage,omitempty"`
	Instance        InstanceSpec    `json:"instance,omitempty"`
	ProvisionMethod ProvisionMethod `json:"provisionMethod,omitempty"`
	NodeSpec        `json:",inline"`
}

type ClusterPhase string

const (
	ClusterPhaseCreating   ClusterPhase = "creating"   // cluster is being created
	ClusterPhaseAssigned   ClusterPhase = "assigned"   // cluster is ready and has been assigned to user
	ClusterPhaseUnassigned ClusterPhase = "unassigned" // cluster is ready and has not been assigned
	ClusterPhaseDeleting   ClusterPhase = "deleting"   // cluster is being deleted
)

type KubespherePhase string

const (
	KubespherePhaseInstalling KubespherePhase = "installing"
	KubespherePhaseInstalled  KubespherePhase = "installed"
)

type KubernetesPhase string

const (
	KubernetesPhaseInstalling KubernetesPhase = "installing"
	KubernetesPhaseInstalled  KubernetesPhase = "installed"
)

// ClusterStatus defines the observed state of Cluster
type ClusterStatus struct {
	Phase                    ClusterPhase    `json:"phase,omitempty"`
	Kubeconfig               string          `json:"kubeconfig,omitempty"`
	ApiServerAddress         string          `json:"apiServerAddress,omitempty"`
	KubesphereConsoleAddress string          `json:"kubesphereConsoleAddress,omitempty"`
	KubespherePhase          KubespherePhase `json:"kubespherePhase,omitempty"`
	KubernetesPhase          KubernetesPhase `json:"kubernetesPhase,omitempty"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:printcolumn:JSONPath=".spec.instance.instanceType",name=InstanceType,type=string
// +kubebuilder:printcolumn:JSONPath=".spec.osImage.osImageID",name=OSImageID,type=string
// +kubebuilder:printcolumn:JSONPath=".spec.specificationID",name=SpecificationID,type=string
// +kubebuilder:printcolumn:JSONPath=".status.phase",name=Phase,type=string
// +kubebuilder:printcolumn:JSONPath=".status.kubespherePhase",name=KubespherePhase,type=string
// +kubebuilder:printcolumn:JSONPath=".status.kubernetesPhase",name=KubernetesPhase,type=string
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// Cluster is the Schema for the clusters API
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterSpec   `json:"spec,omitempty"`
	Status ClusterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterList contains a list of Cluster
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}
