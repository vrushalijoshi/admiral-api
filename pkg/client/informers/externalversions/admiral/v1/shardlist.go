/*
Copyright The Kubernetes Authors.

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
	time "time"

	admiralv1 "github.com/istio-ecosystem/admiral-api/pkg/apis/admiral/v1"
	versioned "github.com/istio-ecosystem/admiral-api/pkg/client/clientset/versioned"
	internalinterfaces "github.com/istio-ecosystem/admiral-api/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/istio-ecosystem/admiral-api/pkg/client/listers/admiral/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ShardListInformer provides access to a shared informer and lister for
// ShardLists.
type ShardListInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ShardListLister
}

type shardListInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewShardListInformer constructs a new informer for ShardList type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewShardListInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredShardListInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredShardListInformer constructs a new informer for ShardList type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredShardListInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AdmiralV1().ShardLists(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AdmiralV1().ShardLists(namespace).Watch(context.TODO(), options)
			},
		},
		&admiralv1.ShardList{},
		resyncPeriod,
		indexers,
	)
}

func (f *shardListInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredShardListInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *shardListInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&admiralv1.ShardList{}, f.defaultInformer)
}

func (f *shardListInformer) Lister() v1.ShardListLister {
	return v1.NewShardListLister(f.Informer().GetIndexer())
}
