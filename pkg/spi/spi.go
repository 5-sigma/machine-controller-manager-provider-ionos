/*
Copyright (c) 2021 SAP SE or an SAP affiliate company. All rights reserved.

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

package spi

// SessionProviderInterface provides an interface to deal with cloud provider session
type SessionProviderInterface interface {
	// NewSession(*corev1.Secret, string) (*session.Session, error)
	// NewEC2API(*session.Session) ec2iface.EC2API
}

// PluginSPIImpl is the real implementation of SPI interface that makes the calls to the provider SDK.
type PluginSPIImpl struct{}
