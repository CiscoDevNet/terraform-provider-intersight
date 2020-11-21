# CapabilityPortRangeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.PortRange"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.PortRange"]
**EndPortId** | Pointer to **int64** | Ending Port ID in this range of ports. | [optional] 
**EndSlotId** | Pointer to **int64** | Ending Slot ID in this range of ports. | [optional] 
**StartPortId** | Pointer to **int64** | Starting Port ID in this range of ports. | [optional] 
**StartSlotId** | Pointer to **int64** | Starting Slot ID in this range of ports. | [optional] 

## Methods

### NewCapabilityPortRangeAllOf

`func NewCapabilityPortRangeAllOf(classId string, objectType string, ) *CapabilityPortRangeAllOf`

NewCapabilityPortRangeAllOf instantiates a new CapabilityPortRangeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityPortRangeAllOfWithDefaults

`func NewCapabilityPortRangeAllOfWithDefaults() *CapabilityPortRangeAllOf`

NewCapabilityPortRangeAllOfWithDefaults instantiates a new CapabilityPortRangeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityPortRangeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityPortRangeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityPortRangeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityPortRangeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityPortRangeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityPortRangeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEndPortId

`func (o *CapabilityPortRangeAllOf) GetEndPortId() int64`

GetEndPortId returns the EndPortId field if non-nil, zero value otherwise.

### GetEndPortIdOk

`func (o *CapabilityPortRangeAllOf) GetEndPortIdOk() (*int64, bool)`

GetEndPortIdOk returns a tuple with the EndPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPortId

`func (o *CapabilityPortRangeAllOf) SetEndPortId(v int64)`

SetEndPortId sets EndPortId field to given value.

### HasEndPortId

`func (o *CapabilityPortRangeAllOf) HasEndPortId() bool`

HasEndPortId returns a boolean if a field has been set.

### GetEndSlotId

`func (o *CapabilityPortRangeAllOf) GetEndSlotId() int64`

GetEndSlotId returns the EndSlotId field if non-nil, zero value otherwise.

### GetEndSlotIdOk

`func (o *CapabilityPortRangeAllOf) GetEndSlotIdOk() (*int64, bool)`

GetEndSlotIdOk returns a tuple with the EndSlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndSlotId

`func (o *CapabilityPortRangeAllOf) SetEndSlotId(v int64)`

SetEndSlotId sets EndSlotId field to given value.

### HasEndSlotId

`func (o *CapabilityPortRangeAllOf) HasEndSlotId() bool`

HasEndSlotId returns a boolean if a field has been set.

### GetStartPortId

`func (o *CapabilityPortRangeAllOf) GetStartPortId() int64`

GetStartPortId returns the StartPortId field if non-nil, zero value otherwise.

### GetStartPortIdOk

`func (o *CapabilityPortRangeAllOf) GetStartPortIdOk() (*int64, bool)`

GetStartPortIdOk returns a tuple with the StartPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPortId

`func (o *CapabilityPortRangeAllOf) SetStartPortId(v int64)`

SetStartPortId sets StartPortId field to given value.

### HasStartPortId

`func (o *CapabilityPortRangeAllOf) HasStartPortId() bool`

HasStartPortId returns a boolean if a field has been set.

### GetStartSlotId

`func (o *CapabilityPortRangeAllOf) GetStartSlotId() int64`

GetStartSlotId returns the StartSlotId field if non-nil, zero value otherwise.

### GetStartSlotIdOk

`func (o *CapabilityPortRangeAllOf) GetStartSlotIdOk() (*int64, bool)`

GetStartSlotIdOk returns a tuple with the StartSlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartSlotId

`func (o *CapabilityPortRangeAllOf) SetStartSlotId(v int64)`

SetStartSlotId sets StartSlotId field to given value.

### HasStartSlotId

`func (o *CapabilityPortRangeAllOf) HasStartSlotId() bool`

HasStartSlotId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


