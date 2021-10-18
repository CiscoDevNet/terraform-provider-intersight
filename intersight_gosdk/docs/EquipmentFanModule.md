# EquipmentFanModule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "equipment.FanModule"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "equipment.FanModule"]
**Description** | Pointer to **string** | This field is to provide description for the fan module. | [optional] [readonly] 
**ModuleId** | Pointer to **int64** | This field acts as the identifier for this particular Module, within the Fabric Interconnect. | [optional] [readonly] 
**OperReason** | Pointer to **[]string** |  | [optional] 
**OperState** | Pointer to **string** | This field is used to indicate this fan module&#39;s operational state. | [optional] [readonly] 
**PartNumber** | Pointer to **string** | This field identifies the Part Number for this Fan Module. | [optional] [readonly] 
**Pid** | Pointer to **string** | This field identifies the Product ID for the fan module. | [optional] [readonly] 
**Sku** | Pointer to **string** | This field identifies the Stockkeeping Unit for this Fan Module. | [optional] [readonly] 
**TrayId** | Pointer to **int64** | Tray identifier for the fan module. | [optional] [readonly] 
**Vid** | Pointer to **string** | This field identifies the Vendor ID for this Fan Module. | [optional] [readonly] 
**ComputeRackUnit** | Pointer to [**ComputeRackUnitRelationship**](ComputeRackUnitRelationship.md) |  | [optional] 
**EquipmentChassis** | Pointer to [**EquipmentChassisRelationship**](EquipmentChassisRelationship.md) |  | [optional] 
**EquipmentExpanderModule** | Pointer to [**EquipmentExpanderModuleRelationship**](EquipmentExpanderModuleRelationship.md) |  | [optional] 
**EquipmentIoCard** | Pointer to [**EquipmentIoCardRelationship**](EquipmentIoCardRelationship.md) |  | [optional] 
**EquipmentRackEnclosure** | Pointer to [**EquipmentRackEnclosureRelationship**](EquipmentRackEnclosureRelationship.md) |  | [optional] 
**Fans** | Pointer to [**[]EquipmentFanRelationship**](EquipmentFanRelationship.md) | An array of relationships to equipmentFan resources. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**NetworkElement** | Pointer to [**NetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewEquipmentFanModule

`func NewEquipmentFanModule(classId string, objectType string, ) *EquipmentFanModule`

NewEquipmentFanModule instantiates a new EquipmentFanModule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentFanModuleWithDefaults

`func NewEquipmentFanModuleWithDefaults() *EquipmentFanModule`

NewEquipmentFanModuleWithDefaults instantiates a new EquipmentFanModule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentFanModule) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentFanModule) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentFanModule) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentFanModule) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentFanModule) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentFanModule) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *EquipmentFanModule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EquipmentFanModule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EquipmentFanModule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EquipmentFanModule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetModuleId

`func (o *EquipmentFanModule) GetModuleId() int64`

GetModuleId returns the ModuleId field if non-nil, zero value otherwise.

### GetModuleIdOk

`func (o *EquipmentFanModule) GetModuleIdOk() (*int64, bool)`

GetModuleIdOk returns a tuple with the ModuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleId

`func (o *EquipmentFanModule) SetModuleId(v int64)`

SetModuleId sets ModuleId field to given value.

### HasModuleId

`func (o *EquipmentFanModule) HasModuleId() bool`

HasModuleId returns a boolean if a field has been set.

### GetOperReason

`func (o *EquipmentFanModule) GetOperReason() []string`

GetOperReason returns the OperReason field if non-nil, zero value otherwise.

### GetOperReasonOk

`func (o *EquipmentFanModule) GetOperReasonOk() (*[]string, bool)`

GetOperReasonOk returns a tuple with the OperReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperReason

`func (o *EquipmentFanModule) SetOperReason(v []string)`

SetOperReason sets OperReason field to given value.

### HasOperReason

`func (o *EquipmentFanModule) HasOperReason() bool`

HasOperReason returns a boolean if a field has been set.

### SetOperReasonNil

`func (o *EquipmentFanModule) SetOperReasonNil(b bool)`

 SetOperReasonNil sets the value for OperReason to be an explicit nil

### UnsetOperReason
`func (o *EquipmentFanModule) UnsetOperReason()`

UnsetOperReason ensures that no value is present for OperReason, not even an explicit nil
### GetOperState

`func (o *EquipmentFanModule) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *EquipmentFanModule) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *EquipmentFanModule) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *EquipmentFanModule) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetPartNumber

`func (o *EquipmentFanModule) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *EquipmentFanModule) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *EquipmentFanModule) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *EquipmentFanModule) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetPid

`func (o *EquipmentFanModule) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *EquipmentFanModule) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *EquipmentFanModule) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *EquipmentFanModule) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetSku

`func (o *EquipmentFanModule) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *EquipmentFanModule) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *EquipmentFanModule) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *EquipmentFanModule) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetTrayId

`func (o *EquipmentFanModule) GetTrayId() int64`

GetTrayId returns the TrayId field if non-nil, zero value otherwise.

### GetTrayIdOk

`func (o *EquipmentFanModule) GetTrayIdOk() (*int64, bool)`

GetTrayIdOk returns a tuple with the TrayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrayId

`func (o *EquipmentFanModule) SetTrayId(v int64)`

SetTrayId sets TrayId field to given value.

### HasTrayId

`func (o *EquipmentFanModule) HasTrayId() bool`

HasTrayId returns a boolean if a field has been set.

### GetVid

`func (o *EquipmentFanModule) GetVid() string`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *EquipmentFanModule) GetVidOk() (*string, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *EquipmentFanModule) SetVid(v string)`

SetVid sets Vid field to given value.

### HasVid

`func (o *EquipmentFanModule) HasVid() bool`

HasVid returns a boolean if a field has been set.

### GetComputeRackUnit

`func (o *EquipmentFanModule) GetComputeRackUnit() ComputeRackUnitRelationship`

GetComputeRackUnit returns the ComputeRackUnit field if non-nil, zero value otherwise.

### GetComputeRackUnitOk

`func (o *EquipmentFanModule) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool)`

GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRackUnit

`func (o *EquipmentFanModule) SetComputeRackUnit(v ComputeRackUnitRelationship)`

SetComputeRackUnit sets ComputeRackUnit field to given value.

### HasComputeRackUnit

`func (o *EquipmentFanModule) HasComputeRackUnit() bool`

HasComputeRackUnit returns a boolean if a field has been set.

### GetEquipmentChassis

`func (o *EquipmentFanModule) GetEquipmentChassis() EquipmentChassisRelationship`

GetEquipmentChassis returns the EquipmentChassis field if non-nil, zero value otherwise.

### GetEquipmentChassisOk

`func (o *EquipmentFanModule) GetEquipmentChassisOk() (*EquipmentChassisRelationship, bool)`

GetEquipmentChassisOk returns a tuple with the EquipmentChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentChassis

`func (o *EquipmentFanModule) SetEquipmentChassis(v EquipmentChassisRelationship)`

SetEquipmentChassis sets EquipmentChassis field to given value.

### HasEquipmentChassis

`func (o *EquipmentFanModule) HasEquipmentChassis() bool`

HasEquipmentChassis returns a boolean if a field has been set.

### GetEquipmentExpanderModule

`func (o *EquipmentFanModule) GetEquipmentExpanderModule() EquipmentExpanderModuleRelationship`

GetEquipmentExpanderModule returns the EquipmentExpanderModule field if non-nil, zero value otherwise.

### GetEquipmentExpanderModuleOk

`func (o *EquipmentFanModule) GetEquipmentExpanderModuleOk() (*EquipmentExpanderModuleRelationship, bool)`

GetEquipmentExpanderModuleOk returns a tuple with the EquipmentExpanderModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentExpanderModule

`func (o *EquipmentFanModule) SetEquipmentExpanderModule(v EquipmentExpanderModuleRelationship)`

SetEquipmentExpanderModule sets EquipmentExpanderModule field to given value.

### HasEquipmentExpanderModule

`func (o *EquipmentFanModule) HasEquipmentExpanderModule() bool`

HasEquipmentExpanderModule returns a boolean if a field has been set.

### GetEquipmentIoCard

`func (o *EquipmentFanModule) GetEquipmentIoCard() EquipmentIoCardRelationship`

GetEquipmentIoCard returns the EquipmentIoCard field if non-nil, zero value otherwise.

### GetEquipmentIoCardOk

`func (o *EquipmentFanModule) GetEquipmentIoCardOk() (*EquipmentIoCardRelationship, bool)`

GetEquipmentIoCardOk returns a tuple with the EquipmentIoCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentIoCard

`func (o *EquipmentFanModule) SetEquipmentIoCard(v EquipmentIoCardRelationship)`

SetEquipmentIoCard sets EquipmentIoCard field to given value.

### HasEquipmentIoCard

`func (o *EquipmentFanModule) HasEquipmentIoCard() bool`

HasEquipmentIoCard returns a boolean if a field has been set.

### GetEquipmentRackEnclosure

`func (o *EquipmentFanModule) GetEquipmentRackEnclosure() EquipmentRackEnclosureRelationship`

GetEquipmentRackEnclosure returns the EquipmentRackEnclosure field if non-nil, zero value otherwise.

### GetEquipmentRackEnclosureOk

`func (o *EquipmentFanModule) GetEquipmentRackEnclosureOk() (*EquipmentRackEnclosureRelationship, bool)`

GetEquipmentRackEnclosureOk returns a tuple with the EquipmentRackEnclosure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentRackEnclosure

`func (o *EquipmentFanModule) SetEquipmentRackEnclosure(v EquipmentRackEnclosureRelationship)`

SetEquipmentRackEnclosure sets EquipmentRackEnclosure field to given value.

### HasEquipmentRackEnclosure

`func (o *EquipmentFanModule) HasEquipmentRackEnclosure() bool`

HasEquipmentRackEnclosure returns a boolean if a field has been set.

### GetFans

`func (o *EquipmentFanModule) GetFans() []EquipmentFanRelationship`

GetFans returns the Fans field if non-nil, zero value otherwise.

### GetFansOk

`func (o *EquipmentFanModule) GetFansOk() (*[]EquipmentFanRelationship, bool)`

GetFansOk returns a tuple with the Fans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFans

`func (o *EquipmentFanModule) SetFans(v []EquipmentFanRelationship)`

SetFans sets Fans field to given value.

### HasFans

`func (o *EquipmentFanModule) HasFans() bool`

HasFans returns a boolean if a field has been set.

### SetFansNil

`func (o *EquipmentFanModule) SetFansNil(b bool)`

 SetFansNil sets the value for Fans to be an explicit nil

### UnsetFans
`func (o *EquipmentFanModule) UnsetFans()`

UnsetFans ensures that no value is present for Fans, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *EquipmentFanModule) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *EquipmentFanModule) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *EquipmentFanModule) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *EquipmentFanModule) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetNetworkElement

`func (o *EquipmentFanModule) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *EquipmentFanModule) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *EquipmentFanModule) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *EquipmentFanModule) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *EquipmentFanModule) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EquipmentFanModule) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EquipmentFanModule) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EquipmentFanModule) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


