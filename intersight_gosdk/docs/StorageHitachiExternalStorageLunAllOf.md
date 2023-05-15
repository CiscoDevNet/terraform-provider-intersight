# StorageHitachiExternalStorageLunAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.HitachiExternalStorageLun"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.HitachiExternalStorageLun"]
**ExternalLun** | Pointer to **int64** | LUN that can be referenced from the port on the external storage system. | [optional] [readonly] 
**ExternalPortId** | Pointer to **string** | Object ID of ports on an external storage system. | [optional] [readonly] 
**ExternalVolumeCapacity** | Pointer to **int64** | Capacity of the external volume on the external storage system (1 block &#x3D; 512 bytes). | [optional] [readonly] 
**ExternalVolumeInfo** | Pointer to **string** | The product ID and the device identification (output in ASCII format) in the SCSI information for the external volume on the external storage system. This information is obtained in a format in which the product ID and the device identification are concatenated by a space. | [optional] [readonly] 
**ExternalWwn** | Pointer to **string** | WWN of the external storage port. | [optional] [readonly] 
**IscsiIpAddress** | Pointer to **string** | The iSCSI IP Address of the external storage port. | [optional] [readonly] 
**IscsiName** | Pointer to **string** | The iSCSI Name of the external storage port. | [optional] [readonly] 
**PortId** | Pointer to **string** | Port ID of the local storage. | [optional] [readonly] 
**VirtualPortId** | Pointer to **int64** | Virtual port ID. This attribute is displayed when an iSCSI port is used and virtual port mode is enabled. | [optional] [readonly] 
**Array** | Pointer to [**StorageHitachiArrayRelationship**](StorageHitachiArrayRelationship.md) |  | [optional] 
**ExternalStoragePort** | Pointer to [**StorageHitachiExternalStoragePortRelationship**](StorageHitachiExternalStoragePortRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewStorageHitachiExternalStorageLunAllOf

`func NewStorageHitachiExternalStorageLunAllOf(classId string, objectType string, ) *StorageHitachiExternalStorageLunAllOf`

NewStorageHitachiExternalStorageLunAllOf instantiates a new StorageHitachiExternalStorageLunAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageHitachiExternalStorageLunAllOfWithDefaults

`func NewStorageHitachiExternalStorageLunAllOfWithDefaults() *StorageHitachiExternalStorageLunAllOf`

NewStorageHitachiExternalStorageLunAllOfWithDefaults instantiates a new StorageHitachiExternalStorageLunAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageHitachiExternalStorageLunAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageHitachiExternalStorageLunAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageHitachiExternalStorageLunAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageHitachiExternalStorageLunAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageHitachiExternalStorageLunAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageHitachiExternalStorageLunAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetExternalLun

`func (o *StorageHitachiExternalStorageLunAllOf) GetExternalLun() int64`

GetExternalLun returns the ExternalLun field if non-nil, zero value otherwise.

### GetExternalLunOk

`func (o *StorageHitachiExternalStorageLunAllOf) GetExternalLunOk() (*int64, bool)`

GetExternalLunOk returns a tuple with the ExternalLun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLun

`func (o *StorageHitachiExternalStorageLunAllOf) SetExternalLun(v int64)`

SetExternalLun sets ExternalLun field to given value.

### HasExternalLun

`func (o *StorageHitachiExternalStorageLunAllOf) HasExternalLun() bool`

HasExternalLun returns a boolean if a field has been set.

### GetExternalPortId

`func (o *StorageHitachiExternalStorageLunAllOf) GetExternalPortId() string`

GetExternalPortId returns the ExternalPortId field if non-nil, zero value otherwise.

### GetExternalPortIdOk

`func (o *StorageHitachiExternalStorageLunAllOf) GetExternalPortIdOk() (*string, bool)`

GetExternalPortIdOk returns a tuple with the ExternalPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPortId

`func (o *StorageHitachiExternalStorageLunAllOf) SetExternalPortId(v string)`

SetExternalPortId sets ExternalPortId field to given value.

### HasExternalPortId

`func (o *StorageHitachiExternalStorageLunAllOf) HasExternalPortId() bool`

HasExternalPortId returns a boolean if a field has been set.

### GetExternalVolumeCapacity

`func (o *StorageHitachiExternalStorageLunAllOf) GetExternalVolumeCapacity() int64`

GetExternalVolumeCapacity returns the ExternalVolumeCapacity field if non-nil, zero value otherwise.

### GetExternalVolumeCapacityOk

`func (o *StorageHitachiExternalStorageLunAllOf) GetExternalVolumeCapacityOk() (*int64, bool)`

GetExternalVolumeCapacityOk returns a tuple with the ExternalVolumeCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalVolumeCapacity

`func (o *StorageHitachiExternalStorageLunAllOf) SetExternalVolumeCapacity(v int64)`

SetExternalVolumeCapacity sets ExternalVolumeCapacity field to given value.

### HasExternalVolumeCapacity

`func (o *StorageHitachiExternalStorageLunAllOf) HasExternalVolumeCapacity() bool`

HasExternalVolumeCapacity returns a boolean if a field has been set.

### GetExternalVolumeInfo

`func (o *StorageHitachiExternalStorageLunAllOf) GetExternalVolumeInfo() string`

GetExternalVolumeInfo returns the ExternalVolumeInfo field if non-nil, zero value otherwise.

### GetExternalVolumeInfoOk

`func (o *StorageHitachiExternalStorageLunAllOf) GetExternalVolumeInfoOk() (*string, bool)`

GetExternalVolumeInfoOk returns a tuple with the ExternalVolumeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalVolumeInfo

`func (o *StorageHitachiExternalStorageLunAllOf) SetExternalVolumeInfo(v string)`

SetExternalVolumeInfo sets ExternalVolumeInfo field to given value.

### HasExternalVolumeInfo

`func (o *StorageHitachiExternalStorageLunAllOf) HasExternalVolumeInfo() bool`

HasExternalVolumeInfo returns a boolean if a field has been set.

### GetExternalWwn

`func (o *StorageHitachiExternalStorageLunAllOf) GetExternalWwn() string`

GetExternalWwn returns the ExternalWwn field if non-nil, zero value otherwise.

### GetExternalWwnOk

`func (o *StorageHitachiExternalStorageLunAllOf) GetExternalWwnOk() (*string, bool)`

GetExternalWwnOk returns a tuple with the ExternalWwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalWwn

`func (o *StorageHitachiExternalStorageLunAllOf) SetExternalWwn(v string)`

SetExternalWwn sets ExternalWwn field to given value.

### HasExternalWwn

`func (o *StorageHitachiExternalStorageLunAllOf) HasExternalWwn() bool`

HasExternalWwn returns a boolean if a field has been set.

### GetIscsiIpAddress

`func (o *StorageHitachiExternalStorageLunAllOf) GetIscsiIpAddress() string`

GetIscsiIpAddress returns the IscsiIpAddress field if non-nil, zero value otherwise.

### GetIscsiIpAddressOk

`func (o *StorageHitachiExternalStorageLunAllOf) GetIscsiIpAddressOk() (*string, bool)`

GetIscsiIpAddressOk returns a tuple with the IscsiIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiIpAddress

`func (o *StorageHitachiExternalStorageLunAllOf) SetIscsiIpAddress(v string)`

SetIscsiIpAddress sets IscsiIpAddress field to given value.

### HasIscsiIpAddress

`func (o *StorageHitachiExternalStorageLunAllOf) HasIscsiIpAddress() bool`

HasIscsiIpAddress returns a boolean if a field has been set.

### GetIscsiName

`func (o *StorageHitachiExternalStorageLunAllOf) GetIscsiName() string`

GetIscsiName returns the IscsiName field if non-nil, zero value otherwise.

### GetIscsiNameOk

`func (o *StorageHitachiExternalStorageLunAllOf) GetIscsiNameOk() (*string, bool)`

GetIscsiNameOk returns a tuple with the IscsiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiName

`func (o *StorageHitachiExternalStorageLunAllOf) SetIscsiName(v string)`

SetIscsiName sets IscsiName field to given value.

### HasIscsiName

`func (o *StorageHitachiExternalStorageLunAllOf) HasIscsiName() bool`

HasIscsiName returns a boolean if a field has been set.

### GetPortId

`func (o *StorageHitachiExternalStorageLunAllOf) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *StorageHitachiExternalStorageLunAllOf) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *StorageHitachiExternalStorageLunAllOf) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *StorageHitachiExternalStorageLunAllOf) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetVirtualPortId

`func (o *StorageHitachiExternalStorageLunAllOf) GetVirtualPortId() int64`

GetVirtualPortId returns the VirtualPortId field if non-nil, zero value otherwise.

### GetVirtualPortIdOk

`func (o *StorageHitachiExternalStorageLunAllOf) GetVirtualPortIdOk() (*int64, bool)`

GetVirtualPortIdOk returns a tuple with the VirtualPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualPortId

`func (o *StorageHitachiExternalStorageLunAllOf) SetVirtualPortId(v int64)`

SetVirtualPortId sets VirtualPortId field to given value.

### HasVirtualPortId

`func (o *StorageHitachiExternalStorageLunAllOf) HasVirtualPortId() bool`

HasVirtualPortId returns a boolean if a field has been set.

### GetArray

`func (o *StorageHitachiExternalStorageLunAllOf) GetArray() StorageHitachiArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StorageHitachiExternalStorageLunAllOf) GetArrayOk() (*StorageHitachiArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StorageHitachiExternalStorageLunAllOf) SetArray(v StorageHitachiArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StorageHitachiExternalStorageLunAllOf) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetExternalStoragePort

`func (o *StorageHitachiExternalStorageLunAllOf) GetExternalStoragePort() StorageHitachiExternalStoragePortRelationship`

GetExternalStoragePort returns the ExternalStoragePort field if non-nil, zero value otherwise.

### GetExternalStoragePortOk

`func (o *StorageHitachiExternalStorageLunAllOf) GetExternalStoragePortOk() (*StorageHitachiExternalStoragePortRelationship, bool)`

GetExternalStoragePortOk returns a tuple with the ExternalStoragePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalStoragePort

`func (o *StorageHitachiExternalStorageLunAllOf) SetExternalStoragePort(v StorageHitachiExternalStoragePortRelationship)`

SetExternalStoragePort sets ExternalStoragePort field to given value.

### HasExternalStoragePort

`func (o *StorageHitachiExternalStorageLunAllOf) HasExternalStoragePort() bool`

HasExternalStoragePort returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageHitachiExternalStorageLunAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageHitachiExternalStorageLunAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageHitachiExternalStorageLunAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageHitachiExternalStorageLunAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


