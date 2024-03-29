/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-15711
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// UuidpoolReservation The UUID reservation object, used to hold reserved identity.
type UuidpoolReservation struct {
	PoolReservation
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// UUID identity to be reserved.
	Identity             *string                               `json:"Identity,omitempty"`
	Block                *UuidpoolBlockRelationship            `json:"Block,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	Pool                 *UuidpoolPoolRelationship             `json:"Pool,omitempty"`
	PoolMember           *UuidpoolPoolMemberRelationship       `json:"PoolMember,omitempty"`
	Universe             *UuidpoolUniverseRelationship         `json:"Universe,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UuidpoolReservation UuidpoolReservation

// NewUuidpoolReservation instantiates a new UuidpoolReservation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUuidpoolReservation(classId string, objectType string) *UuidpoolReservation {
	this := UuidpoolReservation{}
	this.ClassId = classId
	this.ObjectType = objectType
	var allocationType string = "dynamic"
	this.AllocationType = &allocationType
	return &this
}

// NewUuidpoolReservationWithDefaults instantiates a new UuidpoolReservation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUuidpoolReservationWithDefaults() *UuidpoolReservation {
	this := UuidpoolReservation{}
	var classId string = "uuidpool.Reservation"
	this.ClassId = classId
	var objectType string = "uuidpool.Reservation"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *UuidpoolReservation) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *UuidpoolReservation) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *UuidpoolReservation) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *UuidpoolReservation) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *UuidpoolReservation) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *UuidpoolReservation) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *UuidpoolReservation) GetIdentity() string {
	if o == nil || o.Identity == nil {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UuidpoolReservation) GetIdentityOk() (*string, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *UuidpoolReservation) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *UuidpoolReservation) SetIdentity(v string) {
	o.Identity = &v
}

// GetBlock returns the Block field value if set, zero value otherwise.
func (o *UuidpoolReservation) GetBlock() UuidpoolBlockRelationship {
	if o == nil || o.Block == nil {
		var ret UuidpoolBlockRelationship
		return ret
	}
	return *o.Block
}

// GetBlockOk returns a tuple with the Block field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UuidpoolReservation) GetBlockOk() (*UuidpoolBlockRelationship, bool) {
	if o == nil || o.Block == nil {
		return nil, false
	}
	return o.Block, true
}

// HasBlock returns a boolean if a field has been set.
func (o *UuidpoolReservation) HasBlock() bool {
	if o != nil && o.Block != nil {
		return true
	}

	return false
}

// SetBlock gets a reference to the given UuidpoolBlockRelationship and assigns it to the Block field.
func (o *UuidpoolReservation) SetBlock(v UuidpoolBlockRelationship) {
	o.Block = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *UuidpoolReservation) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UuidpoolReservation) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *UuidpoolReservation) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *UuidpoolReservation) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *UuidpoolReservation) GetPool() UuidpoolPoolRelationship {
	if o == nil || o.Pool == nil {
		var ret UuidpoolPoolRelationship
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UuidpoolReservation) GetPoolOk() (*UuidpoolPoolRelationship, bool) {
	if o == nil || o.Pool == nil {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *UuidpoolReservation) HasPool() bool {
	if o != nil && o.Pool != nil {
		return true
	}

	return false
}

// SetPool gets a reference to the given UuidpoolPoolRelationship and assigns it to the Pool field.
func (o *UuidpoolReservation) SetPool(v UuidpoolPoolRelationship) {
	o.Pool = &v
}

// GetPoolMember returns the PoolMember field value if set, zero value otherwise.
func (o *UuidpoolReservation) GetPoolMember() UuidpoolPoolMemberRelationship {
	if o == nil || o.PoolMember == nil {
		var ret UuidpoolPoolMemberRelationship
		return ret
	}
	return *o.PoolMember
}

// GetPoolMemberOk returns a tuple with the PoolMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UuidpoolReservation) GetPoolMemberOk() (*UuidpoolPoolMemberRelationship, bool) {
	if o == nil || o.PoolMember == nil {
		return nil, false
	}
	return o.PoolMember, true
}

// HasPoolMember returns a boolean if a field has been set.
func (o *UuidpoolReservation) HasPoolMember() bool {
	if o != nil && o.PoolMember != nil {
		return true
	}

	return false
}

// SetPoolMember gets a reference to the given UuidpoolPoolMemberRelationship and assigns it to the PoolMember field.
func (o *UuidpoolReservation) SetPoolMember(v UuidpoolPoolMemberRelationship) {
	o.PoolMember = &v
}

// GetUniverse returns the Universe field value if set, zero value otherwise.
func (o *UuidpoolReservation) GetUniverse() UuidpoolUniverseRelationship {
	if o == nil || o.Universe == nil {
		var ret UuidpoolUniverseRelationship
		return ret
	}
	return *o.Universe
}

// GetUniverseOk returns a tuple with the Universe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UuidpoolReservation) GetUniverseOk() (*UuidpoolUniverseRelationship, bool) {
	if o == nil || o.Universe == nil {
		return nil, false
	}
	return o.Universe, true
}

// HasUniverse returns a boolean if a field has been set.
func (o *UuidpoolReservation) HasUniverse() bool {
	if o != nil && o.Universe != nil {
		return true
	}

	return false
}

// SetUniverse gets a reference to the given UuidpoolUniverseRelationship and assigns it to the Universe field.
func (o *UuidpoolReservation) SetUniverse(v UuidpoolUniverseRelationship) {
	o.Universe = &v
}

func (o UuidpoolReservation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPoolReservation, errPoolReservation := json.Marshal(o.PoolReservation)
	if errPoolReservation != nil {
		return []byte{}, errPoolReservation
	}
	errPoolReservation = json.Unmarshal([]byte(serializedPoolReservation), &toSerialize)
	if errPoolReservation != nil {
		return []byte{}, errPoolReservation
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Identity != nil {
		toSerialize["Identity"] = o.Identity
	}
	if o.Block != nil {
		toSerialize["Block"] = o.Block
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Pool != nil {
		toSerialize["Pool"] = o.Pool
	}
	if o.PoolMember != nil {
		toSerialize["PoolMember"] = o.PoolMember
	}
	if o.Universe != nil {
		toSerialize["Universe"] = o.Universe
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UuidpoolReservation) UnmarshalJSON(bytes []byte) (err error) {
	type UuidpoolReservationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// UUID identity to be reserved.
		Identity     *string                               `json:"Identity,omitempty"`
		Block        *UuidpoolBlockRelationship            `json:"Block,omitempty"`
		Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
		Pool         *UuidpoolPoolRelationship             `json:"Pool,omitempty"`
		PoolMember   *UuidpoolPoolMemberRelationship       `json:"PoolMember,omitempty"`
		Universe     *UuidpoolUniverseRelationship         `json:"Universe,omitempty"`
	}

	varUuidpoolReservationWithoutEmbeddedStruct := UuidpoolReservationWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varUuidpoolReservationWithoutEmbeddedStruct)
	if err == nil {
		varUuidpoolReservation := _UuidpoolReservation{}
		varUuidpoolReservation.ClassId = varUuidpoolReservationWithoutEmbeddedStruct.ClassId
		varUuidpoolReservation.ObjectType = varUuidpoolReservationWithoutEmbeddedStruct.ObjectType
		varUuidpoolReservation.Identity = varUuidpoolReservationWithoutEmbeddedStruct.Identity
		varUuidpoolReservation.Block = varUuidpoolReservationWithoutEmbeddedStruct.Block
		varUuidpoolReservation.Organization = varUuidpoolReservationWithoutEmbeddedStruct.Organization
		varUuidpoolReservation.Pool = varUuidpoolReservationWithoutEmbeddedStruct.Pool
		varUuidpoolReservation.PoolMember = varUuidpoolReservationWithoutEmbeddedStruct.PoolMember
		varUuidpoolReservation.Universe = varUuidpoolReservationWithoutEmbeddedStruct.Universe
		*o = UuidpoolReservation(varUuidpoolReservation)
	} else {
		return err
	}

	varUuidpoolReservation := _UuidpoolReservation{}

	err = json.Unmarshal(bytes, &varUuidpoolReservation)
	if err == nil {
		o.PoolReservation = varUuidpoolReservation.PoolReservation
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Identity")
		delete(additionalProperties, "Block")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Pool")
		delete(additionalProperties, "PoolMember")
		delete(additionalProperties, "Universe")

		// remove fields from embedded structs
		reflectPoolReservation := reflect.ValueOf(o.PoolReservation)
		for i := 0; i < reflectPoolReservation.Type().NumField(); i++ {
			t := reflectPoolReservation.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUuidpoolReservation struct {
	value *UuidpoolReservation
	isSet bool
}

func (v NullableUuidpoolReservation) Get() *UuidpoolReservation {
	return v.value
}

func (v *NullableUuidpoolReservation) Set(val *UuidpoolReservation) {
	v.value = val
	v.isSet = true
}

func (v NullableUuidpoolReservation) IsSet() bool {
	return v.isSet
}

func (v *NullableUuidpoolReservation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUuidpoolReservation(val *UuidpoolReservation) *NullableUuidpoolReservation {
	return &NullableUuidpoolReservation{value: val, isSet: true}
}

func (v NullableUuidpoolReservation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUuidpoolReservation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
