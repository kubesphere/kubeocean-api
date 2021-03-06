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

	clusterapiproviderqcv1beta1 "github.com/kubesphere/kubeocean-api/apis/clusterapiproviderqc/v1beta1"
	versioned "github.com/kubesphere/kubeocean-api/client/clientset/versioned"
	internalinterfaces "github.com/kubesphere/kubeocean-api/client/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/kubesphere/kubeocean-api/client/listers/clusterapiproviderqc/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// QCClusterInformer provides access to a shared informer and lister for
// QCClusters.
type QCClusterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.QCClusterLister
}

type qCClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewQCClusterInformer constructs a new informer for QCCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewQCClusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredQCClusterInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredQCClusterInformer constructs a new informer for QCCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredQCClusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ClusterapiproviderqcV1beta1().QCClusters(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ClusterapiproviderqcV1beta1().QCClusters(namespace).Watch(context.TODO(), options)
			},
		},
		&clusterapiproviderqcv1beta1.QCCluster{},
		resyncPeriod,
		indexers,
	)
}

func (f *qCClusterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredQCClusterInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *qCClusterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&clusterapiproviderqcv1beta1.QCCluster{}, f.defaultInformer)
}

func (f *qCClusterInformer) Lister() v1beta1.QCClusterLister {
	return v1beta1.NewQCClusterLister(f.Informer().GetIndexer())
}
