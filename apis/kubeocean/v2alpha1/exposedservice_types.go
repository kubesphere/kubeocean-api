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

package v2alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type ServiceProtocol string

const (
	ServiceProtocolHttp  ServiceProtocol = "http"
	ServiceProtocolHttps ServiceProtocol = "https"
)

type Service struct {
	Namespace string          `json:"namespace,omitempty"`
	Name      string          `json:"name,omitempty"`
	Port      int32           `json:"port,omitempty"`
	Protocol  ServiceProtocol `json:"protocol,omitempty"`
}

// ExposedServiceSpec defines the desired state of ExposedService
type ExposedServiceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Cluster string  `json:"cluster,omitempty"`
	Service Service `json:"service,omitempty"`

	// ingress class name
	// +optional
	IngressClass *string `json:"ingressClass,omitempty"`

	// ingress path, if not set, will use "/"
	// +optional
	Path *string `json:"path,omitempty"`
}

// ExposedServiceStatus defines the observed state of ExposedService
type ExposedServiceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	VisitAddr             string `json:"visitAddr,omitempty"`
	IngressName           string `json:"ingressName,omitempty"`
	InnerServiceClusterIP string `json:"InnerServiceClusterIP,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:scope=Namespaced,shortName=es
//+kubebuilder:printcolumn:name="Cluster",type="string",JSONPath=".spec.cluster",description="The target cluster"
//+kubebuilder:printcolumn:name="Namespace",type="string",JSONPath=".spec.service.namespace",description="The target service namespace"
//+kubebuilder:printcolumn:name="Name",type="string",JSONPath=".spec.service.name",description="The target service name"
//+kubebuilder:printcolumn:name="VisitAddr",type="string",JSONPath=".status.visitAddr",description="The target service visit address"
//+kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// ExposedService is the Schema for the exposedservices API
type ExposedService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ExposedServiceSpec   `json:"spec,omitempty"`
	Status ExposedServiceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//+kubebuilder:object:root=true

// ExposedServiceList contains a list of ExposedService
type ExposedServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExposedService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ExposedService{}, &ExposedServiceList{})
}
