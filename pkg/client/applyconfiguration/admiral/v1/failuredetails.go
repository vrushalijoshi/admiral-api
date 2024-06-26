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

package v1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// FailureDetailsApplyConfiguration represents an declarative configuration of the FailureDetails type for use
// with apply.
type FailureDetailsApplyConfiguration struct {
	LastUpdatedTime *v1.Time                          `json:"lastUpdateTime,omitempty"`
	FailedClusters  []FailedClusterApplyConfiguration `json:"clusters,omitempty"`
}

// FailureDetailsApplyConfiguration constructs an declarative configuration of the FailureDetails type for use with
// apply.
func FailureDetails() *FailureDetailsApplyConfiguration {
	return &FailureDetailsApplyConfiguration{}
}

// WithLastUpdatedTime sets the LastUpdatedTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastUpdatedTime field is set to the value of the last call.
func (b *FailureDetailsApplyConfiguration) WithLastUpdatedTime(value v1.Time) *FailureDetailsApplyConfiguration {
	b.LastUpdatedTime = &value
	return b
}

// WithFailedClusters adds the given value to the FailedClusters field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the FailedClusters field.
func (b *FailureDetailsApplyConfiguration) WithFailedClusters(values ...*FailedClusterApplyConfiguration) *FailureDetailsApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithFailedClusters")
		}
		b.FailedClusters = append(b.FailedClusters, *values[i])
	}
	return b
}
