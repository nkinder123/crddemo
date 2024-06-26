/*
Copyright 2024 The Kubernetes authors.

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

package v1

import (
	"context"
	crddemov1 "crddemo/pkg/apis/crddemo/v1"
	versioned "crddemo/pkg/generated/clientset/versioned"
	internalinterfaces "crddemo/pkg/generated/informers/externalversions/internalinterfaces"
	v1 "crddemo/pkg/generated/listers/crddemo/v1"
	time "time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// MydemoInformer provides access to a shared informer and lister for
// Mydemos.
type MydemoInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.MydemoLister
}

type mydemoInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewMydemoInformer constructs a new informer for Mydemo type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMydemoInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMydemoInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredMydemoInformer constructs a new informer for Mydemo type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredMydemoInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CrddemoV1().Mydemos(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CrddemoV1().Mydemos(namespace).Watch(context.TODO(), options)
			},
		},
		&crddemov1.Mydemo{},
		resyncPeriod,
		indexers,
	)
}

func (f *mydemoInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMydemoInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *mydemoInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&crddemov1.Mydemo{}, f.defaultInformer)
}

func (f *mydemoInformer) Lister() v1.MydemoLister {
	return v1.NewMydemoLister(f.Informer().GetIndexer())
}
