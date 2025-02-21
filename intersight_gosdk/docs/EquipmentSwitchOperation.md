# EquipmentSwitchOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "equipment.SwitchOperation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "equipment.SwitchOperation"]
**AdminEvacState** | Pointer to **string** | Sets evacuation state of the switch. When evacuation is enabled, data traffic flowing through this switch will be suspended for all the servers. Fabric evacuation can be enabled during any maintenance activity on the switch in order to gracefully failover data flows to the peer switch. * &#x60;Disabled&#x60; - Admin configured Disabled State. * &#x60;Enabled&#x60; - Admin configured Enabled State. | [optional] [default to "Disabled"]
**AdminLocatorLedAction** | Pointer to **string** | Action performed on the locator LED of the switch. * &#x60;None&#x60; - No operation action for the Locator Led of an equipment. * &#x60;TurnOn&#x60; - Turn on the Locator Led of an equipment. * &#x60;TurnOff&#x60; - Turn off the Locator Led of an equipment. | [optional] [default to "None"]
**AdminLocatorLedActionState** | Pointer to **string** | Defines status of action performed on AdminLocatorLedState. * &#x60;None&#x60; - Nil value when no action has been triggered by the user. * &#x60;Applied&#x60; - User configured settings are in applied state. * &#x60;Applying&#x60; - User settings are being applied on the target server. * &#x60;Failed&#x60; - User configured settings could not be applied. | [optional] [default to "None"]
**ConfigEvacState** | Pointer to **string** | Captures the status of evacuation on this switch. * &#x60;None&#x60; - Nil value when no action has been triggered by the user. * &#x60;Applied&#x60; - User configured settings are in applied state. * &#x60;Applying&#x60; - User settings are being applied on the target server. * &#x60;Failed&#x60; - User configured settings could not be applied. | [optional] [default to "None"]
**ForceEvac** | Pointer to **bool** | Evacuation is blocked by the system if it can cause a traffic outage in the domain. Select \&quot;Force Evacuation\&quot; only if system rejects the operation and you want to override that. | [optional] 
**ResetAction** | Pointer to [**NullableEquipmentResetAction**](EquipmentResetAction.md) |  | [optional] 
**ResetActionState** | Pointer to **string** | Current status of the reset operation executed on the Fabric Interconnect. * &#x60;None&#x60; - Nil value when no action has been triggered by the user. * &#x60;Applied&#x60; - User configured settings are in applied state. * &#x60;Applying&#x60; - User settings are being applied on the target server. * &#x60;Failed&#x60; - User configured settings could not be applied. | [optional] [default to "None"]
**SwitchName** | Pointer to **string** | Name of the switch on which the switch operation is performed. | [optional] [readonly] 
**DeviceRegistration** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**NetworkElement** | Pointer to [**NullableNetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 

## Methods

### NewEquipmentSwitchOperation

`func NewEquipmentSwitchOperation(classId string, objectType string, ) *EquipmentSwitchOperation`

NewEquipmentSwitchOperation instantiates a new EquipmentSwitchOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentSwitchOperationWithDefaults

`func NewEquipmentSwitchOperationWithDefaults() *EquipmentSwitchOperation`

NewEquipmentSwitchOperationWithDefaults instantiates a new EquipmentSwitchOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentSwitchOperation) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentSwitchOperation) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentSwitchOperation) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentSwitchOperation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentSwitchOperation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentSwitchOperation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminEvacState

`func (o *EquipmentSwitchOperation) GetAdminEvacState() string`

GetAdminEvacState returns the AdminEvacState field if non-nil, zero value otherwise.

### GetAdminEvacStateOk

`func (o *EquipmentSwitchOperation) GetAdminEvacStateOk() (*string, bool)`

GetAdminEvacStateOk returns a tuple with the AdminEvacState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminEvacState

`func (o *EquipmentSwitchOperation) SetAdminEvacState(v string)`

SetAdminEvacState sets AdminEvacState field to given value.

### HasAdminEvacState

`func (o *EquipmentSwitchOperation) HasAdminEvacState() bool`

HasAdminEvacState returns a boolean if a field has been set.

### GetAdminLocatorLedAction

`func (o *EquipmentSwitchOperation) GetAdminLocatorLedAction() string`

GetAdminLocatorLedAction returns the AdminLocatorLedAction field if non-nil, zero value otherwise.

### GetAdminLocatorLedActionOk

`func (o *EquipmentSwitchOperation) GetAdminLocatorLedActionOk() (*string, bool)`

GetAdminLocatorLedActionOk returns a tuple with the AdminLocatorLedAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminLocatorLedAction

`func (o *EquipmentSwitchOperation) SetAdminLocatorLedAction(v string)`

SetAdminLocatorLedAction sets AdminLocatorLedAction field to given value.

### HasAdminLocatorLedAction

`func (o *EquipmentSwitchOperation) HasAdminLocatorLedAction() bool`

HasAdminLocatorLedAction returns a boolean if a field has been set.

### GetAdminLocatorLedActionState

`func (o *EquipmentSwitchOperation) GetAdminLocatorLedActionState() string`

GetAdminLocatorLedActionState returns the AdminLocatorLedActionState field if non-nil, zero value otherwise.

### GetAdminLocatorLedActionStateOk

`func (o *EquipmentSwitchOperation) GetAdminLocatorLedActionStateOk() (*string, bool)`

GetAdminLocatorLedActionStateOk returns a tuple with the AdminLocatorLedActionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminLocatorLedActionState

`func (o *EquipmentSwitchOperation) SetAdminLocatorLedActionState(v string)`

SetAdminLocatorLedActionState sets AdminLocatorLedActionState field to given value.

### HasAdminLocatorLedActionState

`func (o *EquipmentSwitchOperation) HasAdminLocatorLedActionState() bool`

HasAdminLocatorLedActionState returns a boolean if a field has been set.

### GetConfigEvacState

`func (o *EquipmentSwitchOperation) GetConfigEvacState() string`

GetConfigEvacState returns the ConfigEvacState field if non-nil, zero value otherwise.

### GetConfigEvacStateOk

`func (o *EquipmentSwitchOperation) GetConfigEvacStateOk() (*string, bool)`

GetConfigEvacStateOk returns a tuple with the ConfigEvacState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigEvacState

`func (o *EquipmentSwitchOperation) SetConfigEvacState(v string)`

SetConfigEvacState sets ConfigEvacState field to given value.

### HasConfigEvacState

`func (o *EquipmentSwitchOperation) HasConfigEvacState() bool`

HasConfigEvacState returns a boolean if a field has been set.

### GetForceEvac

`func (o *EquipmentSwitchOperation) GetForceEvac() bool`

GetForceEvac returns the ForceEvac field if non-nil, zero value otherwise.

### GetForceEvacOk

`func (o *EquipmentSwitchOperation) GetForceEvacOk() (*bool, bool)`

GetForceEvacOk returns a tuple with the ForceEvac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceEvac

`func (o *EquipmentSwitchOperation) SetForceEvac(v bool)`

SetForceEvac sets ForceEvac field to given value.

### HasForceEvac

`func (o *EquipmentSwitchOperation) HasForceEvac() bool`

HasForceEvac returns a boolean if a field has been set.

### GetResetAction

`func (o *EquipmentSwitchOperation) GetResetAction() EquipmentResetAction`

GetResetAction returns the ResetAction field if non-nil, zero value otherwise.

### GetResetActionOk

`func (o *EquipmentSwitchOperation) GetResetActionOk() (*EquipmentResetAction, bool)`

GetResetActionOk returns a tuple with the ResetAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetAction

`func (o *EquipmentSwitchOperation) SetResetAction(v EquipmentResetAction)`

SetResetAction sets ResetAction field to given value.

### HasResetAction

`func (o *EquipmentSwitchOperation) HasResetAction() bool`

HasResetAction returns a boolean if a field has been set.

### SetResetActionNil

`func (o *EquipmentSwitchOperation) SetResetActionNil(b bool)`

 SetResetActionNil sets the value for ResetAction to be an explicit nil

### UnsetResetAction
`func (o *EquipmentSwitchOperation) UnsetResetAction()`

UnsetResetAction ensures that no value is present for ResetAction, not even an explicit nil
### GetResetActionState

`func (o *EquipmentSwitchOperation) GetResetActionState() string`

GetResetActionState returns the ResetActionState field if non-nil, zero value otherwise.

### GetResetActionStateOk

`func (o *EquipmentSwitchOperation) GetResetActionStateOk() (*string, bool)`

GetResetActionStateOk returns a tuple with the ResetActionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetActionState

`func (o *EquipmentSwitchOperation) SetResetActionState(v string)`

SetResetActionState sets ResetActionState field to given value.

### HasResetActionState

`func (o *EquipmentSwitchOperation) HasResetActionState() bool`

HasResetActionState returns a boolean if a field has been set.

### GetSwitchName

`func (o *EquipmentSwitchOperation) GetSwitchName() string`

GetSwitchName returns the SwitchName field if non-nil, zero value otherwise.

### GetSwitchNameOk

`func (o *EquipmentSwitchOperation) GetSwitchNameOk() (*string, bool)`

GetSwitchNameOk returns a tuple with the SwitchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchName

`func (o *EquipmentSwitchOperation) SetSwitchName(v string)`

SetSwitchName sets SwitchName field to given value.

### HasSwitchName

`func (o *EquipmentSwitchOperation) HasSwitchName() bool`

HasSwitchName returns a boolean if a field has been set.

### GetDeviceRegistration

`func (o *EquipmentSwitchOperation) GetDeviceRegistration() AssetDeviceRegistrationRelationship`

GetDeviceRegistration returns the DeviceRegistration field if non-nil, zero value otherwise.

### GetDeviceRegistrationOk

`func (o *EquipmentSwitchOperation) GetDeviceRegistrationOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceRegistrationOk returns a tuple with the DeviceRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRegistration

`func (o *EquipmentSwitchOperation) SetDeviceRegistration(v AssetDeviceRegistrationRelationship)`

SetDeviceRegistration sets DeviceRegistration field to given value.

### HasDeviceRegistration

`func (o *EquipmentSwitchOperation) HasDeviceRegistration() bool`

HasDeviceRegistration returns a boolean if a field has been set.

### SetDeviceRegistrationNil

`func (o *EquipmentSwitchOperation) SetDeviceRegistrationNil(b bool)`

 SetDeviceRegistrationNil sets the value for DeviceRegistration to be an explicit nil

### UnsetDeviceRegistration
`func (o *EquipmentSwitchOperation) UnsetDeviceRegistration()`

UnsetDeviceRegistration ensures that no value is present for DeviceRegistration, not even an explicit nil
### GetNetworkElement

`func (o *EquipmentSwitchOperation) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *EquipmentSwitchOperation) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *EquipmentSwitchOperation) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *EquipmentSwitchOperation) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### SetNetworkElementNil

`func (o *EquipmentSwitchOperation) SetNetworkElementNil(b bool)`

 SetNetworkElementNil sets the value for NetworkElement to be an explicit nil

### UnsetNetworkElement
`func (o *EquipmentSwitchOperation) UnsetNetworkElement()`

UnsetNetworkElement ensures that no value is present for NetworkElement, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


