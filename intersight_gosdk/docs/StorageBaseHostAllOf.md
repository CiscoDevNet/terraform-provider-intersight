# StorageBaseHostAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Short description about the host. | [optional] [readonly] 
**Initiators** | Pointer to [**[]StorageBaseInitiator**](storage.BaseInitiator.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the host in storage array. | [optional] [readonly] 
**OsType** | Pointer to **string** | Operating system running on the host. | [optional] [readonly] 
**StorageUtilization** | Pointer to [**StorageBaseCapacity**](storage.BaseCapacity.md) |  | [optional] 

## Methods

### NewStorageBaseHostAllOf

`func NewStorageBaseHostAllOf() *StorageBaseHostAllOf`

NewStorageBaseHostAllOf instantiates a new StorageBaseHostAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageBaseHostAllOfWithDefaults

`func NewStorageBaseHostAllOfWithDefaults() *StorageBaseHostAllOf`

NewStorageBaseHostAllOfWithDefaults instantiates a new StorageBaseHostAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *StorageBaseHostAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StorageBaseHostAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StorageBaseHostAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StorageBaseHostAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInitiators

`func (o *StorageBaseHostAllOf) GetInitiators() []StorageBaseInitiator`

GetInitiators returns the Initiators field if non-nil, zero value otherwise.

### GetInitiatorsOk

`func (o *StorageBaseHostAllOf) GetInitiatorsOk() (*[]StorageBaseInitiator, bool)`

GetInitiatorsOk returns a tuple with the Initiators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiators

`func (o *StorageBaseHostAllOf) SetInitiators(v []StorageBaseInitiator)`

SetInitiators sets Initiators field to given value.

### HasInitiators

`func (o *StorageBaseHostAllOf) HasInitiators() bool`

HasInitiators returns a boolean if a field has been set.

### GetName

`func (o *StorageBaseHostAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageBaseHostAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageBaseHostAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageBaseHostAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOsType

`func (o *StorageBaseHostAllOf) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *StorageBaseHostAllOf) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *StorageBaseHostAllOf) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *StorageBaseHostAllOf) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetStorageUtilization

`func (o *StorageBaseHostAllOf) GetStorageUtilization() StorageBaseCapacity`

GetStorageUtilization returns the StorageUtilization field if non-nil, zero value otherwise.

### GetStorageUtilizationOk

`func (o *StorageBaseHostAllOf) GetStorageUtilizationOk() (*StorageBaseCapacity, bool)`

GetStorageUtilizationOk returns a tuple with the StorageUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUtilization

`func (o *StorageBaseHostAllOf) SetStorageUtilization(v StorageBaseCapacity)`

SetStorageUtilization sets StorageUtilization field to given value.

### HasStorageUtilization

`func (o *StorageBaseHostAllOf) HasStorageUtilization() bool`

HasStorageUtilization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


