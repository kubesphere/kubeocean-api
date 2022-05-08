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

package v1alpha1

import (
	"context"
	time "time"

	kubeoceanv1alpha1 "github.com/kubesphere/kubeocean-api/apis/kubeocean/v1alpha1"
	versioned "github.com/kubesphere/kubeocean-api/client/clientset/versioned"
	internalinterfaces "github.com/kubesphere/kubeocean-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/kubesphere/kubeocean-api/client/listers/kubeocean/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// DolphinClusterInformer provides access to a shared informer and lister for
// DolphinClusters.
type DolphinClusterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.DolphinClusterLister
}

type dolphinClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewDolphinClusterInformer constructs a new informer for DolphinCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDolphinClusterInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDolphinClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredDolphinClusterInformer constructs a new informer for DolphinCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDolphinClusterInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeoceanV1alpha1().DolphinClusters().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeoceanV1alpha1().DolphinClusters().Watch(context.TODO(), options)
			},
		},
		&kubeoceanv1alpha1.DolphinCluster{},
		resyncPeriod,
		indexers,
	)
}

func (f *dolphinClusterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDolphinClusterInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *dolphinClusterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&kubeoceanv1alpha1.DolphinCluster{}, f.defaultInformer)
}

func (f *dolphinClusterInformer) Lister() v1alpha1.DolphinClusterLister {
	return v1alpha1.NewDolphinClusterLister(f.Informer().GetIndexer())
}