# EquipmentFanAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | This field is to provide description for the fan. | [optional] [readonly] 
**FanId** | Pointer to **int64** | This field acts as the identifier for this particular Fan, within the Fabric Interconnect. | [optional] [readonly] 
**FanModuleId** | Pointer to **int64** | This field is used to identify the Fan Module to which this Fan belongs. | [optional] [readonly] 
**ModuleId** | Pointer to **int64** | Fan module Identifier for the fan. | [optional] [readonly] 
**OperState** | Pointer to **string** | This field is used to indicate this fan unit&#39;s operational state. | [optional] [readonly] 
**PartNumber** | Pointer to **string** | This field identifies the Part Number for this Fan Unit. | [optional] [readonly] 
**Pid** | Pointer to **string** | This field identifies the Product ID for the fans. | [optional] [readonly] 
**Presence** | Pointer to **string** | This field is used to indicate this fan unit&#39;s presence. | [optional] [readonly] 
**Sku** | Pointer to **string** | This field identifies the Stockkeeping Unit for this Fan Unit. | [optional] [readonly] 
**TrayId** | Pointer to **int64** | Tray identifier for the fan module. | [optional] [readonly] 
**Vid** | Pointer to **string** | This field identifies the Vendor ID for this Fan Unit. | [optional] [readonly] 
**EquipmentFanModule** | Pointer to [**EquipmentFanModuleRelationship**](equipment.FanModule.Relationship.md) |  | [optional] 
**EquipmentFex** | Pointer to [**EquipmentFexRelationship**](equipment.Fex.Relationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewEquipmentFanAllOf

`func NewEquipmentFanAllOf() *EquipmentFanAllOf`

NewEquipmentFanAllOf instantiates a new EquipmentFanAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentFanAllOfWithDefaults

`func NewEquipmentFanAllOfWithDefaults() *EquipmentFanAllOf`

NewEquipmentFanAllOfWithDefaults instantiates a new EquipmentFanAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *EquipmentFanAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EquipmentFanAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EquipmentFanAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EquipmentFanAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFanId

`func (o *EquipmentFanAllOf) GetFanId() int64`

GetFanId returns the FanId field if non-nil, zero value otherwise.

### GetFanIdOk

`func (o *EquipmentFanAllOf) GetFanIdOk() (*int64, bool)`

GetFanIdOk returns a tuple with the FanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanId

`func (o *EquipmentFanAllOf) SetFanId(v int64)`

SetFanId sets FanId field to given value.

### HasFanId

`func (o *EquipmentFanAllOf) HasFanId() bool`

HasFanId returns a boolean if a field has been set.

### GetFanModuleId

`func (o *EquipmentFanAllOf) GetFanModuleId() int64`

GetFanModuleId returns the FanModuleId field if non-nil, zero value otherwise.

### GetFanModuleIdOk

`func (o *EquipmentFanAllOf) GetFanModuleIdOk() (*int64, bool)`

GetFanModuleIdOk returns a tuple with the FanModuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanModuleId

`func (o *EquipmentFanAllOf) SetFanModuleId(v int64)`

SetFanModuleId sets FanModuleId field to given value.

### HasFanModuleId

`func (o *EquipmentFanAllOf) HasFanModuleId() bool`

HasFanModuleId returns a boolean if a field has been set.

### GetModuleId

`func (o *EquipmentFanAllOf) GetModuleId() int64`

GetModuleId returns the ModuleId field if non-nil, zero value otherwise.

### GetModuleIdOk

`func (o *EquipmentFanAllOf) GetModuleIdOk() (*int64, bool)`

GetModuleIdOk returns a tuple with the ModuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleId

`func (o *EquipmentFanAllOf) SetModuleId(v int64)`

SetModuleId sets ModuleId field to given value.

### HasModuleId

`func (o *EquipmentFanAllOf) HasModuleId() bool`

HasModuleId returns a boolean if a field has been set.

### GetOperState

`func (o *EquipmentFanAllOf) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *EquipmentFanAllOf) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *EquipmentFanAllOf) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *EquipmentFanAllOf) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetPartNumber

`func (o *EquipmentFanAllOf) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *EquipmentFanAllOf) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *EquipmentFanAllOf) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *EquipmentFanAllOf) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetPid

`func (o *EquipmentFanAllOf) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *EquipmentFanAllOf) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *EquipmentFanAllOf) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *EquipmentFanAllOf) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetPresence

`func (o *EquipmentFanAllOf) GetPresence() string`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *EquipmentFanAllOf) GetPresenceOk() (*string, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *EquipmentFanAllOf) SetPresence(v string)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *EquipmentFanAllOf) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetSku

`func (o *EquipmentFanAllOf) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *EquipmentFanAllOf) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *EquipmentFanAllOf) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *EquipmentFanAllOf) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetTrayId

`func (o *EquipmentFanAllOf) GetTrayId() int64`

GetTrayId returns the TrayId field if non-nil, zero value otherwise.

### GetTrayIdOk

`func (o *EquipmentFanAllOf) GetTrayIdOk() (*int64, bool)`

GetTrayIdOk returns a tuple with the TrayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrayId

`func (o *EquipmentFanAllOf) SetTrayId(v int64)`

SetTrayId sets TrayId field to given value.

### HasTrayId

`func (o *EquipmentFanAllOf) HasTrayId() bool`

HasTrayId returns a boolean if a field has been set.

### GetVid

`func (o *EquipmentFanAllOf) GetVid() string`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *EquipmentFanAllOf) GetVidOk() (*string, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *EquipmentFanAllOf) SetVid(v string)`

SetVid sets Vid field to given value.

### HasVid

`func (o *EquipmentFanAllOf) HasVid() bool`

HasVid returns a boolean if a field has been set.

### GetEquipmentFanModule

`func (o *EquipmentFanAllOf) GetEquipmentFanModule() EquipmentFanModuleRelationship`

GetEquipmentFanModule returns the EquipmentFanModule field if non-nil, zero value otherwise.

### GetEquipmentFanModuleOk

`func (o *EquipmentFanAllOf) GetEquipmentFanModuleOk() (*EquipmentFanModuleRelationship, bool)`

GetEquipmentFanModuleOk returns a tuple with the EquipmentFanModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentFanModule

`func (o *EquipmentFanAllOf) SetEquipmentFanModule(v EquipmentFanModuleRelationship)`

SetEquipmentFanModule sets EquipmentFanModule field to given value.

### HasEquipmentFanModule

`func (o *EquipmentFanAllOf) HasEquipmentFanModule() bool`

HasEquipmentFanModule returns a boolean if a field has been set.

### GetEquipmentFex

`func (o *EquipmentFanAllOf) GetEquipmentFex() EquipmentFexRelationship`

GetEquipmentFex returns the EquipmentFex field if non-nil, zero value otherwise.

### GetEquipmentFexOk

`func (o *EquipmentFanAllOf) GetEquipmentFexOk() (*EquipmentFexRelationship, bool)`

GetEquipmentFexOk returns a tuple with the EquipmentFex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentFex

`func (o *EquipmentFanAllOf) SetEquipmentFex(v EquipmentFexRelationship)`

SetEquipmentFex sets EquipmentFex field to given value.

### HasEquipmentFex

`func (o *EquipmentFanAllOf) HasEquipmentFex() bool`

HasEquipmentFex returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *EquipmentFanAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *EquipmentFanAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *EquipmentFanAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *EquipmentFanAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *EquipmentFanAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EquipmentFanAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EquipmentFanAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EquipmentFanAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


