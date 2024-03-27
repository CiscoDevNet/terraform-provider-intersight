# EquipmentHybridDriveSlotAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "equipment.HybridDriveSlot"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "equipment.HybridDriveSlot"]
**CurrentMode** | Pointer to **string** | The Configured Mode of the Hybrid Drive slot. * &#x60;&#x60; - Hybrid Drive slot  mode is not applicable. * &#x60;RAID&#x60; - Hybrid Drive slot mode is RAID. * &#x60;Direct&#x60; - Hybrid Drive slot mode is Direct. | [optional] [readonly] [default to ""]
**RequestedMode** | Pointer to **string** | The Requested Mode for the Hybrid Drive slot. * &#x60;&#x60; - Hybrid Drive slot  mode is not applicable. * &#x60;RAID&#x60; - Hybrid Drive slot mode is RAID. * &#x60;Direct&#x60; - Hybrid Drive slot mode is Direct. | [optional] [readonly] [default to ""]
**ComputeBlade** | Pointer to [**ComputeBladeRelationship**](ComputeBladeRelationship.md) |  | [optional] 
**ComputeBoard** | Pointer to [**ComputeBoardRelationship**](ComputeBoardRelationship.md) |  | [optional] 
**ComputeRackUnit** | Pointer to [**ComputeRackUnitRelationship**](ComputeRackUnitRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewEquipmentHybridDriveSlotAllOf

`func NewEquipmentHybridDriveSlotAllOf(classId string, objectType string, ) *EquipmentHybridDriveSlotAllOf`

NewEquipmentHybridDriveSlotAllOf instantiates a new EquipmentHybridDriveSlotAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentHybridDriveSlotAllOfWithDefaults

`func NewEquipmentHybridDriveSlotAllOfWithDefaults() *EquipmentHybridDriveSlotAllOf`

NewEquipmentHybridDriveSlotAllOfWithDefaults instantiates a new EquipmentHybridDriveSlotAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentHybridDriveSlotAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentHybridDriveSlotAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentHybridDriveSlotAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentHybridDriveSlotAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentHybridDriveSlotAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentHybridDriveSlotAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCurrentMode

`func (o *EquipmentHybridDriveSlotAllOf) GetCurrentMode() string`

GetCurrentMode returns the CurrentMode field if non-nil, zero value otherwise.

### GetCurrentModeOk

`func (o *EquipmentHybridDriveSlotAllOf) GetCurrentModeOk() (*string, bool)`

GetCurrentModeOk returns a tuple with the CurrentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentMode

`func (o *EquipmentHybridDriveSlotAllOf) SetCurrentMode(v string)`

SetCurrentMode sets CurrentMode field to given value.

### HasCurrentMode

`func (o *EquipmentHybridDriveSlotAllOf) HasCurrentMode() bool`

HasCurrentMode returns a boolean if a field has been set.

### GetRequestedMode

`func (o *EquipmentHybridDriveSlotAllOf) GetRequestedMode() string`

GetRequestedMode returns the RequestedMode field if non-nil, zero value otherwise.

### GetRequestedModeOk

`func (o *EquipmentHybridDriveSlotAllOf) GetRequestedModeOk() (*string, bool)`

GetRequestedModeOk returns a tuple with the RequestedMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedMode

`func (o *EquipmentHybridDriveSlotAllOf) SetRequestedMode(v string)`

SetRequestedMode sets RequestedMode field to given value.

### HasRequestedMode

`func (o *EquipmentHybridDriveSlotAllOf) HasRequestedMode() bool`

HasRequestedMode returns a boolean if a field has been set.

### GetComputeBlade

`func (o *EquipmentHybridDriveSlotAllOf) GetComputeBlade() ComputeBladeRelationship`

GetComputeBlade returns the ComputeBlade field if non-nil, zero value otherwise.

### GetComputeBladeOk

`func (o *EquipmentHybridDriveSlotAllOf) GetComputeBladeOk() (*ComputeBladeRelationship, bool)`

GetComputeBladeOk returns a tuple with the ComputeBlade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBlade

`func (o *EquipmentHybridDriveSlotAllOf) SetComputeBlade(v ComputeBladeRelationship)`

SetComputeBlade sets ComputeBlade field to given value.

### HasComputeBlade

`func (o *EquipmentHybridDriveSlotAllOf) HasComputeBlade() bool`

HasComputeBlade returns a boolean if a field has been set.

### GetComputeBoard

`func (o *EquipmentHybridDriveSlotAllOf) GetComputeBoard() ComputeBoardRelationship`

GetComputeBoard returns the ComputeBoard field if non-nil, zero value otherwise.

### GetComputeBoardOk

`func (o *EquipmentHybridDriveSlotAllOf) GetComputeBoardOk() (*ComputeBoardRelationship, bool)`

GetComputeBoardOk returns a tuple with the ComputeBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBoard

`func (o *EquipmentHybridDriveSlotAllOf) SetComputeBoard(v ComputeBoardRelationship)`

SetComputeBoard sets ComputeBoard field to given value.

### HasComputeBoard

`func (o *EquipmentHybridDriveSlotAllOf) HasComputeBoard() bool`

HasComputeBoard returns a boolean if a field has been set.

### GetComputeRackUnit

`func (o *EquipmentHybridDriveSlotAllOf) GetComputeRackUnit() ComputeRackUnitRelationship`

GetComputeRackUnit returns the ComputeRackUnit field if non-nil, zero value otherwise.

### GetComputeRackUnitOk

`func (o *EquipmentHybridDriveSlotAllOf) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool)`

GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRackUnit

`func (o *EquipmentHybridDriveSlotAllOf) SetComputeRackUnit(v ComputeRackUnitRelationship)`

SetComputeRackUnit sets ComputeRackUnit field to given value.

### HasComputeRackUnit

`func (o *EquipmentHybridDriveSlotAllOf) HasComputeRackUnit() bool`

HasComputeRackUnit returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *EquipmentHybridDriveSlotAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EquipmentHybridDriveSlotAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EquipmentHybridDriveSlotAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EquipmentHybridDriveSlotAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


