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

type OsImageID string

type Specification struct {
	OsImageID             OsImageID              `json:"osImageID"`
	InstanceType          qingcloud.InstanceType `json:"instanceType"`
	WaitToAssignThreshold int32                  `json:"waitToAssignThreshold"`

	// Below fields are determined by OsImageID
	SshKey            string `json:"sshKey"`
	DiskSize          int32  `json:"diskSize"`
	KubernetesVersion string `json:"kubernetesVersion"`
	KubesphereVersion string `json:"kubesphereVersion"`
	KubespherePlugins string `json:"kubespherePlugins"` // minimal/full/etc.
	Masters           int32  `json:"masters"`
	Workers           int32  `json:"workers"`

	// Below fields are determined by InstanceType
	CpuCurrent    int32 `json:"cpuCurrent"`
	MemoryCurrent int32 `json:"memoryCurrent"`
}

type SpecificationID string

// KindClusterPoolSpec defines the desired state of KindClusterPool
type KindClusterPoolSpec struct {
	Specifications                   map[SpecificationID]Specification `json:"specifications"`
	TargetClusterMaxKindClusterCount int32                             `json:"targetClusterMaxKindClusterCount"`
}

type KindClusterID string
type SpecificationStatus map[KindClusterPhase][]KindClusterID

// KindClusterPoolStatus defines the observed state of KindClusterPool
type KindClusterPoolStatus struct {
	SpecificationStatuses map[SpecificationID]SpecificationStatus `json:"specificationStatuses"`
	CountByPhase          map[KindClusterPhase]int32              `json:"countByPhase"`
	TotalCount            int32                                   `json:"totalCount"`
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
	metav1.ObjectMeta `json:"metadata"`

	Spec   KindClusterPoolSpec   `json:"spec"`
	Status KindClusterPoolStatus `json:"status"`
}

//+kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KindClusterPoolList contains a list of KindClusterPool
type KindClusterPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []KindClusterPool `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KindClusterPool{}, &KindClusterPoolList{})
}
