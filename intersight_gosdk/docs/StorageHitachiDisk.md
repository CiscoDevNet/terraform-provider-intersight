# StorageHitachiDisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.HitachiDisk"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.HitachiDisk"]
**DriveTypeCode** | Pointer to **string** | Drive type code. | [optional] [readonly] 
**ParityGroupId** | Pointer to **string** | Parity group number. When the drive does not belong to any parity group, an empty character string is output. | [optional] [readonly] 
**TypeDetail** | Pointer to **string** | Drive type. * &#x60;N/A&#x60; - Drive Type is not available. * &#x60;SAS&#x60; - SAS. * &#x60;SSD(MLC)&#x60; - SSD (MLC). * &#x60;SSD(FMC)&#x60; - SSD (FMC). * &#x60;SSD(FMD)&#x60; - SSD (FMD). * &#x60;SSD(SLC)&#x60; - SSD (SLC). * &#x60;SSD&#x60; - SSD. * &#x60;SSD(RI)&#x60; - SSD (RI). | [optional] [readonly] [default to "N/A"]
**Usage** | Pointer to **string** | Purpose for which the drive is used. | [optional] [readonly] 
**Array** | Pointer to [**StorageHitachiArrayRelationship**](StorageHitachiArrayRelationship.md) |  | [optional] 
**ParityGroup** | Pointer to [**StorageHitachiParityGroupRelationship**](StorageHitachiParityGroupRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewStorageHitachiDisk

`func NewStorageHitachiDisk(classId string, objectType string, ) *StorageHitachiDisk`

NewStorageHitachiDisk instantiates a new StorageHitachiDisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageHitachiDiskWithDefaults

`func NewStorageHitachiDiskWithDefaults() *StorageHitachiDisk`

NewStorageHitachiDiskWithDefaults instantiates a new StorageHitachiDisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageHitachiDisk) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageHitachiDisk) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageHitachiDisk) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageHitachiDisk) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageHitachiDisk) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageHitachiDisk) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDriveTypeCode

`func (o *StorageHitachiDisk) GetDriveTypeCode() string`

GetDriveTypeCode returns the DriveTypeCode field if non-nil, zero value otherwise.

### GetDriveTypeCodeOk

`func (o *StorageHitachiDisk) GetDriveTypeCodeOk() (*string, bool)`

GetDriveTypeCodeOk returns a tuple with the DriveTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveTypeCode

`func (o *StorageHitachiDisk) SetDriveTypeCode(v string)`

SetDriveTypeCode sets DriveTypeCode field to given value.

### HasDriveTypeCode

`func (o *StorageHitachiDisk) HasDriveTypeCode() bool`

HasDriveTypeCode returns a boolean if a field has been set.

### GetParityGroupId

`func (o *StorageHitachiDisk) GetParityGroupId() string`

GetParityGroupId returns the ParityGroupId field if non-nil, zero value otherwise.

### GetParityGroupIdOk

`func (o *StorageHitachiDisk) GetParityGroupIdOk() (*string, bool)`

GetParityGroupIdOk returns a tuple with the ParityGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParityGroupId

`func (o *StorageHitachiDisk) SetParityGroupId(v string)`

SetParityGroupId sets ParityGroupId field to given value.

### HasParityGroupId

`func (o *StorageHitachiDisk) HasParityGroupId() bool`

HasParityGroupId returns a boolean if a field has been set.

### GetTypeDetail

`func (o *StorageHitachiDisk) GetTypeDetail() string`

GetTypeDetail returns the TypeDetail field if non-nil, zero value otherwise.

### GetTypeDetailOk

`func (o *StorageHitachiDisk) GetTypeDetailOk() (*string, bool)`

GetTypeDetailOk returns a tuple with the TypeDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeDetail

`func (o *StorageHitachiDisk) SetTypeDetail(v string)`

SetTypeDetail sets TypeDetail field to given value.

### HasTypeDetail

`func (o *StorageHitachiDisk) HasTypeDetail() bool`

HasTypeDetail returns a boolean if a field has been set.

### GetUsage

`func (o *StorageHitachiDisk) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *StorageHitachiDisk) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *StorageHitachiDisk) SetUsage(v string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *StorageHitachiDisk) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetArray

`func (o *StorageHitachiDisk) GetArray() StorageHitachiArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StorageHitachiDisk) GetArrayOk() (*StorageHitachiArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StorageHitachiDisk) SetArray(v StorageHitachiArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StorageHitachiDisk) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetParityGroup

`func (o *StorageHitachiDisk) GetParityGroup() StorageHitachiParityGroupRelationship`

GetParityGroup returns the ParityGroup field if non-nil, zero value otherwise.

### GetParityGroupOk

`func (o *StorageHitachiDisk) GetParityGroupOk() (*StorageHitachiParityGroupRelationship, bool)`

GetParityGroupOk returns a tuple with the ParityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParityGroup

`func (o *StorageHitachiDisk) SetParityGroup(v StorageHitachiParityGroupRelationship)`

SetParityGroup sets ParityGroup field to given value.

### HasParityGroup

`func (o *StorageHitachiDisk) HasParityGroup() bool`

HasParityGroup returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageHitachiDisk) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageHitachiDisk) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageHitachiDisk) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageHitachiDisk) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


