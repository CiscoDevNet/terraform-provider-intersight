# BiosBootMode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "bios.BootMode"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "bios.BootMode"]
**ActualBootMode** | Pointer to **string** | The actual BIOS boot mode as reported by the platform BIOS. | [optional] 
**ComputeBlade** | Pointer to [**NullableComputeBladeRelationship**](ComputeBladeRelationship.md) |  | [optional] 
**ComputeRackUnit** | Pointer to [**NullableComputeRackUnitRelationship**](ComputeRackUnitRelationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**NullableInventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewBiosBootMode

`func NewBiosBootMode(classId string, objectType string, ) *BiosBootMode`

NewBiosBootMode instantiates a new BiosBootMode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiosBootModeWithDefaults

`func NewBiosBootModeWithDefaults() *BiosBootMode`

NewBiosBootModeWithDefaults instantiates a new BiosBootMode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BiosBootMode) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BiosBootMode) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BiosBootMode) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BiosBootMode) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BiosBootMode) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BiosBootMode) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetActualBootMode

`func (o *BiosBootMode) GetActualBootMode() string`

GetActualBootMode returns the ActualBootMode field if non-nil, zero value otherwise.

### GetActualBootModeOk

`func (o *BiosBootMode) GetActualBootModeOk() (*string, bool)`

GetActualBootModeOk returns a tuple with the ActualBootMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualBootMode

`func (o *BiosBootMode) SetActualBootMode(v string)`

SetActualBootMode sets ActualBootMode field to given value.

### HasActualBootMode

`func (o *BiosBootMode) HasActualBootMode() bool`

HasActualBootMode returns a boolean if a field has been set.

### GetComputeBlade

`func (o *BiosBootMode) GetComputeBlade() ComputeBladeRelationship`

GetComputeBlade returns the ComputeBlade field if non-nil, zero value otherwise.

### GetComputeBladeOk

`func (o *BiosBootMode) GetComputeBladeOk() (*ComputeBladeRelationship, bool)`

GetComputeBladeOk returns a tuple with the ComputeBlade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBlade

`func (o *BiosBootMode) SetComputeBlade(v ComputeBladeRelationship)`

SetComputeBlade sets ComputeBlade field to given value.

### HasComputeBlade

`func (o *BiosBootMode) HasComputeBlade() bool`

HasComputeBlade returns a boolean if a field has been set.

### SetComputeBladeNil

`func (o *BiosBootMode) SetComputeBladeNil(b bool)`

 SetComputeBladeNil sets the value for ComputeBlade to be an explicit nil

### UnsetComputeBlade
`func (o *BiosBootMode) UnsetComputeBlade()`

UnsetComputeBlade ensures that no value is present for ComputeBlade, not even an explicit nil
### GetComputeRackUnit

`func (o *BiosBootMode) GetComputeRackUnit() ComputeRackUnitRelationship`

GetComputeRackUnit returns the ComputeRackUnit field if non-nil, zero value otherwise.

### GetComputeRackUnitOk

`func (o *BiosBootMode) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool)`

GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRackUnit

`func (o *BiosBootMode) SetComputeRackUnit(v ComputeRackUnitRelationship)`

SetComputeRackUnit sets ComputeRackUnit field to given value.

### HasComputeRackUnit

`func (o *BiosBootMode) HasComputeRackUnit() bool`

HasComputeRackUnit returns a boolean if a field has been set.

### SetComputeRackUnitNil

`func (o *BiosBootMode) SetComputeRackUnitNil(b bool)`

 SetComputeRackUnitNil sets the value for ComputeRackUnit to be an explicit nil

### UnsetComputeRackUnit
`func (o *BiosBootMode) UnsetComputeRackUnit()`

UnsetComputeRackUnit ensures that no value is present for ComputeRackUnit, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *BiosBootMode) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *BiosBootMode) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *BiosBootMode) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *BiosBootMode) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### SetInventoryDeviceInfoNil

`func (o *BiosBootMode) SetInventoryDeviceInfoNil(b bool)`

 SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil

### UnsetInventoryDeviceInfo
`func (o *BiosBootMode) UnsetInventoryDeviceInfo()`

UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
### GetRegisteredDevice

`func (o *BiosBootMode) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *BiosBootMode) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *BiosBootMode) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *BiosBootMode) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *BiosBootMode) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *BiosBootMode) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


