# BiosVfSelectMemoryRasConfigurationAllOf

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

### NewBiosVfSelectMemoryRasConfigurationAllOf

`func NewBiosVfSelectMemoryRasConfigurationAllOf(classId string, objectType string, ) *BiosVfSelectMemoryRasConfigurationAllOf`

NewBiosVfSelectMemoryRasConfigurationAllOf instantiates a new BiosVfSelectMemoryRasConfigurationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiosVfSelectMemoryRasConfigurationAllOfWithDefaults

`func NewBiosVfSelectMemoryRasConfigurationAllOfWithDefaults() *BiosVfSelectMemoryRasConfigurationAllOf`

NewBiosVfSelectMemoryRasConfigurationAllOfWithDefaults instantiates a new BiosVfSelectMemoryRasConfigurationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BiosVfSelectMemoryRasConfigurationAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BiosVfSelectMemoryRasConfigurationAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BiosVfSelectMemoryRasConfigurationAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BiosVfSelectMemoryRasConfigurationAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BiosVfSelectMemoryRasConfigurationAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BiosVfSelectMemoryRasConfigurationAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSerial

`func (o *BiosVfSelectMemoryRasConfigurationAllOf) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *BiosVfSelectMemoryRasConfigurationAllOf) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *BiosVfSelectMemoryRasConfigurationAllOf) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *BiosVfSelectMemoryRasConfigurationAllOf) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetVpSelectMemoryRasConfiguration

`func (o *BiosVfSelectMemoryRasConfigurationAllOf) GetVpSelectMemoryRasConfiguration() string`

GetVpSelectMemoryRasConfiguration returns the VpSelectMemoryRasConfiguration field if non-nil, zero value otherwise.

### GetVpSelectMemoryRasConfigurationOk

`func (o *BiosVfSelectMemoryRasConfigurationAllOf) GetVpSelectMemoryRasConfigurationOk() (*string, bool)`

GetVpSelectMemoryRasConfigurationOk returns a tuple with the VpSelectMemoryRasConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpSelectMemoryRasConfiguration

`func (o *BiosVfSelectMemoryRasConfigurationAllOf) SetVpSelectMemoryRasConfiguration(v string)`

SetVpSelectMemoryRasConfiguration sets VpSelectMemoryRasConfiguration field to given value.

### HasVpSelectMemoryRasConfiguration

`func (o *BiosVfSelectMemoryRasConfigurationAllOf) HasVpSelectMemoryRasConfiguration() bool`

HasVpSelectMemoryRasConfiguration returns a boolean if a field has been set.

### GetComputeBlade

`func (o *BiosVfSelectMemoryRasConfigurationAllOf) GetComputeBlade() ComputeBladeRelationship`

GetComputeBlade returns the ComputeBlade field if non-nil, zero value otherwise.

### GetComputeBladeOk

`func (o *BiosVfSelectMemoryRasConfigurationAllOf) GetComputeBladeOk() (*ComputeBladeRelationship, bool)`

GetComputeBladeOk returns a tuple with the ComputeBlade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBlade

`func (o *BiosVfSelectMemoryRasConfigurationAllOf) SetComputeBlade(v ComputeBladeRelationship)`

SetComputeBlade sets ComputeBlade field to given value.

### HasComputeBlade

`func (o *BiosVfSelectMemoryRasConfigurationAllOf) HasComputeBlade() bool`

HasComputeBlade returns a boolean if a field has been set.

### GetComputeRackUnit

`func (o *BiosVfSelectMemoryRasConfigurationAllOf) GetComputeRackUnit() ComputeRackUnitRelationship`

GetComputeRackUnit returns the ComputeRackUnit field if non-nil, zero value otherwise.

### GetComputeRackUnitOk

`func (o *BiosVfSelectMemoryRasConfigurationAllOf) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool)`

GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRackUnit

`func (o *BiosVfSelectMemoryRasConfigurationAllOf) SetComputeRackUnit(v ComputeRackUnitRelationship)`

SetComputeRackUnit sets ComputeRackUnit field to given value.

### HasComputeRackUnit

`func (o *BiosVfSelectMemoryRasConfigurationAllOf) HasComputeRackUnit() bool`

HasComputeRackUnit returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *BiosVfSelectMemoryRasConfigurationAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *BiosVfSelectMemoryRasConfigurationAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *BiosVfSelectMemoryRasConfigurationAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *BiosVfSelectMemoryRasConfigurationAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *BiosVfSelectMemoryRasConfigurationAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *BiosVfSelectMemoryRasConfigurationAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *BiosVfSelectMemoryRasConfigurationAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *BiosVfSelectMemoryRasConfigurationAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


