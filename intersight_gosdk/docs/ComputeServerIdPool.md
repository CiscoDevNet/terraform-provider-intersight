# ComputeServerIdPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "compute.ServerIdPool"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "compute.ServerIdPool"]
**GapAvailableIds** | Pointer to **[]int64** |  | [optional] 
**NextAvailableId** | Pointer to **int64** | Shows the next available server ID to be allocated. | [optional] 
**DeviceRegistration** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewComputeServerIdPool

`func NewComputeServerIdPool(classId string, objectType string, ) *ComputeServerIdPool`

NewComputeServerIdPool instantiates a new ComputeServerIdPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeServerIdPoolWithDefaults

`func NewComputeServerIdPoolWithDefaults() *ComputeServerIdPool`

NewComputeServerIdPoolWithDefaults instantiates a new ComputeServerIdPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ComputeServerIdPool) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ComputeServerIdPool) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ComputeServerIdPool) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ComputeServerIdPool) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputeServerIdPool) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputeServerIdPool) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetGapAvailableIds

`func (o *ComputeServerIdPool) GetGapAvailableIds() []int64`

GetGapAvailableIds returns the GapAvailableIds field if non-nil, zero value otherwise.

### GetGapAvailableIdsOk

`func (o *ComputeServerIdPool) GetGapAvailableIdsOk() (*[]int64, bool)`

GetGapAvailableIdsOk returns a tuple with the GapAvailableIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGapAvailableIds

`func (o *ComputeServerIdPool) SetGapAvailableIds(v []int64)`

SetGapAvailableIds sets GapAvailableIds field to given value.

### HasGapAvailableIds

`func (o *ComputeServerIdPool) HasGapAvailableIds() bool`

HasGapAvailableIds returns a boolean if a field has been set.

### SetGapAvailableIdsNil

`func (o *ComputeServerIdPool) SetGapAvailableIdsNil(b bool)`

 SetGapAvailableIdsNil sets the value for GapAvailableIds to be an explicit nil

### UnsetGapAvailableIds
`func (o *ComputeServerIdPool) UnsetGapAvailableIds()`

UnsetGapAvailableIds ensures that no value is present for GapAvailableIds, not even an explicit nil
### GetNextAvailableId

`func (o *ComputeServerIdPool) GetNextAvailableId() int64`

GetNextAvailableId returns the NextAvailableId field if non-nil, zero value otherwise.

### GetNextAvailableIdOk

`func (o *ComputeServerIdPool) GetNextAvailableIdOk() (*int64, bool)`

GetNextAvailableIdOk returns a tuple with the NextAvailableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAvailableId

`func (o *ComputeServerIdPool) SetNextAvailableId(v int64)`

SetNextAvailableId sets NextAvailableId field to given value.

### HasNextAvailableId

`func (o *ComputeServerIdPool) HasNextAvailableId() bool`

HasNextAvailableId returns a boolean if a field has been set.

### GetDeviceRegistration

`func (o *ComputeServerIdPool) GetDeviceRegistration() AssetDeviceRegistrationRelationship`

GetDeviceRegistration returns the DeviceRegistration field if non-nil, zero value otherwise.

### GetDeviceRegistrationOk

`func (o *ComputeServerIdPool) GetDeviceRegistrationOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceRegistrationOk returns a tuple with the DeviceRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRegistration

`func (o *ComputeServerIdPool) SetDeviceRegistration(v AssetDeviceRegistrationRelationship)`

SetDeviceRegistration sets DeviceRegistration field to given value.

### HasDeviceRegistration

`func (o *ComputeServerIdPool) HasDeviceRegistration() bool`

HasDeviceRegistration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


