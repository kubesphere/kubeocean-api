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
// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"

	v1beta1 "github.com/kubesphere/kubeocean-api/apis/bootstrap/v1beta1"
	clusterapiv1beta1 "github.com/kubesphere/kubeocean-api/apis/clusterapi/v1beta1"
	clusterapiproviderqcv1beta1 "github.com/kubesphere/kubeocean-api/apis/clusterapiproviderqc/v1beta1"
	controlplanev1beta1 "github.com/kubesphere/kubeocean-api/apis/controlplane/v1beta1"
	v1alpha1 "github.com/kubesphere/kubeocean-api/apis/kubeocean/v1alpha1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=bootstrap, Version=v1beta1
	case v1beta1.SchemeGroupVersion.WithResource("kubeadmconfigs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Bootstrap().V1beta1().KubeadmConfigs().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("kubeadmconfigtemplates"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Bootstrap().V1beta1().KubeadmConfigTemplates().Informer()}, nil

		// Group=clusterapi, Version=v1beta1
	case clusterapiv1beta1.SchemeGroupVersion.WithResource("clusters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Clusterapi().V1beta1().Clusters().Informer()}, nil
	case clusterapiv1beta1.SchemeGroupVersion.WithResource("clusterclasses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Clusterapi().V1beta1().ClusterClasses().Informer()}, nil
	case clusterapiv1beta1.SchemeGroupVersion.WithResource("machines"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Clusterapi().V1beta1().Machines().Informer()}, nil
	case clusterapiv1beta1.SchemeGroupVersion.WithResource("machinedeployments"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Clusterapi().V1beta1().MachineDeployments().Informer()}, nil
	case clusterapiv1beta1.SchemeGroupVersion.WithResource("machinehealthchecks"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Clusterapi().V1beta1().MachineHealthChecks().Informer()}, nil
	case clusterapiv1beta1.SchemeGroupVersion.WithResource("machinesets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Clusterapi().V1beta1().MachineSets().Informer()}, nil

		// Group=clusterapiproviderqc, Version=v1beta1
	case clusterapiproviderqcv1beta1.SchemeGroupVersion.WithResource("qcclusters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Clusterapiproviderqc().V1beta1().QCClusters().Informer()}, nil
	case clusterapiproviderqcv1beta1.SchemeGroupVersion.WithResource("qcclustertemplates"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Clusterapiproviderqc().V1beta1().QCClusterTemplates().Informer()}, nil
	case clusterapiproviderqcv1beta1.SchemeGroupVersion.WithResource("qcmachines"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Clusterapiproviderqc().V1beta1().QCMachines().Informer()}, nil
	case clusterapiproviderqcv1beta1.SchemeGroupVersion.WithResource("qcmachinetemplates"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Clusterapiproviderqc().V1beta1().QCMachineTemplates().Informer()}, nil

		// Group=controlplane, Version=v1beta1
	case controlplanev1beta1.SchemeGroupVersion.WithResource("kubeadmcontrolplanes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Controlplane().V1beta1().KubeadmControlPlanes().Informer()}, nil
	case controlplanev1beta1.SchemeGroupVersion.WithResource("kubeadmcontrolplanetemplates"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Controlplane().V1beta1().KubeadmControlPlaneTemplates().Informer()}, nil

		// Group=kubeocean.kubesphere.io, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("dolphinclusters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubeocean().V1alpha1().DolphinClusters().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("dolphinclusterspecifications"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubeocean().V1alpha1().DolphinClusterSpecifications().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("dolphinservices"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubeocean().V1alpha1().DolphinServices().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
