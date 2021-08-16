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

// Package ionos contains the IONOS provider specific implementations to manage machines
package ionos

import (
	"github.com/23technologies/machine-controller-manager-provider-ionos/pkg/spi"
	"github.com/gardener/machine-controller-manager/pkg/util/provider/driver"
)

// MachineProvider is the struct that implements the driver interface
type MachineProvider struct {
	SPI spi.SessionProviderInterface
}

// NewIonosProvider returns a provider object.
//
// PARAMETERS
// spi spi.SessionProviderInterface Session provider interface to attach
func NewIonosProvider(spi spi.SessionProviderInterface) driver.Driver {
	return &MachineProvider{
		SPI: spi,
	}
}
