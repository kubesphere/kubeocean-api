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

	clusterapiv1beta1 "github.com/kubesphere/kubeocean-api/apis/clusterapi/v1beta1"
	versioned "github.com/kubesphere/kubeocean-api/client/clientset/versioned"
	internalinterfaces "github.com/kubesphere/kubeocean-api/client/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/kubesphere/kubeocean-api/client/listers/clusterapi/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterClassInformer provides access to a shared informer and lister for
// ClusterClasses.
type ClusterClassInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.ClusterClassLister
}

type clusterClassInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewClusterClassInformer constructs a new informer for ClusterClass type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterClassInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterClassInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredClusterClassInformer constructs a new informer for ClusterClass type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterClassInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ClusterapiV1beta1().ClusterClasses(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ClusterapiV1beta1().ClusterClasses(namespace).Watch(context.TODO(), options)
			},
		},
		&clusterapiv1beta1.ClusterClass{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterClassInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterClassInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterClassInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&clusterapiv1beta1.ClusterClass{}, f.defaultInformer)
}

func (f *clusterClassInformer) Lister() v1beta1.ClusterClassLister {
	return v1beta1.NewClusterClassLister(f.Informer().GetIndexer())
}
