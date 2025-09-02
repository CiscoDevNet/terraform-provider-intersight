# EtherHostPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "ether.HostPort"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "ether.HostPort"]
**AdminState** | Pointer to **string** | Administratively configured state (enabled/disabled) for this port. | [optional] [readonly] 
**AggregatePortId** | Pointer to **int64** | Breakout port member in the fabric extender. | [optional] [readonly] 
**ModuleId** | Pointer to **int64** | Fabric extender identifier for this port. | [optional] [readonly] 
**Speed** | Pointer to **string** | Host Port Speed of IO card or fabric extender. | [optional] [readonly] 
**EquipmentIoCardBase** | Pointer to [**NullableEquipmentIoCardBaseRelationship**](EquipmentIoCardBaseRelationship.md) |  | [optional] 
**EquipmentSwitchCard** | Pointer to [**NullableEquipmentSwitchCardRelationship**](EquipmentSwitchCardRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewEtherHostPort

`func NewEtherHostPort(classId string, objectType string, ) *EtherHostPort`

NewEtherHostPort instantiates a new EtherHostPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEtherHostPortWithDefaults

`func NewEtherHostPortWithDefaults() *EtherHostPort`

NewEtherHostPortWithDefaults instantiates a new EtherHostPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EtherHostPort) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EtherHostPort) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EtherHostPort) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EtherHostPort) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EtherHostPort) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EtherHostPort) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminState

`func (o *EtherHostPort) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *EtherHostPort) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *EtherHostPort) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *EtherHostPort) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetAggregatePortId

`func (o *EtherHostPort) GetAggregatePortId() int64`

GetAggregatePortId returns the AggregatePortId field if non-nil, zero value otherwise.

### GetAggregatePortIdOk

`func (o *EtherHostPort) GetAggregatePortIdOk() (*int64, bool)`

GetAggregatePortIdOk returns a tuple with the AggregatePortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatePortId

`func (o *EtherHostPort) SetAggregatePortId(v int64)`

SetAggregatePortId sets AggregatePortId field to given value.

### HasAggregatePortId

`func (o *EtherHostPort) HasAggregatePortId() bool`

HasAggregatePortId returns a boolean if a field has been set.

### GetModuleId

`func (o *EtherHostPort) GetModuleId() int64`

GetModuleId returns the ModuleId field if non-nil, zero value otherwise.

### GetModuleIdOk

`func (o *EtherHostPort) GetModuleIdOk() (*int64, bool)`

GetModuleIdOk returns a tuple with the ModuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleId

`func (o *EtherHostPort) SetModuleId(v int64)`

SetModuleId sets ModuleId field to given value.

### HasModuleId

`func (o *EtherHostPort) HasModuleId() bool`

HasModuleId returns a boolean if a field has been set.

### GetSpeed

`func (o *EtherHostPort) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *EtherHostPort) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *EtherHostPort) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *EtherHostPort) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetEquipmentIoCardBase

`func (o *EtherHostPort) GetEquipmentIoCardBase() EquipmentIoCardBaseRelationship`

GetEquipmentIoCardBase returns the EquipmentIoCardBase field if non-nil, zero value otherwise.

### GetEquipmentIoCardBaseOk

`func (o *EtherHostPort) GetEquipmentIoCardBaseOk() (*EquipmentIoCardBaseRelationship, bool)`

GetEquipmentIoCardBaseOk returns a tuple with the EquipmentIoCardBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentIoCardBase

`func (o *EtherHostPort) SetEquipmentIoCardBase(v EquipmentIoCardBaseRelationship)`

SetEquipmentIoCardBase sets EquipmentIoCardBase field to given value.

### HasEquipmentIoCardBase

`func (o *EtherHostPort) HasEquipmentIoCardBase() bool`

HasEquipmentIoCardBase returns a boolean if a field has been set.

### SetEquipmentIoCardBaseNil

`func (o *EtherHostPort) SetEquipmentIoCardBaseNil(b bool)`

 SetEquipmentIoCardBaseNil sets the value for EquipmentIoCardBase to be an explicit nil

### UnsetEquipmentIoCardBase
`func (o *EtherHostPort) UnsetEquipmentIoCardBase()`

UnsetEquipmentIoCardBase ensures that no value is present for EquipmentIoCardBase, not even an explicit nil
### GetEquipmentSwitchCard

`func (o *EtherHostPort) GetEquipmentSwitchCard() EquipmentSwitchCardRelationship`

GetEquipmentSwitchCard returns the EquipmentSwitchCard field if non-nil, zero value otherwise.

### GetEquipmentSwitchCardOk

`func (o *EtherHostPort) GetEquipmentSwitchCardOk() (*EquipmentSwitchCardRelationship, bool)`

GetEquipmentSwitchCardOk returns a tuple with the EquipmentSwitchCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentSwitchCard

`func (o *EtherHostPort) SetEquipmentSwitchCard(v EquipmentSwitchCardRelationship)`

SetEquipmentSwitchCard sets EquipmentSwitchCard field to given value.

### HasEquipmentSwitchCard

`func (o *EtherHostPort) HasEquipmentSwitchCard() bool`

HasEquipmentSwitchCard returns a boolean if a field has been set.

### SetEquipmentSwitchCardNil

`func (o *EtherHostPort) SetEquipmentSwitchCardNil(b bool)`

 SetEquipmentSwitchCardNil sets the value for EquipmentSwitchCard to be an explicit nil

### UnsetEquipmentSwitchCard
`func (o *EtherHostPort) UnsetEquipmentSwitchCard()`

UnsetEquipmentSwitchCard ensures that no value is present for EquipmentSwitchCard, not even an explicit nil
### GetRegisteredDevice

`func (o *EtherHostPort) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EtherHostPort) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EtherHostPort) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EtherHostPort) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *EtherHostPort) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *EtherHostPort) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


