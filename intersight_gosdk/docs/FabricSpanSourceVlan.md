# FabricSpanSourceVlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.SpanSourceVlan"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.SpanSourceVlan"]
**VlanId** | Pointer to **int64** | SPAN source VLAN Identifier. | [optional] 
**SpanSession** | Pointer to [**NullableFabricSpanSessionRelationship**](FabricSpanSessionRelationship.md) |  | [optional] 

## Methods

### NewFabricSpanSourceVlan

`func NewFabricSpanSourceVlan(classId string, objectType string, ) *FabricSpanSourceVlan`

NewFabricSpanSourceVlan instantiates a new FabricSpanSourceVlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricSpanSourceVlanWithDefaults

`func NewFabricSpanSourceVlanWithDefaults() *FabricSpanSourceVlan`

NewFabricSpanSourceVlanWithDefaults instantiates a new FabricSpanSourceVlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricSpanSourceVlan) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricSpanSourceVlan) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricSpanSourceVlan) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricSpanSourceVlan) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricSpanSourceVlan) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricSpanSourceVlan) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetVlanId

`func (o *FabricSpanSourceVlan) GetVlanId() int64`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *FabricSpanSourceVlan) GetVlanIdOk() (*int64, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *FabricSpanSourceVlan) SetVlanId(v int64)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *FabricSpanSourceVlan) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetSpanSession

`func (o *FabricSpanSourceVlan) GetSpanSession() FabricSpanSessionRelationship`

GetSpanSession returns the SpanSession field if non-nil, zero value otherwise.

### GetSpanSessionOk

`func (o *FabricSpanSourceVlan) GetSpanSessionOk() (*FabricSpanSessionRelationship, bool)`

GetSpanSessionOk returns a tuple with the SpanSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanSession

`func (o *FabricSpanSourceVlan) SetSpanSession(v FabricSpanSessionRelationship)`

SetSpanSession sets SpanSession field to given value.

### HasSpanSession

`func (o *FabricSpanSourceVlan) HasSpanSession() bool`

HasSpanSession returns a boolean if a field has been set.

### SetSpanSessionNil

`func (o *FabricSpanSourceVlan) SetSpanSessionNil(b bool)`

 SetSpanSessionNil sets the value for SpanSession to be an explicit nil

### UnsetSpanSession
`func (o *FabricSpanSourceVlan) UnsetSpanSession()`

UnsetSpanSession ensures that no value is present for SpanSession, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


