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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfiguration

import (
	v1 "github.com/istio-ecosystem/admiral-api/pkg/apis/admiral/v1"
	admiralv1 "github.com/istio-ecosystem/admiral-api/pkg/client/applyconfiguration/admiral/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=admiral.io, Version=v1
	case v1.SchemeGroupVersion.WithKind("AdaptiveConcurrency"):
		return &admiralv1.AdaptiveConcurrencyApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("AppOverride"):
		return &admiralv1.AppOverrideApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("AppQuotaGroup"):
		return &admiralv1.AppQuotaGroupApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("AssetDB"):
		return &admiralv1.AssetDBApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("AssetDBSpec"):
		return &admiralv1.AssetDBSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ClusterItem"):
		return &admiralv1.ClusterItemApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ClusterShards"):
		return &admiralv1.ClusterShardsApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Config"):
		return &admiralv1.ConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Dependency"):
		return &admiralv1.DependencyApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("DependencySpec"):
		return &admiralv1.DependencySpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("DependencyStatus"):
		return &admiralv1.DependencyStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("DynamicRouting"):
		return &admiralv1.DynamicRoutingApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("EdgeService"):
		return &admiralv1.EdgeServiceApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("FailedCluster"):
		return &admiralv1.FailedClusterApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("FailedIdentity"):
		return &admiralv1.FailedIdentityApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("FailureDetails"):
		return &admiralv1.FailureDetailsApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Filter"):
		return &admiralv1.FilterApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Header"):
		return &admiralv1.HeaderApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IdentityItem"):
		return &admiralv1.IdentityItemApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Quota"):
		return &admiralv1.QuotaApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("QuotaGroup"):
		return &admiralv1.QuotaGroupApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Retry"):
		return &admiralv1.RetryApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Route"):
		return &admiralv1.RouteApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Shard"):
		return &admiralv1.ShardApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ShardSpec"):
		return &admiralv1.ShardSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ShardStatus"):
		return &admiralv1.ShardStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ShardStatusCondition"):
		return &admiralv1.ShardStatusConditionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Target"):
		return &admiralv1.TargetApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("TargetGroup"):
		return &admiralv1.TargetGroupApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("TotalQuotaGroup"):
		return &admiralv1.TotalQuotaGroupApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("TrafficConfig"):
		return &admiralv1.TrafficConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("TrafficConfigSpec"):
		return &admiralv1.TrafficConfigSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("TrafficConfigStatus"):
		return &admiralv1.TrafficConfigStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Weight"):
		return &admiralv1.WeightApplyConfiguration{}

	}
	return nil
}
