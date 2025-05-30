/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-2025051220
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"fmt"
)

// ResourcepoolLeaseResourceRelationship - A relationship to the 'resourcepool.LeaseResource' resource, or the expanded 'resourcepool.LeaseResource' resource, or the 'null' value.
type ResourcepoolLeaseResourceRelationship struct {
	MoMoRef                   *MoMoRef
	ResourcepoolLeaseResource *ResourcepoolLeaseResource
}

// MoMoRefAsResourcepoolLeaseResourceRelationship is a convenience function that returns MoMoRef wrapped in ResourcepoolLeaseResourceRelationship
func MoMoRefAsResourcepoolLeaseResourceRelationship(v *MoMoRef) ResourcepoolLeaseResourceRelationship {
	return ResourcepoolLeaseResourceRelationship{
		MoMoRef: v,
	}
}

// ResourcepoolLeaseResourceAsResourcepoolLeaseResourceRelationship is a convenience function that returns ResourcepoolLeaseResource wrapped in ResourcepoolLeaseResourceRelationship
func ResourcepoolLeaseResourceAsResourcepoolLeaseResourceRelationship(v *ResourcepoolLeaseResource) ResourcepoolLeaseResourceRelationship {
	return ResourcepoolLeaseResourceRelationship{
		ResourcepoolLeaseResource: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ResourcepoolLeaseResourceRelationship) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'mo.MoRef'
	if jsonDict["ClassId"] == "mo.MoRef" {
		// try to unmarshal JSON data into MoMoRef
		err = json.Unmarshal(data, &dst.MoMoRef)
		if err == nil {
			return nil // data stored in dst.MoMoRef, return on the first match
		} else {
			dst.MoMoRef = nil
			return fmt.Errorf("failed to unmarshal ResourcepoolLeaseResourceRelationship as MoMoRef: %s", err.Error())
		}
	}

	// check if the discriminator value is 'resourcepool.LeaseResource'
	if jsonDict["ClassId"] == "resourcepool.LeaseResource" {
		// try to unmarshal JSON data into ResourcepoolLeaseResource
		err = json.Unmarshal(data, &dst.ResourcepoolLeaseResource)
		if err == nil {
			return nil // data stored in dst.ResourcepoolLeaseResource, return on the first match
		} else {
			dst.ResourcepoolLeaseResource = nil
			return fmt.Errorf("failed to unmarshal ResourcepoolLeaseResourceRelationship as ResourcepoolLeaseResource: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ResourcepoolLeaseResourceRelationship) MarshalJSON() ([]byte, error) {
	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	if src.ResourcepoolLeaseResource != nil {
		return json.Marshal(&src.ResourcepoolLeaseResource)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ResourcepoolLeaseResourceRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	if obj.ResourcepoolLeaseResource != nil {
		return obj.ResourcepoolLeaseResource
	}

	// all schemas are nil
	return nil
}

type NullableResourcepoolLeaseResourceRelationship struct {
	value *ResourcepoolLeaseResourceRelationship
	isSet bool
}

func (v NullableResourcepoolLeaseResourceRelationship) Get() *ResourcepoolLeaseResourceRelationship {
	return v.value
}

func (v *NullableResourcepoolLeaseResourceRelationship) Set(val *ResourcepoolLeaseResourceRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableResourcepoolLeaseResourceRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableResourcepoolLeaseResourceRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourcepoolLeaseResourceRelationship(val *ResourcepoolLeaseResourceRelationship) *NullableResourcepoolLeaseResourceRelationship {
	return &NullableResourcepoolLeaseResourceRelationship{value: val, isSet: true}
}

func (v NullableResourcepoolLeaseResourceRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourcepoolLeaseResourceRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
