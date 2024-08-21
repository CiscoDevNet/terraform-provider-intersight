# FabricAbstractSpanSourceVirtualIf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "fabric.SpanSourceVnicEthIf"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "fabric.SpanSourceVnicEthIf"]
**Name** | Pointer to **string** | Name of the VNIC referenced by vnic relationship. Vnic name is not updated if it gets updated by the user. | [optional] [readonly] 
**VifId** | Pointer to **int64** | VIF ID of the VNIC placed on Fabric Interconnect associated to the SPAN session. | [optional] [readonly] 
**SpanSession** | Pointer to [**NullableFabricSpanSessionRelationship**](FabricSpanSessionRelationship.md) |  | [optional] 

## Methods

### NewFabricAbstractSpanSourceVirtualIf

`func NewFabricAbstractSpanSourceVirtualIf(classId string, objectType string, ) *FabricAbstractSpanSourceVirtualIf`

NewFabricAbstractSpanSourceVirtualIf instantiates a new FabricAbstractSpanSourceVirtualIf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricAbstractSpanSourceVirtualIfWithDefaults

`func NewFabricAbstractSpanSourceVirtualIfWithDefaults() *FabricAbstractSpanSourceVirtualIf`

NewFabricAbstractSpanSourceVirtualIfWithDefaults instantiates a new FabricAbstractSpanSourceVirtualIf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricAbstractSpanSourceVirtualIf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricAbstractSpanSourceVirtualIf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricAbstractSpanSourceVirtualIf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricAbstractSpanSourceVirtualIf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricAbstractSpanSourceVirtualIf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricAbstractSpanSourceVirtualIf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *FabricAbstractSpanSourceVirtualIf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FabricAbstractSpanSourceVirtualIf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FabricAbstractSpanSourceVirtualIf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FabricAbstractSpanSourceVirtualIf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVifId

`func (o *FabricAbstractSpanSourceVirtualIf) GetVifId() int64`

GetVifId returns the VifId field if non-nil, zero value otherwise.

### GetVifIdOk

`func (o *FabricAbstractSpanSourceVirtualIf) GetVifIdOk() (*int64, bool)`

GetVifIdOk returns a tuple with the VifId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVifId

`func (o *FabricAbstractSpanSourceVirtualIf) SetVifId(v int64)`

SetVifId sets VifId field to given value.

### HasVifId

`func (o *FabricAbstractSpanSourceVirtualIf) HasVifId() bool`

HasVifId returns a boolean if a field has been set.

### GetSpanSession

`func (o *FabricAbstractSpanSourceVirtualIf) GetSpanSession() FabricSpanSessionRelationship`

GetSpanSession returns the SpanSession field if non-nil, zero value otherwise.

### GetSpanSessionOk

`func (o *FabricAbstractSpanSourceVirtualIf) GetSpanSessionOk() (*FabricSpanSessionRelationship, bool)`

GetSpanSessionOk returns a tuple with the SpanSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanSession

`func (o *FabricAbstractSpanSourceVirtualIf) SetSpanSession(v FabricSpanSessionRelationship)`

SetSpanSession sets SpanSession field to given value.

### HasSpanSession

`func (o *FabricAbstractSpanSourceVirtualIf) HasSpanSession() bool`

HasSpanSession returns a boolean if a field has been set.

### SetSpanSessionNil

`func (o *FabricAbstractSpanSourceVirtualIf) SetSpanSessionNil(b bool)`

 SetSpanSessionNil sets the value for SpanSession to be an explicit nil

### UnsetSpanSession
`func (o *FabricAbstractSpanSourceVirtualIf) UnsetSpanSession()`

UnsetSpanSession ensures that no value is present for SpanSession, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


