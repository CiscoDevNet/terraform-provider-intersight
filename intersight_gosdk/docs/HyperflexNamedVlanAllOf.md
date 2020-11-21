# HyperflexNamedVlanAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.NamedVlan"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.NamedVlan"]
**Name** | Pointer to **string** | The name of the VLAN. The name can be from 1 to 32 characters long and can contain a combination of alphanumeric characters, underscores, and hyphens. | [optional] 
**VlanId** | Pointer to **int64** | The ID of the named VLAN. An ID of 0 means the traffic is untagged. The ID can be any number between 0 and 4095, inclusive. | [optional] 

## Methods

### NewHyperflexNamedVlanAllOf

`func NewHyperflexNamedVlanAllOf(classId string, objectType string, ) *HyperflexNamedVlanAllOf`

NewHyperflexNamedVlanAllOf instantiates a new HyperflexNamedVlanAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexNamedVlanAllOfWithDefaults

`func NewHyperflexNamedVlanAllOfWithDefaults() *HyperflexNamedVlanAllOf`

NewHyperflexNamedVlanAllOfWithDefaults instantiates a new HyperflexNamedVlanAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexNamedVlanAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexNamedVlanAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexNamedVlanAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexNamedVlanAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexNamedVlanAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexNamedVlanAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *HyperflexNamedVlanAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HyperflexNamedVlanAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HyperflexNamedVlanAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HyperflexNamedVlanAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVlanId

`func (o *HyperflexNamedVlanAllOf) GetVlanId() int64`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *HyperflexNamedVlanAllOf) GetVlanIdOk() (*int64, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *HyperflexNamedVlanAllOf) SetVlanId(v int64)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *HyperflexNamedVlanAllOf) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


