// Copyright 2020 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/kubeflow/tf-operator/pkg/apis/tensorflow/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TFJobLister helps list TFJobs.
type TFJobLister interface {
	// List lists all TFJobs in the indexer.
	List(selector labels.Selector) (ret []*v1.TFJob, err error)
	// TFJobs returns an object that can list and get TFJobs.
	TFJobs(namespace string) TFJobNamespaceLister
	TFJobListerExpansion
}

// tFJobLister implements the TFJobLister interface.
type tFJobLister struct {
	indexer cache.Indexer
}

// NewTFJobLister returns a new TFJobLister.
func NewTFJobLister(indexer cache.Indexer) TFJobLister {
	return &tFJobLister{indexer: indexer}
}

// List lists all TFJobs in the indexer.
func (s *tFJobLister) List(selector labels.Selector) (ret []*v1.TFJob, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.TFJob))
	})
	return ret, err
}

// TFJobs returns an object that can list and get TFJobs.
func (s *tFJobLister) TFJobs(namespace string) TFJobNamespaceLister {
	return tFJobNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TFJobNamespaceLister helps list and get TFJobs.
type TFJobNamespaceLister interface {
	// List lists all TFJobs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.TFJob, err error)
	// Get retrieves the TFJob from the indexer for a given namespace and name.
	Get(name string) (*v1.TFJob, error)
	TFJobNamespaceListerExpansion
}

// tFJobNamespaceLister implements the TFJobNamespaceLister
// interface.
type tFJobNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TFJobs in the indexer for a given namespace.
func (s tFJobNamespaceLister) List(selector labels.Selector) (ret []*v1.TFJob, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.TFJob))
	})
	return ret, err
}

// Get retrieves the TFJob from the indexer for a given namespace and name.
func (s tFJobNamespaceLister) Get(name string) (*v1.TFJob, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("tfjob"), name)
	}
	return obj.(*v1.TFJob), nil
}
