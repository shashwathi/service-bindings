/*
Copyright 2020 VMware, Inc.
SPDX-License-Identifier: Apache-2.0
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha2

import (
	time "time"

	serviceinternalv1alpha2 "github.com/vmware-labs/service-bindings/pkg/apis/serviceinternal/v1alpha2"
	versioned "github.com/vmware-labs/service-bindings/pkg/client/clientset/versioned"
	internalinterfaces "github.com/vmware-labs/service-bindings/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha2 "github.com/vmware-labs/service-bindings/pkg/client/listers/serviceinternal/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ServiceBindingProjectionInformer provides access to a shared informer and lister for
// ServiceBindingProjections.
type ServiceBindingProjectionInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha2.ServiceBindingProjectionLister
}

type serviceBindingProjectionInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewServiceBindingProjectionInformer constructs a new informer for ServiceBindingProjection type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewServiceBindingProjectionInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredServiceBindingProjectionInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredServiceBindingProjectionInformer constructs a new informer for ServiceBindingProjection type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredServiceBindingProjectionInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.InternalV1alpha2().ServiceBindingProjections(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.InternalV1alpha2().ServiceBindingProjections(namespace).Watch(options)
			},
		},
		&serviceinternalv1alpha2.ServiceBindingProjection{},
		resyncPeriod,
		indexers,
	)
}

func (f *serviceBindingProjectionInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredServiceBindingProjectionInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *serviceBindingProjectionInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&serviceinternalv1alpha2.ServiceBindingProjection{}, f.defaultInformer)
}

func (f *serviceBindingProjectionInformer) Lister() v1alpha2.ServiceBindingProjectionLister {
	return v1alpha2.NewServiceBindingProjectionLister(f.Informer().GetIndexer())
}
