# BiosUnitAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InitSeq** | Pointer to **string** | The initSeq of the equipment. | [optional] [readonly] 
**InitTs** | Pointer to **string** | The initTs of the equipment. | [optional] [readonly] 
**ComputeBlade** | Pointer to [**ComputeBladeRelationship**](compute.Blade.Relationship.md) |  | [optional] 
**ComputeRackUnit** | Pointer to [**ComputeRackUnitRelationship**](compute.RackUnit.Relationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**RunningFirmware** | Pointer to [**[]FirmwareRunningFirmwareRelationship**](firmware.RunningFirmware.Relationship.md) | An array of relationships to firmwareRunningFirmware resources. | [optional] [readonly] 
**SystemBootOrder** | Pointer to [**BiosSystemBootOrderRelationship**](bios.SystemBootOrder.Relationship.md) |  | [optional] 

## Methods

### NewBiosUnitAllOf

`func NewBiosUnitAllOf() *BiosUnitAllOf`

NewBiosUnitAllOf instantiates a new BiosUnitAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiosUnitAllOfWithDefaults

`func NewBiosUnitAllOfWithDefaults() *BiosUnitAllOf`

NewBiosUnitAllOfWithDefaults instantiates a new BiosUnitAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInitSeq

`func (o *BiosUnitAllOf) GetInitSeq() string`

GetInitSeq returns the InitSeq field if non-nil, zero value otherwise.

### GetInitSeqOk

`func (o *BiosUnitAllOf) GetInitSeqOk() (*string, bool)`

GetInitSeqOk returns a tuple with the InitSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitSeq

`func (o *BiosUnitAllOf) SetInitSeq(v string)`

SetInitSeq sets InitSeq field to given value.

### HasInitSeq

`func (o *BiosUnitAllOf) HasInitSeq() bool`

HasInitSeq returns a boolean if a field has been set.

### GetInitTs

`func (o *BiosUnitAllOf) GetInitTs() string`

GetInitTs returns the InitTs field if non-nil, zero value otherwise.

### GetInitTsOk

`func (o *BiosUnitAllOf) GetInitTsOk() (*string, bool)`

GetInitTsOk returns a tuple with the InitTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitTs

`func (o *BiosUnitAllOf) SetInitTs(v string)`

SetInitTs sets InitTs field to given value.

### HasInitTs

`func (o *BiosUnitAllOf) HasInitTs() bool`

HasInitTs returns a boolean if a field has been set.

### GetComputeBlade

`func (o *BiosUnitAllOf) GetComputeBlade() ComputeBladeRelationship`

GetComputeBlade returns the ComputeBlade field if non-nil, zero value otherwise.

### GetComputeBladeOk

`func (o *BiosUnitAllOf) GetComputeBladeOk() (*ComputeBladeRelationship, bool)`

GetComputeBladeOk returns a tuple with the ComputeBlade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBlade

`func (o *BiosUnitAllOf) SetComputeBlade(v ComputeBladeRelationship)`

SetComputeBlade sets ComputeBlade field to given value.

### HasComputeBlade

`func (o *BiosUnitAllOf) HasComputeBlade() bool`

HasComputeBlade returns a boolean if a field has been set.

### GetComputeRackUnit

`func (o *BiosUnitAllOf) GetComputeRackUnit() ComputeRackUnitRelationship`

GetComputeRackUnit returns the ComputeRackUnit field if non-nil, zero value otherwise.

### GetComputeRackUnitOk

`func (o *BiosUnitAllOf) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool)`

GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRackUnit

`func (o *BiosUnitAllOf) SetComputeRackUnit(v ComputeRackUnitRelationship)`

SetComputeRackUnit sets ComputeRackUnit field to given value.

### HasComputeRackUnit

`func (o *BiosUnitAllOf) HasComputeRackUnit() bool`

HasComputeRackUnit returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *BiosUnitAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *BiosUnitAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *BiosUnitAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *BiosUnitAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *BiosUnitAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *BiosUnitAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *BiosUnitAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *BiosUnitAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetRunningFirmware

`func (o *BiosUnitAllOf) GetRunningFirmware() []FirmwareRunningFirmwareRelationship`

GetRunningFirmware returns the RunningFirmware field if non-nil, zero value otherwise.

### GetRunningFirmwareOk

`func (o *BiosUnitAllOf) GetRunningFirmwareOk() (*[]FirmwareRunningFirmwareRelationship, bool)`

GetRunningFirmwareOk returns a tuple with the RunningFirmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningFirmware

`func (o *BiosUnitAllOf) SetRunningFirmware(v []FirmwareRunningFirmwareRelationship)`

SetRunningFirmware sets RunningFirmware field to given value.

### HasRunningFirmware

`func (o *BiosUnitAllOf) HasRunningFirmware() bool`

HasRunningFirmware returns a boolean if a field has been set.

### SetRunningFirmwareNil

`func (o *BiosUnitAllOf) SetRunningFirmwareNil(b bool)`

 SetRunningFirmwareNil sets the value for RunningFirmware to be an explicit nil

### UnsetRunningFirmware
`func (o *BiosUnitAllOf) UnsetRunningFirmware()`

UnsetRunningFirmware ensures that no value is present for RunningFirmware, not even an explicit nil
### GetSystemBootOrder

`func (o *BiosUnitAllOf) GetSystemBootOrder() BiosSystemBootOrderRelationship`

GetSystemBootOrder returns the SystemBootOrder field if non-nil, zero value otherwise.

### GetSystemBootOrderOk

`func (o *BiosUnitAllOf) GetSystemBootOrderOk() (*BiosSystemBootOrderRelationship, bool)`

GetSystemBootOrderOk returns a tuple with the SystemBootOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemBootOrder

`func (o *BiosUnitAllOf) SetSystemBootOrder(v BiosSystemBootOrderRelationship)`

SetSystemBootOrder sets SystemBootOrder field to given value.

### HasSystemBootOrder

`func (o *BiosUnitAllOf) HasSystemBootOrder() bool`

HasSystemBootOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


