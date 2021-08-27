# BiosSystemBootOrderAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "bios.SystemBootOrder"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "bios.SystemBootOrder"]
**BootMode** | Pointer to **string** | The BIOS boot mode. UEFI uses the GUID Partition Table (GPT) whereas Legacy mode uses the Master Boot Record (MBR) partitioning scheme. * &#x60;Legacy&#x60; - Legacy mode refers to the traditional process of booting from BIOS. Legacy mode uses the Master Boot Record (MBR) to locate the bootloader. * &#x60;Uefi&#x60; - UEFI mode uses the GUID Partition Table (GPT) to locate EFI Service Partitions to boot from. | [optional] [readonly] [default to "Legacy"]
**Dn** | Pointer to **string** | The Distinguished Name for this object, used to uniquely identify this object. | [optional] [readonly] 
**SecureBoot** | Pointer to **string** | Secure boot if set to enabled, enforces that device boots using only software that is trusted by the Original Equipment Manufacturer (OEM). * &#x60;NotAvailable&#x60; - Set the state of Secure Boot to Not Available. * &#x60;Disabled&#x60; - Set the state of Secure Boot to Disabled. * &#x60;Enabled&#x60; - Set the state of Secure Boot to Enabled. | [optional] [readonly] [default to "NotAvailable"]
**BiosUnit** | Pointer to [**BiosUnitRelationship**](BiosUnitRelationship.md) |  | [optional] 
**BootDevices** | Pointer to [**[]BiosBootDeviceRelationship**](BiosBootDeviceRelationship.md) | An array of relationships to biosBootDevice resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewBiosSystemBootOrderAllOf

`func NewBiosSystemBootOrderAllOf(classId string, objectType string, ) *BiosSystemBootOrderAllOf`

NewBiosSystemBootOrderAllOf instantiates a new BiosSystemBootOrderAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiosSystemBootOrderAllOfWithDefaults

`func NewBiosSystemBootOrderAllOfWithDefaults() *BiosSystemBootOrderAllOf`

NewBiosSystemBootOrderAllOfWithDefaults instantiates a new BiosSystemBootOrderAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BiosSystemBootOrderAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BiosSystemBootOrderAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BiosSystemBootOrderAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BiosSystemBootOrderAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BiosSystemBootOrderAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BiosSystemBootOrderAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBootMode

`func (o *BiosSystemBootOrderAllOf) GetBootMode() string`

GetBootMode returns the BootMode field if non-nil, zero value otherwise.

### GetBootModeOk

`func (o *BiosSystemBootOrderAllOf) GetBootModeOk() (*string, bool)`

GetBootModeOk returns a tuple with the BootMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootMode

`func (o *BiosSystemBootOrderAllOf) SetBootMode(v string)`

SetBootMode sets BootMode field to given value.

### HasBootMode

`func (o *BiosSystemBootOrderAllOf) HasBootMode() bool`

HasBootMode returns a boolean if a field has been set.

### GetDn

`func (o *BiosSystemBootOrderAllOf) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *BiosSystemBootOrderAllOf) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *BiosSystemBootOrderAllOf) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *BiosSystemBootOrderAllOf) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetSecureBoot

`func (o *BiosSystemBootOrderAllOf) GetSecureBoot() string`

GetSecureBoot returns the SecureBoot field if non-nil, zero value otherwise.

### GetSecureBootOk

`func (o *BiosSystemBootOrderAllOf) GetSecureBootOk() (*string, bool)`

GetSecureBootOk returns a tuple with the SecureBoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureBoot

`func (o *BiosSystemBootOrderAllOf) SetSecureBoot(v string)`

SetSecureBoot sets SecureBoot field to given value.

### HasSecureBoot

`func (o *BiosSystemBootOrderAllOf) HasSecureBoot() bool`

HasSecureBoot returns a boolean if a field has been set.

### GetBiosUnit

`func (o *BiosSystemBootOrderAllOf) GetBiosUnit() BiosUnitRelationship`

GetBiosUnit returns the BiosUnit field if non-nil, zero value otherwise.

### GetBiosUnitOk

`func (o *BiosSystemBootOrderAllOf) GetBiosUnitOk() (*BiosUnitRelationship, bool)`

GetBiosUnitOk returns a tuple with the BiosUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosUnit

`func (o *BiosSystemBootOrderAllOf) SetBiosUnit(v BiosUnitRelationship)`

SetBiosUnit sets BiosUnit field to given value.

### HasBiosUnit

`func (o *BiosSystemBootOrderAllOf) HasBiosUnit() bool`

HasBiosUnit returns a boolean if a field has been set.

### GetBootDevices

`func (o *BiosSystemBootOrderAllOf) GetBootDevices() []BiosBootDeviceRelationship`

GetBootDevices returns the BootDevices field if non-nil, zero value otherwise.

### GetBootDevicesOk

`func (o *BiosSystemBootOrderAllOf) GetBootDevicesOk() (*[]BiosBootDeviceRelationship, bool)`

GetBootDevicesOk returns a tuple with the BootDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootDevices

`func (o *BiosSystemBootOrderAllOf) SetBootDevices(v []BiosBootDeviceRelationship)`

SetBootDevices sets BootDevices field to given value.

### HasBootDevices

`func (o *BiosSystemBootOrderAllOf) HasBootDevices() bool`

HasBootDevices returns a boolean if a field has been set.

### SetBootDevicesNil

`func (o *BiosSystemBootOrderAllOf) SetBootDevicesNil(b bool)`

 SetBootDevicesNil sets the value for BootDevices to be an explicit nil

### UnsetBootDevices
`func (o *BiosSystemBootOrderAllOf) UnsetBootDevices()`

UnsetBootDevices ensures that no value is present for BootDevices, not even an explicit nil
### GetRegisteredDevice

`func (o *BiosSystemBootOrderAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *BiosSystemBootOrderAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *BiosSystemBootOrderAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *BiosSystemBootOrderAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


