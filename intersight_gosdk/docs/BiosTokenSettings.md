# BiosTokenSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "bios.TokenSettings"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "bios.TokenSettings"]
**IsAssigned** | Pointer to **string** | Value that represents if the BIOS configuration is active. Possible values are \&quot;yes\&quot; and \&quot;no\&quot;. | [optional] [readonly] 
**Serial** | Pointer to **string** | Parent server serial number. | [optional] 
**SettingsMoRn** | Pointer to **string** | BIOS configuration name as found in dn. Possible values are \&quot;ADDDC-Sparing\&quot;, \&quot;Maximum-Performance\&quot;, \&quot;Partial-Mirror-mode-1LM\&quot;, \&quot;Mirror-Mode-1LM\&quot;, \&quot;Mirroring\&quot;, \&quot;Lockstep\&quot;, \&quot;Sparing\&quot;, \&quot;Platform-Default\&quot;. | [optional] [readonly] 
**ComputeBlade** | Pointer to [**ComputeBladeRelationship**](compute.Blade.Relationship.md) |  | [optional] 
**ComputeRackUnit** | Pointer to [**ComputeRackUnitRelationship**](compute.RackUnit.Relationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewBiosTokenSettings

`func NewBiosTokenSettings(classId string, objectType string, ) *BiosTokenSettings`

NewBiosTokenSettings instantiates a new BiosTokenSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiosTokenSettingsWithDefaults

`func NewBiosTokenSettingsWithDefaults() *BiosTokenSettings`

NewBiosTokenSettingsWithDefaults instantiates a new BiosTokenSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BiosTokenSettings) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BiosTokenSettings) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BiosTokenSettings) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BiosTokenSettings) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BiosTokenSettings) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BiosTokenSettings) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIsAssigned

`func (o *BiosTokenSettings) GetIsAssigned() string`

GetIsAssigned returns the IsAssigned field if non-nil, zero value otherwise.

### GetIsAssignedOk

`func (o *BiosTokenSettings) GetIsAssignedOk() (*string, bool)`

GetIsAssignedOk returns a tuple with the IsAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAssigned

`func (o *BiosTokenSettings) SetIsAssigned(v string)`

SetIsAssigned sets IsAssigned field to given value.

### HasIsAssigned

`func (o *BiosTokenSettings) HasIsAssigned() bool`

HasIsAssigned returns a boolean if a field has been set.

### GetSerial

`func (o *BiosTokenSettings) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *BiosTokenSettings) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *BiosTokenSettings) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *BiosTokenSettings) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetSettingsMoRn

`func (o *BiosTokenSettings) GetSettingsMoRn() string`

GetSettingsMoRn returns the SettingsMoRn field if non-nil, zero value otherwise.

### GetSettingsMoRnOk

`func (o *BiosTokenSettings) GetSettingsMoRnOk() (*string, bool)`

GetSettingsMoRnOk returns a tuple with the SettingsMoRn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingsMoRn

`func (o *BiosTokenSettings) SetSettingsMoRn(v string)`

SetSettingsMoRn sets SettingsMoRn field to given value.

### HasSettingsMoRn

`func (o *BiosTokenSettings) HasSettingsMoRn() bool`

HasSettingsMoRn returns a boolean if a field has been set.

### GetComputeBlade

`func (o *BiosTokenSettings) GetComputeBlade() ComputeBladeRelationship`

GetComputeBlade returns the ComputeBlade field if non-nil, zero value otherwise.

### GetComputeBladeOk

`func (o *BiosTokenSettings) GetComputeBladeOk() (*ComputeBladeRelationship, bool)`

GetComputeBladeOk returns a tuple with the ComputeBlade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBlade

`func (o *BiosTokenSettings) SetComputeBlade(v ComputeBladeRelationship)`

SetComputeBlade sets ComputeBlade field to given value.

### HasComputeBlade

`func (o *BiosTokenSettings) HasComputeBlade() bool`

HasComputeBlade returns a boolean if a field has been set.

### GetComputeRackUnit

`func (o *BiosTokenSettings) GetComputeRackUnit() ComputeRackUnitRelationship`

GetComputeRackUnit returns the ComputeRackUnit field if non-nil, zero value otherwise.

### GetComputeRackUnitOk

`func (o *BiosTokenSettings) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool)`

GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRackUnit

`func (o *BiosTokenSettings) SetComputeRackUnit(v ComputeRackUnitRelationship)`

SetComputeRackUnit sets ComputeRackUnit field to given value.

### HasComputeRackUnit

`func (o *BiosTokenSettings) HasComputeRackUnit() bool`

HasComputeRackUnit returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *BiosTokenSettings) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *BiosTokenSettings) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *BiosTokenSettings) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *BiosTokenSettings) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *BiosTokenSettings) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *BiosTokenSettings) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *BiosTokenSettings) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *BiosTokenSettings) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


