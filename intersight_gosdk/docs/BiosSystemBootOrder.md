# BiosSystemBootOrder

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

### NewBiosSystemBootOrder

`func NewBiosSystemBootOrder(classId string, objectType string, ) *BiosSystemBootOrder`

NewBiosSystemBootOrder instantiates a new BiosSystemBootOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiosSystemBootOrderWithDefaults

`func NewBiosSystemBootOrderWithDefaults() *BiosSystemBootOrder`

NewBiosSystemBootOrderWithDefaults instantiates a new BiosSystemBootOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BiosSystemBootOrder) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BiosSystemBootOrder) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BiosSystemBootOrder) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BiosSystemBootOrder) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BiosSystemBootOrder) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BiosSystemBootOrder) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBootMode

`func (o *BiosSystemBootOrder) GetBootMode() string`

GetBootMode returns the BootMode field if non-nil, zero value otherwise.

### GetBootModeOk

`func (o *BiosSystemBootOrder) GetBootModeOk() (*string, bool)`

GetBootModeOk returns a tuple with the BootMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootMode

`func (o *BiosSystemBootOrder) SetBootMode(v string)`

SetBootMode sets BootMode field to given value.

### HasBootMode

`func (o *BiosSystemBootOrder) HasBootMode() bool`

HasBootMode returns a boolean if a field has been set.

### GetDn

`func (o *BiosSystemBootOrder) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *BiosSystemBootOrder) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *BiosSystemBootOrder) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *BiosSystemBootOrder) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetSecureBoot

`func (o *BiosSystemBootOrder) GetSecureBoot() string`

GetSecureBoot returns the SecureBoot field if non-nil, zero value otherwise.

### GetSecureBootOk

`func (o *BiosSystemBootOrder) GetSecureBootOk() (*string, bool)`

GetSecureBootOk returns a tuple with the SecureBoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureBoot

`func (o *BiosSystemBootOrder) SetSecureBoot(v string)`

SetSecureBoot sets SecureBoot field to given value.

### HasSecureBoot

`func (o *BiosSystemBootOrder) HasSecureBoot() bool`

HasSecureBoot returns a boolean if a field has been set.

### GetBiosUnit

`func (o *BiosSystemBootOrder) GetBiosUnit() BiosUnitRelationship`

GetBiosUnit returns the BiosUnit field if non-nil, zero value otherwise.

### GetBiosUnitOk

`func (o *BiosSystemBootOrder) GetBiosUnitOk() (*BiosUnitRelationship, bool)`

GetBiosUnitOk returns a tuple with the BiosUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosUnit

`func (o *BiosSystemBootOrder) SetBiosUnit(v BiosUnitRelationship)`

SetBiosUnit sets BiosUnit field to given value.

### HasBiosUnit

`func (o *BiosSystemBootOrder) HasBiosUnit() bool`

HasBiosUnit returns a boolean if a field has been set.

### GetBootDevices

`func (o *BiosSystemBootOrder) GetBootDevices() []BiosBootDeviceRelationship`

GetBootDevices returns the BootDevices field if non-nil, zero value otherwise.

### GetBootDevicesOk

`func (o *BiosSystemBootOrder) GetBootDevicesOk() (*[]BiosBootDeviceRelationship, bool)`

GetBootDevicesOk returns a tuple with the BootDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootDevices

`func (o *BiosSystemBootOrder) SetBootDevices(v []BiosBootDeviceRelationship)`

SetBootDevices sets BootDevices field to given value.

### HasBootDevices

`func (o *BiosSystemBootOrder) HasBootDevices() bool`

HasBootDevices returns a boolean if a field has been set.

### SetBootDevicesNil

`func (o *BiosSystemBootOrder) SetBootDevicesNil(b bool)`

 SetBootDevicesNil sets the value for BootDevices to be an explicit nil

### UnsetBootDevices
`func (o *BiosSystemBootOrder) UnsetBootDevices()`

UnsetBootDevices ensures that no value is present for BootDevices, not even an explicit nil
### GetRegisteredDevice

`func (o *BiosSystemBootOrder) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *BiosSystemBootOrder) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *BiosSystemBootOrder) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *BiosSystemBootOrder) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


