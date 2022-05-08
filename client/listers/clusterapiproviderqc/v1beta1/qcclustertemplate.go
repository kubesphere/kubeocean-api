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
// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/kubesphere/kubeocean-api/apis/clusterapiproviderqc/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// QCClusterTemplateLister helps list QCClusterTemplates.
// All objects returned here must be treated as read-only.
type QCClusterTemplateLister interface {
	// List lists all QCClusterTemplates in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.QCClusterTemplate, err error)
	// QCClusterTemplates returns an object that can list and get QCClusterTemplates.
	QCClusterTemplates(namespace string) QCClusterTemplateNamespaceLister
	QCClusterTemplateListerExpansion
}

// qCClusterTemplateLister implements the QCClusterTemplateLister interface.
type qCClusterTemplateLister struct {
	indexer cache.Indexer
}

// NewQCClusterTemplateLister returns a new QCClusterTemplateLister.
func NewQCClusterTemplateLister(indexer cache.Indexer) QCClusterTemplateLister {
	return &qCClusterTemplateLister{indexer: indexer}
}

// List lists all QCClusterTemplates in the indexer.
func (s *qCClusterTemplateLister) List(selector labels.Selector) (ret []*v1beta1.QCClusterTemplate, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.QCClusterTemplate))
	})
	return ret, err
}

// QCClusterTemplates returns an object that can list and get QCClusterTemplates.
func (s *qCClusterTemplateLister) QCClusterTemplates(namespace string) QCClusterTemplateNamespaceLister {
	return qCClusterTemplateNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// QCClusterTemplateNamespaceLister helps list and get QCClusterTemplates.
// All objects returned here must be treated as read-only.
type QCClusterTemplateNamespaceLister interface {
	// List lists all QCClusterTemplates in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.QCClusterTemplate, err error)
	// Get retrieves the QCClusterTemplate from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.QCClusterTemplate, error)
	QCClusterTemplateNamespaceListerExpansion
}

// qCClusterTemplateNamespaceLister implements the QCClusterTemplateNamespaceLister
// interface.
type qCClusterTemplateNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all QCClusterTemplates in the indexer for a given namespace.
func (s qCClusterTemplateNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.QCClusterTemplate, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.QCClusterTemplate))
	})
	return ret, err
}

// Get retrieves the QCClusterTemplate from the indexer for a given namespace and name.
func (s qCClusterTemplateNamespaceLister) Get(name string) (*v1beta1.QCClusterTemplate, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("qcclustertemplate"), name)
	}
	return obj.(*v1beta1.QCClusterTemplate), nil
}