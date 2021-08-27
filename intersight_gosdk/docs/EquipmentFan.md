# EquipmentFan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "equipment.Fan"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "equipment.Fan"]
**Description** | Pointer to **string** | This field is to provide description for the fan. | [optional] [readonly] 
**FanId** | Pointer to **int64** | This field acts as the identifier for this particular Fan, within the Fabric Interconnect. | [optional] [readonly] 
**FanModuleId** | Pointer to **int64** | This field is used to identify the Fan Module to which this Fan belongs. | [optional] [readonly] 
**ModuleId** | Pointer to **int64** | Fan module Identifier for the fan. | [optional] [readonly] 
**OperReason** | Pointer to **[]string** |  | [optional] 
**OperState** | Pointer to **string** | This field is used to indicate this fan unit&#39;s operational state. | [optional] [readonly] 
**PartNumber** | Pointer to **string** | This field identifies the Part Number for this Fan Unit. | [optional] [readonly] 
**Pid** | Pointer to **string** | This field identifies the Product ID for the fans. | [optional] [readonly] 
**Sku** | Pointer to **string** | This field identifies the Stockkeeping Unit for this Fan Unit. | [optional] [readonly] 
**TrayId** | Pointer to **int64** | Tray identifier for the fan module. | [optional] [readonly] 
**Vid** | Pointer to **string** | This field identifies the Vendor ID for this Fan Unit. | [optional] [readonly] 
**EquipmentFanModule** | Pointer to [**EquipmentFanModuleRelationship**](EquipmentFanModuleRelationship.md) |  | [optional] 
**EquipmentFex** | Pointer to [**EquipmentFexRelationship**](EquipmentFexRelationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewEquipmentFan

`func NewEquipmentFan(classId string, objectType string, ) *EquipmentFan`

NewEquipmentFan instantiates a new EquipmentFan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentFanWithDefaults

`func NewEquipmentFanWithDefaults() *EquipmentFan`

NewEquipmentFanWithDefaults instantiates a new EquipmentFan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentFan) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentFan) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentFan) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentFan) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentFan) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentFan) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *EquipmentFan) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EquipmentFan) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EquipmentFan) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EquipmentFan) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFanId

`func (o *EquipmentFan) GetFanId() int64`

GetFanId returns the FanId field if non-nil, zero value otherwise.

### GetFanIdOk

`func (o *EquipmentFan) GetFanIdOk() (*int64, bool)`

GetFanIdOk returns a tuple with the FanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanId

`func (o *EquipmentFan) SetFanId(v int64)`

SetFanId sets FanId field to given value.

### HasFanId

`func (o *EquipmentFan) HasFanId() bool`

HasFanId returns a boolean if a field has been set.

### GetFanModuleId

`func (o *EquipmentFan) GetFanModuleId() int64`

GetFanModuleId returns the FanModuleId field if non-nil, zero value otherwise.

### GetFanModuleIdOk

`func (o *EquipmentFan) GetFanModuleIdOk() (*int64, bool)`

GetFanModuleIdOk returns a tuple with the FanModuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanModuleId

`func (o *EquipmentFan) SetFanModuleId(v int64)`

SetFanModuleId sets FanModuleId field to given value.

### HasFanModuleId

`func (o *EquipmentFan) HasFanModuleId() bool`

HasFanModuleId returns a boolean if a field has been set.

### GetModuleId

`func (o *EquipmentFan) GetModuleId() int64`

GetModuleId returns the ModuleId field if non-nil, zero value otherwise.

### GetModuleIdOk

`func (o *EquipmentFan) GetModuleIdOk() (*int64, bool)`

GetModuleIdOk returns a tuple with the ModuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleId

`func (o *EquipmentFan) SetModuleId(v int64)`

SetModuleId sets ModuleId field to given value.

### HasModuleId

`func (o *EquipmentFan) HasModuleId() bool`

HasModuleId returns a boolean if a field has been set.

### GetOperReason

`func (o *EquipmentFan) GetOperReason() []string`

GetOperReason returns the OperReason field if non-nil, zero value otherwise.

### GetOperReasonOk

`func (o *EquipmentFan) GetOperReasonOk() (*[]string, bool)`

GetOperReasonOk returns a tuple with the OperReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperReason

`func (o *EquipmentFan) SetOperReason(v []string)`

SetOperReason sets OperReason field to given value.

### HasOperReason

`func (o *EquipmentFan) HasOperReason() bool`

HasOperReason returns a boolean if a field has been set.

### SetOperReasonNil

`func (o *EquipmentFan) SetOperReasonNil(b bool)`

 SetOperReasonNil sets the value for OperReason to be an explicit nil

### UnsetOperReason
`func (o *EquipmentFan) UnsetOperReason()`

UnsetOperReason ensures that no value is present for OperReason, not even an explicit nil
### GetOperState

`func (o *EquipmentFan) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *EquipmentFan) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *EquipmentFan) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *EquipmentFan) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetPartNumber

`func (o *EquipmentFan) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *EquipmentFan) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *EquipmentFan) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *EquipmentFan) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetPid

`func (o *EquipmentFan) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *EquipmentFan) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *EquipmentFan) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *EquipmentFan) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetSku

`func (o *EquipmentFan) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *EquipmentFan) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *EquipmentFan) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *EquipmentFan) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetTrayId

`func (o *EquipmentFan) GetTrayId() int64`

GetTrayId returns the TrayId field if non-nil, zero value otherwise.

### GetTrayIdOk

`func (o *EquipmentFan) GetTrayIdOk() (*int64, bool)`

GetTrayIdOk returns a tuple with the TrayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrayId

`func (o *EquipmentFan) SetTrayId(v int64)`

SetTrayId sets TrayId field to given value.

### HasTrayId

`func (o *EquipmentFan) HasTrayId() bool`

HasTrayId returns a boolean if a field has been set.

### GetVid

`func (o *EquipmentFan) GetVid() string`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *EquipmentFan) GetVidOk() (*string, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *EquipmentFan) SetVid(v string)`

SetVid sets Vid field to given value.

### HasVid

`func (o *EquipmentFan) HasVid() bool`

HasVid returns a boolean if a field has been set.

### GetEquipmentFanModule

`func (o *EquipmentFan) GetEquipmentFanModule() EquipmentFanModuleRelationship`

GetEquipmentFanModule returns the EquipmentFanModule field if non-nil, zero value otherwise.

### GetEquipmentFanModuleOk

`func (o *EquipmentFan) GetEquipmentFanModuleOk() (*EquipmentFanModuleRelationship, bool)`

GetEquipmentFanModuleOk returns a tuple with the EquipmentFanModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentFanModule

`func (o *EquipmentFan) SetEquipmentFanModule(v EquipmentFanModuleRelationship)`

SetEquipmentFanModule sets EquipmentFanModule field to given value.

### HasEquipmentFanModule

`func (o *EquipmentFan) HasEquipmentFanModule() bool`

HasEquipmentFanModule returns a boolean if a field has been set.

### GetEquipmentFex

`func (o *EquipmentFan) GetEquipmentFex() EquipmentFexRelationship`

GetEquipmentFex returns the EquipmentFex field if non-nil, zero value otherwise.

### GetEquipmentFexOk

`func (o *EquipmentFan) GetEquipmentFexOk() (*EquipmentFexRelationship, bool)`

GetEquipmentFexOk returns a tuple with the EquipmentFex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentFex

`func (o *EquipmentFan) SetEquipmentFex(v EquipmentFexRelationship)`

SetEquipmentFex sets EquipmentFex field to given value.

### HasEquipmentFex

`func (o *EquipmentFan) HasEquipmentFex() bool`

HasEquipmentFex returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *EquipmentFan) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *EquipmentFan) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *EquipmentFan) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *EquipmentFan) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *EquipmentFan) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EquipmentFan) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EquipmentFan) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EquipmentFan) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


