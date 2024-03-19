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

// ShardInformer provides access to a shared informer and lister for
// Shards.
type ShardInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ShardLister
}

type shardInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewShardInformer constructs a new informer for Shard type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewShardInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredShardInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredShardInformer constructs a new informer for Shard type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredShardInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AdmiralV1().Shards(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AdmiralV1().Shards(namespace).Watch(context.TODO(), options)
			},
		},
		&admiralv1.Shard{},
		resyncPeriod,
		indexers,
	)
}

func (f *shardInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredShardInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *shardInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&admiralv1.Shard{}, f.defaultInformer)
}

func (f *shardInformer) Lister() v1.ShardLister {
	return v1.NewShardLister(f.Informer().GetIndexer())
}