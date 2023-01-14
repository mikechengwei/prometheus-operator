// Copyright The prometheus-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	apismonitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	v1 "k8s.io/api/core/v1"
)

// ThanosSpecApplyConfiguration represents an declarative configuration of the ThanosSpec type for use
// with apply.
type ThanosSpecApplyConfiguration struct {
	Image                   *string                      `json:"image,omitempty"`
	Version                 *string                      `json:"version,omitempty"`
	Tag                     *string                      `json:"tag,omitempty"`
	SHA                     *string                      `json:"sha,omitempty"`
	BaseImage               *string                      `json:"baseImage,omitempty"`
	Resources               *v1.ResourceRequirements     `json:"resources,omitempty"`
	ObjectStorageConfig     *v1.SecretKeySelector        `json:"objectStorageConfig,omitempty"`
	ObjectStorageConfigFile *string                      `json:"objectStorageConfigFile,omitempty"`
	ListenLocal             *bool                        `json:"listenLocal,omitempty"`
	GRPCListenLocal         *bool                        `json:"grpcListenLocal,omitempty"`
	HTTPListenLocal         *bool                        `json:"httpListenLocal,omitempty"`
	TracingConfig           *v1.SecretKeySelector        `json:"tracingConfig,omitempty"`
	TracingConfigFile       *string                      `json:"tracingConfigFile,omitempty"`
	GRPCServerTLSConfig     *TLSConfigApplyConfiguration `json:"grpcServerTlsConfig,omitempty"`
	LogLevel                *string                      `json:"logLevel,omitempty"`
	LogFormat               *string                      `json:"logFormat,omitempty"`
	MinTime                 *string                      `json:"minTime,omitempty"`
	ReadyTimeout            *apismonitoringv1.Duration   `json:"readyTimeout,omitempty"`
	VolumeMounts            []v1.VolumeMount             `json:"volumeMounts,omitempty"`
	AdditionalArgs          []ArgumentApplyConfiguration `json:"additionalArgs,omitempty"`
}

// ThanosSpecApplyConfiguration constructs an declarative configuration of the ThanosSpec type for use with
// apply.
func ThanosSpec() *ThanosSpecApplyConfiguration {
	return &ThanosSpecApplyConfiguration{}
}

// WithImage sets the Image field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Image field is set to the value of the last call.
func (b *ThanosSpecApplyConfiguration) WithImage(value string) *ThanosSpecApplyConfiguration {
	b.Image = &value
	return b
}

// WithVersion sets the Version field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Version field is set to the value of the last call.
func (b *ThanosSpecApplyConfiguration) WithVersion(value string) *ThanosSpecApplyConfiguration {
	b.Version = &value
	return b
}

// WithTag sets the Tag field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Tag field is set to the value of the last call.
func (b *ThanosSpecApplyConfiguration) WithTag(value string) *ThanosSpecApplyConfiguration {
	b.Tag = &value
	return b
}

// WithSHA sets the SHA field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SHA field is set to the value of the last call.
func (b *ThanosSpecApplyConfiguration) WithSHA(value string) *ThanosSpecApplyConfiguration {
	b.SHA = &value
	return b
}

// WithBaseImage sets the BaseImage field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BaseImage field is set to the value of the last call.
func (b *ThanosSpecApplyConfiguration) WithBaseImage(value string) *ThanosSpecApplyConfiguration {
	b.BaseImage = &value
	return b
}

// WithResources sets the Resources field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Resources field is set to the value of the last call.
func (b *ThanosSpecApplyConfiguration) WithResources(value v1.ResourceRequirements) *ThanosSpecApplyConfiguration {
	b.Resources = &value
	return b
}

// WithObjectStorageConfig sets the ObjectStorageConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ObjectStorageConfig field is set to the value of the last call.
func (b *ThanosSpecApplyConfiguration) WithObjectStorageConfig(value v1.SecretKeySelector) *ThanosSpecApplyConfiguration {
	b.ObjectStorageConfig = &value
	return b
}

// WithObjectStorageConfigFile sets the ObjectStorageConfigFile field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ObjectStorageConfigFile field is set to the value of the last call.
func (b *ThanosSpecApplyConfiguration) WithObjectStorageConfigFile(value string) *ThanosSpecApplyConfiguration {
	b.ObjectStorageConfigFile = &value
	return b
}

// WithListenLocal sets the ListenLocal field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ListenLocal field is set to the value of the last call.
func (b *ThanosSpecApplyConfiguration) WithListenLocal(value bool) *ThanosSpecApplyConfiguration {
	b.ListenLocal = &value
	return b
}

// WithGRPCListenLocal sets the GRPCListenLocal field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GRPCListenLocal field is set to the value of the last call.
func (b *ThanosSpecApplyConfiguration) WithGRPCListenLocal(value bool) *ThanosSpecApplyConfiguration {
	b.GRPCListenLocal = &value
	return b
}

// WithHTTPListenLocal sets the HTTPListenLocal field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HTTPListenLocal field is set to the value of the last call.
func (b *ThanosSpecApplyConfiguration) WithHTTPListenLocal(value bool) *ThanosSpecApplyConfiguration {
	b.HTTPListenLocal = &value
	return b
}

// WithTracingConfig sets the TracingConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TracingConfig field is set to the value of the last call.
func (b *ThanosSpecApplyConfiguration) WithTracingConfig(value v1.SecretKeySelector) *ThanosSpecApplyConfiguration {
	b.TracingConfig = &value
	return b
}

// WithTracingConfigFile sets the TracingConfigFile field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TracingConfigFile field is set to the value of the last call.
func (b *ThanosSpecApplyConfiguration) WithTracingConfigFile(value string) *ThanosSpecApplyConfiguration {
	b.TracingConfigFile = &value
	return b
}

// WithGRPCServerTLSConfig sets the GRPCServerTLSConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GRPCServerTLSConfig field is set to the value of the last call.
func (b *ThanosSpecApplyConfiguration) WithGRPCServerTLSConfig(value *TLSConfigApplyConfiguration) *ThanosSpecApplyConfiguration {
	b.GRPCServerTLSConfig = value
	return b
}

// WithLogLevel sets the LogLevel field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LogLevel field is set to the value of the last call.
func (b *ThanosSpecApplyConfiguration) WithLogLevel(value string) *ThanosSpecApplyConfiguration {
	b.LogLevel = &value
	return b
}

// WithLogFormat sets the LogFormat field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LogFormat field is set to the value of the last call.
func (b *ThanosSpecApplyConfiguration) WithLogFormat(value string) *ThanosSpecApplyConfiguration {
	b.LogFormat = &value
	return b
}

// WithMinTime sets the MinTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MinTime field is set to the value of the last call.
func (b *ThanosSpecApplyConfiguration) WithMinTime(value string) *ThanosSpecApplyConfiguration {
	b.MinTime = &value
	return b
}

// WithReadyTimeout sets the ReadyTimeout field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReadyTimeout field is set to the value of the last call.
func (b *ThanosSpecApplyConfiguration) WithReadyTimeout(value apismonitoringv1.Duration) *ThanosSpecApplyConfiguration {
	b.ReadyTimeout = &value
	return b
}

// WithVolumeMounts adds the given value to the VolumeMounts field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the VolumeMounts field.
func (b *ThanosSpecApplyConfiguration) WithVolumeMounts(values ...v1.VolumeMount) *ThanosSpecApplyConfiguration {
	for i := range values {
		b.VolumeMounts = append(b.VolumeMounts, values[i])
	}
	return b
}

// WithAdditionalArgs adds the given value to the AdditionalArgs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the AdditionalArgs field.
func (b *ThanosSpecApplyConfiguration) WithAdditionalArgs(values ...*ArgumentApplyConfiguration) *ThanosSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithAdditionalArgs")
		}
		b.AdditionalArgs = append(b.AdditionalArgs, *values[i])
	}
	return b
}
