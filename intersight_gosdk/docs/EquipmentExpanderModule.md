# EquipmentExpanderModule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "equipment.ExpanderModule"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "equipment.ExpanderModule"]
**ModuleId** | Pointer to **int64** | Module identifier for the expander module. | [optional] [readonly] 
**OperReason** | Pointer to **[]string** |  | [optional] 
**OperState** | Pointer to **string** | Operational state of expander module. | [optional] [readonly] 
**PartNumber** | Pointer to **string** | Part number identifier for the expander module. | [optional] 
**EquipmentChassis** | Pointer to [**NullableEquipmentChassisRelationship**](EquipmentChassisRelationship.md) |  | [optional] 
**FanModules** | Pointer to [**[]EquipmentFanModuleRelationship**](EquipmentFanModuleRelationship.md) | An array of relationships to equipmentFanModule resources. | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewEquipmentExpanderModule

`func NewEquipmentExpanderModule(classId string, objectType string, ) *EquipmentExpanderModule`

NewEquipmentExpanderModule instantiates a new EquipmentExpanderModule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentExpanderModuleWithDefaults

`func NewEquipmentExpanderModuleWithDefaults() *EquipmentExpanderModule`

NewEquipmentExpanderModuleWithDefaults instantiates a new EquipmentExpanderModule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentExpanderModule) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentExpanderModule) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentExpanderModule) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentExpanderModule) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentExpanderModule) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentExpanderModule) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetModuleId

`func (o *EquipmentExpanderModule) GetModuleId() int64`

GetModuleId returns the ModuleId field if non-nil, zero value otherwise.

### GetModuleIdOk

`func (o *EquipmentExpanderModule) GetModuleIdOk() (*int64, bool)`

GetModuleIdOk returns a tuple with the ModuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleId

`func (o *EquipmentExpanderModule) SetModuleId(v int64)`

SetModuleId sets ModuleId field to given value.

### HasModuleId

`func (o *EquipmentExpanderModule) HasModuleId() bool`

HasModuleId returns a boolean if a field has been set.

### GetOperReason

`func (o *EquipmentExpanderModule) GetOperReason() []string`

GetOperReason returns the OperReason field if non-nil, zero value otherwise.

### GetOperReasonOk

`func (o *EquipmentExpanderModule) GetOperReasonOk() (*[]string, bool)`

GetOperReasonOk returns a tuple with the OperReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperReason

`func (o *EquipmentExpanderModule) SetOperReason(v []string)`

SetOperReason sets OperReason field to given value.

### HasOperReason

`func (o *EquipmentExpanderModule) HasOperReason() bool`

HasOperReason returns a boolean if a field has been set.

### SetOperReasonNil

`func (o *EquipmentExpanderModule) SetOperReasonNil(b bool)`

 SetOperReasonNil sets the value for OperReason to be an explicit nil

### UnsetOperReason
`func (o *EquipmentExpanderModule) UnsetOperReason()`

UnsetOperReason ensures that no value is present for OperReason, not even an explicit nil
### GetOperState

`func (o *EquipmentExpanderModule) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *EquipmentExpanderModule) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *EquipmentExpanderModule) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *EquipmentExpanderModule) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetPartNumber

`func (o *EquipmentExpanderModule) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *EquipmentExpanderModule) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *EquipmentExpanderModule) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *EquipmentExpanderModule) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetEquipmentChassis

`func (o *EquipmentExpanderModule) GetEquipmentChassis() EquipmentChassisRelationship`

GetEquipmentChassis returns the EquipmentChassis field if non-nil, zero value otherwise.

### GetEquipmentChassisOk

`func (o *EquipmentExpanderModule) GetEquipmentChassisOk() (*EquipmentChassisRelationship, bool)`

GetEquipmentChassisOk returns a tuple with the EquipmentChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentChassis

`func (o *EquipmentExpanderModule) SetEquipmentChassis(v EquipmentChassisRelationship)`

SetEquipmentChassis sets EquipmentChassis field to given value.

### HasEquipmentChassis

`func (o *EquipmentExpanderModule) HasEquipmentChassis() bool`

HasEquipmentChassis returns a boolean if a field has been set.

### SetEquipmentChassisNil

`func (o *EquipmentExpanderModule) SetEquipmentChassisNil(b bool)`

 SetEquipmentChassisNil sets the value for EquipmentChassis to be an explicit nil

### UnsetEquipmentChassis
`func (o *EquipmentExpanderModule) UnsetEquipmentChassis()`

UnsetEquipmentChassis ensures that no value is present for EquipmentChassis, not even an explicit nil
### GetFanModules

`func (o *EquipmentExpanderModule) GetFanModules() []EquipmentFanModuleRelationship`

GetFanModules returns the FanModules field if non-nil, zero value otherwise.

### GetFanModulesOk

`func (o *EquipmentExpanderModule) GetFanModulesOk() (*[]EquipmentFanModuleRelationship, bool)`

GetFanModulesOk returns a tuple with the FanModules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanModules

`func (o *EquipmentExpanderModule) SetFanModules(v []EquipmentFanModuleRelationship)`

SetFanModules sets FanModules field to given value.

### HasFanModules

`func (o *EquipmentExpanderModule) HasFanModules() bool`

HasFanModules returns a boolean if a field has been set.

### SetFanModulesNil

`func (o *EquipmentExpanderModule) SetFanModulesNil(b bool)`

 SetFanModulesNil sets the value for FanModules to be an explicit nil

### UnsetFanModules
`func (o *EquipmentExpanderModule) UnsetFanModules()`

UnsetFanModules ensures that no value is present for FanModules, not even an explicit nil
### GetRegisteredDevice

`func (o *EquipmentExpanderModule) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EquipmentExpanderModule) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EquipmentExpanderModule) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EquipmentExpanderModule) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *EquipmentExpanderModule) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *EquipmentExpanderModule) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


