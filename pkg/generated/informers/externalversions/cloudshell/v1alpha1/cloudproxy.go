/*
Copyright 2022.

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

	cloudshellv1alpha1 "github.com/cloudtty/cloudtty/pkg/apis/cloudshell/v1alpha1"
	versioned "github.com/cloudtty/cloudtty/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/cloudtty/cloudtty/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/cloudtty/cloudtty/pkg/generated/listers/cloudshell/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CloudProxyInformer provides access to a shared informer and lister for
// CloudProxies.
type CloudProxyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.CloudProxyLister
}

type cloudProxyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewCloudProxyInformer constructs a new informer for CloudProxy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCloudProxyInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCloudProxyInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredCloudProxyInformer constructs a new informer for CloudProxy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCloudProxyInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CloudshellV1alpha1().CloudProxies().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CloudshellV1alpha1().CloudProxies().Watch(context.TODO(), options)
			},
		},
		&cloudshellv1alpha1.CloudProxy{},
		resyncPeriod,
		indexers,
	)
}

func (f *cloudProxyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCloudProxyInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *cloudProxyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&cloudshellv1alpha1.CloudProxy{}, f.defaultInformer)
}

func (f *cloudProxyInformer) Lister() v1alpha1.CloudProxyLister {
	return v1alpha1.NewCloudProxyLister(f.Informer().GetIndexer())
}
