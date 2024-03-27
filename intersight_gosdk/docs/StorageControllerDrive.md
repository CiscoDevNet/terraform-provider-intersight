# StorageControllerDrive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.ControllerDrive"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.ControllerDrive"]
**ControllerType** | Pointer to **string** | The type of Local Storage like FlexMMC, USB Drive. * &#x60;Unknown&#x60; - Storage partition type is not known. * &#x60;FlexMMC&#x60; - The FlexMMC type of storage local drive. * &#x60;USB&#x60; - The USB type of storage drive. | [optional] [readonly] [default to "Unknown"]
**Description** | Pointer to **string** | The description of local storage drive. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the local Storage. | [optional] [readonly] 
**PartitionCount** | Pointer to **int64** | Total Partition count in local storage. | [optional] [readonly] 
**StorageId** | Pointer to **string** | The Id of the local Storage. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of storage like internal or external. * &#x60;Unknown&#x60; - Not any of the known Storage Types. * &#x60;Internal&#x60; - The internal storage type. * &#x60;External&#x60; - The external storage type. | [optional] [readonly] [default to "Unknown"]
**ComputeBoard** | Pointer to [**ComputeBoardRelationship**](ComputeBoardRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**StorageItem** | Pointer to [**[]StorageItemRelationship**](StorageItemRelationship.md) | An array of relationships to storageItem resources. | [optional] [readonly] 

## Methods

### NewStorageControllerDrive

`func NewStorageControllerDrive(classId string, objectType string, ) *StorageControllerDrive`

NewStorageControllerDrive instantiates a new StorageControllerDrive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageControllerDriveWithDefaults

`func NewStorageControllerDriveWithDefaults() *StorageControllerDrive`

NewStorageControllerDriveWithDefaults instantiates a new StorageControllerDrive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageControllerDrive) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageControllerDrive) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageControllerDrive) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageControllerDrive) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageControllerDrive) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageControllerDrive) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetControllerType

`func (o *StorageControllerDrive) GetControllerType() string`

GetControllerType returns the ControllerType field if non-nil, zero value otherwise.

### GetControllerTypeOk

`func (o *StorageControllerDrive) GetControllerTypeOk() (*string, bool)`

GetControllerTypeOk returns a tuple with the ControllerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerType

`func (o *StorageControllerDrive) SetControllerType(v string)`

SetControllerType sets ControllerType field to given value.

### HasControllerType

`func (o *StorageControllerDrive) HasControllerType() bool`

HasControllerType returns a boolean if a field has been set.

### GetDescription

`func (o *StorageControllerDrive) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StorageControllerDrive) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StorageControllerDrive) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StorageControllerDrive) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *StorageControllerDrive) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageControllerDrive) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageControllerDrive) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageControllerDrive) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPartitionCount

`func (o *StorageControllerDrive) GetPartitionCount() int64`

GetPartitionCount returns the PartitionCount field if non-nil, zero value otherwise.

### GetPartitionCountOk

`func (o *StorageControllerDrive) GetPartitionCountOk() (*int64, bool)`

GetPartitionCountOk returns a tuple with the PartitionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionCount

`func (o *StorageControllerDrive) SetPartitionCount(v int64)`

SetPartitionCount sets PartitionCount field to given value.

### HasPartitionCount

`func (o *StorageControllerDrive) HasPartitionCount() bool`

HasPartitionCount returns a boolean if a field has been set.

### GetStorageId

`func (o *StorageControllerDrive) GetStorageId() string`

GetStorageId returns the StorageId field if non-nil, zero value otherwise.

### GetStorageIdOk

`func (o *StorageControllerDrive) GetStorageIdOk() (*string, bool)`

GetStorageIdOk returns a tuple with the StorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageId

`func (o *StorageControllerDrive) SetStorageId(v string)`

SetStorageId sets StorageId field to given value.

### HasStorageId

`func (o *StorageControllerDrive) HasStorageId() bool`

HasStorageId returns a boolean if a field has been set.

### GetType

`func (o *StorageControllerDrive) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageControllerDrive) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageControllerDrive) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StorageControllerDrive) HasType() bool`

HasType returns a boolean if a field has been set.

### GetComputeBoard

`func (o *StorageControllerDrive) GetComputeBoard() ComputeBoardRelationship`

GetComputeBoard returns the ComputeBoard field if non-nil, zero value otherwise.

### GetComputeBoardOk

`func (o *StorageControllerDrive) GetComputeBoardOk() (*ComputeBoardRelationship, bool)`

GetComputeBoardOk returns a tuple with the ComputeBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBoard

`func (o *StorageControllerDrive) SetComputeBoard(v ComputeBoardRelationship)`

SetComputeBoard sets ComputeBoard field to given value.

### HasComputeBoard

`func (o *StorageControllerDrive) HasComputeBoard() bool`

HasComputeBoard returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageControllerDrive) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageControllerDrive) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageControllerDrive) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageControllerDrive) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetStorageItem

`func (o *StorageControllerDrive) GetStorageItem() []StorageItemRelationship`

GetStorageItem returns the StorageItem field if non-nil, zero value otherwise.

### GetStorageItemOk

`func (o *StorageControllerDrive) GetStorageItemOk() (*[]StorageItemRelationship, bool)`

GetStorageItemOk returns a tuple with the StorageItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageItem

`func (o *StorageControllerDrive) SetStorageItem(v []StorageItemRelationship)`

SetStorageItem sets StorageItem field to given value.

### HasStorageItem

`func (o *StorageControllerDrive) HasStorageItem() bool`

HasStorageItem returns a boolean if a field has been set.

### SetStorageItemNil

`func (o *StorageControllerDrive) SetStorageItemNil(b bool)`

 SetStorageItemNil sets the value for StorageItem to be an explicit nil

### UnsetStorageItem
`func (o *StorageControllerDrive) UnsetStorageItem()`

UnsetStorageItem ensures that no value is present for StorageItem, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


