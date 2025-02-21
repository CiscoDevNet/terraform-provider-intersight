# EquipmentHybridDriveSlot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "equipment.HybridDriveSlot"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "equipment.HybridDriveSlot"]
**CurrentMode** | Pointer to **string** | The Configured Mode of the Hybrid Drive slot. * &#x60;&#x60; - Hybrid Drive slot mode is not applicable. * &#x60;Controller&#x60; - Hybrid Drive slot mode is Controller. * &#x60;Direct&#x60; - Hybrid Drive slot mode is Direct. | [optional] [readonly] [default to ""]
**RequestedMode** | Pointer to **string** | The Requested Mode for the Hybrid Drive slot. * &#x60;&#x60; - Hybrid Drive slot mode is not applicable. * &#x60;Controller&#x60; - Hybrid Drive slot mode is Controller. * &#x60;Direct&#x60; - Hybrid Drive slot mode is Direct. | [optional] [readonly] [default to ""]
**ComputeBlade** | Pointer to [**NullableComputeBladeRelationship**](ComputeBladeRelationship.md) |  | [optional] 
**ComputeBoard** | Pointer to [**NullableComputeBoardRelationship**](ComputeBoardRelationship.md) |  | [optional] 
**ComputeRackUnit** | Pointer to [**NullableComputeRackUnitRelationship**](ComputeRackUnitRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewEquipmentHybridDriveSlot

`func NewEquipmentHybridDriveSlot(classId string, objectType string, ) *EquipmentHybridDriveSlot`

NewEquipmentHybridDriveSlot instantiates a new EquipmentHybridDriveSlot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentHybridDriveSlotWithDefaults

`func NewEquipmentHybridDriveSlotWithDefaults() *EquipmentHybridDriveSlot`

NewEquipmentHybridDriveSlotWithDefaults instantiates a new EquipmentHybridDriveSlot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentHybridDriveSlot) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentHybridDriveSlot) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentHybridDriveSlot) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentHybridDriveSlot) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentHybridDriveSlot) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentHybridDriveSlot) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCurrentMode

`func (o *EquipmentHybridDriveSlot) GetCurrentMode() string`

GetCurrentMode returns the CurrentMode field if non-nil, zero value otherwise.

### GetCurrentModeOk

`func (o *EquipmentHybridDriveSlot) GetCurrentModeOk() (*string, bool)`

GetCurrentModeOk returns a tuple with the CurrentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentMode

`func (o *EquipmentHybridDriveSlot) SetCurrentMode(v string)`

SetCurrentMode sets CurrentMode field to given value.

### HasCurrentMode

`func (o *EquipmentHybridDriveSlot) HasCurrentMode() bool`

HasCurrentMode returns a boolean if a field has been set.

### GetRequestedMode

`func (o *EquipmentHybridDriveSlot) GetRequestedMode() string`

GetRequestedMode returns the RequestedMode field if non-nil, zero value otherwise.

### GetRequestedModeOk

`func (o *EquipmentHybridDriveSlot) GetRequestedModeOk() (*string, bool)`

GetRequestedModeOk returns a tuple with the RequestedMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedMode

`func (o *EquipmentHybridDriveSlot) SetRequestedMode(v string)`

SetRequestedMode sets RequestedMode field to given value.

### HasRequestedMode

`func (o *EquipmentHybridDriveSlot) HasRequestedMode() bool`

HasRequestedMode returns a boolean if a field has been set.

### GetComputeBlade

`func (o *EquipmentHybridDriveSlot) GetComputeBlade() ComputeBladeRelationship`

GetComputeBlade returns the ComputeBlade field if non-nil, zero value otherwise.

### GetComputeBladeOk

`func (o *EquipmentHybridDriveSlot) GetComputeBladeOk() (*ComputeBladeRelationship, bool)`

GetComputeBladeOk returns a tuple with the ComputeBlade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBlade

`func (o *EquipmentHybridDriveSlot) SetComputeBlade(v ComputeBladeRelationship)`

SetComputeBlade sets ComputeBlade field to given value.

### HasComputeBlade

`func (o *EquipmentHybridDriveSlot) HasComputeBlade() bool`

HasComputeBlade returns a boolean if a field has been set.

### SetComputeBladeNil

`func (o *EquipmentHybridDriveSlot) SetComputeBladeNil(b bool)`

 SetComputeBladeNil sets the value for ComputeBlade to be an explicit nil

### UnsetComputeBlade
`func (o *EquipmentHybridDriveSlot) UnsetComputeBlade()`

UnsetComputeBlade ensures that no value is present for ComputeBlade, not even an explicit nil
### GetComputeBoard

`func (o *EquipmentHybridDriveSlot) GetComputeBoard() ComputeBoardRelationship`

GetComputeBoard returns the ComputeBoard field if non-nil, zero value otherwise.

### GetComputeBoardOk

`func (o *EquipmentHybridDriveSlot) GetComputeBoardOk() (*ComputeBoardRelationship, bool)`

GetComputeBoardOk returns a tuple with the ComputeBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBoard

`func (o *EquipmentHybridDriveSlot) SetComputeBoard(v ComputeBoardRelationship)`

SetComputeBoard sets ComputeBoard field to given value.

### HasComputeBoard

`func (o *EquipmentHybridDriveSlot) HasComputeBoard() bool`

HasComputeBoard returns a boolean if a field has been set.

### SetComputeBoardNil

`func (o *EquipmentHybridDriveSlot) SetComputeBoardNil(b bool)`

 SetComputeBoardNil sets the value for ComputeBoard to be an explicit nil

### UnsetComputeBoard
`func (o *EquipmentHybridDriveSlot) UnsetComputeBoard()`

UnsetComputeBoard ensures that no value is present for ComputeBoard, not even an explicit nil
### GetComputeRackUnit

`func (o *EquipmentHybridDriveSlot) GetComputeRackUnit() ComputeRackUnitRelationship`

GetComputeRackUnit returns the ComputeRackUnit field if non-nil, zero value otherwise.

### GetComputeRackUnitOk

`func (o *EquipmentHybridDriveSlot) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool)`

GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRackUnit

`func (o *EquipmentHybridDriveSlot) SetComputeRackUnit(v ComputeRackUnitRelationship)`

SetComputeRackUnit sets ComputeRackUnit field to given value.

### HasComputeRackUnit

`func (o *EquipmentHybridDriveSlot) HasComputeRackUnit() bool`

HasComputeRackUnit returns a boolean if a field has been set.

### SetComputeRackUnitNil

`func (o *EquipmentHybridDriveSlot) SetComputeRackUnitNil(b bool)`

 SetComputeRackUnitNil sets the value for ComputeRackUnit to be an explicit nil

### UnsetComputeRackUnit
`func (o *EquipmentHybridDriveSlot) UnsetComputeRackUnit()`

UnsetComputeRackUnit ensures that no value is present for ComputeRackUnit, not even an explicit nil
### GetRegisteredDevice

`func (o *EquipmentHybridDriveSlot) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EquipmentHybridDriveSlot) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EquipmentHybridDriveSlot) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EquipmentHybridDriveSlot) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *EquipmentHybridDriveSlot) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *EquipmentHybridDriveSlot) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


