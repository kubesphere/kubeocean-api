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

// DolphinServiceInformer provides access to a shared informer and lister for
// DolphinServices.
type DolphinServiceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.DolphinServiceLister
}

type dolphinServiceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewDolphinServiceInformer constructs a new informer for DolphinService type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDolphinServiceInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDolphinServiceInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredDolphinServiceInformer constructs a new informer for DolphinService type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDolphinServiceInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeoceanV1alpha1().DolphinServices().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeoceanV1alpha1().DolphinServices().Watch(context.TODO(), options)
			},
		},
		&kubeoceanv1alpha1.DolphinService{},
		resyncPeriod,
		indexers,
	)
}

func (f *dolphinServiceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDolphinServiceInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *dolphinServiceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&kubeoceanv1alpha1.DolphinService{}, f.defaultInformer)
}

func (f *dolphinServiceInformer) Lister() v1alpha1.DolphinServiceLister {
	return v1alpha1.NewDolphinServiceLister(f.Informer().GetIndexer())
}
