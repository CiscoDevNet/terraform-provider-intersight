# BiosVfSelectMemoryRasConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "bios.VfSelectMemoryRasConfiguration"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "bios.VfSelectMemoryRasConfiguration"]
**Serial** | Pointer to **string** | Parent server serial number. | [optional] 
**VpSelectMemoryRasConfiguration** | Pointer to **string** | The actual BIOS memory RAS configuration as reported by the platform BIOS. Possible values are \&quot;maximum-performance\&quot;, \&quot;mirror-mode-1lm\&quot;, \&quot;adddc-sparing\&quot;, \&quot;platform-default\&quot;, \&quot;lockstep\&quot;, \&quot;sparing\&quot;, \&quot;mirroring\&quot;. | [optional] [readonly] 
**ComputeBlade** | Pointer to [**ComputeBladeRelationship**](compute.Blade.Relationship.md) |  | [optional] 
**ComputeRackUnit** | Pointer to [**ComputeRackUnitRelationship**](compute.RackUnit.Relationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewBiosVfSelectMemoryRasConfiguration

`func NewBiosVfSelectMemoryRasConfiguration(classId string, objectType string, ) *BiosVfSelectMemoryRasConfiguration`

NewBiosVfSelectMemoryRasConfiguration instantiates a new BiosVfSelectMemoryRasConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiosVfSelectMemoryRasConfigurationWithDefaults

`func NewBiosVfSelectMemoryRasConfigurationWithDefaults() *BiosVfSelectMemoryRasConfiguration`

NewBiosVfSelectMemoryRasConfigurationWithDefaults instantiates a new BiosVfSelectMemoryRasConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BiosVfSelectMemoryRasConfiguration) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BiosVfSelectMemoryRasConfiguration) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BiosVfSelectMemoryRasConfiguration) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BiosVfSelectMemoryRasConfiguration) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BiosVfSelectMemoryRasConfiguration) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BiosVfSelectMemoryRasConfiguration) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSerial

`func (o *BiosVfSelectMemoryRasConfiguration) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *BiosVfSelectMemoryRasConfiguration) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *BiosVfSelectMemoryRasConfiguration) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *BiosVfSelectMemoryRasConfiguration) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetVpSelectMemoryRasConfiguration

`func (o *BiosVfSelectMemoryRasConfiguration) GetVpSelectMemoryRasConfiguration() string`

GetVpSelectMemoryRasConfiguration returns the VpSelectMemoryRasConfiguration field if non-nil, zero value otherwise.

### GetVpSelectMemoryRasConfigurationOk

`func (o *BiosVfSelectMemoryRasConfiguration) GetVpSelectMemoryRasConfigurationOk() (*string, bool)`

GetVpSelectMemoryRasConfigurationOk returns a tuple with the VpSelectMemoryRasConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpSelectMemoryRasConfiguration

`func (o *BiosVfSelectMemoryRasConfiguration) SetVpSelectMemoryRasConfiguration(v string)`

SetVpSelectMemoryRasConfiguration sets VpSelectMemoryRasConfiguration field to given value.

### HasVpSelectMemoryRasConfiguration

`func (o *BiosVfSelectMemoryRasConfiguration) HasVpSelectMemoryRasConfiguration() bool`

HasVpSelectMemoryRasConfiguration returns a boolean if a field has been set.

### GetComputeBlade

`func (o *BiosVfSelectMemoryRasConfiguration) GetComputeBlade() ComputeBladeRelationship`

GetComputeBlade returns the ComputeBlade field if non-nil, zero value otherwise.

### GetComputeBladeOk

`func (o *BiosVfSelectMemoryRasConfiguration) GetComputeBladeOk() (*ComputeBladeRelationship, bool)`

GetComputeBladeOk returns a tuple with the ComputeBlade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBlade

`func (o *BiosVfSelectMemoryRasConfiguration) SetComputeBlade(v ComputeBladeRelationship)`

SetComputeBlade sets ComputeBlade field to given value.

### HasComputeBlade

`func (o *BiosVfSelectMemoryRasConfiguration) HasComputeBlade() bool`

HasComputeBlade returns a boolean if a field has been set.

### GetComputeRackUnit

`func (o *BiosVfSelectMemoryRasConfiguration) GetComputeRackUnit() ComputeRackUnitRelationship`

GetComputeRackUnit returns the ComputeRackUnit field if non-nil, zero value otherwise.

### GetComputeRackUnitOk

`func (o *BiosVfSelectMemoryRasConfiguration) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool)`

GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRackUnit

`func (o *BiosVfSelectMemoryRasConfiguration) SetComputeRackUnit(v ComputeRackUnitRelationship)`

SetComputeRackUnit sets ComputeRackUnit field to given value.

### HasComputeRackUnit

`func (o *BiosVfSelectMemoryRasConfiguration) HasComputeRackUnit() bool`

HasComputeRackUnit returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *BiosVfSelectMemoryRasConfiguration) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *BiosVfSelectMemoryRasConfiguration) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *BiosVfSelectMemoryRasConfiguration) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *BiosVfSelectMemoryRasConfiguration) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *BiosVfSelectMemoryRasConfiguration) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *BiosVfSelectMemoryRasConfiguration) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *BiosVfSelectMemoryRasConfiguration) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *BiosVfSelectMemoryRasConfiguration) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


