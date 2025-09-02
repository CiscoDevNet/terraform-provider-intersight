# CapabilityPcIdRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.PcIdRange"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.PcIdRange"]
**EndPcId** | Pointer to **int64** | Ending Port-Channel ID in this range of port-channels. | [optional] 
**StartPcId** | Pointer to **int64** | Starting Port-Channel ID in this range of port-channels. | [optional] 

## Methods

### NewCapabilityPcIdRange

`func NewCapabilityPcIdRange(classId string, objectType string, ) *CapabilityPcIdRange`

NewCapabilityPcIdRange instantiates a new CapabilityPcIdRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityPcIdRangeWithDefaults

`func NewCapabilityPcIdRangeWithDefaults() *CapabilityPcIdRange`

NewCapabilityPcIdRangeWithDefaults instantiates a new CapabilityPcIdRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityPcIdRange) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityPcIdRange) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityPcIdRange) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityPcIdRange) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityPcIdRange) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityPcIdRange) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEndPcId

`func (o *CapabilityPcIdRange) GetEndPcId() int64`

GetEndPcId returns the EndPcId field if non-nil, zero value otherwise.

### GetEndPcIdOk

`func (o *CapabilityPcIdRange) GetEndPcIdOk() (*int64, bool)`

GetEndPcIdOk returns a tuple with the EndPcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPcId

`func (o *CapabilityPcIdRange) SetEndPcId(v int64)`

SetEndPcId sets EndPcId field to given value.

### HasEndPcId

`func (o *CapabilityPcIdRange) HasEndPcId() bool`

HasEndPcId returns a boolean if a field has been set.

### GetStartPcId

`func (o *CapabilityPcIdRange) GetStartPcId() int64`

GetStartPcId returns the StartPcId field if non-nil, zero value otherwise.

### GetStartPcIdOk

`func (o *CapabilityPcIdRange) GetStartPcIdOk() (*int64, bool)`

GetStartPcIdOk returns a tuple with the StartPcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPcId

`func (o *CapabilityPcIdRange) SetStartPcId(v int64)`

SetStartPcId sets StartPcId field to given value.

### HasStartPcId

`func (o *CapabilityPcIdRange) HasStartPcId() bool`

HasStartPcId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


