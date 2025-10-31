# CapabilityGpuConnectionPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.GpuConnectionPoint"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.GpuConnectionPoint"]
**EndPointIds** | Pointer to **[]int64** |  | [optional] 
**PcieNodeSlots** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewCapabilityGpuConnectionPoint

`func NewCapabilityGpuConnectionPoint(classId string, objectType string, ) *CapabilityGpuConnectionPoint`

NewCapabilityGpuConnectionPoint instantiates a new CapabilityGpuConnectionPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityGpuConnectionPointWithDefaults

`func NewCapabilityGpuConnectionPointWithDefaults() *CapabilityGpuConnectionPoint`

NewCapabilityGpuConnectionPointWithDefaults instantiates a new CapabilityGpuConnectionPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityGpuConnectionPoint) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityGpuConnectionPoint) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityGpuConnectionPoint) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityGpuConnectionPoint) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityGpuConnectionPoint) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityGpuConnectionPoint) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEndPointIds

`func (o *CapabilityGpuConnectionPoint) GetEndPointIds() []int64`

GetEndPointIds returns the EndPointIds field if non-nil, zero value otherwise.

### GetEndPointIdsOk

`func (o *CapabilityGpuConnectionPoint) GetEndPointIdsOk() (*[]int64, bool)`

GetEndPointIdsOk returns a tuple with the EndPointIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPointIds

`func (o *CapabilityGpuConnectionPoint) SetEndPointIds(v []int64)`

SetEndPointIds sets EndPointIds field to given value.

### HasEndPointIds

`func (o *CapabilityGpuConnectionPoint) HasEndPointIds() bool`

HasEndPointIds returns a boolean if a field has been set.

### SetEndPointIdsNil

`func (o *CapabilityGpuConnectionPoint) SetEndPointIdsNil(b bool)`

 SetEndPointIdsNil sets the value for EndPointIds to be an explicit nil

### UnsetEndPointIds
`func (o *CapabilityGpuConnectionPoint) UnsetEndPointIds()`

UnsetEndPointIds ensures that no value is present for EndPointIds, not even an explicit nil
### GetPcieNodeSlots

`func (o *CapabilityGpuConnectionPoint) GetPcieNodeSlots() []int64`

GetPcieNodeSlots returns the PcieNodeSlots field if non-nil, zero value otherwise.

### GetPcieNodeSlotsOk

`func (o *CapabilityGpuConnectionPoint) GetPcieNodeSlotsOk() (*[]int64, bool)`

GetPcieNodeSlotsOk returns a tuple with the PcieNodeSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcieNodeSlots

`func (o *CapabilityGpuConnectionPoint) SetPcieNodeSlots(v []int64)`

SetPcieNodeSlots sets PcieNodeSlots field to given value.

### HasPcieNodeSlots

`func (o *CapabilityGpuConnectionPoint) HasPcieNodeSlots() bool`

HasPcieNodeSlots returns a boolean if a field has been set.

### SetPcieNodeSlotsNil

`func (o *CapabilityGpuConnectionPoint) SetPcieNodeSlotsNil(b bool)`

 SetPcieNodeSlotsNil sets the value for PcieNodeSlots to be an explicit nil

### UnsetPcieNodeSlots
`func (o *CapabilityGpuConnectionPoint) UnsetPcieNodeSlots()`

UnsetPcieNodeSlots ensures that no value is present for PcieNodeSlots, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


