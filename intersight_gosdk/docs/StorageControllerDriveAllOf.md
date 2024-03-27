# StorageControllerDriveAllOf

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

### NewStorageControllerDriveAllOf

`func NewStorageControllerDriveAllOf(classId string, objectType string, ) *StorageControllerDriveAllOf`

NewStorageControllerDriveAllOf instantiates a new StorageControllerDriveAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageControllerDriveAllOfWithDefaults

`func NewStorageControllerDriveAllOfWithDefaults() *StorageControllerDriveAllOf`

NewStorageControllerDriveAllOfWithDefaults instantiates a new StorageControllerDriveAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageControllerDriveAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageControllerDriveAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageControllerDriveAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageControllerDriveAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageControllerDriveAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageControllerDriveAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetControllerType

`func (o *StorageControllerDriveAllOf) GetControllerType() string`

GetControllerType returns the ControllerType field if non-nil, zero value otherwise.

### GetControllerTypeOk

`func (o *StorageControllerDriveAllOf) GetControllerTypeOk() (*string, bool)`

GetControllerTypeOk returns a tuple with the ControllerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerType

`func (o *StorageControllerDriveAllOf) SetControllerType(v string)`

SetControllerType sets ControllerType field to given value.

### HasControllerType

`func (o *StorageControllerDriveAllOf) HasControllerType() bool`

HasControllerType returns a boolean if a field has been set.

### GetDescription

`func (o *StorageControllerDriveAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StorageControllerDriveAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StorageControllerDriveAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StorageControllerDriveAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *StorageControllerDriveAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageControllerDriveAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageControllerDriveAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageControllerDriveAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPartitionCount

`func (o *StorageControllerDriveAllOf) GetPartitionCount() int64`

GetPartitionCount returns the PartitionCount field if non-nil, zero value otherwise.

### GetPartitionCountOk

`func (o *StorageControllerDriveAllOf) GetPartitionCountOk() (*int64, bool)`

GetPartitionCountOk returns a tuple with the PartitionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionCount

`func (o *StorageControllerDriveAllOf) SetPartitionCount(v int64)`

SetPartitionCount sets PartitionCount field to given value.

### HasPartitionCount

`func (o *StorageControllerDriveAllOf) HasPartitionCount() bool`

HasPartitionCount returns a boolean if a field has been set.

### GetStorageId

`func (o *StorageControllerDriveAllOf) GetStorageId() string`

GetStorageId returns the StorageId field if non-nil, zero value otherwise.

### GetStorageIdOk

`func (o *StorageControllerDriveAllOf) GetStorageIdOk() (*string, bool)`

GetStorageIdOk returns a tuple with the StorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageId

`func (o *StorageControllerDriveAllOf) SetStorageId(v string)`

SetStorageId sets StorageId field to given value.

### HasStorageId

`func (o *StorageControllerDriveAllOf) HasStorageId() bool`

HasStorageId returns a boolean if a field has been set.

### GetType

`func (o *StorageControllerDriveAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageControllerDriveAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageControllerDriveAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StorageControllerDriveAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetComputeBoard

`func (o *StorageControllerDriveAllOf) GetComputeBoard() ComputeBoardRelationship`

GetComputeBoard returns the ComputeBoard field if non-nil, zero value otherwise.

### GetComputeBoardOk

`func (o *StorageControllerDriveAllOf) GetComputeBoardOk() (*ComputeBoardRelationship, bool)`

GetComputeBoardOk returns a tuple with the ComputeBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBoard

`func (o *StorageControllerDriveAllOf) SetComputeBoard(v ComputeBoardRelationship)`

SetComputeBoard sets ComputeBoard field to given value.

### HasComputeBoard

`func (o *StorageControllerDriveAllOf) HasComputeBoard() bool`

HasComputeBoard returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageControllerDriveAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageControllerDriveAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageControllerDriveAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageControllerDriveAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetStorageItem

`func (o *StorageControllerDriveAllOf) GetStorageItem() []StorageItemRelationship`

GetStorageItem returns the StorageItem field if non-nil, zero value otherwise.

### GetStorageItemOk

`func (o *StorageControllerDriveAllOf) GetStorageItemOk() (*[]StorageItemRelationship, bool)`

GetStorageItemOk returns a tuple with the StorageItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageItem

`func (o *StorageControllerDriveAllOf) SetStorageItem(v []StorageItemRelationship)`

SetStorageItem sets StorageItem field to given value.

### HasStorageItem

`func (o *StorageControllerDriveAllOf) HasStorageItem() bool`

HasStorageItem returns a boolean if a field has been set.

### SetStorageItemNil

`func (o *StorageControllerDriveAllOf) SetStorageItemNil(b bool)`

 SetStorageItemNil sets the value for StorageItem to be an explicit nil

### UnsetStorageItem
`func (o *StorageControllerDriveAllOf) UnsetStorageItem()`

UnsetStorageItem ensures that no value is present for StorageItem, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


