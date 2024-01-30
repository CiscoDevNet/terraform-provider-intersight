# EquipmentChassisOperationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "equipment.ChassisOperation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "equipment.ChassisOperation"]
**AdminLocatorLedAction** | Pointer to **string** | User configured state of the locator LED for the Chassis. * &#x60;None&#x60; - No operation action for the Locator Led of an equipment. * &#x60;TurnOn&#x60; - Turn on the Locator Led of an equipment. * &#x60;TurnOff&#x60; - Turn off the Locator Led of an equipment. | [optional] [default to "None"]
**AdminPowerCycleSlotId** | Pointer to **int64** | Slot id of the chassis slot that needs to be power cycled. | [optional] 
**ChassisOperationStatus** | Pointer to [**[]EquipmentChassisOperationStatus**](EquipmentChassisOperationStatus.md) |  | [optional] 
**ConfigState** | Pointer to **string** | The configured state of these settings in the target chassis. The value is any one of Applied, Applying, Failed. Applied - This state denotes that the settings are applied successfully in the target chassis. Applying - This state denotes that the settings are being applied in the target chassis. Failed - This state denotes that the settings could not be applied in the target chassis. * &#x60;None&#x60; - Nil value when no action has been triggered by the user. * &#x60;Applied&#x60; - User configured settings are in applied state. * &#x60;Applying&#x60; - User settings are being applied on the target server. * &#x60;Failed&#x60; - User configured settings could not be applied. | [optional] [readonly] [default to "None"]
**Chassis** | Pointer to [**EquipmentChassisRelationship**](EquipmentChassisRelationship.md) |  | [optional] 
**DeviceRegistration** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewEquipmentChassisOperationAllOf

`func NewEquipmentChassisOperationAllOf(classId string, objectType string, ) *EquipmentChassisOperationAllOf`

NewEquipmentChassisOperationAllOf instantiates a new EquipmentChassisOperationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentChassisOperationAllOfWithDefaults

`func NewEquipmentChassisOperationAllOfWithDefaults() *EquipmentChassisOperationAllOf`

NewEquipmentChassisOperationAllOfWithDefaults instantiates a new EquipmentChassisOperationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentChassisOperationAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentChassisOperationAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentChassisOperationAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentChassisOperationAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentChassisOperationAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentChassisOperationAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminLocatorLedAction

`func (o *EquipmentChassisOperationAllOf) GetAdminLocatorLedAction() string`

GetAdminLocatorLedAction returns the AdminLocatorLedAction field if non-nil, zero value otherwise.

### GetAdminLocatorLedActionOk

`func (o *EquipmentChassisOperationAllOf) GetAdminLocatorLedActionOk() (*string, bool)`

GetAdminLocatorLedActionOk returns a tuple with the AdminLocatorLedAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminLocatorLedAction

`func (o *EquipmentChassisOperationAllOf) SetAdminLocatorLedAction(v string)`

SetAdminLocatorLedAction sets AdminLocatorLedAction field to given value.

### HasAdminLocatorLedAction

`func (o *EquipmentChassisOperationAllOf) HasAdminLocatorLedAction() bool`

HasAdminLocatorLedAction returns a boolean if a field has been set.

### GetAdminPowerCycleSlotId

`func (o *EquipmentChassisOperationAllOf) GetAdminPowerCycleSlotId() int64`

GetAdminPowerCycleSlotId returns the AdminPowerCycleSlotId field if non-nil, zero value otherwise.

### GetAdminPowerCycleSlotIdOk

`func (o *EquipmentChassisOperationAllOf) GetAdminPowerCycleSlotIdOk() (*int64, bool)`

GetAdminPowerCycleSlotIdOk returns a tuple with the AdminPowerCycleSlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPowerCycleSlotId

`func (o *EquipmentChassisOperationAllOf) SetAdminPowerCycleSlotId(v int64)`

SetAdminPowerCycleSlotId sets AdminPowerCycleSlotId field to given value.

### HasAdminPowerCycleSlotId

`func (o *EquipmentChassisOperationAllOf) HasAdminPowerCycleSlotId() bool`

HasAdminPowerCycleSlotId returns a boolean if a field has been set.

### GetChassisOperationStatus

`func (o *EquipmentChassisOperationAllOf) GetChassisOperationStatus() []EquipmentChassisOperationStatus`

GetChassisOperationStatus returns the ChassisOperationStatus field if non-nil, zero value otherwise.

### GetChassisOperationStatusOk

`func (o *EquipmentChassisOperationAllOf) GetChassisOperationStatusOk() (*[]EquipmentChassisOperationStatus, bool)`

GetChassisOperationStatusOk returns a tuple with the ChassisOperationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisOperationStatus

`func (o *EquipmentChassisOperationAllOf) SetChassisOperationStatus(v []EquipmentChassisOperationStatus)`

SetChassisOperationStatus sets ChassisOperationStatus field to given value.

### HasChassisOperationStatus

`func (o *EquipmentChassisOperationAllOf) HasChassisOperationStatus() bool`

HasChassisOperationStatus returns a boolean if a field has been set.

### SetChassisOperationStatusNil

`func (o *EquipmentChassisOperationAllOf) SetChassisOperationStatusNil(b bool)`

 SetChassisOperationStatusNil sets the value for ChassisOperationStatus to be an explicit nil

### UnsetChassisOperationStatus
`func (o *EquipmentChassisOperationAllOf) UnsetChassisOperationStatus()`

UnsetChassisOperationStatus ensures that no value is present for ChassisOperationStatus, not even an explicit nil
### GetConfigState

`func (o *EquipmentChassisOperationAllOf) GetConfigState() string`

GetConfigState returns the ConfigState field if non-nil, zero value otherwise.

### GetConfigStateOk

`func (o *EquipmentChassisOperationAllOf) GetConfigStateOk() (*string, bool)`

GetConfigStateOk returns a tuple with the ConfigState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigState

`func (o *EquipmentChassisOperationAllOf) SetConfigState(v string)`

SetConfigState sets ConfigState field to given value.

### HasConfigState

`func (o *EquipmentChassisOperationAllOf) HasConfigState() bool`

HasConfigState returns a boolean if a field has been set.

### GetChassis

`func (o *EquipmentChassisOperationAllOf) GetChassis() EquipmentChassisRelationship`

GetChassis returns the Chassis field if non-nil, zero value otherwise.

### GetChassisOk

`func (o *EquipmentChassisOperationAllOf) GetChassisOk() (*EquipmentChassisRelationship, bool)`

GetChassisOk returns a tuple with the Chassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassis

`func (o *EquipmentChassisOperationAllOf) SetChassis(v EquipmentChassisRelationship)`

SetChassis sets Chassis field to given value.

### HasChassis

`func (o *EquipmentChassisOperationAllOf) HasChassis() bool`

HasChassis returns a boolean if a field has been set.

### GetDeviceRegistration

`func (o *EquipmentChassisOperationAllOf) GetDeviceRegistration() AssetDeviceRegistrationRelationship`

GetDeviceRegistration returns the DeviceRegistration field if non-nil, zero value otherwise.

### GetDeviceRegistrationOk

`func (o *EquipmentChassisOperationAllOf) GetDeviceRegistrationOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceRegistrationOk returns a tuple with the DeviceRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRegistration

`func (o *EquipmentChassisOperationAllOf) SetDeviceRegistration(v AssetDeviceRegistrationRelationship)`

SetDeviceRegistration sets DeviceRegistration field to given value.

### HasDeviceRegistration

`func (o *EquipmentChassisOperationAllOf) HasDeviceRegistration() bool`

HasDeviceRegistration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


