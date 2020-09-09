# EquipmentDeviceSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dn** | Pointer to **string** | The distinguished name for the Network Element. | [optional] [readonly] 
**Model** | Pointer to **string** | The model information of the Network Element. | [optional] [readonly] 
**Serial** | Pointer to **string** | The serial number for the Network Element. | [optional] [readonly] 
**SourceObjectType** | Pointer to **string** | The source object type of this view MO. | [optional] [readonly] 
**ComputeBlade** | Pointer to [**ComputeBladeRelationship**](compute.Blade.Relationship.md) |  | [optional] 
**ComputeRackUnit** | Pointer to [**ComputeRackUnitRelationship**](compute.RackUnit.Relationship.md) |  | [optional] 
**EquipmentChassis** | Pointer to [**EquipmentChassisRelationship**](equipment.Chassis.Relationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewEquipmentDeviceSummary

`func NewEquipmentDeviceSummary() *EquipmentDeviceSummary`

NewEquipmentDeviceSummary instantiates a new EquipmentDeviceSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentDeviceSummaryWithDefaults

`func NewEquipmentDeviceSummaryWithDefaults() *EquipmentDeviceSummary`

NewEquipmentDeviceSummaryWithDefaults instantiates a new EquipmentDeviceSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDn

`func (o *EquipmentDeviceSummary) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *EquipmentDeviceSummary) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *EquipmentDeviceSummary) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *EquipmentDeviceSummary) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetModel

`func (o *EquipmentDeviceSummary) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *EquipmentDeviceSummary) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *EquipmentDeviceSummary) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *EquipmentDeviceSummary) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetSerial

`func (o *EquipmentDeviceSummary) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *EquipmentDeviceSummary) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *EquipmentDeviceSummary) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *EquipmentDeviceSummary) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetSourceObjectType

`func (o *EquipmentDeviceSummary) GetSourceObjectType() string`

GetSourceObjectType returns the SourceObjectType field if non-nil, zero value otherwise.

### GetSourceObjectTypeOk

`func (o *EquipmentDeviceSummary) GetSourceObjectTypeOk() (*string, bool)`

GetSourceObjectTypeOk returns a tuple with the SourceObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceObjectType

`func (o *EquipmentDeviceSummary) SetSourceObjectType(v string)`

SetSourceObjectType sets SourceObjectType field to given value.

### HasSourceObjectType

`func (o *EquipmentDeviceSummary) HasSourceObjectType() bool`

HasSourceObjectType returns a boolean if a field has been set.

### GetComputeBlade

`func (o *EquipmentDeviceSummary) GetComputeBlade() ComputeBladeRelationship`

GetComputeBlade returns the ComputeBlade field if non-nil, zero value otherwise.

### GetComputeBladeOk

`func (o *EquipmentDeviceSummary) GetComputeBladeOk() (*ComputeBladeRelationship, bool)`

GetComputeBladeOk returns a tuple with the ComputeBlade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBlade

`func (o *EquipmentDeviceSummary) SetComputeBlade(v ComputeBladeRelationship)`

SetComputeBlade sets ComputeBlade field to given value.

### HasComputeBlade

`func (o *EquipmentDeviceSummary) HasComputeBlade() bool`

HasComputeBlade returns a boolean if a field has been set.

### GetComputeRackUnit

`func (o *EquipmentDeviceSummary) GetComputeRackUnit() ComputeRackUnitRelationship`

GetComputeRackUnit returns the ComputeRackUnit field if non-nil, zero value otherwise.

### GetComputeRackUnitOk

`func (o *EquipmentDeviceSummary) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool)`

GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRackUnit

`func (o *EquipmentDeviceSummary) SetComputeRackUnit(v ComputeRackUnitRelationship)`

SetComputeRackUnit sets ComputeRackUnit field to given value.

### HasComputeRackUnit

`func (o *EquipmentDeviceSummary) HasComputeRackUnit() bool`

HasComputeRackUnit returns a boolean if a field has been set.

### GetEquipmentChassis

`func (o *EquipmentDeviceSummary) GetEquipmentChassis() EquipmentChassisRelationship`

GetEquipmentChassis returns the EquipmentChassis field if non-nil, zero value otherwise.

### GetEquipmentChassisOk

`func (o *EquipmentDeviceSummary) GetEquipmentChassisOk() (*EquipmentChassisRelationship, bool)`

GetEquipmentChassisOk returns a tuple with the EquipmentChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentChassis

`func (o *EquipmentDeviceSummary) SetEquipmentChassis(v EquipmentChassisRelationship)`

SetEquipmentChassis sets EquipmentChassis field to given value.

### HasEquipmentChassis

`func (o *EquipmentDeviceSummary) HasEquipmentChassis() bool`

HasEquipmentChassis returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *EquipmentDeviceSummary) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *EquipmentDeviceSummary) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *EquipmentDeviceSummary) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *EquipmentDeviceSummary) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *EquipmentDeviceSummary) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EquipmentDeviceSummary) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EquipmentDeviceSummary) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EquipmentDeviceSummary) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


