# EquipmentPsu

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "equipment.Psu"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "equipment.Psu"]
**Description** | Pointer to **string** | This field is to provide description for the power supply unit. | [optional] [readonly] 
**OperReason** | Pointer to **[]string** |  | [optional] 
**OperState** | Pointer to **string** | This field identifies the psu operational state. | [optional] [readonly] 
**PartNumber** | Pointer to **string** | This field identifies the Part Number for this Power Supply Unit. | [optional] [readonly] 
**Pid** | Pointer to **string** | This field identifies the Product ID for the Power Supply. | [optional] [readonly] 
**PsuFwVersion** | Pointer to **string** | This field identifies the Firmware Version of the Power Supply. | [optional] [readonly] 
**PsuId** | Pointer to **int64** | This represents power supply unit identifier in chassis/server. | [optional] [readonly] 
**PsuInputSrc** | Pointer to **string** | This field identifies the input source for the Power Supply. | [optional] [readonly] 
**PsuType** | Pointer to **string** | This field identifies the type of the Power Supply. | [optional] [readonly] 
**PsuWattage** | Pointer to **string** | This field identifies the Wattage of the Power Supply. | [optional] [readonly] 
**Sku** | Pointer to **string** | This field identifies the Stockkeeping Unit for this Power Supply. | [optional] [readonly] 
**Vid** | Pointer to **string** | This field identifies the Vendor ID for this Power Supply Unit. | [optional] [readonly] 
**Voltage** | Pointer to **string** | This field is used to indicate the voltage state for this Power Supply. | [optional] [readonly] 
**ComputeRackUnit** | Pointer to [**ComputeRackUnitRelationship**](ComputeRackUnitRelationship.md) |  | [optional] 
**EquipmentChassis** | Pointer to [**EquipmentChassisRelationship**](EquipmentChassisRelationship.md) |  | [optional] 
**EquipmentFex** | Pointer to [**EquipmentFexRelationship**](EquipmentFexRelationship.md) |  | [optional] 
**EquipmentRackEnclosure** | Pointer to [**EquipmentRackEnclosureRelationship**](EquipmentRackEnclosureRelationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**NetworkElement** | Pointer to [**NetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewEquipmentPsu

`func NewEquipmentPsu(classId string, objectType string, ) *EquipmentPsu`

NewEquipmentPsu instantiates a new EquipmentPsu object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentPsuWithDefaults

`func NewEquipmentPsuWithDefaults() *EquipmentPsu`

NewEquipmentPsuWithDefaults instantiates a new EquipmentPsu object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentPsu) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentPsu) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentPsu) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentPsu) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentPsu) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentPsu) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *EquipmentPsu) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EquipmentPsu) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EquipmentPsu) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EquipmentPsu) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOperReason

`func (o *EquipmentPsu) GetOperReason() []string`

GetOperReason returns the OperReason field if non-nil, zero value otherwise.

### GetOperReasonOk

`func (o *EquipmentPsu) GetOperReasonOk() (*[]string, bool)`

GetOperReasonOk returns a tuple with the OperReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperReason

`func (o *EquipmentPsu) SetOperReason(v []string)`

SetOperReason sets OperReason field to given value.

### HasOperReason

`func (o *EquipmentPsu) HasOperReason() bool`

HasOperReason returns a boolean if a field has been set.

### SetOperReasonNil

`func (o *EquipmentPsu) SetOperReasonNil(b bool)`

 SetOperReasonNil sets the value for OperReason to be an explicit nil

### UnsetOperReason
`func (o *EquipmentPsu) UnsetOperReason()`

UnsetOperReason ensures that no value is present for OperReason, not even an explicit nil
### GetOperState

`func (o *EquipmentPsu) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *EquipmentPsu) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *EquipmentPsu) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *EquipmentPsu) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetPartNumber

`func (o *EquipmentPsu) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *EquipmentPsu) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *EquipmentPsu) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *EquipmentPsu) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetPid

`func (o *EquipmentPsu) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *EquipmentPsu) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *EquipmentPsu) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *EquipmentPsu) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetPsuFwVersion

`func (o *EquipmentPsu) GetPsuFwVersion() string`

GetPsuFwVersion returns the PsuFwVersion field if non-nil, zero value otherwise.

### GetPsuFwVersionOk

`func (o *EquipmentPsu) GetPsuFwVersionOk() (*string, bool)`

GetPsuFwVersionOk returns a tuple with the PsuFwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsuFwVersion

`func (o *EquipmentPsu) SetPsuFwVersion(v string)`

SetPsuFwVersion sets PsuFwVersion field to given value.

### HasPsuFwVersion

`func (o *EquipmentPsu) HasPsuFwVersion() bool`

HasPsuFwVersion returns a boolean if a field has been set.

### GetPsuId

`func (o *EquipmentPsu) GetPsuId() int64`

GetPsuId returns the PsuId field if non-nil, zero value otherwise.

### GetPsuIdOk

`func (o *EquipmentPsu) GetPsuIdOk() (*int64, bool)`

GetPsuIdOk returns a tuple with the PsuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsuId

`func (o *EquipmentPsu) SetPsuId(v int64)`

SetPsuId sets PsuId field to given value.

### HasPsuId

`func (o *EquipmentPsu) HasPsuId() bool`

HasPsuId returns a boolean if a field has been set.

### GetPsuInputSrc

`func (o *EquipmentPsu) GetPsuInputSrc() string`

GetPsuInputSrc returns the PsuInputSrc field if non-nil, zero value otherwise.

### GetPsuInputSrcOk

`func (o *EquipmentPsu) GetPsuInputSrcOk() (*string, bool)`

GetPsuInputSrcOk returns a tuple with the PsuInputSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsuInputSrc

`func (o *EquipmentPsu) SetPsuInputSrc(v string)`

SetPsuInputSrc sets PsuInputSrc field to given value.

### HasPsuInputSrc

`func (o *EquipmentPsu) HasPsuInputSrc() bool`

HasPsuInputSrc returns a boolean if a field has been set.

### GetPsuType

`func (o *EquipmentPsu) GetPsuType() string`

GetPsuType returns the PsuType field if non-nil, zero value otherwise.

### GetPsuTypeOk

`func (o *EquipmentPsu) GetPsuTypeOk() (*string, bool)`

GetPsuTypeOk returns a tuple with the PsuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsuType

`func (o *EquipmentPsu) SetPsuType(v string)`

SetPsuType sets PsuType field to given value.

### HasPsuType

`func (o *EquipmentPsu) HasPsuType() bool`

HasPsuType returns a boolean if a field has been set.

### GetPsuWattage

`func (o *EquipmentPsu) GetPsuWattage() string`

GetPsuWattage returns the PsuWattage field if non-nil, zero value otherwise.

### GetPsuWattageOk

`func (o *EquipmentPsu) GetPsuWattageOk() (*string, bool)`

GetPsuWattageOk returns a tuple with the PsuWattage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsuWattage

`func (o *EquipmentPsu) SetPsuWattage(v string)`

SetPsuWattage sets PsuWattage field to given value.

### HasPsuWattage

`func (o *EquipmentPsu) HasPsuWattage() bool`

HasPsuWattage returns a boolean if a field has been set.

### GetSku

`func (o *EquipmentPsu) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *EquipmentPsu) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *EquipmentPsu) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *EquipmentPsu) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetVid

`func (o *EquipmentPsu) GetVid() string`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *EquipmentPsu) GetVidOk() (*string, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *EquipmentPsu) SetVid(v string)`

SetVid sets Vid field to given value.

### HasVid

`func (o *EquipmentPsu) HasVid() bool`

HasVid returns a boolean if a field has been set.

### GetVoltage

`func (o *EquipmentPsu) GetVoltage() string`

GetVoltage returns the Voltage field if non-nil, zero value otherwise.

### GetVoltageOk

`func (o *EquipmentPsu) GetVoltageOk() (*string, bool)`

GetVoltageOk returns a tuple with the Voltage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoltage

`func (o *EquipmentPsu) SetVoltage(v string)`

SetVoltage sets Voltage field to given value.

### HasVoltage

`func (o *EquipmentPsu) HasVoltage() bool`

HasVoltage returns a boolean if a field has been set.

### GetComputeRackUnit

`func (o *EquipmentPsu) GetComputeRackUnit() ComputeRackUnitRelationship`

GetComputeRackUnit returns the ComputeRackUnit field if non-nil, zero value otherwise.

### GetComputeRackUnitOk

`func (o *EquipmentPsu) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool)`

GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRackUnit

`func (o *EquipmentPsu) SetComputeRackUnit(v ComputeRackUnitRelationship)`

SetComputeRackUnit sets ComputeRackUnit field to given value.

### HasComputeRackUnit

`func (o *EquipmentPsu) HasComputeRackUnit() bool`

HasComputeRackUnit returns a boolean if a field has been set.

### GetEquipmentChassis

`func (o *EquipmentPsu) GetEquipmentChassis() EquipmentChassisRelationship`

GetEquipmentChassis returns the EquipmentChassis field if non-nil, zero value otherwise.

### GetEquipmentChassisOk

`func (o *EquipmentPsu) GetEquipmentChassisOk() (*EquipmentChassisRelationship, bool)`

GetEquipmentChassisOk returns a tuple with the EquipmentChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentChassis

`func (o *EquipmentPsu) SetEquipmentChassis(v EquipmentChassisRelationship)`

SetEquipmentChassis sets EquipmentChassis field to given value.

### HasEquipmentChassis

`func (o *EquipmentPsu) HasEquipmentChassis() bool`

HasEquipmentChassis returns a boolean if a field has been set.

### GetEquipmentFex

`func (o *EquipmentPsu) GetEquipmentFex() EquipmentFexRelationship`

GetEquipmentFex returns the EquipmentFex field if non-nil, zero value otherwise.

### GetEquipmentFexOk

`func (o *EquipmentPsu) GetEquipmentFexOk() (*EquipmentFexRelationship, bool)`

GetEquipmentFexOk returns a tuple with the EquipmentFex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentFex

`func (o *EquipmentPsu) SetEquipmentFex(v EquipmentFexRelationship)`

SetEquipmentFex sets EquipmentFex field to given value.

### HasEquipmentFex

`func (o *EquipmentPsu) HasEquipmentFex() bool`

HasEquipmentFex returns a boolean if a field has been set.

### GetEquipmentRackEnclosure

`func (o *EquipmentPsu) GetEquipmentRackEnclosure() EquipmentRackEnclosureRelationship`

GetEquipmentRackEnclosure returns the EquipmentRackEnclosure field if non-nil, zero value otherwise.

### GetEquipmentRackEnclosureOk

`func (o *EquipmentPsu) GetEquipmentRackEnclosureOk() (*EquipmentRackEnclosureRelationship, bool)`

GetEquipmentRackEnclosureOk returns a tuple with the EquipmentRackEnclosure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentRackEnclosure

`func (o *EquipmentPsu) SetEquipmentRackEnclosure(v EquipmentRackEnclosureRelationship)`

SetEquipmentRackEnclosure sets EquipmentRackEnclosure field to given value.

### HasEquipmentRackEnclosure

`func (o *EquipmentPsu) HasEquipmentRackEnclosure() bool`

HasEquipmentRackEnclosure returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *EquipmentPsu) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *EquipmentPsu) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *EquipmentPsu) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *EquipmentPsu) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetNetworkElement

`func (o *EquipmentPsu) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *EquipmentPsu) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *EquipmentPsu) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *EquipmentPsu) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *EquipmentPsu) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EquipmentPsu) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EquipmentPsu) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EquipmentPsu) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


