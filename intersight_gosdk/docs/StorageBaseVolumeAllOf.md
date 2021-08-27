# StorageBaseVolumeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Description** | Pointer to **string** | Short description about the volume. | [optional] [readonly] 
**NaaId** | Pointer to **string** | NAA id of volume. It is a significant number to identify corresponding lun path in hypervisor. | [optional] [readonly] 
**Name** | Pointer to **string** | Named entity of the volume. | [optional] [readonly] 
**Size** | Pointer to **int64** | User provisioned volume size. It is the size exposed to host. | [optional] [readonly] 
**StorageUtilization** | Pointer to [**NullableStorageBaseCapacity**](StorageBaseCapacity.md) |  | [optional] 

## Methods

### NewStorageBaseVolumeAllOf

`func NewStorageBaseVolumeAllOf(classId string, objectType string, ) *StorageBaseVolumeAllOf`

NewStorageBaseVolumeAllOf instantiates a new StorageBaseVolumeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageBaseVolumeAllOfWithDefaults

`func NewStorageBaseVolumeAllOfWithDefaults() *StorageBaseVolumeAllOf`

NewStorageBaseVolumeAllOfWithDefaults instantiates a new StorageBaseVolumeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageBaseVolumeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageBaseVolumeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageBaseVolumeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageBaseVolumeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageBaseVolumeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageBaseVolumeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *StorageBaseVolumeAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StorageBaseVolumeAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StorageBaseVolumeAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StorageBaseVolumeAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNaaId

`func (o *StorageBaseVolumeAllOf) GetNaaId() string`

GetNaaId returns the NaaId field if non-nil, zero value otherwise.

### GetNaaIdOk

`func (o *StorageBaseVolumeAllOf) GetNaaIdOk() (*string, bool)`

GetNaaIdOk returns a tuple with the NaaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaaId

`func (o *StorageBaseVolumeAllOf) SetNaaId(v string)`

SetNaaId sets NaaId field to given value.

### HasNaaId

`func (o *StorageBaseVolumeAllOf) HasNaaId() bool`

HasNaaId returns a boolean if a field has been set.

### GetName

`func (o *StorageBaseVolumeAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageBaseVolumeAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageBaseVolumeAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageBaseVolumeAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *StorageBaseVolumeAllOf) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *StorageBaseVolumeAllOf) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *StorageBaseVolumeAllOf) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *StorageBaseVolumeAllOf) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStorageUtilization

`func (o *StorageBaseVolumeAllOf) GetStorageUtilization() StorageBaseCapacity`

GetStorageUtilization returns the StorageUtilization field if non-nil, zero value otherwise.

### GetStorageUtilizationOk

`func (o *StorageBaseVolumeAllOf) GetStorageUtilizationOk() (*StorageBaseCapacity, bool)`

GetStorageUtilizationOk returns a tuple with the StorageUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUtilization

`func (o *StorageBaseVolumeAllOf) SetStorageUtilization(v StorageBaseCapacity)`

SetStorageUtilization sets StorageUtilization field to given value.

### HasStorageUtilization

`func (o *StorageBaseVolumeAllOf) HasStorageUtilization() bool`

HasStorageUtilization returns a boolean if a field has been set.

### SetStorageUtilizationNil

`func (o *StorageBaseVolumeAllOf) SetStorageUtilizationNil(b bool)`

 SetStorageUtilizationNil sets the value for StorageUtilization to be an explicit nil

### UnsetStorageUtilization
`func (o *StorageBaseVolumeAllOf) UnsetStorageUtilization()`

UnsetStorageUtilization ensures that no value is present for StorageUtilization, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


