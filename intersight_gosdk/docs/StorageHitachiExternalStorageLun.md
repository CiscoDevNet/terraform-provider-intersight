# StorageHitachiExternalStorageLun

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
**Array** | Pointer to [**NullableStorageHitachiArrayRelationship**](StorageHitachiArrayRelationship.md) |  | [optional] 
**ExternalStoragePort** | Pointer to [**NullableStorageHitachiExternalStoragePortRelationship**](StorageHitachiExternalStoragePortRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewStorageHitachiExternalStorageLun

`func NewStorageHitachiExternalStorageLun(classId string, objectType string, ) *StorageHitachiExternalStorageLun`

NewStorageHitachiExternalStorageLun instantiates a new StorageHitachiExternalStorageLun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageHitachiExternalStorageLunWithDefaults

`func NewStorageHitachiExternalStorageLunWithDefaults() *StorageHitachiExternalStorageLun`

NewStorageHitachiExternalStorageLunWithDefaults instantiates a new StorageHitachiExternalStorageLun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageHitachiExternalStorageLun) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageHitachiExternalStorageLun) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageHitachiExternalStorageLun) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageHitachiExternalStorageLun) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageHitachiExternalStorageLun) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageHitachiExternalStorageLun) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetExternalLun

`func (o *StorageHitachiExternalStorageLun) GetExternalLun() int64`

GetExternalLun returns the ExternalLun field if non-nil, zero value otherwise.

### GetExternalLunOk

`func (o *StorageHitachiExternalStorageLun) GetExternalLunOk() (*int64, bool)`

GetExternalLunOk returns a tuple with the ExternalLun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLun

`func (o *StorageHitachiExternalStorageLun) SetExternalLun(v int64)`

SetExternalLun sets ExternalLun field to given value.

### HasExternalLun

`func (o *StorageHitachiExternalStorageLun) HasExternalLun() bool`

HasExternalLun returns a boolean if a field has been set.

### GetExternalPortId

`func (o *StorageHitachiExternalStorageLun) GetExternalPortId() string`

GetExternalPortId returns the ExternalPortId field if non-nil, zero value otherwise.

### GetExternalPortIdOk

`func (o *StorageHitachiExternalStorageLun) GetExternalPortIdOk() (*string, bool)`

GetExternalPortIdOk returns a tuple with the ExternalPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPortId

`func (o *StorageHitachiExternalStorageLun) SetExternalPortId(v string)`

SetExternalPortId sets ExternalPortId field to given value.

### HasExternalPortId

`func (o *StorageHitachiExternalStorageLun) HasExternalPortId() bool`

HasExternalPortId returns a boolean if a field has been set.

### GetExternalVolumeCapacity

`func (o *StorageHitachiExternalStorageLun) GetExternalVolumeCapacity() int64`

GetExternalVolumeCapacity returns the ExternalVolumeCapacity field if non-nil, zero value otherwise.

### GetExternalVolumeCapacityOk

`func (o *StorageHitachiExternalStorageLun) GetExternalVolumeCapacityOk() (*int64, bool)`

GetExternalVolumeCapacityOk returns a tuple with the ExternalVolumeCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalVolumeCapacity

`func (o *StorageHitachiExternalStorageLun) SetExternalVolumeCapacity(v int64)`

SetExternalVolumeCapacity sets ExternalVolumeCapacity field to given value.

### HasExternalVolumeCapacity

`func (o *StorageHitachiExternalStorageLun) HasExternalVolumeCapacity() bool`

HasExternalVolumeCapacity returns a boolean if a field has been set.

### GetExternalVolumeInfo

`func (o *StorageHitachiExternalStorageLun) GetExternalVolumeInfo() string`

GetExternalVolumeInfo returns the ExternalVolumeInfo field if non-nil, zero value otherwise.

### GetExternalVolumeInfoOk

`func (o *StorageHitachiExternalStorageLun) GetExternalVolumeInfoOk() (*string, bool)`

GetExternalVolumeInfoOk returns a tuple with the ExternalVolumeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalVolumeInfo

`func (o *StorageHitachiExternalStorageLun) SetExternalVolumeInfo(v string)`

SetExternalVolumeInfo sets ExternalVolumeInfo field to given value.

### HasExternalVolumeInfo

`func (o *StorageHitachiExternalStorageLun) HasExternalVolumeInfo() bool`

HasExternalVolumeInfo returns a boolean if a field has been set.

### GetExternalWwn

`func (o *StorageHitachiExternalStorageLun) GetExternalWwn() string`

GetExternalWwn returns the ExternalWwn field if non-nil, zero value otherwise.

### GetExternalWwnOk

`func (o *StorageHitachiExternalStorageLun) GetExternalWwnOk() (*string, bool)`

GetExternalWwnOk returns a tuple with the ExternalWwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalWwn

`func (o *StorageHitachiExternalStorageLun) SetExternalWwn(v string)`

SetExternalWwn sets ExternalWwn field to given value.

### HasExternalWwn

`func (o *StorageHitachiExternalStorageLun) HasExternalWwn() bool`

HasExternalWwn returns a boolean if a field has been set.

### GetIscsiIpAddress

`func (o *StorageHitachiExternalStorageLun) GetIscsiIpAddress() string`

GetIscsiIpAddress returns the IscsiIpAddress field if non-nil, zero value otherwise.

### GetIscsiIpAddressOk

`func (o *StorageHitachiExternalStorageLun) GetIscsiIpAddressOk() (*string, bool)`

GetIscsiIpAddressOk returns a tuple with the IscsiIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiIpAddress

`func (o *StorageHitachiExternalStorageLun) SetIscsiIpAddress(v string)`

SetIscsiIpAddress sets IscsiIpAddress field to given value.

### HasIscsiIpAddress

`func (o *StorageHitachiExternalStorageLun) HasIscsiIpAddress() bool`

HasIscsiIpAddress returns a boolean if a field has been set.

### GetIscsiName

`func (o *StorageHitachiExternalStorageLun) GetIscsiName() string`

GetIscsiName returns the IscsiName field if non-nil, zero value otherwise.

### GetIscsiNameOk

`func (o *StorageHitachiExternalStorageLun) GetIscsiNameOk() (*string, bool)`

GetIscsiNameOk returns a tuple with the IscsiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiName

`func (o *StorageHitachiExternalStorageLun) SetIscsiName(v string)`

SetIscsiName sets IscsiName field to given value.

### HasIscsiName

`func (o *StorageHitachiExternalStorageLun) HasIscsiName() bool`

HasIscsiName returns a boolean if a field has been set.

### GetPortId

`func (o *StorageHitachiExternalStorageLun) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *StorageHitachiExternalStorageLun) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *StorageHitachiExternalStorageLun) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *StorageHitachiExternalStorageLun) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetVirtualPortId

`func (o *StorageHitachiExternalStorageLun) GetVirtualPortId() int64`

GetVirtualPortId returns the VirtualPortId field if non-nil, zero value otherwise.

### GetVirtualPortIdOk

`func (o *StorageHitachiExternalStorageLun) GetVirtualPortIdOk() (*int64, bool)`

GetVirtualPortIdOk returns a tuple with the VirtualPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualPortId

`func (o *StorageHitachiExternalStorageLun) SetVirtualPortId(v int64)`

SetVirtualPortId sets VirtualPortId field to given value.

### HasVirtualPortId

`func (o *StorageHitachiExternalStorageLun) HasVirtualPortId() bool`

HasVirtualPortId returns a boolean if a field has been set.

### GetArray

`func (o *StorageHitachiExternalStorageLun) GetArray() StorageHitachiArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StorageHitachiExternalStorageLun) GetArrayOk() (*StorageHitachiArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StorageHitachiExternalStorageLun) SetArray(v StorageHitachiArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StorageHitachiExternalStorageLun) HasArray() bool`

HasArray returns a boolean if a field has been set.

### SetArrayNil

`func (o *StorageHitachiExternalStorageLun) SetArrayNil(b bool)`

 SetArrayNil sets the value for Array to be an explicit nil

### UnsetArray
`func (o *StorageHitachiExternalStorageLun) UnsetArray()`

UnsetArray ensures that no value is present for Array, not even an explicit nil
### GetExternalStoragePort

`func (o *StorageHitachiExternalStorageLun) GetExternalStoragePort() StorageHitachiExternalStoragePortRelationship`

GetExternalStoragePort returns the ExternalStoragePort field if non-nil, zero value otherwise.

### GetExternalStoragePortOk

`func (o *StorageHitachiExternalStorageLun) GetExternalStoragePortOk() (*StorageHitachiExternalStoragePortRelationship, bool)`

GetExternalStoragePortOk returns a tuple with the ExternalStoragePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalStoragePort

`func (o *StorageHitachiExternalStorageLun) SetExternalStoragePort(v StorageHitachiExternalStoragePortRelationship)`

SetExternalStoragePort sets ExternalStoragePort field to given value.

### HasExternalStoragePort

`func (o *StorageHitachiExternalStorageLun) HasExternalStoragePort() bool`

HasExternalStoragePort returns a boolean if a field has been set.

### SetExternalStoragePortNil

`func (o *StorageHitachiExternalStorageLun) SetExternalStoragePortNil(b bool)`

 SetExternalStoragePortNil sets the value for ExternalStoragePort to be an explicit nil

### UnsetExternalStoragePort
`func (o *StorageHitachiExternalStorageLun) UnsetExternalStoragePort()`

UnsetExternalStoragePort ensures that no value is present for ExternalStoragePort, not even an explicit nil
### GetRegisteredDevice

`func (o *StorageHitachiExternalStorageLun) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageHitachiExternalStorageLun) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageHitachiExternalStorageLun) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageHitachiExternalStorageLun) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *StorageHitachiExternalStorageLun) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *StorageHitachiExternalStorageLun) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


