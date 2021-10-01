/*
 * CLOUD API
 *
 * An enterprise-grade Infrastructure is provided as a Service (IaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.   The API allows you to perform a variety of management tasks such as spinning up additional servers, adding volumes, adjusting networking, and so forth. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 5.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// KubernetesCluster struct for KubernetesCluster
type KubernetesCluster struct {
	// The resource's unique identifier.
	Id *string `json:"id,omitempty"`
	// The type of object
	Type *string `json:"type,omitempty"`
	// URL to the object representation (absolute path)
	Href *string `json:"href,omitempty"`
	Metadata *DatacenterElementMetadata `json:"metadata,omitempty"`
	Properties *KubernetesClusterProperties `json:"properties"`
	Entities *KubernetesClusterEntities `json:"entities,omitempty"`
}



// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KubernetesCluster) GetId() *string {
	if o == nil {
		return nil
	}


	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesCluster) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Id, true
}

// SetId sets field value
func (o *KubernetesCluster) SetId(v string) {


	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *KubernetesCluster) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}


// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KubernetesCluster) GetType() *string {
	if o == nil {
		return nil
	}


	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesCluster) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Type, true
}

// SetType sets field value
func (o *KubernetesCluster) SetType(v string) {


	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *KubernetesCluster) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}


// GetHref returns the Href field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KubernetesCluster) GetHref() *string {
	if o == nil {
		return nil
	}


	return o.Href

}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesCluster) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Href, true
}

// SetHref sets field value
func (o *KubernetesCluster) SetHref(v string) {


	o.Href = &v

}

// HasHref returns a boolean if a field has been set.
func (o *KubernetesCluster) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}


// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for DatacenterElementMetadata will be returned
func (o *KubernetesCluster) GetMetadata() *DatacenterElementMetadata {
	if o == nil {
		return nil
	}


	return o.Metadata

}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesCluster) GetMetadataOk() (*DatacenterElementMetadata, bool) {
	if o == nil {
		return nil, false
	}


	return o.Metadata, true
}

// SetMetadata sets field value
func (o *KubernetesCluster) SetMetadata(v DatacenterElementMetadata) {


	o.Metadata = &v

}

// HasMetadata returns a boolean if a field has been set.
func (o *KubernetesCluster) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}


// GetProperties returns the Properties field value
// If the value is explicit nil, the zero value for KubernetesClusterProperties will be returned
func (o *KubernetesCluster) GetProperties() *KubernetesClusterProperties {
	if o == nil {
		return nil
	}


	return o.Properties

}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesCluster) GetPropertiesOk() (*KubernetesClusterProperties, bool) {
	if o == nil {
		return nil, false
	}


	return o.Properties, true
}

// SetProperties sets field value
func (o *KubernetesCluster) SetProperties(v KubernetesClusterProperties) {


	o.Properties = &v

}

// HasProperties returns a boolean if a field has been set.
func (o *KubernetesCluster) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}


// GetEntities returns the Entities field value
// If the value is explicit nil, the zero value for KubernetesClusterEntities will be returned
func (o *KubernetesCluster) GetEntities() *KubernetesClusterEntities {
	if o == nil {
		return nil
	}


	return o.Entities

}

// GetEntitiesOk returns a tuple with the Entities field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesCluster) GetEntitiesOk() (*KubernetesClusterEntities, bool) {
	if o == nil {
		return nil, false
	}


	return o.Entities, true
}

// SetEntities sets field value
func (o *KubernetesCluster) SetEntities(v KubernetesClusterEntities) {


	o.Entities = &v

}

// HasEntities returns a boolean if a field has been set.
func (o *KubernetesCluster) HasEntities() bool {
	if o != nil && o.Entities != nil {
		return true
	}

	return false
}

func (o KubernetesCluster) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	if o.Href != nil {
		toSerialize["href"] = o.Href
	}

	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}

	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}

	if o.Entities != nil {
		toSerialize["entities"] = o.Entities
	}
	return json.Marshal(toSerialize)
}

type NullableKubernetesCluster struct {
	value *KubernetesCluster
	isSet bool
}

func (v NullableKubernetesCluster) Get() *KubernetesCluster {
	return v.value
}

func (v *NullableKubernetesCluster) Set(val *KubernetesCluster) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesCluster(val *KubernetesCluster) *NullableKubernetesCluster {
	return &NullableKubernetesCluster{value: val, isSet: true}
}

func (v NullableKubernetesCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


