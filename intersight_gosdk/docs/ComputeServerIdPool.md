# ComputeServerIdPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "compute.ServerIdPool"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "compute.ServerIdPool"]
**PreferredIds** | Pointer to **[]int64** |  | [optional] 
**DeviceRegistration** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

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


### GetPreferredIds

`func (o *ComputeServerIdPool) GetPreferredIds() []int64`

GetPreferredIds returns the PreferredIds field if non-nil, zero value otherwise.

### GetPreferredIdsOk

`func (o *ComputeServerIdPool) GetPreferredIdsOk() (*[]int64, bool)`

GetPreferredIdsOk returns a tuple with the PreferredIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredIds

`func (o *ComputeServerIdPool) SetPreferredIds(v []int64)`

SetPreferredIds sets PreferredIds field to given value.

### HasPreferredIds

`func (o *ComputeServerIdPool) HasPreferredIds() bool`

HasPreferredIds returns a boolean if a field has been set.

### SetPreferredIdsNil

`func (o *ComputeServerIdPool) SetPreferredIdsNil(b bool)`

 SetPreferredIdsNil sets the value for PreferredIds to be an explicit nil

### UnsetPreferredIds
`func (o *ComputeServerIdPool) UnsetPreferredIds()`

UnsetPreferredIds ensures that no value is present for PreferredIds, not even an explicit nil
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

### SetDeviceRegistrationNil

`func (o *ComputeServerIdPool) SetDeviceRegistrationNil(b bool)`

 SetDeviceRegistrationNil sets the value for DeviceRegistration to be an explicit nil

### UnsetDeviceRegistration
`func (o *ComputeServerIdPool) UnsetDeviceRegistration()`

UnsetDeviceRegistration ensures that no value is present for DeviceRegistration, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


