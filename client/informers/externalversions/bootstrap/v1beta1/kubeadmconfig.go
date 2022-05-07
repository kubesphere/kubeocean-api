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
	"context"
	time "time"

	bootstrapv1beta1 "github.com/kubesphere/kubeocean-api/apis/bootstrap/v1beta1"
	versioned "github.com/kubesphere/kubeocean-api/client/clientset/versioned"
	internalinterfaces "github.com/kubesphere/kubeocean-api/client/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/kubesphere/kubeocean-api/client/listers/bootstrap/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// KubeadmConfigInformer provides access to a shared informer and lister for
// KubeadmConfigs.
type KubeadmConfigInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.KubeadmConfigLister
}

type kubeadmConfigInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewKubeadmConfigInformer constructs a new informer for KubeadmConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewKubeadmConfigInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredKubeadmConfigInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredKubeadmConfigInformer constructs a new informer for KubeadmConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredKubeadmConfigInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BootstrapV1beta1().KubeadmConfigs(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BootstrapV1beta1().KubeadmConfigs(namespace).Watch(context.TODO(), options)
			},
		},
		&bootstrapv1beta1.KubeadmConfig{},
		resyncPeriod,
		indexers,
	)
}

func (f *kubeadmConfigInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredKubeadmConfigInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *kubeadmConfigInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&bootstrapv1beta1.KubeadmConfig{}, f.defaultInformer)
}

func (f *kubeadmConfigInformer) Lister() v1beta1.KubeadmConfigLister {
	return v1beta1.NewKubeadmConfigLister(f.Informer().GetIndexer())
}
