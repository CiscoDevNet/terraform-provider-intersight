# EquipmentPsuAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | This field is to provide description for the power supply unit. | [optional] [readonly] 
**OperState** | Pointer to **string** | This field identifies the psu operational state. | [optional] [readonly] 
**PartNumber** | Pointer to **string** | This field identifies the Part Number for this Power Supply Unit. | [optional] [readonly] 
**Pid** | Pointer to **string** | This field identifies the Product ID for the Power Supply. | [optional] [readonly] 
**Presence** | Pointer to **string** | This field identifies the presence state of the psu. | [optional] [readonly] 
**PsuFwVersion** | Pointer to **string** | This field identifies the Firmware Version of the Power Supply. | [optional] [readonly] 
**PsuId** | Pointer to **int64** | This represents power supply unit identifier in chassis/server. | [optional] [readonly] 
**PsuInputSrc** | Pointer to **string** | This field identifies the input source for the Power Supply. | [optional] [readonly] 
**PsuType** | Pointer to **string** | This field identifies the type of the Power Supply. | [optional] [readonly] 
**PsuWattage** | Pointer to **string** | This field identifies the Wattage of the Power Supply. | [optional] [readonly] 
**Sku** | Pointer to **string** | This field identifies the Stockkeeping Unit for this Power Supply. | [optional] [readonly] 
**Vid** | Pointer to **string** | This field identifies the Vendor ID for this Power Supply Unit. | [optional] [readonly] 
**Voltage** | Pointer to **string** | This field is used to indicate the Voltage for this Power Supply. | [optional] [readonly] 
**ComputeRackUnit** | Pointer to [**ComputeRackUnitRelationship**](compute.RackUnit.Relationship.md) |  | [optional] 
**EquipmentChassis** | Pointer to [**EquipmentChassisRelationship**](equipment.Chassis.Relationship.md) |  | [optional] 
**EquipmentFex** | Pointer to [**EquipmentFexRelationship**](equipment.Fex.Relationship.md) |  | [optional] 
**EquipmentRackEnclosure** | Pointer to [**EquipmentRackEnclosureRelationship**](equipment.RackEnclosure.Relationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**NetworkElement** | Pointer to [**NetworkElementRelationship**](network.Element.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewEquipmentPsuAllOf

`func NewEquipmentPsuAllOf() *EquipmentPsuAllOf`

NewEquipmentPsuAllOf instantiates a new EquipmentPsuAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentPsuAllOfWithDefaults

`func NewEquipmentPsuAllOfWithDefaults() *EquipmentPsuAllOf`

NewEquipmentPsuAllOfWithDefaults instantiates a new EquipmentPsuAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *EquipmentPsuAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EquipmentPsuAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EquipmentPsuAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EquipmentPsuAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOperState

`func (o *EquipmentPsuAllOf) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *EquipmentPsuAllOf) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *EquipmentPsuAllOf) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *EquipmentPsuAllOf) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetPartNumber

`func (o *EquipmentPsuAllOf) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *EquipmentPsuAllOf) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *EquipmentPsuAllOf) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *EquipmentPsuAllOf) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetPid

`func (o *EquipmentPsuAllOf) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *EquipmentPsuAllOf) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *EquipmentPsuAllOf) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *EquipmentPsuAllOf) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetPresence

`func (o *EquipmentPsuAllOf) GetPresence() string`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *EquipmentPsuAllOf) GetPresenceOk() (*string, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *EquipmentPsuAllOf) SetPresence(v string)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *EquipmentPsuAllOf) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetPsuFwVersion

`func (o *EquipmentPsuAllOf) GetPsuFwVersion() string`

GetPsuFwVersion returns the PsuFwVersion field if non-nil, zero value otherwise.

### GetPsuFwVersionOk

`func (o *EquipmentPsuAllOf) GetPsuFwVersionOk() (*string, bool)`

GetPsuFwVersionOk returns a tuple with the PsuFwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsuFwVersion

`func (o *EquipmentPsuAllOf) SetPsuFwVersion(v string)`

SetPsuFwVersion sets PsuFwVersion field to given value.

### HasPsuFwVersion

`func (o *EquipmentPsuAllOf) HasPsuFwVersion() bool`

HasPsuFwVersion returns a boolean if a field has been set.

### GetPsuId

`func (o *EquipmentPsuAllOf) GetPsuId() int64`

GetPsuId returns the PsuId field if non-nil, zero value otherwise.

### GetPsuIdOk

`func (o *EquipmentPsuAllOf) GetPsuIdOk() (*int64, bool)`

GetPsuIdOk returns a tuple with the PsuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsuId

`func (o *EquipmentPsuAllOf) SetPsuId(v int64)`

SetPsuId sets PsuId field to given value.

### HasPsuId

`func (o *EquipmentPsuAllOf) HasPsuId() bool`

HasPsuId returns a boolean if a field has been set.

### GetPsuInputSrc

`func (o *EquipmentPsuAllOf) GetPsuInputSrc() string`

GetPsuInputSrc returns the PsuInputSrc field if non-nil, zero value otherwise.

### GetPsuInputSrcOk

`func (o *EquipmentPsuAllOf) GetPsuInputSrcOk() (*string, bool)`

GetPsuInputSrcOk returns a tuple with the PsuInputSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsuInputSrc

`func (o *EquipmentPsuAllOf) SetPsuInputSrc(v string)`

SetPsuInputSrc sets PsuInputSrc field to given value.

### HasPsuInputSrc

`func (o *EquipmentPsuAllOf) HasPsuInputSrc() bool`

HasPsuInputSrc returns a boolean if a field has been set.

### GetPsuType

`func (o *EquipmentPsuAllOf) GetPsuType() string`

GetPsuType returns the PsuType field if non-nil, zero value otherwise.

### GetPsuTypeOk

`func (o *EquipmentPsuAllOf) GetPsuTypeOk() (*string, bool)`

GetPsuTypeOk returns a tuple with the PsuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsuType

`func (o *EquipmentPsuAllOf) SetPsuType(v string)`

SetPsuType sets PsuType field to given value.

### HasPsuType

`func (o *EquipmentPsuAllOf) HasPsuType() bool`

HasPsuType returns a boolean if a field has been set.

### GetPsuWattage

`func (o *EquipmentPsuAllOf) GetPsuWattage() string`

GetPsuWattage returns the PsuWattage field if non-nil, zero value otherwise.

### GetPsuWattageOk

`func (o *EquipmentPsuAllOf) GetPsuWattageOk() (*string, bool)`

GetPsuWattageOk returns a tuple with the PsuWattage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsuWattage

`func (o *EquipmentPsuAllOf) SetPsuWattage(v string)`

SetPsuWattage sets PsuWattage field to given value.

### HasPsuWattage

`func (o *EquipmentPsuAllOf) HasPsuWattage() bool`

HasPsuWattage returns a boolean if a field has been set.

### GetSku

`func (o *EquipmentPsuAllOf) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *EquipmentPsuAllOf) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *EquipmentPsuAllOf) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *EquipmentPsuAllOf) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetVid

`func (o *EquipmentPsuAllOf) GetVid() string`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *EquipmentPsuAllOf) GetVidOk() (*string, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *EquipmentPsuAllOf) SetVid(v string)`

SetVid sets Vid field to given value.

### HasVid

`func (o *EquipmentPsuAllOf) HasVid() bool`

HasVid returns a boolean if a field has been set.

### GetVoltage

`func (o *EquipmentPsuAllOf) GetVoltage() string`

GetVoltage returns the Voltage field if non-nil, zero value otherwise.

### GetVoltageOk

`func (o *EquipmentPsuAllOf) GetVoltageOk() (*string, bool)`

GetVoltageOk returns a tuple with the Voltage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoltage

`func (o *EquipmentPsuAllOf) SetVoltage(v string)`

SetVoltage sets Voltage field to given value.

### HasVoltage

`func (o *EquipmentPsuAllOf) HasVoltage() bool`

HasVoltage returns a boolean if a field has been set.

### GetComputeRackUnit

`func (o *EquipmentPsuAllOf) GetComputeRackUnit() ComputeRackUnitRelationship`

GetComputeRackUnit returns the ComputeRackUnit field if non-nil, zero value otherwise.

### GetComputeRackUnitOk

`func (o *EquipmentPsuAllOf) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool)`

GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRackUnit

`func (o *EquipmentPsuAllOf) SetComputeRackUnit(v ComputeRackUnitRelationship)`

SetComputeRackUnit sets ComputeRackUnit field to given value.

### HasComputeRackUnit

`func (o *EquipmentPsuAllOf) HasComputeRackUnit() bool`

HasComputeRackUnit returns a boolean if a field has been set.

### GetEquipmentChassis

`func (o *EquipmentPsuAllOf) GetEquipmentChassis() EquipmentChassisRelationship`

GetEquipmentChassis returns the EquipmentChassis field if non-nil, zero value otherwise.

### GetEquipmentChassisOk

`func (o *EquipmentPsuAllOf) GetEquipmentChassisOk() (*EquipmentChassisRelationship, bool)`

GetEquipmentChassisOk returns a tuple with the EquipmentChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentChassis

`func (o *EquipmentPsuAllOf) SetEquipmentChassis(v EquipmentChassisRelationship)`

SetEquipmentChassis sets EquipmentChassis field to given value.

### HasEquipmentChassis

`func (o *EquipmentPsuAllOf) HasEquipmentChassis() bool`

HasEquipmentChassis returns a boolean if a field has been set.

### GetEquipmentFex

`func (o *EquipmentPsuAllOf) GetEquipmentFex() EquipmentFexRelationship`

GetEquipmentFex returns the EquipmentFex field if non-nil, zero value otherwise.

### GetEquipmentFexOk

`func (o *EquipmentPsuAllOf) GetEquipmentFexOk() (*EquipmentFexRelationship, bool)`

GetEquipmentFexOk returns a tuple with the EquipmentFex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentFex

`func (o *EquipmentPsuAllOf) SetEquipmentFex(v EquipmentFexRelationship)`

SetEquipmentFex sets EquipmentFex field to given value.

### HasEquipmentFex

`func (o *EquipmentPsuAllOf) HasEquipmentFex() bool`

HasEquipmentFex returns a boolean if a field has been set.

### GetEquipmentRackEnclosure

`func (o *EquipmentPsuAllOf) GetEquipmentRackEnclosure() EquipmentRackEnclosureRelationship`

GetEquipmentRackEnclosure returns the EquipmentRackEnclosure field if non-nil, zero value otherwise.

### GetEquipmentRackEnclosureOk

`func (o *EquipmentPsuAllOf) GetEquipmentRackEnclosureOk() (*EquipmentRackEnclosureRelationship, bool)`

GetEquipmentRackEnclosureOk returns a tuple with the EquipmentRackEnclosure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentRackEnclosure

`func (o *EquipmentPsuAllOf) SetEquipmentRackEnclosure(v EquipmentRackEnclosureRelationship)`

SetEquipmentRackEnclosure sets EquipmentRackEnclosure field to given value.

### HasEquipmentRackEnclosure

`func (o *EquipmentPsuAllOf) HasEquipmentRackEnclosure() bool`

HasEquipmentRackEnclosure returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *EquipmentPsuAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *EquipmentPsuAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *EquipmentPsuAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *EquipmentPsuAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetNetworkElement

`func (o *EquipmentPsuAllOf) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *EquipmentPsuAllOf) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *EquipmentPsuAllOf) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *EquipmentPsuAllOf) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *EquipmentPsuAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EquipmentPsuAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EquipmentPsuAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EquipmentPsuAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


