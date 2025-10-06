# VnicVifIdPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.VifIdPool"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.VifIdPool"]
**GapAvailableIds** | Pointer to **[]int64** |  | [optional] 
**NextAvailableId** | Pointer to **int64** | Shows the next available channel number ID to be allocated for a vNIC. | [optional] 
**NetworkElement** | Pointer to [**NullableNetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 

## Methods

### NewVnicVifIdPool

`func NewVnicVifIdPool(classId string, objectType string, ) *VnicVifIdPool`

NewVnicVifIdPool instantiates a new VnicVifIdPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicVifIdPoolWithDefaults

`func NewVnicVifIdPoolWithDefaults() *VnicVifIdPool`

NewVnicVifIdPoolWithDefaults instantiates a new VnicVifIdPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicVifIdPool) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicVifIdPool) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicVifIdPool) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicVifIdPool) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicVifIdPool) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicVifIdPool) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetGapAvailableIds

`func (o *VnicVifIdPool) GetGapAvailableIds() []int64`

GetGapAvailableIds returns the GapAvailableIds field if non-nil, zero value otherwise.

### GetGapAvailableIdsOk

`func (o *VnicVifIdPool) GetGapAvailableIdsOk() (*[]int64, bool)`

GetGapAvailableIdsOk returns a tuple with the GapAvailableIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGapAvailableIds

`func (o *VnicVifIdPool) SetGapAvailableIds(v []int64)`

SetGapAvailableIds sets GapAvailableIds field to given value.

### HasGapAvailableIds

`func (o *VnicVifIdPool) HasGapAvailableIds() bool`

HasGapAvailableIds returns a boolean if a field has been set.

### SetGapAvailableIdsNil

`func (o *VnicVifIdPool) SetGapAvailableIdsNil(b bool)`

 SetGapAvailableIdsNil sets the value for GapAvailableIds to be an explicit nil

### UnsetGapAvailableIds
`func (o *VnicVifIdPool) UnsetGapAvailableIds()`

UnsetGapAvailableIds ensures that no value is present for GapAvailableIds, not even an explicit nil
### GetNextAvailableId

`func (o *VnicVifIdPool) GetNextAvailableId() int64`

GetNextAvailableId returns the NextAvailableId field if non-nil, zero value otherwise.

### GetNextAvailableIdOk

`func (o *VnicVifIdPool) GetNextAvailableIdOk() (*int64, bool)`

GetNextAvailableIdOk returns a tuple with the NextAvailableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAvailableId

`func (o *VnicVifIdPool) SetNextAvailableId(v int64)`

SetNextAvailableId sets NextAvailableId field to given value.

### HasNextAvailableId

`func (o *VnicVifIdPool) HasNextAvailableId() bool`

HasNextAvailableId returns a boolean if a field has been set.

### GetNetworkElement

`func (o *VnicVifIdPool) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *VnicVifIdPool) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *VnicVifIdPool) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *VnicVifIdPool) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### SetNetworkElementNil

`func (o *VnicVifIdPool) SetNetworkElementNil(b bool)`

 SetNetworkElementNil sets the value for NetworkElement to be an explicit nil

### UnsetNetworkElement
`func (o *VnicVifIdPool) UnsetNetworkElement()`

UnsetNetworkElement ensures that no value is present for NetworkElement, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


