# StorageHitachiExternalStoragePort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.HitachiExternalStoragePort"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.HitachiExternalStoragePort"]
**ExternalIsUsed** | Pointer to **bool** | Is Used of the external storage. | [optional] [readonly] 
**ExternalPathMode** | Pointer to **string** | Path Mode of the external storage. | [optional] [readonly] 
**ExternalPortId** | Pointer to **string** | Object ID of ports on an external storage system. | [optional] [readonly] 
**ExternalSerialNumber** | Pointer to **string** | Serial Number of the external storage. | [optional] [readonly] 
**ExternalStorageInfo** | Pointer to **string** | Storage Information of the external storage. | [optional] [readonly] 
**ExternalWwn** | Pointer to **string** | WWN of the external storage port. | [optional] [readonly] 
**IscsiIpAddress** | Pointer to **string** | The iSCSI IP Address of the external storage port. | [optional] [readonly] 
**IscsiName** | Pointer to **string** | The iSCSI Name of the external storage port. | [optional] [readonly] 
**PortId** | Pointer to **string** | Port ID of the local storage. | [optional] [readonly] 
**VirtualPortId** | Pointer to **int64** | Virtual port ID. This attribute is displayed when an iSCSI port is used and virtual port mode is enabled. | [optional] [readonly] 
**Array** | Pointer to [**NullableStorageHitachiArrayRelationship**](StorageHitachiArrayRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewStorageHitachiExternalStoragePort

`func NewStorageHitachiExternalStoragePort(classId string, objectType string, ) *StorageHitachiExternalStoragePort`

NewStorageHitachiExternalStoragePort instantiates a new StorageHitachiExternalStoragePort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageHitachiExternalStoragePortWithDefaults

`func NewStorageHitachiExternalStoragePortWithDefaults() *StorageHitachiExternalStoragePort`

NewStorageHitachiExternalStoragePortWithDefaults instantiates a new StorageHitachiExternalStoragePort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageHitachiExternalStoragePort) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageHitachiExternalStoragePort) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageHitachiExternalStoragePort) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageHitachiExternalStoragePort) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageHitachiExternalStoragePort) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageHitachiExternalStoragePort) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetExternalIsUsed

`func (o *StorageHitachiExternalStoragePort) GetExternalIsUsed() bool`

GetExternalIsUsed returns the ExternalIsUsed field if non-nil, zero value otherwise.

### GetExternalIsUsedOk

`func (o *StorageHitachiExternalStoragePort) GetExternalIsUsedOk() (*bool, bool)`

GetExternalIsUsedOk returns a tuple with the ExternalIsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIsUsed

`func (o *StorageHitachiExternalStoragePort) SetExternalIsUsed(v bool)`

SetExternalIsUsed sets ExternalIsUsed field to given value.

### HasExternalIsUsed

`func (o *StorageHitachiExternalStoragePort) HasExternalIsUsed() bool`

HasExternalIsUsed returns a boolean if a field has been set.

### GetExternalPathMode

`func (o *StorageHitachiExternalStoragePort) GetExternalPathMode() string`

GetExternalPathMode returns the ExternalPathMode field if non-nil, zero value otherwise.

### GetExternalPathModeOk

`func (o *StorageHitachiExternalStoragePort) GetExternalPathModeOk() (*string, bool)`

GetExternalPathModeOk returns a tuple with the ExternalPathMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPathMode

`func (o *StorageHitachiExternalStoragePort) SetExternalPathMode(v string)`

SetExternalPathMode sets ExternalPathMode field to given value.

### HasExternalPathMode

`func (o *StorageHitachiExternalStoragePort) HasExternalPathMode() bool`

HasExternalPathMode returns a boolean if a field has been set.

### GetExternalPortId

`func (o *StorageHitachiExternalStoragePort) GetExternalPortId() string`

GetExternalPortId returns the ExternalPortId field if non-nil, zero value otherwise.

### GetExternalPortIdOk

`func (o *StorageHitachiExternalStoragePort) GetExternalPortIdOk() (*string, bool)`

GetExternalPortIdOk returns a tuple with the ExternalPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPortId

`func (o *StorageHitachiExternalStoragePort) SetExternalPortId(v string)`

SetExternalPortId sets ExternalPortId field to given value.

### HasExternalPortId

`func (o *StorageHitachiExternalStoragePort) HasExternalPortId() bool`

HasExternalPortId returns a boolean if a field has been set.

### GetExternalSerialNumber

`func (o *StorageHitachiExternalStoragePort) GetExternalSerialNumber() string`

GetExternalSerialNumber returns the ExternalSerialNumber field if non-nil, zero value otherwise.

### GetExternalSerialNumberOk

`func (o *StorageHitachiExternalStoragePort) GetExternalSerialNumberOk() (*string, bool)`

GetExternalSerialNumberOk returns a tuple with the ExternalSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSerialNumber

`func (o *StorageHitachiExternalStoragePort) SetExternalSerialNumber(v string)`

SetExternalSerialNumber sets ExternalSerialNumber field to given value.

### HasExternalSerialNumber

`func (o *StorageHitachiExternalStoragePort) HasExternalSerialNumber() bool`

HasExternalSerialNumber returns a boolean if a field has been set.

### GetExternalStorageInfo

`func (o *StorageHitachiExternalStoragePort) GetExternalStorageInfo() string`

GetExternalStorageInfo returns the ExternalStorageInfo field if non-nil, zero value otherwise.

### GetExternalStorageInfoOk

`func (o *StorageHitachiExternalStoragePort) GetExternalStorageInfoOk() (*string, bool)`

GetExternalStorageInfoOk returns a tuple with the ExternalStorageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalStorageInfo

`func (o *StorageHitachiExternalStoragePort) SetExternalStorageInfo(v string)`

SetExternalStorageInfo sets ExternalStorageInfo field to given value.

### HasExternalStorageInfo

`func (o *StorageHitachiExternalStoragePort) HasExternalStorageInfo() bool`

HasExternalStorageInfo returns a boolean if a field has been set.

### GetExternalWwn

`func (o *StorageHitachiExternalStoragePort) GetExternalWwn() string`

GetExternalWwn returns the ExternalWwn field if non-nil, zero value otherwise.

### GetExternalWwnOk

`func (o *StorageHitachiExternalStoragePort) GetExternalWwnOk() (*string, bool)`

GetExternalWwnOk returns a tuple with the ExternalWwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalWwn

`func (o *StorageHitachiExternalStoragePort) SetExternalWwn(v string)`

SetExternalWwn sets ExternalWwn field to given value.

### HasExternalWwn

`func (o *StorageHitachiExternalStoragePort) HasExternalWwn() bool`

HasExternalWwn returns a boolean if a field has been set.

### GetIscsiIpAddress

`func (o *StorageHitachiExternalStoragePort) GetIscsiIpAddress() string`

GetIscsiIpAddress returns the IscsiIpAddress field if non-nil, zero value otherwise.

### GetIscsiIpAddressOk

`func (o *StorageHitachiExternalStoragePort) GetIscsiIpAddressOk() (*string, bool)`

GetIscsiIpAddressOk returns a tuple with the IscsiIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiIpAddress

`func (o *StorageHitachiExternalStoragePort) SetIscsiIpAddress(v string)`

SetIscsiIpAddress sets IscsiIpAddress field to given value.

### HasIscsiIpAddress

`func (o *StorageHitachiExternalStoragePort) HasIscsiIpAddress() bool`

HasIscsiIpAddress returns a boolean if a field has been set.

### GetIscsiName

`func (o *StorageHitachiExternalStoragePort) GetIscsiName() string`

GetIscsiName returns the IscsiName field if non-nil, zero value otherwise.

### GetIscsiNameOk

`func (o *StorageHitachiExternalStoragePort) GetIscsiNameOk() (*string, bool)`

GetIscsiNameOk returns a tuple with the IscsiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiName

`func (o *StorageHitachiExternalStoragePort) SetIscsiName(v string)`

SetIscsiName sets IscsiName field to given value.

### HasIscsiName

`func (o *StorageHitachiExternalStoragePort) HasIscsiName() bool`

HasIscsiName returns a boolean if a field has been set.

### GetPortId

`func (o *StorageHitachiExternalStoragePort) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *StorageHitachiExternalStoragePort) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *StorageHitachiExternalStoragePort) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *StorageHitachiExternalStoragePort) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetVirtualPortId

`func (o *StorageHitachiExternalStoragePort) GetVirtualPortId() int64`

GetVirtualPortId returns the VirtualPortId field if non-nil, zero value otherwise.

### GetVirtualPortIdOk

`func (o *StorageHitachiExternalStoragePort) GetVirtualPortIdOk() (*int64, bool)`

GetVirtualPortIdOk returns a tuple with the VirtualPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualPortId

`func (o *StorageHitachiExternalStoragePort) SetVirtualPortId(v int64)`

SetVirtualPortId sets VirtualPortId field to given value.

### HasVirtualPortId

`func (o *StorageHitachiExternalStoragePort) HasVirtualPortId() bool`

HasVirtualPortId returns a boolean if a field has been set.

### GetArray

`func (o *StorageHitachiExternalStoragePort) GetArray() StorageHitachiArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StorageHitachiExternalStoragePort) GetArrayOk() (*StorageHitachiArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StorageHitachiExternalStoragePort) SetArray(v StorageHitachiArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StorageHitachiExternalStoragePort) HasArray() bool`

HasArray returns a boolean if a field has been set.

### SetArrayNil

`func (o *StorageHitachiExternalStoragePort) SetArrayNil(b bool)`

 SetArrayNil sets the value for Array to be an explicit nil

### UnsetArray
`func (o *StorageHitachiExternalStoragePort) UnsetArray()`

UnsetArray ensures that no value is present for Array, not even an explicit nil
### GetRegisteredDevice

`func (o *StorageHitachiExternalStoragePort) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageHitachiExternalStoragePort) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageHitachiExternalStoragePort) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageHitachiExternalStoragePort) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *StorageHitachiExternalStoragePort) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *StorageHitachiExternalStoragePort) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


