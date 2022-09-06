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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// KindClusterSpec defines the desired state of KindCluster
type KindClusterSpec struct {
	Plugins        SpecificationPluginType `json:"plugins,omitempty"`
	MachineOsImage OsImageType             `json:"machineOsImage,omitempty"`
	InstanceType   InstanceTypeType        `json:"instanceType,omitempty"`
	SshKey         string                  `json:"sshKey,omitempty"`
}

type KindClusterPhase string

const (
	KindClusterCreatingPhase KindClusterPhase = "creating"
	KindClusterCreatedPhase  KindClusterPhase = "created"
)

// KindClusterStatus defines the observed state of KindCluster
type KindClusterStatus struct {
	Phase      KindClusterPhase `json:"phase,omitempty"`
	Kubeconfig string           `json:"kubeconfig,omitempty"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:scope=Cluster,shortName=kc
//+kubebuilder:printcolumn:JSONPath=".spec.instanceType",name=InstanceType,type=string
//+kubebuilder:printcolumn:JSONPath=".spec.plugins",name=Plugins,type=string
//+kubebuilder:printcolumn:JSONPath=".spec.machineOsImage",name=OsImage,type=string
//+kubebuilder:printcolumn:JSONPath=".status.phase",name=Phase,type=string
//+kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// KindCluster is the Schema for the kindclusters API
type KindCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KindClusterSpec   `json:"spec,omitempty"`
	Status KindClusterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KindClusterList contains a list of KindCluster
type KindClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KindCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KindCluster{}, &KindClusterList{})
}
