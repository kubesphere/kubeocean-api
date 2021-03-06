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

package v1beta1

import (
	internalinterfaces "github.com/kubesphere/kubeocean-api/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// QCClusters returns a QCClusterInformer.
	QCClusters() QCClusterInformer
	// QCClusterTemplates returns a QCClusterTemplateInformer.
	QCClusterTemplates() QCClusterTemplateInformer
	// QCMachines returns a QCMachineInformer.
	QCMachines() QCMachineInformer
	// QCMachineTemplates returns a QCMachineTemplateInformer.
	QCMachineTemplates() QCMachineTemplateInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// QCClusters returns a QCClusterInformer.
func (v *version) QCClusters() QCClusterInformer {
	return &qCClusterInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// QCClusterTemplates returns a QCClusterTemplateInformer.
func (v *version) QCClusterTemplates() QCClusterTemplateInformer {
	return &qCClusterTemplateInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// QCMachines returns a QCMachineInformer.
func (v *version) QCMachines() QCMachineInformer {
	return &qCMachineInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// QCMachineTemplates returns a QCMachineTemplateInformer.
func (v *version) QCMachineTemplates() QCMachineTemplateInformer {
	return &qCMachineTemplateInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
