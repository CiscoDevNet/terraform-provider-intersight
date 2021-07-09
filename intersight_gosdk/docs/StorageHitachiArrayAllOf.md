# StorageHitachiArrayAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.HitachiArray"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.HitachiArray"]
**Ctl1Ip** | Pointer to **string** | IP address of controller 1 of the storage system. | [optional] [readonly] 
**Ctl1MicroVersion** | Pointer to **string** | GUM (Gateway for Unified Management) version of the controller 1. | [optional] [readonly] 
**Ctl2Ip** | Pointer to **string** | IP address of controller 2 of the storage system. | [optional] [readonly] 
**Ctl2MicroVersion** | Pointer to **string** | GUM (Gateway for Unified Management) version of the controller 2. | [optional] [readonly] 
**DeviceId** | Pointer to **string** | ID of the Storage device. | [optional] [readonly] 
**SvpIp** | Pointer to **string** | IP address of the SVP (Service Processor). The SVP provides out‑of‑band configuration and management of the storage system, and collects performance data for key components to enable diagnostic testing and analysis. | [optional] [readonly] 
**TargetCtl** | Pointer to **string** | Controller operated by the REST API. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewStorageHitachiArrayAllOf

`func NewStorageHitachiArrayAllOf(classId string, objectType string, ) *StorageHitachiArrayAllOf`

NewStorageHitachiArrayAllOf instantiates a new StorageHitachiArrayAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageHitachiArrayAllOfWithDefaults

`func NewStorageHitachiArrayAllOfWithDefaults() *StorageHitachiArrayAllOf`

NewStorageHitachiArrayAllOfWithDefaults instantiates a new StorageHitachiArrayAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageHitachiArrayAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageHitachiArrayAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageHitachiArrayAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageHitachiArrayAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageHitachiArrayAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageHitachiArrayAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCtl1Ip

`func (o *StorageHitachiArrayAllOf) GetCtl1Ip() string`

GetCtl1Ip returns the Ctl1Ip field if non-nil, zero value otherwise.

### GetCtl1IpOk

`func (o *StorageHitachiArrayAllOf) GetCtl1IpOk() (*string, bool)`

GetCtl1IpOk returns a tuple with the Ctl1Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtl1Ip

`func (o *StorageHitachiArrayAllOf) SetCtl1Ip(v string)`

SetCtl1Ip sets Ctl1Ip field to given value.

### HasCtl1Ip

`func (o *StorageHitachiArrayAllOf) HasCtl1Ip() bool`

HasCtl1Ip returns a boolean if a field has been set.

### GetCtl1MicroVersion

`func (o *StorageHitachiArrayAllOf) GetCtl1MicroVersion() string`

GetCtl1MicroVersion returns the Ctl1MicroVersion field if non-nil, zero value otherwise.

### GetCtl1MicroVersionOk

`func (o *StorageHitachiArrayAllOf) GetCtl1MicroVersionOk() (*string, bool)`

GetCtl1MicroVersionOk returns a tuple with the Ctl1MicroVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtl1MicroVersion

`func (o *StorageHitachiArrayAllOf) SetCtl1MicroVersion(v string)`

SetCtl1MicroVersion sets Ctl1MicroVersion field to given value.

### HasCtl1MicroVersion

`func (o *StorageHitachiArrayAllOf) HasCtl1MicroVersion() bool`

HasCtl1MicroVersion returns a boolean if a field has been set.

### GetCtl2Ip

`func (o *StorageHitachiArrayAllOf) GetCtl2Ip() string`

GetCtl2Ip returns the Ctl2Ip field if non-nil, zero value otherwise.

### GetCtl2IpOk

`func (o *StorageHitachiArrayAllOf) GetCtl2IpOk() (*string, bool)`

GetCtl2IpOk returns a tuple with the Ctl2Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtl2Ip

`func (o *StorageHitachiArrayAllOf) SetCtl2Ip(v string)`

SetCtl2Ip sets Ctl2Ip field to given value.

### HasCtl2Ip

`func (o *StorageHitachiArrayAllOf) HasCtl2Ip() bool`

HasCtl2Ip returns a boolean if a field has been set.

### GetCtl2MicroVersion

`func (o *StorageHitachiArrayAllOf) GetCtl2MicroVersion() string`

GetCtl2MicroVersion returns the Ctl2MicroVersion field if non-nil, zero value otherwise.

### GetCtl2MicroVersionOk

`func (o *StorageHitachiArrayAllOf) GetCtl2MicroVersionOk() (*string, bool)`

GetCtl2MicroVersionOk returns a tuple with the Ctl2MicroVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtl2MicroVersion

`func (o *StorageHitachiArrayAllOf) SetCtl2MicroVersion(v string)`

SetCtl2MicroVersion sets Ctl2MicroVersion field to given value.

### HasCtl2MicroVersion

`func (o *StorageHitachiArrayAllOf) HasCtl2MicroVersion() bool`

HasCtl2MicroVersion returns a boolean if a field has been set.

### GetDeviceId

`func (o *StorageHitachiArrayAllOf) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *StorageHitachiArrayAllOf) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *StorageHitachiArrayAllOf) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *StorageHitachiArrayAllOf) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetSvpIp

`func (o *StorageHitachiArrayAllOf) GetSvpIp() string`

GetSvpIp returns the SvpIp field if non-nil, zero value otherwise.

### GetSvpIpOk

`func (o *StorageHitachiArrayAllOf) GetSvpIpOk() (*string, bool)`

GetSvpIpOk returns a tuple with the SvpIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvpIp

`func (o *StorageHitachiArrayAllOf) SetSvpIp(v string)`

SetSvpIp sets SvpIp field to given value.

### HasSvpIp

`func (o *StorageHitachiArrayAllOf) HasSvpIp() bool`

HasSvpIp returns a boolean if a field has been set.

### GetTargetCtl

`func (o *StorageHitachiArrayAllOf) GetTargetCtl() string`

GetTargetCtl returns the TargetCtl field if non-nil, zero value otherwise.

### GetTargetCtlOk

`func (o *StorageHitachiArrayAllOf) GetTargetCtlOk() (*string, bool)`

GetTargetCtlOk returns a tuple with the TargetCtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetCtl

`func (o *StorageHitachiArrayAllOf) SetTargetCtl(v string)`

SetTargetCtl sets TargetCtl field to given value.

### HasTargetCtl

`func (o *StorageHitachiArrayAllOf) HasTargetCtl() bool`

HasTargetCtl returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageHitachiArrayAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageHitachiArrayAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageHitachiArrayAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageHitachiArrayAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


